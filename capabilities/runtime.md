---
displayName: "Serverless Runtime"
id: "serverless"
epicsLabels:
  - area/serverless
---

<!-- above metadata will be used on kyma.project.io page to display nice name of capability and have a reference to label that should be used while fetching from ZenHub/GitHub the information about related Epics and their delivery plan   -->

## Scope

The Serverless Runtime capability offering is a central part of the serverless strategy of Kyma. It is the easiest way to run custom code in Kyma and to integrate different services provided by the Services Brokers and the Application Connector. Following the "batteries included" rule of Kyma, the function-as-a-service (FaaS) solution is provided inside the OSS Project.

## Vision

The goal of Serverless Runtime capability is to:

- Provide a simple FaaS runtime as a part of Kyma.
- Enable an easy deployment from version control systems.
- Allow fast feedback loops for development and testing.
- Allow managing FaaS workload via UI and CLI commands.
- Integrate third-party serverless providers.
- Integrate both Kyma and third-party FaaS solutions with other services provided by Kyma, such as Eventing, API Gateway, Service Catalog.
- Provide an easy way to run containerized workloads inside and outside Kyma.
- Provide helpers/sdk/utilities to eliminate boilerplate code.
