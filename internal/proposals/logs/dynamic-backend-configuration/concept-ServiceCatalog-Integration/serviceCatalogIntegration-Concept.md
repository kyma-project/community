
# Concept

Goal: Telemetry operator should be extended to a kind of automatic templating to be used in combination with the ServiceCatalog.


The original workflow for a user creating a ServiceBinding remains unchanged, which keeps the user workflow easier. There is no advantage of changing the user workflow. We also thought of changing the workflow to have a more automated setup, but this would require more knowledge of the user to actually perform the ServiceBinding with the automated FluentBit Configuration.
Instead, the workflow will be expanded by one additional step to create the FluentBit configuration for the corresponding ServiceBinding.

This concept covers the basic requirements with a minimal setup and reusing the Telemetry Operator for this purpose. At a later point, the Telemetry Operator could be split up to have a cleaner operator architecture, but this needs to be discussed further in the future.


## General Architecture / Workflow:

The following Workflow is explained in the context of the BTP Operator's ServiceBinding. The user can also refer to a Secret which is created by another instance or by the user. In this case, step one and two are skipped, and in step three the user refers to a custom Secret.

1. User creates ServiceBinding using CLI, BTP, or Busola.
2. BTP Operator watches ServiceBinding-CRs and creates corresponding Secret.
3. User Creates LoggingConfigurationInstance-CR and references to the Secret created by the BTP Operator. Furthermore, the user must specify which LoggingConfigurationTemplate should be used.
4. Telemetry Operator watches LoggingConfigurationInstance-CRs. If a new CR is created, it creates a new LoggingConfiguration-CR, using the LoggingConfigurationTemplate and the information given by the referenced Secret.
5. Telemetry operator creates new FluentBitConfig based on LoggingConfiguration-CR.

![Workflow Architecture](images/workflow-overview.svg)

## Templating

To map the key-value pairs given by the referenced Secret, we need a CRD that maps the keys of the Secret to the corresponding FluentBit Output keys. Thus, the `LoggingConfigurationTemplate`-CRD is needed. This CRD is defined by:
- Name
- Mapping from Secret keys to FluentBit keys
- The filter and output plugins of FluentBit that are to be used
- Configuration of these filters and outputs

Kyma will then have predefined LoggingConfigurationTemplate-CRs, which the customer can use to create a FluentBit configuration based on the customer's created ServiceBinding. In this way, the users does not have to care about maintaining filters (i.e. `Lua` scripts), configuration of outputs (i.e. `http`-Plugin), etc.


## Secret Rotation

In this concept, the user refers to the Secret created by the BTP Operator, which is then mounted by the Telemetry Operator in the corresponding container. Therefore, the Secret rotation happens automatically by the general function of the Telemetry Operator.

> **NOTE:** The Secret rotation is not implemented in the Telemetry Operator until now (as of today; 15.10.2021).
