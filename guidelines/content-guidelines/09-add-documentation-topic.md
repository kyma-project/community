---
title: Add new documentation to the website
---
​
This document explains how to add a new topic to the Kyma documentation to render it on the [Kyma website](https://kyma-project.io). It also describes how to modify an existing topic if you want OpenAPI specifications to show in a given topic's documentation.
​
## Add a new documentation topic
​
Follow these steps:

1. Create a pull request with `.md` files for the new documentation topic. Place the `.md` files under a new `docs` subfolder in the `kyma` repository, such as `docs/api-gateway-v2/`.

2. In the same PR, create a `.yaml` file under the [`templates`](https://github.com/kyma-project/kyma/tree/master/resources/core/charts/docs/charts/content-ui/templates) folder to add a [ClusterDocsTopic CR](https://kyma-project.io/docs/components/headless-cms/#custom-resource-cluster-docs-topic) for your topic. For example, if you add a ClusterDocsTopic CR for the API Gateway v2 component, name it `docs-components-api-gateway-v2-cdt.yaml`. ​

    See the example definition:
    ​
    ``` yaml
    apiVersion: cms.kyma-project.io/v1alpha1
    kind: ClusterDocsTopic
    metadata:
      labels:
        cms.kyma-project.io/view-context: docs-ui
        cms.kyma-project.io/group-name: components
        cms.kyma-project.io/order: "11"
      name: api-gateway-v2
    spec:
      displayName: "API Gateway v2"
      description: "Overall documentation for API Gateway v2"
      sources:
       - type: markdown
         name: docs
         mode: package
         url: https://github.com/{{ .Values.global.kymaOrgName }}/kyma/archive/{{ .Values.global.docs.clusterDocsTopicsVersion }}.zip
         filter: /docs/api-gateway-v2/
    ```
​
3. Adjust values for these fields:
​
    - **cms.kyma-project.io/order** defines the number of your topic on the list in the left navigation, such as `"11"`. To add it correctly, check in other `.yaml` files which number is assigned to the last documentation topic in the navigation on the website, and add a consecutive number to your component. If you decide to modify the existing topic order, change values for this parameter in all other `.yaml` files accordingly to avoid duplicates.
    - **metadata.name** defines the CR name, such as `api-gateway-v2`.
    - **spec.displayname** defines the component name displayed on the website, such as `"API Gateway v2"`.
    - **spec.sources.filter** defines the location of the new topic's document sources, such as `/docs/api-gateway-v2/`.
​
4. Merge the changes and wait until the website is rebuilt.

​
## Add a single OpenAPI specification
​
In addition to Kyma documentation, there are also [OpenAPI](https://swagger.io/specification/) specifications rendered on the [Kyma website](https://kyma-project.io). You can find these specifications under the **API Consoles** type in the right navigation panel of a given component topic
​

To add a new specification, follow these steps:
​
1. Go to the [`templates`](https://github.com/kyma-project/kyma/tree/master/resources/core/charts/docs/charts/content-ui/templates) folder and locate an existing ClusterDocsTopic CR that you want to modify.
​
2. Add a new source entry in the **sources** field:
​
    ``` yaml
    sources:
      ...
      - type: {SPECIFICATION_TYPE}
        name: {SPECIFICATION_NAME}
        mode: single
        url: {SPECIFICATION_URL}
    ```
   ​
    where:
  ​
    - **{SPECIFICATION_TYPE}** defines a type of a given specification. Currently, only [OpenAPI](https://swagger.io/specification/) specifications are supported and they are defined under the `openapi` type.
  ​
    - **{SPECIFICATION_NAME}** defines a unique identifier of a given specification. This field defines the URL on https://kyma-project.io/docs under which the specification is displayed. For example, if the specification is added in the `application-connector` ClusterDocsTopic CR with the `connectorapi` value of the **name** field, its URL is `https://kyma-project.io/docs/{VERSION_OF_DOCS}/components/application-connector/specifications/connectorapi/`.
  ​
    - **{SPECIFICATION_URL}** defines the location of the specification. It may contain directives using values defined in `values.yaml` files. For internal specifications defined in the [`kyma`](https://github.com/kyma-project/kyma) repository, it is recommended to use the directive with a Kyma version and the organization name, such as:
 ​
    ``` yaml
    url: https://raw.githubusercontent.com/{{ .Values.global.kymaOrgName }}/kyma/{{ .Values.global.docs.clusterDocsTopicsVersion }}/docs/application-connector/assets/connectorapi.yaml
    ```

    See the example:

    ``` yaml
    sources:
      ...
      - type: openapi
        name: connectorapi
        mode: single
        url: https://raw.githubusercontent.com/{{ .Values.global.kymaOrgName }}/kyma/{{ .Values.global.docs.clusterDocsTopicsVersion }}/docs/application-connector/assets/connectorapi.yaml
    ```
​
3. Merge the changes and wait until the website is rebuilt.
