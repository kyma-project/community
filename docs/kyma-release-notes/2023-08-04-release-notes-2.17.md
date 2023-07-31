## General
Cleanup script to remove logging component

## Observability

### Monitoring
- New production profile settings: https://github.com/kyma-project/kyma/pull/17652
- The dashboard/datasource reloader has been updated, so that the logs to stdout are reduced to a minimum (https://github.com/kyma-project/kyma/pull/17812)
- We have updated the monitoring stack (https://github.com/kyma-project/kyma/pull/17877)
  - Prometheus 2.45.0 LTS
  - Prometheus-operator 0.66.0

### Removal of logging component
The logging component including the in-cluster Loki stack has been removed (https://github.com/kyma-project/kyma/issues/15827) as announced in detail with blog post (https://blogs.sap.com/2023/06/02/removal-of-loki-based-application-logs-in-sap-btp-kyma-runtime/). You must follow the cleanup instructions because no updates to the component will be shipped anymore.

## Telemetry
The Telemetry stack has been upgraded:
- Otel-collector 0.81.0 https://github.com/kyma-project/kyma/pull/17807
- Fluent Bit 2.1.7 https://github.com/kyma-project/kyma/pull/17878

We have implemented bug fixes for: 
- A single TracePipeline referencing a non-existent secret resulted in a crashing trace collector https://github.com/kyma-project/telemetry-manager/issues/272
- A Logpipeline referencing a non-existent secret resulted in a broken Fluent Bit configuration https://github.com/kyma-project/telemetry-manager/issues/137
	
## Service Mesh
A bug got fixed where the istio sidecars tried to send spans to an unknown cluster-local address causing unneeded stress on CodeDNS (https://github.com/kyma-project/kyma/pull/17811)

### Istio upgraded to 1.18.2
In this release, we upgraded Istio from 1.18.1 to 1.18.2. For more details on the changes, read the official [Istio 1.18.2 release notes](https://istio.io/latest/news/releases/1.18.x/announcing-1.18/upgrade-notes/).

## API Gateway
Ory Hydra component has been deprecated since Kyma 2.2 and is planned to be removed with Kyma 2.19. Follow the procedure outlined in this blog post to migrate from ORY Hydra to other providers. For more information on the ongoing changes, read about SAP BTP, Kyma Runtime API Gateway future architecture based on Istio.
