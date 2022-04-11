# Dynamic Fluent Bit Configuration

This document proposes and evaluates a valid design for a custom operator. It is needed to enable a dynamic in-cluster logging configuration with `Fluent Bit`, as outlined in [Spike: Dynamic configuration of logging backend #11105](https://github.com/kyma-project/kyma/issues/11105).

## Criteria
- Customers can place fluent-bit config snippets as k8s resource in any namespace.
- The customers' configuration must not be reset during reconciliation, even though the central configuration might be overwritten at any time.
- A new configuration is picked up dynamically without any user interaction (for example, no need that users restart pods).
- Have basic validation to give early feedback on errors, which supports debugging a startup problem.
- The proposal should be extensible to solve the same issue for OpenTelemetry collector.
- There's a way to provide the auth details for a backend config in a secure way.

## Proposal

### Architecture

To build the operator, we use an SDK for a quick and easy start-up, among other reasons. We decided to use [`kubebuilder`](https://github.com/kubernetes-sigs/kubebuilder), an upstream project of `Kubernetes`. `kubebuilder` provides the required features for our operator with the smallest overhead. In addition, this `kubebuilder` is also used in other parts of Kyma, ensuring consistent design.

An alternative for `kubebuilder` is the [`Operator-SDK`](https://github.com/operator-framework/operator-sdk) from `Red Hat`. `Operator-SDK` has similar features to the `kubebuilder`, with some additions like the integration of [`operatorhub`](https://operatorhub.io/). We decided against it because such features aren't needed for our purpose, and because it isn't an upstream project of Kubernetes.

To have a simple API, one `CRD` for Fluent Bit configuration is created. This CRD has a field that holds the status of the CR, called `Status`, as well as a struct of the type `LoggingConfigurationSpec`, called `Spec`, holding a list of configuration sections, called `Sections`. Each `Section` has a `content` attribute that holds a raw Fluent Bit configuration section. A `Section` can have optional `Files` and an `Environment`. `Files` are mounted into the Fluent Bit pods and can be referenced in configurations, for example, in Lua filters. The environment is a list of Kubernetes secret references.

Using this structure supports the full Fluent Bit syntax without needing to maintain new features of various plugins. Furthermore, the users get a clear overview of their sequence of applied filters and outputs. Using the Kyma documentation, we could also lead the users to think more in a way of pipelines, in such that they create one CR for each Fluent Bit pipeline.

<details>
<summary><b>Pipeline Overview</b> for User - Click to expand</summary>

![Thank you](assets/fluentbit_CR_overview.drawio.svg)
</details>

We're proposing the following constraints for the custom operator:
- It doesn't support dynamic plugins that must be loaded into the Fluent Bit image.

![Fluent Bit Pipeline Architecture](assets/fluentbit_dynamic_config.drawio.svg)

1. The logs are fetched using a Fluent Bit, which is created by this custom operator. This INPUT tags all logs with the following tag scheme: `kube.<namespace_name>.<pod_name>.<container_name>`.
2. The logs are enriched with Kubernetes information using the `kubernetes` filter of Fluent Bit.
3. The filter `rewrite_tag` is used to split the pipeline:
   - The original logs are forwarded to the Kyma Loki backend.
   - Users can use the new copy with another tag and configure new filters and outputs with the provided CRD.

Using this approach, we avoid having an unused INPUT or other overhead.
If users want to use more than one pipeline for the log processing, they can use the 'rewrite_tag' filter on their pipeline to create more pipelines. Alternatively, they can configure an output plugin to process them with another log processor, as mentioned before.

Additionally, when creating a `CR`, a [webhook](https://book.kubebuilder.io/cronjob-tutorial/webhook-implementation.html) validates the correctness of the Fluent Bit configuration based on the Fluent Bit `dry-run` feature.

To actually apply the changes of the users, the operator creates or adapts the ConfigMap for FluentBit and restarts Fluent Bit by deleting the pods of the Fluent Bit deployment.

To make sure that the configuration by the users won't be overwritten by reconciliation, the basic configuration (`kubernetes` filter, `rewrite_tag`, etc.) is written into a ConfigMap. This ConfigMap is embedded by an `@INCLUDE` statement in the chart of this operator.

The following example demonstrates the CRD that will be used by the telemetry operator:

```YAML
kind: LogPipeline
apiVersion: telemetry.kyma-project.io/v1alpha1
metadata:
  name: ElasticService-instanceXYZ
spec:
  parsers: {}
  multilineParsers:
    - content: |
        # Example from https://docs.fluentbit.io/manual/pipeline/filters/multiline-stacktrace
        name          multiline-custom-regex
        type          regex
        flush_timeout 1000
        rule      "start_state"   "/(Dec \d+ \d+\:\d+\:\d+)(.*)/"  "cont"
        rule      "cont"          "/^\s+at.*/"                     "cont" 
  filters:
    - content: |
        name                  multiline
        match                 *
        multiline.key_content log
        multiline.parser      go, multiline-custom-regex
    - content: |
        # Generated from selector in LogPresetBinding
        Name    grep
        Match   *
        Regex   $kubernetes['labels']['app'] my-deployment 
    - content: |
        Name    record_modifier
        Match   *
        Record  cluster_identifier ${KUBERNETES_SERVICE_HOST}
  outputs:
    - content: |
        Name               es
        Alias              es-output
        Match              *
        Host               ${ES_ENDPOINT} # Defined in Secret
        HTTP_User          ${ES_USER} # Defined in Secret
        HTTP_Password      ${ES_PASSWORD} # Defined in Secret
        LabelMapPath       /files/labelmap.json
  files:
    - name: labelmap.json
      content: |
      {
          "kubernetes": {
            "namespace_name": "namespace",
            "pod_name": "pod"
          },
          "stream": "stream"
      }
  secretRefs:
    - name: my-elastic-credentials
      namespace: default
    - name: ElasticService-static # Created by the Telemetry Operator to store static values from the LogPresetBinding
      namespace: default
```

The CRD contains a separate section for the different parts for the Fluent Bit configuration. The operator validates the configuration and merges static parts like predefined inputs, parsers or filters.

### Workflow for the User

To configure Fluent Bit, the user must create a new CR regarding to the CRD of this operator. Then, this operator will notice the new or changed CR, and will create or update a ConfigMap for Fluent Bit. Before the ConfigMap is applied, the operator uses the `dry-run` feature of Fluent Bit to validate the new configuration. If the check was successful, the new ConfigMap is applied and Fluent Bit is restarted.

## Fluent-Bit Log Routing

The LogPipeline CRD enables the user to define filter and output elements of a Fluent Bit pipeline. Therefore, a tag that can be consumed by user defined pipeline elements has to be provided by Kyma (either the static Fluent Bit configuration or the telemetry-operator). The different options to provide a log stream under a specific tag will be described in this section.

Log pipelines should be isolated in a way that a dysfunctional pipeline should not affect the availability of other pipelines. A pipeline can for instance become unavailable by an unavailable backend or faulty configuration.

The following requirements for the telemetry-operator should be considered when choosing a sufficient Fluent Bit configuration:
* Fluent Bit pods and the used image should be managed by Kyma
* User-provided configuration parts should be as close as possible to the Fluent Bit configuration format
* Filters can be used to
  * Parse workload specific log formats (for instance, multi-line exceptions)
  * Include or exclude parts of the logs (ship only a specific namespace to a backend)
  * Prepare the log metadata to comply with the backend's requirements (de-dotting for Elastic)

### Tags managed by the telemetry-operator

Properties:
* The user-defined LogPipeline elements must not contain any "match" attributes
* Pipeline elements that emit new tags (for instance rewrite_tag filters) must not be used
* The telemetry-operator adds ad "match" attribute to all filters and outputs

Consequences:
* Fill control over resource consumption since no additional buffers can be created
* Design goal to allow pasting existing Fluent Bit config samples is partially lost

### Managed tag per pipeline

Properties:
* The telemetry-operator provides a specific tag for each LogPipeline that has to be consumed by the pipeline sections
  * The telemetry-operator creates either a dedicated input plugin or buffered rewrite_tag filter to isolate the pipeline
* There is a defined way to consume the operator-provided log stream (for instance a tag placeholder or documented naming pattern)
* The LogPipeline can emit new tags using the rewrite_tag filter
* Potential overhead since complex pipelines have to be split to multiple simple pipelines, each having an own buffer

Consequences:
* Full flexibility to use all Fluent Bit concepts for the user
* The "contract" gives flexibility to change the implementation afterwards (for instance switch to an own input per pipeline or even an own DaemonSet)
* User can create elements that increase the resource consumption
* Allows describing complex pipelines and thus reduce the overall resource consumption

### Managed Fluent Bit DaemonSet per LogPipeline

Properties:
* The telemetry-operator creates a dedicated Fluent Bit DaemonSet per LogPipeline
* All DaemonSets have configured the same input and Kubernetes filter sections

Consequences:
* Best possible isolation between different pipelines
* Highest resource consumption
* Telemetry-operator has to manage the DaemonSets and not only ConfigMaps
* Support for custom output plugins might be added by the option to specify an own Fluent Bit image
 
### All LogPipelines use the same base-tag

Properties
* Current state of telemetry chart (2022-04-12)
* All LogPipelines match to kube.*

Consequences:
* Additional filters can be injected to any pipeline
  * Allows to modify default Loki output
* Low resource consumption, but user has the control to add additional buffers
* A dysfunctional output stalls all pipelines (no isolation)
