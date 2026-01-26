# Autogenerate CRD Documentation

Autogenerate CRD documentation directly from code rather than maintaining it manually. This approach reduces maintenance effort and ensures documentation stays in sync with code changes.

While you can choose any generator that suits your needs, prioritize solutions already tested and used across Kyma teams. Regardless of the tool you select, ensure the generated documentation adheres to the [Custom Resource Topic template](./01-docu-templates.md#custom-resource-topics) to maintain consistency across Kyma documentation.

## Kyma Table Generator (Recommended)

The table generator is the solution developed by the Kyma team. It allows you to generate one table with all the resource's parameters and add it to any document. 

For example, see [Telemetry](https://github.com/kyma-project/telemetry-manager/blob/main/docs/user/resources/01-telemetry.md?plain=1).

### Generate Documentation Tables

Each table consists of three columns:

- **Parameter** - with the paramater name
- **Type** - with the parameter type
- **Description** - with the parameter description

Generate the table automatically from the CRD specification file using the following steps:

1. Prepare the parameters' descriptions in the CR's specification file. For example, for the APIRule CR, prepare the description in [`apirules.gateway.crd.yaml`](https://github.com/kyma-project/kyma/blob/main/installation/resources/crds/api-gateway/apirules.gateway.crd.yaml).

2. Add the following mappings to [this Makefile](https://github.com/kyma-project/kyma/blob/main/hack/table-gen/Makefile):

- `--crd-filename` - full or relative path to the `.yaml` file with the CRD
- `--md-filename` - full or relative path to the `.md` file in which you want to generate the table

   See the existing examples for further details.

3. Set up the table generator in the `.md` file in which you want to generate the table. Add the `TABLE-START` and `TABLE-END` tags in the exact place in the document where you want to generate the table.

```
   <!-- TABLE-START -->

   <!-- TABLE-END -->
```

4. In the terminal, go to the [`hack/table-gen`](https://github.com/kyma-project/kyma/tree/main/hack/table-gen) directory and run the following command:

   ```bash
   make generate
   ```

For more details on the table generator, read the [`table-gen` documentation](https://github.com/kyma-project/kyma/blob/main/hack/table-gen/README.md) in the `kyma` repository.

## CRD Reference Documentation Generator

Some Kyma teams use the [crd-ref-docs](https://github.com/elastic/crd-ref-docs/tree/master) tool to autogenerate CRD documentation. Unlike the Kyma Table Generator, which creates a single table containing all fields, this tool generates multiple tables—one for each CRD schema. This structure allows direct linking to specific sections and makes it easier to navigate to where a particular schema is defined.

For example usage, see [Istio](https://github.com/kyma-project/istio/blob/main/docs/user/04-00-istio-custom-resource.md).


### Generate Documentation Tables


