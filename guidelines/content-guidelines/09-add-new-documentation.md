---
title: Add new documentation to the website
---
​
This document explains how to add a new topic to the documentation of repository under Kyma Project to render it on the [Kyma website](https://kyma-project.io). It also describes how to modify an existing topic if you want OpenAPI specifications to show in a given topic's documentation and additionally what to do to display the documentation from a given Kyma's repository on the [documentation view](https://kyma-project.io/docs)

## Add a new documentation topic

Follow these steps:

<div tabs name="documentation-topic" group="new-documentation">
  <details>
  <summary label="kyma-repository">
  Kyma repository
  </summary>

1. Create a pull request with `.md` files for the new documentation topic. Place the `.md` files under a new `docs` subfolder in the repository, such as `docs/serverless/`.

2. In the same PR, create a `.yaml` file under the [`templates`](https://github.com/kyma-project/kyma/tree/master/resources/core/charts/docs/charts/content-ui/templates) folder to add a [ClusterAssetGroup CR](/components/rafter/#custom-resource-cluster-asset-group) for your topic. For example, if you add a ClusterAssetGroup CR for the Serverless component, name it `docs-components-serverless-cag.yaml`. ​

   See the example definition:
    ​
   ```yaml
   apiVersion: rafter.kyma-project.io/v1beta1
   kind: ClusterAssetGroup
   metadata:
     labels:
       rafter.kyma-project.io/view-context: docs-ui
       rafter.kyma-project.io/group-name: components
       rafter.kyma-project.io/order: "11"
     name: serverless
   spec:
     displayName: "Serverless"
     description: "Overall documentation for Serverless"
     sources:
       - type: markdown
         name: docs
         mode: package
         url: https://github.com/{{ .Values.global.kymaOrgName }}/kyma/archive/{{ .Values.global.docs.clusterAssetGroupVersion }}.zip
         filter: /docs/serverless/
   ```

3. Adjust values for these fields:

- **rafter.kyma-project.io/order** defines the number of your topic on the list in the left navigation, such as `"11"`. To add it correctly, check in other `.yaml` files which number is assigned to the last documentation topic in the navigation on the website, and add a consecutive number to your component. If you decide to modify the existing topic order, change values for this parameter in all other `.yaml` files accordingly to avoid duplicates.
- **metadata.name** defines the CR name, such as `serverless`.
- **spec.displayname** defines the component name displayed on the website, such as `"Serverless"`.
- **spec.sources.filter** defines the location of the new topic's document sources, such as `/docs/serverless/`.

4. Merge the changes and wait until the website is rebuilt.

  </details>
  <details>
  <summary label="other-repository">
  Other repository
  </summary>

1. Create a pull request with `.md` files for the new documentation topic. Place the `.md` files under a new `docs` subfolder in the repository, such as `docs/commands/`.

2. In the same PR, create a `.yaml` file under the `.kyma-project-io` folder to add a [ClusterAssetGroup CR](/components/rafter/#custom-resource-cluster-asset-group) for your topic. Format of file name should be `{topic-name}-cag.yaml`.

   See the example definition:
    ​
   ```yaml
   apiVersion: rafter.kyma-project.io/v1beta1
   kind: ClusterAssetGroup
   metadata:
     labels:
       rafter.kyma-project.io/view-context: cli
       rafter.kyma-project.io/group-name: cli
       rafter.kyma-project.io/order: "2"
     name: commands
   spec:
     displayName: "Commands"
     description: "Overall documentation for Kyma CLI Commands"
     sources:
       - type: markdown
         name: docs
         mode: package
         filter: /docs/commands/
   ```

3. Adjust values for these fields:

- **rafter.kyma-project.io/order** defines the number of your topic on the list in the left navigation, such as `"2"`. To add it correctly, check in other `.yaml` files which number is assigned to the last documentation topic in the navigation on the website, and add a consecutive number to your component. If you decide to modify the existing topic order, change values for this parameter in all other `.yaml` files accordingly to avoid duplicates.
- **metadata.name** defines the CR name, such as `commands`.
- **spec.displayname** defines the component name displayed on the website, such as `"Commands"`.
- **spec.sources.filter** defines the location of the new topic's document sources, such as `/docs/commands/`.

4. Merge the changes and wait until the website is rebuilt.

  </details>
</div>

> **NOTE:** Before merge, you can check that everything is good with your changes, checking rendered documentation thanks to [`docs-preview`](#documentation-preview-documentation-preview).

## Add a single OpenAPI specification

In addition to documentation, there are also [OpenAPI](https://swagger.io/specification/) specifications rendered on the [Kyma website](https://kyma-project.io). You can find these specifications under the **API Consoles** type in the right navigation panel of a given documentation topic.
​
To add a new specification, follow these steps:

<div tabs name="openapi-specification" group="new-documentation">
  <details>
  <summary label="kyma-repository">
  Kyma repository
  </summary>

1. Go to the [`templates`](https://github.com/kyma-project/kyma/tree/master/resources/core/charts/docs/charts/content-ui/templates) folder and locate an existing ClusterAssetGroup CR that you want to modify.
​
2. Add a new source entry in the **sources** field:

   ``` yaml
   sources:
     ...
     - type: {SPECIFICATION_TYPE}
       name: {SPECIFICATION_NAME}
       mode: single
       url: {SPECIFICATION_URL}
   ```

   where:
  
   - **{SPECIFICATION_TYPE}** defines a type of a given specification. Currently, only [OpenAPI](https://swagger.io/specification/) specifications are supported and they are defined under the `openapi` type.
   - **{SPECIFICATION_NAME}** defines a unique identifier of a given specification. This field defines the URL on https://kyma-project.io/docs under which the specification is displayed. For example, if the specification is added in the `application-connector` ClusterAssetGroup CR with the `connectorapi` value of the **name** field, its URL is `https://kyma-project.io/docs/{VERSION_OF_DOCS}/components/application-connector/specifications/connectorapi/`.
   - **{SPECIFICATION_URL}** defines the location of the specification. It may contain directives using values defined in `values.yaml` files. For internal specifications defined in the [`kyma`](https://github.com/kyma-project/kyma) repository, it is recommended to use the directive with a Kyma version and the organization name, such as:

   ``` yaml
   url: https://raw.githubusercontent.com/{{ .Values.global.kymaOrgName }}/kyma/{{ .Values.global.docs.clusterAssetGroupsVersion }}/docs/application-connector/assets/connectorapi.yaml
   ```

   See the example:

   ``` yaml
   sources:
     ...
     - type: openapi
       name: connectorapi
       mode: single
       url: https://raw.githubusercontent.com/{{ .Values.global.kymaOrgName }}/kyma/{{ .Values.global.docs.clusterAssetGroupsVersion }}/docs/application-connector/assets/connectorapi.yaml
   ```

3. Merge the changes and wait until the website is rebuilt.

  </details>
  <details>
  <summary label="other-repository">
  Other repository
  </summary>

1. Go to the `.kyma-project-io` folder in repository and locate an existing ClusterAssetGroup CR that you want to modify.
​
2. Add a new source entry in the **sources** field:

   ``` yaml
   sources:
     ...
     - type: {SPECIFICATION_TYPE}
       name: {SPECIFICATION_NAME}
       mode: single
       url: {SPECIFICATION_URL}
   ```

   where:
  
   - **{SPECIFICATION_TYPE}** defines a type of a given specification. Currently, only [OpenAPI](https://swagger.io/specification/) specifications are supported and they are defined under the `openapi` type.
   - **{SPECIFICATION_NAME}** defines a unique identifier of a given specification. This field defines the URL on https://kyma-project.io/docs under which the specification is displayed. For example, if the specification is added in the `commands` ClusterAssetGroup CR with the `provision` value of the **name** field, its URL is `https://kyma-project.io/docs/{VERSION_OF_DOCS}/cli/commands/specifications/provision/`.
   - **{SPECIFICATION_URL}** defines the location of the specification. It may contain directives using values defined in `values.yaml` files.

   ``` yaml
   url: https://raw.githubusercontent.com/kyma-project/cli/{VERSION_OF_DOCS}/docs/commands/assets/provision.yaml
   ```

   See the example:

   ``` yaml
   sources:
     ...
     - type: openapi
       name: connectorapi
       mode: single
       url: https://raw.githubusercontent.com/kyma-project/cli/{VERSION_OF_DOCS}/docs/commands/assets/provision.yaml
   ```

3. Merge the changes and wait until the website is rebuilt.

  </details>
</div>

> **NOTE:** Before merge, you can check that everything is good with your changes, checking rendered documentation thanks to [`docs-preview`](#documentation-preview-documentation-preview).

## Add new repository to render its documentation

Follow these steps:

1. Add in repository `.kyma-project-io` folder and appropriate [documentation topics](#add-new-documentation-to-the-website-add-new-documentation-to-the-website-add-a-new-documentation-topic) inside created directory.

2. Add the new entry in the **docs** field in [`config.json`](https://github.com/kyma-project/website/blob/master/config.json):

   ```json
   "{REPOSITORY_NAME}": {
      "displayName": "{DISPLAY_NAME}",
      "organization": "{ORGANIZATION_NAME}",
      "repository": "{REPOSITORY_NAME}",
      "branches": {BRANCHES},
      "lastReleases": {LAST_RELEASES},
      "navPath": "{NAV_PATH}",
      "rootPath": {
        "docsType": "{DOCS_TYPE}",
        "docsTopic": "{DOCS_TOPIC}"
      }
   }
   ```

   where:

   - **{REPOSITORY_NAME}** - name of the repository.
   - **{DISPLAY_NAME}** - name which will be visible under **DOCS** dropdown in main navigation in the [Kyma website](https://kyma-project.io).
   - **{ORGANIZATION_NAME}** - name of the organization.
   - **{BRANCHES}** - branch names that will be rendered. It should contains at least `master` branch.
   - **{LAST_RELEASES}** - number of the last releases that will be rendered. If only branches should be displayed, the value should be 0.
   - **{NAV_PATH}** - navigation path of the entry under **DOCS** dropdown in main navigation in the [Kyma website](https://kyma-project.io).
   - **{DOCS_TYPE}** - documentation type which should be visible under navigation path.
   - **{DOCS_TOPIC}** - documentation topic which should be visible under navigation path.

   Example:

   ```json
   "{REPOSITORY_NAME}": {
      "displayName": "CLI",
      "organization": "kyma-project",
      "repository": "cli",
      "branches": ["master"],
      "lastReleases": 0,
      "navPath": "cli",
      "rootPath": {
        "docsType": "cli",
        "docsTopic": "overview"
      }
   }
   ``` 

3. Create the Pull Request, wait for review, merge the changes and wait until the website is rebuilt. New entry under **DOCS** dropdown in main navigation in the [Kyma website](https://kyma-project.io) should appear with new documentation.
