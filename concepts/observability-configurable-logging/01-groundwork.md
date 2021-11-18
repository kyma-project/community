# Current Situation and Motivation

In the current shape of Kyma in year 2021, the observability stack is focussing on providing oppinionated solutions, working out of the box, to solve basic requirements for application operators. By that, it is not focussing on integration aspects in order to cover a broader and richer usage scenario.

![a](./assets/current_all.drawio.svg)

In the diagram you see that all three observability aspects (log, trace, metric) are providing a pre-configured backend with visualizations. However, they are not providing a neutral and unified way of integration of backends outside of the cluster. The tracing stack provides no way to centrally push trace data to the outside. Logging can be configured much more flexible and neutral, however, the configuration needs to be done at installation time to not get lost at the next kyma upgrade process. Furthermore it is not easy to use to mix and match different integrations as you need to deal with one centralized configuration (the fluent-bit config).

Integration (and with that changing the focus away from in-cluster backends) is the key to open up the stack for a broad range of use cases. Users can simply bring there own backends as they already use a commercial offering or run own infrastructure. The data can be stored outside of the cluster in a managed offering, shared with the data of multiple clusters, away from any tampering/deletion attempt of a hacker, ...

This concept proposes how to open up to that new scenarios by making integration possible at runtime in a convenient way. For that it will focus on the logging scenario only being ready to include the other data types at a later time.

# Requirements

The basic requirements can be condensed to this list:

- Basic backend configuration
  - Have a vendor-neutral agent layer being responsible for collecting and shipping the telemetry data, but not permanently storing it (as a backend)
  - Support configuration of the selected agent at runtime (no need to run a kyma upgrade process) in a scenario focused approach
  - Support multiple configurations at the same time in individual resources to enable easy activation of dedicated scenarios
  - one vendor-neutral input and output support is enough as a minimum, it should be possible to chain your custom agent for specific conversions (like for traces, supporting the OLTP protocol will support most of the vendors already. Chaining a custom otel-collector can do custom conversion to a specific protocol)
  - all typical settings for supported outputs of the used agent are supported (do not hide/abstract them)
  - the agent should run stable at anytime, bad configuration should be pre-validated and rejected, fast feedback is welcome
  - secrets must be kept secret

- Template definitions
  - Have a mechanism to provide templates/best practices for typical scenarios which can be innstantiated at runtime
  - Such template provides same feature-richness as a configuration scenario
  - If specific templates are bundled with kyma, updates of a template in use should propogate to the actual agent configuration
  - Supports placeholder definitions with default values and descriptions

- Template instantiation
  - A template gets instantiated by usually binding it to a secret, fullfilling placeholders like URL and credentials
  - Placeholders can be satisfied with configMaps and static values as well
  - It validates fullfilment of placeholders
  - A template can be instantiated for specific workload and/or namespaces only

# Proposed Solution

The idea of the proposal is to introduce a pre-configured agent layer being responsible for collecting all telemetry data. That agents can be dynamically configured at runtime with different configuration scenarios so that the agents will start shipping the data to the configured backends. The dynamic configuration and management of the agent will happen by a new operator which can be configured via k8s resources. The agents and the new operator will be bundled in a new core package called `telemetry`. The bundled backends and UIs will become just one possible solution which can be installed optional and will stay in the existing modules like `logging`.

![b](./assets/future_all.drawio.svg)

# Focus on Logging

While the section above was more general, in the following it will be focused on the logging aspect only. This is mainly to have an emerging approach and solving the most urgent topics first. As logs have the longest tradition, everyone is used and expects to have solutions for logging place.
The following documents will outline what agent technology gets used, what operator gets used and how the general API will look like.
