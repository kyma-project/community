Colorful leaves dropping on the ground signify high autumn. With Kyma 2.20, we say goodbye to one of our components that falls down the Kyma tree. However, it is a natural process. See two other components getting reborn in their modular versions, and watch them in full bloom again. Read on to learn more about changes in the new Kyma release.

## Application Connector

### Migration of the Application Connector (AC) component
With the introduction of the Application Connector module, the AC component is no longer installed by default with Kyma.
To enable Application Connector on new clusters, follow the [installation instructions](https://github.com/kyma-project/application-connector-manager/blob/main/docs/contributor/01-10-installation.md).
 
## Observability

### Removal of the Monitoring component
The Monitoring component, including the in-cluster Prometheus/Grafana stack, [has been removed](https://github.com/kyma-project/kyma/issues/16306) as announced in the [Removal of Prometheus/Grafana-based monitoring in SAP BTP, Kyma runtime](https://blogs.sap.com/2023/09/07/removal-of-prometheus-grafana-based-monitoring-in-sap-btp-kyma-runtime) blog post. You must follow the cleanup instructions because no updates to the component will be shipped anymore.

## Serverless

### Migration of the Serverless component
As of Kyma 2.20, Serverless is no longer installed by default as a Kyma component. It is now developed and released as an independent module. Adding the module to your Kyma runtime, you extend the runtime with the same Serverless capabilities, that you have already known so well from the previous Kyma versions.
To enable Serverless on new Kyma runtime instances, you must install Serverless Operator. Serverless Operator will install and continue reconciling Serverless installation based on the Serverless custom resource.
For more details, follow the [installation instructions](https://github.com/kyma-project/serverless-manager/tree/main#install).
