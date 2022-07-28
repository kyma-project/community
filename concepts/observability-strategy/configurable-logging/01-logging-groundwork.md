# Configurable Logging: Groundwork

## Current Situation and Motivation

![a](./assets/logging-current.drawio.svg)

In the current setup, the logging component provides a feature-rich end-to-end logging solution with Loki as backend and Grafana as log browser. Users can configure Fluent Bit, the log collector, flexibly and neutrally - but only during installation, otherwise the configuration is lost at the next Kyma upgrade. Mixing and matching different integrations is hard, because users must deal with one centralized configuration (the Fluent Bit config) at deployment.
Furthermore, Loki is provided in a non-scalabale setup and cannot be configured at runtime.

As outlined in the [general strategy](../strategy.md), integration (and with that, changing the focus away from in-cluster backends) is the key to open up the stack for a broad range of use cases. Users can simply bring their own backends if they already use a commercial offering or run their own infrastructure. To name just a few advantages, the data can be stored outside the cluster in a managed offering, shared with the data of multiple clusters, and kept away from any tampering or deletion attempts by hackers.

This concept proposes how to open up to those new scenarios by supporting convenient integration at runtime.

## Requirements

### Basic backend configuration
- Have a vendor-neutral layer of collectors that collects and ships the telemetry data, but does not permanently store it (as a backend)
- Support configuration of the selected collector at runtime (no need to run a Kyma upgrade process) in a scenario-focused approach
- Support multiple configurations at the same time in individual resources to enable easy activation of dedicated scenarios
- As a minimum, support one vendor-neutral input and output. It should be possible to chain your custom collector for specific conversions. For example, for traces, supporting the OLTP protocol will support most of the vendors already. Chaining a custom OpenTelemetry Collector can do custom conversion to a specific protocol.
- Support all typical settings for the supported outputs of the used collector, do not hide or abstract them
- The collector must run stable at any time. Bad configuration must be prevalidated and rejected. Fast feedback is welcome.
- Secrets must be kept secret.
- Scenarios must be isolated and have their own buffer management. If a backend is in a bad shape and cannot process any data anymore, data should still continue to be pushed to other backends.
- Typical authentication mechanisms for the integration must be supported, especially solutions based on client certificates.
- For logging: de-dotting for elasticsearch must be possible
- Filtering of irrelevant data (like dropping logs of kyma-system Namespace) must be possible.

### Template definitions
- Have a mechanism to provide templates and best practices for typical scenarios, which can be instantiated at runtime
- Such templates provide the same feature-richness as a configuration scenario.
- If specific templates are bundled with Kyma, updates of a template in use must propagate to the actual collector configuration.
  - Support placeholder definitions with default values and descriptions

### Template instantiation
- A template is instantiated by binding it to a secret that provides input for placeholders like URL and credentials. Alternatively, placeholders can be filled out with ConfigMaps and static values.
  - The template validates whether all placeholders have been replaced with values.
- A template can be instantiated for a specific workload and/or Namespace only.

### Local backend
- Kyma will provide a blueprint based on Helm for installing the typical Loki stack as example.
- The setup is not meant to be HA and scalable.

## Proposed solution

The proposal introduces a new preconfigured layer of collectors that's responsible for collecting all telemetry data. Users can configure those collectors dynamically at runtime with different configuration scenarios, so that the collectors start shipping the data to the configured backends. The dynamic configuration and management of a collector is handled by a new operator, which is configured with Kubernetes resources. The collectors and the new operator are bundled in a new core package called `telemetry`. The existing Kyma backends and UIs will be just one possible solution to integrate with. The user can install them manually following a blueprint.

![b](./assets/logging-future.drawio.svg)
