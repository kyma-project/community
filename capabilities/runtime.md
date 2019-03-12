---
displayName: Serverless Runtime
epicsLabels:
  - area/serverless
---
<!-- above metadata will be used on kyma.project.io page to display nice name of capability and have a reference to label that should be used while fetching from ZenHub/GitHub the information about related Epics and their delivery plan   -->

## Scope

The Serverless Runtime capability offering is a central part of the serverless strategy of kyma. It is the easiest way to to run custom code in kyma and to integrate different services provided by the services broker and application connector. Based on the Batteries included strategy of Kyma a FaaS Solution based on Knative is provided inside the OSS Project. Beside the included FaaS Solution it is possible to schedule Workloads on 3rd Party Offerings using knative.

## Vision

The goal of Serverless Runtime capability is to:

- Provide a simple FaaS runtime as part of kyma.
- Enable easy deployment from version control systems.
- Allow fast feedback loops for development and testing.
- Integrate 3rd Party Serverless providers using knative.
- Integrate FaaS (kyma FaaS and 3rd Party FaaS) with other services provided by Kyma (Eventing, API, Service Catalog)
- Provide a easy way to run Containerized workloads inside and outside of kyme (using knative)
- Provide a easy way to host static file for frontend workloads
