---
displayName: Logging-Tracing-Monitoring
epicsLabels:
  - area/logging
  - area/tracing
  - area/monitoring
---

 ## Scope
Enable operator and developer to easily observe the health of a kyma cluster and the health of distributed applications running on kyma; on base of application logs and metrics as well as transaction traces; on development and productive environments.

The assurance that kyma components are creating application logs and metrics as well as trace data in a consistent way and in a sufficent detail level is not in scope of the capability.

 ## Vision
In a cloud-native microservices architecture, a request triggered by some end user action often flows through dozens or maybe more microservcies. Tooling such as logging and monitoring will provide insights about the health of a particular component. On base of that, pro-active and re-active actions can be taken to keep or recover the health of a component. Tracing will enrich that observability picture by adding the facet of transactions traces across the different components.

To support that aspects of observability and following the *batteries-included* principle, kyma should provide leightweight and cloud-native solutions for logging, tracing and monitoring. They should enable developers and operators to easily query all application health data across the different workloads, in an development environment as well as on productive setups.

This includes:
* Setup and maintain optional, lightweight and cloud-native solutions for logging, tracing and monitoring
* As much as possible data derived out-of-the-box from Kubernetes and the Service Mesh
* Support for local development (minikube)
* Integration of all 3 aspects into solutions of the supported cloud providers, especially DynaTrace for monitoring
* Easy accessible and secure API and UI/CLI to query logs
* Namespace separation support
* (Logging) Support of transactional event logs (like audit logs)
* (Monitoring) Support for auto-scaling on base of application metrics
* (Monitoring) Pre-integration with typical notification systems like VictorOps and Slack
