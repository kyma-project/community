---
displayName: Tracing
epicsLabels:
  - area/tracing
---

 ## Scope
Enable operator and developer to observe business transactions as it propagates across the different workloads running in a kyma cluster, in development environments and productive systems.

Assurance of a consistent approach for custom tracing data of kyma components is not in scope of the capability.

 ## Vision
Good observability tools should serve the purpose of telling clear stories. In a cloud-native microservices architecture, a request triggered by some end user action often flows through dozens or maybe more microservcies. Tooling such as logging, monitoring etc. have their own value. However, they consider each component or microservice in isolation. This is an issue from operational perspective.
Distributed Tracing charts out the stories about the transactions in cloud-native systems.

Distributed tracing is a observability tool for developers building cloud native applications in Kyma.

It enables developers to:
* Trace the distributed transaction as it propagetes across various components and microservices using an optional, lightweight and cloud-native tracing system
* Trace end-to-end event flows including outer-kyma applications by integrating into the tracing solutions of all supported cloud providers
* Support the identification of culprit components for a failed distributed transaction.
* Support identification of latency and performance bottlenecks along the request flow
* As much as possible out-of-the-box tracing data get derived from Kubernetes and the Service Mesh
* Support for local development (minikube)
* Easy accessible and secure API and UI/CLI to query metrics
* Namespace separation support

