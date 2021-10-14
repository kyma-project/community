# Concept

Goal: Telemetry operator should be extended to a kind of templating to be used in combination with the ServiceCatalog.

## Concept A

### Workflow
1. User creates ServiceBinding as usual
2. BTP Operator creates corresponding secret
3. User needs to fetch newly created secret name
4. User creates new CR providing only secret and plan name
5. Telemetry operator creates new FluentBitConfig based on service plan incl. new secret

#### Templating

To configure FluentBit we need a templating mechanism, which makes use of the provided secret and service plan.

User input for this CR:
- Service Plan: TO know which service the user is refering to
- Secret: Secret created by BTP Operator which needs to be imported to the FluentBit config
- Unique Identifier: 


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

## Concept B
### Workflow
1. User creates new `CustomResource` based on `FluentBitBinding`-CRD
2. This CR creates ServiceBinding CR (`services.cloud.sap.com/v1alpha1`)
   1. BTP Operator fetches this CR and creates a ServiceBinding to this Service
   2. BTP Operator creates a new secret with given name, specified in `FluentBitBinding`
3. Secret gets used to configure FluentBit instance

### CRD's

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
              serviceInstanceName: 
                description: '...'
                type: string
              externalName:
                description: '...'
                type: string
              secretName:
                description: '...'
                type: string
              parameters:
                description: '...'
                type: object 
```
</details>


<details>
<summary><b>ServiceBinding CR</b> already implemented in BTP Operator - Click to expand</summary>

```yaml
apiVersion: services.cloud.sap.com/v1alpha1
kind: ServiceBinding
metadata:
  name: my-binding
spec:
  serviceInstanceName: my-service-instance
  externalName: my-binding-external
  secretName: mySecret
  parameters:
    key1: val1
    key2: val2      
```
</details>

## Comparison
| | Concept A | Concept B|
|----|----|----|
|Workflow| Workflow for user is same as old, but requires additional step to configure FluentBit.| Instead of creating ServiceBinding, user creates FluentBitServiceBinding, no additional step needed|
|Upsides|<ul><li>Same Workflow as before to create ServiceBinding</li><li>No need to check if ServiceBinding is valid</li></ul> |<ul><li>No additional step needed &#8594; whole configuration automatic</li></ul>|
|Downsides|<ul><li>Additional step needed</li></ul>|<ul><li>Check for valid ServiceBinding is needed</li><li></li></ul>|


secret rotation wenn secret alt bitte neues erstellen und rotieren

ddotting scripte in cls doku 

lua script und dann weiter auf http output

Konzept:
CRD für templates
Daraus eine CLS CR template
User erstellt "mapping" und referenziert auf service, secret, und welches template