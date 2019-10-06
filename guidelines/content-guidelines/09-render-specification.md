---
title: Render specifications
---

There are guidelines for defining specifications that are rendered on the [Kyma website](https://kyma-project.io).

> **NOTE:** Currently, only the [OpenAPI](https://swagger.io/specification/) specification is supported and only in the [Docs](https://kyma-project.io/docs) pages.

## Add specification

To add specification to render, add new source in `sources` field of specific [ClusterDocsTopic CR](/docs/master/components/headless-cms/#custom-resource-cluster-docs-topic) in [this](https://github.com/kyma-project/kyma/tree/master/resources/core/charts/docs/charts/content-ui/templates) place as below:

``` yaml
sources:
  ...
  - type: {SPECIFICATION_TYPE}
    name: {SPECIFICATION_NAME}
    mode: single
    url: {SPECIFICATION_URL}
```

where:

- `SPECIFICATION_TYPE` defines a type of a given specification. Currently, only the [OpenAPI](https://swagger.io/specification/) specification is supported under `openapi` type.

- `SPECIFICATION_NAME` defines a unique identifier of a given specification. Based on this field, defined is the URL on https://kyma-project.io/docs under which the specification is displayed. For example: if specification is added in `application-connector` ClusterDocsTopic CR with `connectorapi` value of `name` field, then URL will be `https://kyma-project.io/docs/{VERSION_OF_DOCS}/components/application-connector/specifications/connectorapi/`. 

- `SPECIFICATION_URL` defines the location of a specification. It may contain directives used values defined in `values.yaml` files. For internal specifications, defined in [Kyma core](https://github.com/kyma-project/kyma) repository, is recommended to use directive with version of Kyma and Kyma organization name, for example:

  ``` yaml
  url: https://raw.githubusercontent.com/{{ .Values.global.kymaOrgName }}/kyma/{{ .Values.global.docs.clusterDocsTopicsVersion }}/docs/application-connector/assets/connectorapi.yaml
  ```

### Example:

``` yaml
sources:
  ...
  - type: openapi
    name: connectorapi
    mode: single
    url: https://raw.githubusercontent.com/{{ .Values.global.kymaOrgName }}/kyma/{{ .Values.global.docs.clusterDocsTopicsVersion }}/docs/application-connector/assets/connectorapi.yaml
```
