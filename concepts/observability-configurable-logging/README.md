# Dynamic Logging Backend Configuration

This folder contains pages for evaluating and proposing a concept for having the logging backends (and general telemetry backends) in a Kyma cluster configurable. It mainly tries to propose a concept for [Configurable logging #11236](https://github.com/kyma-project/kyma/issues/11236).

For that, the following pages guide you through the different steps of the proposal:
* [Ground work - current state, requirements, goal](./01-groundwork.md)
* [Comparison of existing log agents](./02-agents.md)
* [Comparison of existing FluentBit operators](./03-fluent-bit-operator1.md)
* [Another Comparison of existing FluentBit operators](./04-fluent-bit-operator2.md)
* [Proposal for a custom FluentBit operator](./05-custom-fluentbit-operator.md)
* [Templating Concept](./06-servicecatalog-integration.md) to enhance the operator with bindings to ServiceCatalog instances
