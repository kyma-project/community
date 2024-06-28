Sailing across [the midnight train going anywhere](https://youtu.be/VcjzHMhBtf0?si=woNdDUXHw0xjPL5b&t=26) we have embarked upon fixes of bugs and insightful updates, all to ensure a strikingly better and smoother user experience.
So, hold on to the rhythm of coding, believe in the power of what this new release can do, and read on to learn more about it!

## Application Connectivity

[Additional logging options have been added](https://github.com/kyma-project/application-connector-manager/tree/main/components/central-application-gateway#debugging) to Central Application Gateway.


## Observability

> **Early warning:** The Monitoring component will be removed with Kyma 2.20, as announced in this [blog post](https://blogs.sap.com/2023/09/07/removal-of-prometheus-grafana-based-monitoring-in-sap-btp-kyma-runtime/).

Update of the software stack:
 - [kube-state-metrics 2.10.0](https://github.com/kyma-project/kyma/pull/18135)
 - [oauth2-proxy 7.5.1](https://github.com/kyma-project/kyma/pull/18222)

## Telemetry

The Telemetry component has been [replaced ](https://github.com/kyma-project/kyma/issues/16301) by the [Telemetry module](https://github.com/kyma-project/telemetry-manager) and won't be part of Kyma releases anymore. Instead, learn about upcoming releases of the Telemetry module on the dedicated [release page](https://github.com/kyma-project/telemetry-manager/releases).
	
## Service Mesh

With Kyma 2.19, we have updated the Istio component to version 1.1.0. It contains the following changes:
 - [Istio updated to version 1.19.0](https://github.com/kyma-project/istio/pull/373)
 - [Enabled Horizontal Pod Autoscaling capability](https://github.com/kyma-project/istio/pull/371) for a smaller cluster installation
 - [Additional Kyma resources and configuration](https://github.com/kyma-project/istio/issues/334), such as Istio Grafana dashboards or PeerAuthentication configuring service-mesh traffic to only allow mTLS

## Security

As announced in the [Kyma 2.18 release notes](https://github.com/kyma-project/kyma/releases/tag/2.18.0), we have removed the following components:
 - Cluster Users
 - ECDHE-RSA-AES256-SHA and ECDHE-RSA-AES128-SHA cipher suites

## API Gateway

We have removed the Ory Hydra component, as announced in the [Kyma 2.17 release notes](https://github.com/kyma-project/kyma/releases/tag/2.17.0).
Also, we have removed APIRule in version `v1alpha1`, as announced in the [Kyma 2.16 release notes](https://github.com/kyma-project/kyma/releases/tag/2.16.0).  
