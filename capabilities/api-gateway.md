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
* Traffic management for APIs
	* Control outbound traffic for APIs - define a list of external services which the API can access
	* Control internal API traffic - to specify which services can access APIs internally
	* Traffic management for different API versions - split traffic between different versions of one API
	* API failure prevention - enable setting circuit breakers for APIs  
* Expose GraphQl service apis as REST apis  
    * Allow applying authorization rules on top
* Enable developers to create and expose APIs separated on namespace level
    * Allow exposure of API with Namespace name as a part of domain
    * Allow blocking communication between services living in different Namespaces
* Expose services running on different providers - not directly on Kyma
	      
	