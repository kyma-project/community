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

## Solution
  
![](assets/main-arch.svg)

### DocsTopic and ClusterDocsTopic
- Documentation should be provided in a zip/tar.gz format and its location with additional spec info should go into the Custom Resource (DocsTopic and Cluster Docs Topic)
- Such a zip package with documentation contains everything. There is a default convention applied if nothing is specified
  - markdown sources
  - assets
  - all kind of supported specs
  
  
### Documentation Controller
- The package is refered in the CR and fetched by a Documentation Controller
- Package is unziped and validated
- Index of all docs and assets is generated and stored in respective folders as `index.yaml` file
  - for assets it contains a name of all the assets. The order is based on filename
  - for docs it contains a name and related markdown metadata. The order is based on filename

