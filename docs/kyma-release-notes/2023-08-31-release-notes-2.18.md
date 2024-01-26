With some of its offspring having left the parent ship as fully-grown, [independent modules](https://kyma-project.io/#/06-modules/README) and more getting ready to follow suit, [Kyma’s sailing, Kyma’s sailing…](https://www.youtube.com/watch?v=FOt3oQ_k008) and still evolving. To understand its journey towards enhanced functionality, look at the latest updates and fixes.

## Application Connectivity

With this release, all 5XX codes passing through Central Application Gateway are now rewritten to `502`. The original error code is returned in a Target-System-Status header.

## Telemetry

Kyma 2.18 brings the following improvements:
- We’ve [fixed the bug that caused problems scraping the metrics of the Fluent Bit component](https://github.com/kyma-project/kyma/issues/17976) by third-party vendors.
- We‘ve added [mTLS support for TracePipeline OTLP outputs](https://github.com/kyma-project/kyma/issues/17995).
- We‘ve [updated the following software stack](https://github.com/kyma-project/kyma/pull/18021):
    - OTel Collector 0.83.0  
    - Fluent Bit 2.1.8

## Service Mesh

As a significant step in our journey towards the Istio module’s release, we have introduced a new, stable, and reliable method of installing Istio - Kyma Istio Operator. This release also includes improvements to the installation and upgrade processes of Istio, as well as a new version of the Istio custom resource that provides additional configuration options. To learn more about it, visit the [Kyma Istio Operator repository](https://github.com/kyma-project/istio/blob/main/docs/user/README.md).

## Security 

### Cluster Users removal

The Cluster Users component [has been deprecated since Kyma 2.7](https://github.com/kyma-project/website/blob/main/content/blog-posts/2022-09-22-release-notes-2.7/index.md#cluster-users-component-deprecated) and will be removed with Kyma 2.19. 
The component includes predefined Kubernetes ClusterRoles such as `kyma-admin`, `kyma-namespace-admin`, `kyma-edit`, `kyma-developer`, `kyma-view`, and `kyma-essentials`. These Roles specify permissions for accessing Kyma resources that you can assign to users. For example, if you bind a user to the `kyma-admin` ClusterRole, it grants them full admin access to the entire cluster, and if you bind a user to the `kyma-view` ClusterRole, they are only allowed to view and list the resources within the cluster. 
Once the component is removed, these Roles will no longer be available for newly created clusters. This means that you won’t be able to use these predefined sets of rights and will be required to specify yourself which users or groups should have access to which of your resources. However, for clusters created before the release of Kyma 2.19, the already-defined Roles will not be deleted.

### Cipher suits removal

As a part of security hardening and Kyma security team recommendations, ECDHE-RSA-AES256-SHA and ECDHE-RSA-AES128-SHA cipher suites used in default Kyma Gateway have been deprecated since the [2.15 Kyma version](https://github.com/kyma-project/kyma/releases/tag/2.15.0). Although we initially planned to remove these cipher suites with Kyma 2.18, we have decided to delay their removal until version 2.19. After the Kyma 2.19 release, clients dependent on the mentioned cipher suites won't be accepted.

## Eventing

### NATS

The following NATS Images have been updated:
- [prometheus-nats-exporter `v0.12.0`](https://github.com/nats-io/prometheus-nats-exporter/releases/tag/v0.12.0)
- [nats-config-reloader `v0.12.0`](https://github.com/nats-io/prometheus-nats-exporter/releases/tag/v0.12.0)

## Serverless

Kyma 2.18 brings more observability into Node.js-based Functions.
They are now exposing a metrics endpoint containing the following auto-instrumented metrics:
- histogram for Function execution duration: function_duration_seconds
- number of calls in total: function_calls_total
- number of exceptions in total: function_failures_total

## User Interface

From now on, you have the opportunity to give feedback about our product directly in Kyma dashboard. To do that, use the shiny new button in the top right corner of the shell bar. Read the [UX Scorecard in Kyma Dashboard](https://blogs.sap.com/2023/08/18/ux-scorecard-in-kyma-dashboard/) blog post for more details.
