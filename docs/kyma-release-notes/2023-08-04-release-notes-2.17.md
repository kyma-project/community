[Hooray! Hooray! It’s a Kyma release day!](https://www.youtube.com/watch?v=ModISbNyQ8I&t=36s) If you’re on your vacation, enjoying your leisure time, cold drinks, and sunny weather, you should ask yourself a few very important questions. Am I familiar with the latest 2.17 version of Kyma? What changes does it bring for Observability and Telemetry? Is the Istio component upgraded? Read on to find answers to all those burning questions!

## Observability

### Monitoring
- We have introduced new [production profile settings](https://github.com/kyma-project/kyma/pull/17652). 
- We have updated the [dashboard/datasource reloader](https://github.com/kyma-project/kyma/pull/17812). The logs to `stdout` are now reduced to a minimum.
- We have updated the [Monitoring stack](https://github.com/kyma-project/kyma/pull/17877):
  - Prometheus to version 2.45.0 LTS
  - Prometheus-operator to version 0.66.0

### Removal of the Logging component
The [Logging component](https://github.com/kyma-project/kyma/issues/15827), including the in-cluster Loki stack, has been removed, as announced in detail in this [blog post](https://blogs.sap.com/2023/06/02/removal-of-loki-based-application-logs-in-sap-btp-kyma-runtime/). Follow the [cleanup instructions](https://github.com/kyma-project/kyma/blob/release-2.17/docs/migration-guide-2.16-2.17.md) because updates to the component will no longer be shipped.

## Telemetry
The Telemetry stack has been upgraded:
- [OTel Collector to version 0.81.0](https://github.com/kyma-project/kyma/pull/17807)
- [Fluent Bit to version 2.1.7](https://github.com/kyma-project/kyma/pull/17878)

We have implemented bug fixes for: 
- A single TracePipeline referencing a non-existent Secret, resulting in a [crashing Trace Collector](https://github.com/kyma-project/telemetry-manager/issues/272).
- A LogPipeline referencing a non-existent Secret, resulting in a [broken Fluent Bit configuration](https://github.com/kyma-project/telemetry-manager/issues/137).
	
## Service Mesh
We have fixed a bug where the Istio sidecars tried to send spans to an unknown cluster-local address, causing [unneeded stress on CodeDNS](https://github.com/kyma-project/kyma/pull/17811).

### Istio upgraded to 1.18.2
In this release, we have upgraded Istio from 1.18.1 to 1.18.2. For more details on the changes, read the official [Istio 1.18.2 release notes](https://istio.io/latest/news/releases/1.18.x/announcing-1.18.2/).

## API Gateway
The Ory Hydra component has been deprecated since Kyma 2.2 and is planned to be removed with Kyma 2.19. Follow the procedure outlined in this [blog post](https://blogs.sap.com/2023/06/06/sap-btp-kyma-runtime-ory-hydra-oauth2-client-migration/) to migrate from ORY Hydra to other providers. For more information on the ongoing changes, read about [SAP BTP, Kyma Runtime API Gateway future architecture based on Istio](https://blogs.sap.com/2023/02/10/sap-btp-kyma-runtime-api-gateway-future-architecture-based-on-istio/).
