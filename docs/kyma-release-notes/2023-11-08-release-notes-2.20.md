Colorful leaves dropping on the ground signify high autumn. With Kyma 2.20, we say goodbye to some of our components that fall down the Kyma tree. However, it is a natural process as they will be reborn in their modular versions, and we will watch them in full bloom again. Read on to learn more about changes in the new Kyma release.

## Application Connector

### Removal of the Application Connector component
With the introduction of the Application Connector module, the AC component is no longer installed by default with Kyma.
To enable Application Connector on new clusters, follow the [installation instructions](https://github.com/kyma-project/application-connector-manager/blob/main/docs/contributor/01-10-installation.md).
https://github.com/kyma-project/application-connector-manager/blob/main/docs/contributor/01-10-installation.md
 
## Observability

### Removal of the Monitoring component
The Monitoring component, including the in-cluster Prometheus/Grafana stack, has been removed (https://github.com/kyma-project/kyma/issues/16306) as announced in the [Removal of Prometheus/Grafana-based monitoring in SAP BTP, Kyma runtime](https://blogs.sap.com/2023/09/07/removal-of-prometheus-grafana-based-monitoring-in-sap-btp-kyma-runtime) blog post. You must follow the cleanup instructions because no updates to the component will be shipped anymore.

## Serverless

### Removal of the Serverless component
As of kyma v2.20 the Serverless component is no longer installed by default as a kyma component.
In order to enable serverless on new kyma runtime instances you would need to install serverless operator that will install and continue reconciling serverless installation based on Serverless Custom Resource.
Please follow the installation instructions:
https://github.com/kyma-project/serverless-manager/tree/main#install   
