---
displayName: Application Connectivity
epicsLabels:
  - area/application-connector
---

## Scope

The Application Connectivity capability supports the extensibility and integration between applications in a cloud-native way using the best available solutions.
It provides the integration patterns and necessary tooling for establishing the trusted connection between Kyma and third-party applications.

The goal is to enable the integration of various systems in a coherent way, where the management and usage of APIs and Events are standardized and allows natural extension of such systems in the cloud-native fashion.
The unification of the connected system enables the fast development which unlocks the brand new possibilities of extending and customizing the existing solution in the modern, cloud-based, way.

The new way of scalability of the system is possible. The event mechanism and simple access to exposed API is a foundation for moving the workload from the legacy system to the Kubernetes.

An additional benefit is that you can mesh different systems using the language of your choice. 

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
    * In the case of integration with the legacy system, which is not exposed using open standards, like REST, or if it is hidden behind a network firewall, the connectors will be provided.
    * The connector ensures that a legacy system will be exposed in the same way as other systems, and the required translation of the API calls or events will be done. 
    * The palette of provided connectors should be kept to the required minimum and wherever possible the industry standards must be used.

* Security
    * The security of the integration is one of the top priority. The trusted relationship between connected application must be ensured, including the authentication and authorization.
    * Various standard security mechanisms are provided to ensure the identity of the application, access to the secure resources, in both communication directions, no matter if it is asynchronous communication using events or synchronous communication using REST calls. The support for OAuth, Basic auth, client certificates, CSRF tokes and more must be provided.