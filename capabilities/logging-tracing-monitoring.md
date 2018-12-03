---
displayName: Logging-Tracing-Monitoring
epicsLabels:
  - area/logging
  - area/tracing
  - area/monitoring
---

 ## Scope
Enable the operator and developer to easily observe the state of the Kyma cluster and the state of distributed applications running on Kyma. The data is made available and collected by pre-bundled infrastructure collecting application logs and metrics as well as transaction traces. This infrastructure integrates with cloud provider-specific tools.

Ensuring that Kyma components create application logs. metrics, and trace data in a consistent way and on a sufficient detail level is not in scope of this capability.

 ## Vision
In a cloud-native microservice architecture, a user request often flows through dozens of different microservices. Tools such as logging and monitoring provide insights on the health and state of a particular component. Based on the results, you can act proactively and reactively to keep or recover the health of a component. Tracing enriches observability by identifying transaction traces across the different components.

To support application observability and follow the batteries included philosophy, Kyma would provide lightweight and cloud-native solutions for logging, tracing and monitoring. Such solutions enable developers and operators to easily query all application health data across the different workloads in development and production environments.

This includes:
* Set up and maintain optional, lightweight and cloud-native solutions for logging, tracing and monitoring.
* As much as possible data derived out of the box from Kubernetes and the Service Mesh.
* Support for local development (Minikube).
* Integration of all 3 aspects into solutions of the supported cloud providers, especially DynaTrace for monitoring.
* Easily accessible and secure API and UI/CLI to query logs.
* Namespace separation support.
* (Logging) Support of transactional event logs such as audit logs.
* (Monitoring) Support for auto-scaling based on application metrics.
* (Monitoring) Pre-integration with notification systems like VictorOps and Slack.

This enables the following:
* Tracing the distributed transaction as it propagates across components and microservices.
* Tracing end-to-end Event flows including external applications.
* Improved debugging in local and remote development.
* Proactive notification for potential unhealthy states.
* Reactive warning about unhealthy components.

