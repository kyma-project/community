..placeholder for preface..

## Application Connector

### Removal of the Application Connector component
With the introduction of the AC module, the AC component will no longer by installed by default with Kyma.
To enable Application Connector on new clusters, follow the installation instructions:
https://github.com/kyma-project/application-connector-manager/blob/main/docs/contributor/01-10-installation.md
 
## Observability

### Removal of monitoring component
The monitoring component including the in-cluster Prometheus/Grafana stack has been removed (https://github.com/kyma-project/kyma/issues/16306) as announced with blog post (https://blogs.sap.com/2023/09/07/removal-of-prometheus-grafana-based-monitoring-in-sap-btp-kyma-runtime). You must follow the cleanup instructions because no updates to the component will be shipped anymore.

## Serverless

### Removal of serverless component
As of kyma v2.20 the Serverless component is no longer installed by default as a kyma component.
In order to enable serverless on new kyma runtime instances you would need to install serverless operator that will install and continue reconciling serverless installation based on Serverless Custom Resource.
Please follow the installation instructions:
https://github.com/kyma-project/serverless-manager/tree/main#install   
