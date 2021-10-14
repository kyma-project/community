
# Concept

Goal: Telemetry operator should be extended to a kind of automatic templating to be used in combination with the ServiceCatalog.


The orginal workflow of a user creating a ServiceBinding is not being adapted, to have an easier workflow for the user plus there are no real upside of why we should change the user workflow. We also though of changing the workflow to have a more automated setup, but this would require more knowledge of the user to actually perform the ServiceBinding with the automated FluentBit Configuration.
Instead, the workflow will be expanded by one additional step to create the FluentBit configuration for the corresponding ServiceBinding.

This concept covers the basic requirements with a minimal setup and reusing the Telemetry Operator for this purpose. At a later point, the Telemetry Operator could be split up to have a cleaner operator architecture, but this needs to be discussed further in the future.


## General Architecture / Workflow:

1. User creates ServiceBinding using CLI, BTP, or Busola.
2. BTP Operator watches ServiceBinding-CRs and creates corresponding secret.
3. User Creates ServiceBindingConfig-CR and references to the secret created by the BTP Operator. Furthermore, the user needs to specify which Template should be used.
4. Telemetry Operator watches ServiceBindingConfig-CRs. If a new CR is created, it creates a new LoggingConfiguration-CR, using the template and the information given by the referenced secret.
5. Telemetry operator creates new FluentBitConfig based on LoggingConfiguration-CR

![Workflow Architecture](images/workflow-overview.svg)

## Templating

To map the key-value-pairs given by the referenced secret, we need a CRD which maps the keys of the secret to the corresponding FluentBit Output keys. Thus, the `Template`-CRD is needed. This CRD is defined by:
- Name
- Mapping from secret keys to FluentBit keys
- Which filter and output plugins of FluentBit need to be used
- Configuration of these filters and outputs

Kyma will then have pre-defined Template-CRs which the customer can us to create a FluentBit configuration based on the customer's created ServiceBinding. In this way, the users does not have to take care about maintaining filters (i.e. `Lua` scripts), configuration of outputs (i.e. `http`-Plugin), etc.


## Secret Rotation

In this concept the user refers to the secret created by the BTP-Operator, which is then mounted by the Telemetry-Operator in the corresponding container. Therefore, the secret rotation will happen automatically by the general function of the Telemetry Operator.

> **NOTE:** The secret rotation is not implemented in the Telemetry Operator until now (as of today; 14.20.2021).
