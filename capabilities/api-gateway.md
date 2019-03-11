---
displayName: API Gateway 
epicsLabels:
  - area/service-mesh
  - quality/security  
---

## Scope


The API Gateway capability aims to provide a set of functionalities allowing developers to expose, secure and manage their API's in an easy way. 
Based on the Service Mesh capability, it provides a common way to plug in security (authorization & authentication), enable routing and accessibility for all created APIs.
An API can be any application (lambda, GO application, etc.)



## Vision


* Extend authorization strategies for API's
	* OAuth2 server issuing access tokens to exposed Kyma APIs (both user and non-user oriented tokens)
	* OAuth2 proxy securing exposed APIs in Kyma, allowing access based on issued access tokens 
	* Enable Open Policy Agent policies for authorization and admission control (https://www.openpolicyagent.org/)
	* Support refreshing Oauth2 tokens	
	
* Traffic management for APIs
	* Control outbound traffic for APIs - define a list of external services which the API can access
	* Control internal API traffic - to specify which services can access APIs internally
	* Traffic management for different API versions - split traffic between different versions of one API
	* API failure prevention - enable setting circuit breakers for APIs
	  
* Enable developers to create and expose APIs separated on the namespace level
    * Allow application to be exposed with Namespace name as a part of the hostname
    * Allow blocking communication between services living in different Namespaces
    
* Enable easy adoption of GraphQl in an API 
    * Allow legacy/microservice/serverless applications to be exposed and visible as GraphQl APIs 
    * Automate configuration and deployment of front end proxies to allow communication using GraphQl
  
* Expose services running on different environments
    * Allow proxying requests to services running outside Kyma - on cloud foundry, other k8s clusters etc.
    * Allow configuring the same authentication and authorization policies as for services deployed in the Kyma
    
	