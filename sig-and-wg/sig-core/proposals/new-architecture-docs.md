# New Architecture for Documentation and Specifications Management

Created on 2018-10-13 by Lukasz Gornicki (@derberg).

## Status

Proposed on 2018-10-16

## Motivation

1. To solve below technological dept
   - Current approach is based on 2 different ways of loading content sources into Minio
     - Throught a docker image to Minio
     - Application Connector through the Metadata Service directly to Minio
   - No single validation component
   - To get content in the UI we have 2 different places to define details, 
     - config located in docs sources
     - navigation manifesto where you define topic name and id
2. To enable single solution for any type of docs and make it possible for users to easily reuse our solution for their needs
3. Enable modularization for minio, so it can be easily replaced by s3 and not maintain a special cache for docs in ui-api-layer
4. Enable modularization of documentation, so you can load documentation only for modules that are installed in Kyma

## Solution
  
![](assets/main-arch.svg)

### DocsTopic and ClusterDocsTopic
- All details of a given documentation topic, including doc soureces are specified with [Custom Resource](/assets/doc-topic-crd-and-example.yaml) (DocsTopic and ClusterDocsTopic)
- Supported formats, markdown + assets, swagger, asyncapi, odata
- Documentation can be provided in different formats:
  - in a zip/tar.gz format with info what can be found in a package, docs and specs. There is a default convention applied if nothing is specified
  - different location of docs or specs can be provided, in case of docs and assets you point to an index with names of the files available under given link
  - mixcure of above is possible

```
---
apiVersion: documentation.kyma-project.io/v1alpha1
kind: ClusterDocsTopic
metadata:
  name: service-catalog #example based on current documentation topic https://github.com/kyma-project/kyma/tree/master/docs/service-catalog
  labels:
    viewContext: docs-view
spec:
  description: Overal documentation for Service Catalog
  order: 1 # helps with ordering of the list in navigation
  displayName: Service Catalog
  group: Components #I think that because of the plural nature of the name it is much more intuitive to expect here plural then in an attribute called "type"
  source:
    package: https://some.domain.com/kyma.zip #zip or tar of package with docs and speci, structure must follow accepted convention
    docs: https://some.domain.com/index.yaml
    specs:
      swagger: 
        url: https://some.domain.com/swagger.yaml
        rewrites: 
          basePath: /test/v2
      asyncapi: 
        url: https://some.domain.com/asyncapi.yaml
      odata: 
        url: https://some.domain.com/odata.xml
status:
  ready: False
  reason: ValidationFailed # or UploadFailed or SourceFetchFailure
  message: "swagger file is not a valid json"
#status:
#  ready: True
#  resource:
#    docsUrl: 
#      index: $LINK-TO-INDEX
#       apiVersion: v1
#         files:
#         - name: 01-overview.md
#           metadata:
#             title: MyOverview
#             type: Overview
#         - name: 02-details.md
#           metadata:
#             title: MyDetails
#             type: Details
#         - name: 03-installation.md
#           metadata:
#             title: MyInstallation
#             type: Tutorial
#         - name: assets/diagram.svg
#    spec:
#      swagger: $LINK-TO-FILE
```

### Documentation Controller

The package or specific files and indexes are refered in the CR and fetched by a Documentation Controller:

Package
- Package is unziped and fetched content validated and rewritten if needed (for example baseUrl rewrite in swagger.json)
- Index of all docs and assets is generated and stored in respective folders as `index.yaml` file
  - for assets it contains a name of all the assets. The order is based on filename
  - for docs it contains a name and related markdown metadata. The order is based on filename
- Together with index files content is uploaded to storage

Direct links
- Content is fetched and validated
- Together with index files content is uploaded to storage

```
#sample of index file with markdown and assets files
apiVersion: v1
files:
  - name: 01-overview.md
    metadata:
      title: MyOverview
      type: Overview
  - name: 02-details.md
    metadata:
      title: MyDetails
      type: Details
  - name: 03-installation.md
    metadata:
      title: MyInstallation
      type: Tutorial
  - name: assets/diagram.svg
```

Controller sets status of the CR:
- in case of successful creation of the CR, controller adds to the CR information about location of the fetched resources and details if index, if such exists
- in case of failure, error message is specified

Controller cleans up storage in case of CR is deleted. 

### Catalog Docs Controller

When you register a ServiceBroker, this controller listens to all newly addedded ServiceClasses to the Catalog and creates for them DocsTopic or ClusterDocsTopic CR. Such ServiceClass on which controller reacts must contain `external.metadata.content` object.

```
package: https://some.domain.com/kyma.zip #zip or tar of package with docs and speci, structure must follow accepted convention
    docs: https://some.domain.com/index.yaml
    specs:
      swagger: 
        url: https://some.domain.com/swagger.yaml
        rewrites: 
          basePath: /test/v2
      asyncapi: 
        url: https://some.domain.com/asyncapi.yaml
      odata: 
        url: https://some.domain.com/odata.xml
```

For docs cleanup reasons (unregister broker case), controller during CR creation specifies an `ownerReference` pointing to the ServiceClass. The controller will make sure that for such use case it will add a finalizer to the DocsTopic CR and not allow its deletion until storage is really cleaned up.

### Storage/Minio

- Minio can be replaced with S3 if needed.
- Bucket with documentation is called `docstopics/VERSION`
- In case of Cluster CR path to sample swagger would be `docstopics/v1alpha1/CR_NAME/swagger.yaml`
- In case of Namespaced CR path to sample swagger would be `docstopics/v1alpha1/namespace/NAMESPACE_NAME/CR_NAME/swagger.yaml`

## Website running outside Kyma cluster

To avoid duplication, `topics.yaml` is basically a list of CRs that point to available locations of documentation for a given Kyma version. Website pipeline acts as a controller, generating indexes if needed and puts files in storage (in this case it is website github page)
