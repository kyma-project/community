# General Strategy

## Current Situation and Motivation

In the current shape of Kyma in 2022, the observability stack is focussing on providing an opinionated and lightweight solution, working out of the box, to solve basic requirements for application operators. It is not focussing on integration aspects in order to enable for example cross-runtime observations, integration of advanced analytic tools or simple re-usage of existing infrastructure. Also it does not provide a guide on how to extend the setup to become HA with historical storage of the data. Having very limited integration possibilities and providing a limited in-cluster solution only, it is missing a lot of usage scenarios and the opportunity to focus on enabling users for modern observability. Overall, the potential audience is very limited.

![a](./assets/current.drawio.svg)

The diagram shows that all three observability aspects (log, trace, metric) provide a preconfigured backend with visualisations. However, they don't provide a neutral and unified way to integrate backends outside of the cluster.
- The tracing stack provides no way to centrally push trace data to the outside.
- Logging can be configured much more flexibly and neutrally. However, users must apply the configuration during installation; otherwise it's lost at the next Kyma upgrade process. Furthermore, it is hard to mix and match different integrations, because you must deal with one centralized configuration (the Fluent Bit config).
- Monitoring is based on Prometheus, which is widely adopted but limited to the Prometheus protocols. Users can configure the monitoring stack only during the Kyma upgrade process.

Integration (and with that, changing the focus away from in-cluster backends) is the key to open up the stack for a broad range of use cases. Users can bring their own backends if they already use a commercial offering or run their own infrastructure. The data can be stored outside the cluster in a managed offering, shared with the data of multiple clusters, away from any tampering or deletion attempt of a hacker, to name just a few.

Providing ready-to-use in-cluster backends is necessarily opinionated, does not cover all usage scenarios, and does not fit into Kyma's goal of providing the Kubernetes building blocks to integrate into the the SAP ecosystem. Also, the licensing issues (particularly with Grafana and Loki, or ElasticSearch as an alternative backend technology) show that an opinionated stack is problematic. It's better to handle opinionation by integrating with actual managed services. 

The following strategy proposes a new approach for the Kyma observability stack by opening up to new scenarios by supporting convenient integration at runtime. At the same time, it reduces the focus on ready-to-use solutions running inside the cluster.

## Goals

### Stages of observability

Observability can be split up into:
1. Instrumentation of the users application
2. The collection and preprocessing of the signals from the user application and the surrounding infrastructure, including metadata enrichment
3. The delivery of the signal to an backend for storage
4. An optional aggregation of the signals
5. Storage of the signals
6. The analysis, querying, and dashboarding of the signals

![a](./assets/stages.drawio.svg)

Kyma is a runtime to operate actual workloads. To fulfill typical operational requirements, users must be able to observe the workload deployed to the runtime .

### Mandatory feature set

The mandatory parts for a Kyma runtime are the stages that must be happen inside the runtime. With that, the following points are the major goal of the Kyma feature set:
1. Instrumentation support: Users usually need custom code for the instrumentation of a workload to expose typical signals like logs, traces, and metrics. Kyma can support common aspects of these tasks with guidelines, best practice guides, and ways of auto-instrumentation (like Istio tracing).
2. Signal collection: The runtime must provide the basic infrastructure to collect the emitted signals (like a unified log collector). If the user has followed the provided guides, that infrastructure should collect the signals instantly. Special cases (like specific protocols not backed by guides and infrastructure) should be supported by simple customizations (plug in a custom trace converter). Furthermore, the infrastructure must already enrich the signals of the workloads with metadata of the infrastructure, because this cannot happen at a later stage.
3. Signal delivery: The signals must be consumed by some party, either running inside or outside the cluster. It can be an aggregating layer or the backend directly. In either case, the shipment of the data must be configurable and should be based on a neutral protocol, so that any aggregation layer or backend can be integrated.

### Optional feature set

Optionally, Kyma can provide features for the stages to the right (aggregation, storage, and analysis). Here, users usually must pick an opinionated solution because there is no common technology available. Kyma can support such opinionated solutions by pre-configuring either a starting point or a full-blown solution. However, these aspects will usually not cover all usage, operations, and scalability scenarios. Furthermore, huge investments are necessary to provide any additional backend solution.


### Goals: Summary

The goal of the Kyma observability stack should be the seamless enablement of the stages on the left, allowing the ingestion of the signals into any backend system, opening up plenty usage scenarios. Kyma should provide at least blueprints as guidance for the stages on the right, especially for enabling the SAP ecosystem.
To sum it up, the goals of Kyma observability should be:
1. Provide guides on instrumentation (with potential helpers and auto-instrumentation options).
2. Collect resulting signals instantly when guides are followed, provide customization options for special cases.
3. Ship the signals reliably to a configurable vendor-neutral destination.
4. Provide blueprints for integration with specific vendors.

## Architecture

The proposal introduces a new preconfigured agent layer that's responsible for collecting all telemetry data. Hereby, the collection is very dependent on the signal type. The agents can be configured at runtime with different signal pipelines with basic filtering (inclusion and exclusion of signals) and outputs, so that the agents start shipping the signals through the pipelines to the configured backends. The dynamic configuration and management of the agent is handled by a new operator, which is configured using Kubernetes resources. The agents and the new operator are bundled in a new core component called `telemetry`. The existing Kyma backends and UIs will be just one possible solution to integrate with. The users will still be able to install them manually with a blueprint, but they will no longer be part of the Kyma offering.

![b](./assets/future.drawio.svg)

As mentioned before, the technology and protocols for the signal collection are different for the signal types: Logs are tailed from container log files, metrics are pulled using the Prometheus format usually, and traces are pushed with OTLP. With that, also the pre-integration (so that typical signals are collected instantly) is different per type.
That's why the specific concepts for the different types are different, and are discussed in more detail in the individual concepts:
* [Concept - Configurable Logging](./configurable-logging/README.md) 
* [Concept - Configurable Monitoring](./configurable-monitoring/README.md)
* [Concept - Configurable Tracing](./configurable-tracing/README.md)

All concepts follow general rules and will provide harmonized user APIs:
- **By default, signals of user workloads are collected**: If a workload is instrumented as outlined in the best practice guide, the telemetry data is instantly available in a pipeline without further configuration. Just define a pipeline with an output and it works.
- **One input protocol per type**: Only one common input protocol (like OTLP) is supported per signal type. Kyma provides best practice guides on how to integrate with it.
- **Custom input protocols can be integrated**: When a workload doesn't use the supported protocol for signal exposure, users can have custom transformation to the supported protocol, usually with additional operational effort (like running a dedicated Otel collector sidecar or deployment in the user's Namespace)
- **Harmonized API**: Signal pipelines are defined by well-defined Kubernetes CRDs that belong to only one group and following similar semantics.
- **One output protocol per type**: Only one common output protocol (like OTLP) is supported per signal type.
- **Custom output protocols can be integrated**: Users can have custom transformation from the supported protocol, usually with additional operational effort (like running a dedicated Otel Collector sidecar or deployment in the users' Namespace). Furthermore, Kyma offers typical output plugins of the used collector technology as part of a pipeline definition; however, only with limited support.
- **Blueprints are fully integrated**: Blueprints for local backend deployments are integrated with the used pipeline mechanism, so that they work instantly.

## Execution

The outlined strategy shifts the Kyma observability stack heavily, and the transformation process must be executed stepwise by priorities:

1. **Transform Kiali component into a blueprint**: Kiali is a very valuable tool for visualizing the Istio service mesh. It is based and fully dependent on Istio metrics from Prometheus and the Kubernetes APIServer, and brings integrations into Jaeger and Grafana. It is very specific to that toolset and needs to run within the cluster. The effort of providing a scalable HA setup fitting all usage scenarios is too high and distracts from the new focus of the Kyma observability stack. Thus, the Kiali component must be transformed into a blueprint based on the upstream kiali-operator Helm chart, providing a `values.yaml` file with instant integration. This step has high priority because it has no further dependency, low investment, and reduces maintenance efforts.
2. **Transform tracing component into a blueprint**: The tracing component bundles a Jaeger-all-in-one installation pre-integrated with Kyma's Istio, Serverless, and Eventing components. The bundle has a very lightweight setup that is mainly for demo purposes. It should be turned into a blueprint to save maintenance efforts. In the same step, Kyma must be adjusted to leverage Istio's telemetry API so that the telemetry can be activated at runtime.
2. [**Configurable Logging**](https://github.com/kyma-project/kyma/issues/11236): Introduce a configurable log collector and support log pipeline configurations at runtime with focus on the SAP ecosystem as MVP. Accessing application logs is the first and most common way of troubleshooting applications and is the minimal feature that users expect.
3. **Transform logging component into a blueprint**: As soon as the application logs can be out-streamed to external systems, Kyma can choose alternatives for the in-cluster logging solution. The Loki stack (Loki and Grafana) can be turned into a blueprint with the advantage of using the latest Loki versions (solving the license problem). Again, users get instant integration with the upstream Loki Helm chart and a provided `values.yaml` file.
4. [**Configurable Monitoring**](https://github.com/kyma-project/kyma/issues/13079): Introduce a configurable metrics collector that instantly scrapes all annotated workloads and supports the shipment of the metrics with a pipeline configuration. To cover a lot of providers instantly, the focus is on integrations based on OTLP. After logs, metrics are the next important feature to gain insights into distributed applications.
5. **Transform the monitoring component into a blueprint**: After finishing the configurable monitoring story, there is a good way to instrument and collect metrics of workloads independent of the Prometheus Operator stack. Now Kyma can introduce a blueprint to turn the actual Prometheus storage and Grafana visualization into a self-hosted component. Kyma-specific Grafana dashboards might be still bundled and maintained. All other dashboards are anyway based on the upstream kube-prometheus-stack, on which the monitoring blueprint will be based.
6. [**Configurable Tracing**](https://github.com/kyma-project/kyma/issues/11231): Switch the official trace propagation protocol to [W3C-tracecontext](https://www.w3.org/TR/trace-context/), and introduce a configurable trace collector based on OTLP. 
