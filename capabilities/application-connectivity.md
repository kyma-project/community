---
displayName: Application Connectivity
epicsLabels:
  - area/application-connector
---

## Scope

The Application Connectivity capability supports the extensibility and integration between applications in a cloud-native way using the best available solutions.

It provides the integration patterns and necessary tooling for establishing the trusted connection between Kyma and third-party applications.

The selection of enterprise integration patterns represented by the documentation, the tooling, and examples are provided for developer convenience and insurance that integrated systems will work fluently, where the development effort is low.

## Vision

* API Registry and Discoverability

    * Integrated application can register its APIs and Event catalog. 
    * The registration contains the configuration of the endpoints together with required security credentials, documentation of the API and Event catalog based on open standards like OpenAPI and AsyncAPI, and additional documentation.
    * Registered entities are integrated with Service Catalog
     
* Orchestration

    * The orchestration functionality allows automatic binding between Kyma runtime and application registered by the customer.
    * The orchestration is tightly coupled with the API registry and discovery.
    
* Events integration
    * The event integration functionality provides required middleware for delivery of the business events to Kyma.
    * The support for the delivery guarantee, monitoring, and tracing is added.
    
* Access to the registered APIs
    * The access to API exposed by an integrated application is provided, and all requests are proxied by the delivered proxy service.
    * The proxy service is handling the authentication and authorization, integration with monitoring and tracing. 
    * The Service Catalog binding controls the access to the proxy service and therefore to the API. The development effort is reduced to the required minimum, and all the boiler-plate code is packed into connectors.
    
* Connectors
    * The set of connectors ensures that plumbing between an external application and SAP CP XF is possible in all cases where the use of standard functionality is not feasible. 
    * The palette of provided connectors should be kept to the required minimum and wherever possible the industry standards must be used.
