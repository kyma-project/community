# Why

The main goal of custom UI modules is to provide a way to extend Busola with custom UI components, views, and features, that do not follow the standard pattern of editing custom resources (list, details, create, update, delete). Examples are:
- testing serverless functions by sending HTTP requests to the function pod through kubernetes apiserver proxy and reading the response and logs
- fetching service catalog entries from a remote service manager API using credentials stored in a secret

# Requirements

1. Access APIs outside of Kubernetes cluster
2. Use user privileges in the cluster (do not escalate privileges by using service accounts)
3. Keep modules code independent of busola (independent releases, loose coupling, etc.)
4. Use the same authentication/authorization mechanism as busola (OIDC, RBAC)
5. List of modules and their versions are cluster-specific (not global)
6. Module UI is deployed together with the module operator and is part of the module release.


# Module UI hosting (centralized vs in-cluster)

Centralized hosting pros and cons:
- :+1: UI can be loaded from a single source (e.g. a CDN) and cached by the browser
- :+1: UI can be updated without updating the module itself (e.g. to fix a bug - no need to update the module operator)
- :-1: UI can be blocked by the browser due to Content Security Policy (CSP) restrictions
- :-1: UI can be blocked by the network in the restricted markets (e.g. if the CDN is blocked)

In-cluster hosting pros and cons:
- :+1: UI is part of the module release and is deployed together with the module operator. We can reuse existing mechanisms for deploying and updating the module.
- :+1: Users can introduce custom modules without the need to host them externally
- :-1: UI is not cached by the browser and is loaded every time the user opens the module 


# Security considerations
Loading custom code from not trusted sources can be a security risk. The custom code can access the same resources as Busola, so it can read and modify any resource in the cluster user has permissions to. It is important to ensure that only approved modules are loaded into the browser. 

