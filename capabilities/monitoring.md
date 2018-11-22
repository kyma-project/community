---
displayName: Monitoring
epicsLabels:
  - area/monitoring
---

 ## Scope
Monitoring = Metrics + Alerts

Enable operator and developer to gain insights about the health of a kyma cluster and workload running on top; by collecting metrics data. Enabling these people to react to unhealthy state; by having ways to get alerted.

Assurance of a consistent approach for custom metrics and alert rules of kyma components is not in scope of the capability.

 ## Vision
Following the *batteries-included* paradigm, kyma should provide a leightweight monitoring solution which enables developers to easily collect and query metrics across workloads, in an development environment as well on productive setups. It allowes to define alert rules and notification systems to detect unhealthy situations proactively but also to detect downtimes and errors.

This includes:
* Setup and maintain an optional, central, lightweight and cloud-native monitoring system
* As much as possible out-of-the-box metrics derived from Kubernetes and the Service Mesh
* Support for local development (minikube)
* Easy accessible and secure API and UI/CLI to query metrics
* Pre-integration with typical notification systems like VictorOps and Slack
* Integration into monitoring solutions of all supported cloud providers, especially DynaTrace
* Namespace separation support
