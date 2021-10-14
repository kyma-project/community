

secret rotation wenn secret alt bitte neues erstellen und rotieren

ddotting scripte in cls doku 

lua script und dann weiter auf http output

Konzept:

User erstellt "mapping" und referenziert auf service, secret, und welches template

No need to check if ServiceBinding is valid

# Concept

Goal: Telemetry operator should be extended to a kind of templating to be used in combination with the ServiceCatalog.


The orginal workflow of a user creating a ServiceBinding is not being adapted, to have an easier workflow for the user plus there are no real upside of why we should change this flow. We also though of changing the workflow to have a more automated setup, but this would require more knowledge of the user to actually perform the ServiceBinding with the automated Fluentbit Configuration.

Instead, the workflow will be expanded by one additional step to create the Fluentbit configuration.

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

Kyma will then have pre-defined Template-CRs which the customer can us to create a FluentBit configuration based on the customer's created ServiceBinding. In this way, the users does not have to take care about maintaining `Lua` scripts, etc.


## Secret Rotation

<details>
<summary><b>ServiceBinding CR</b> needs to be implemented in Telemetry Operator - Click to expand</summary>

```yaml
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: binding.fluentbit.kyma-project.io
spec:
  group: fluentbit.kyma-project.io
  names:
    kind: FluentBitBinding
    listKind: fluentBitBindingList
    plural: fluentBitBindings
    singular: fluentBitBinding
  scope: Namespaced
  versions:
    - name: v1alpha1
      schema:
        description: FluentbitBinding is the Schema for the FluentbitBindings
          API.
        properties:
          [...]
          spec:
            description: FluentbitBindingSpec defines the desired state of FluentbitBinding.
            properties:
              servicePlanName: 
                description: '...'
                type: string
              secretName:
                description: '...'
                type: string
```
</details>