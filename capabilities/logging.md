---
displayName: Logging
epicsLabels:
  - area/logging
---

 ## Scope
Assures that people which are operating kyma or running workload on kyma are having a central place to access and query application logs, for any kind of kyma deployment.

Assurance of a consistent logging approach for kyma components is not in scope of the capability.

 ## Vision
Following the *batteries-included* paradigm, kyma should provide a leightweight logging solution which enables developers to easily query all application logs across workloads, in an development environment as well on productive setups.

This includes:
* Setup and maintain an optional, central, lightweight and cloud-native logging solution
* Support for local development (minikube)
* Integration into logging solutions of all supported cloud providers
* Support of transactional event logs (like audit logs)
* Easy accessible and secure API and UI/CLI to query logs
* Namespace separation support
