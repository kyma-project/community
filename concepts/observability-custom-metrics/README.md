# Custom workload metrics

The following document presents possible alternatives for a Kyma user to store and analyze custom worload metrics. We want to address the following aspects:

* How easy it is for a Kyma user to enable and adjust worload scraping?
* Can we guarantee monitoring stack stability in case of a user mistake (e.g. which results in cardinality explosion). Monitoring and recoverability are important

## Shared Prometheus

Kyma comes with a pre-installed Prometheus. Currently, it is only used for monitoring Kyma components and it is not possible for it to scrape custom workload metrics. 
The easiest way to let the customer monitor her worloads would be lifting the restrictions of Kyma Prometheus. However, in this case it's quite hard to guarantee its stability.

One problem that can occur is so called cardinality explosion when the time series count grows because of high-cardinality labels. A good example is metrics exposed by Istio when elastic workloads scale up and down. Another very common use case is bad code deploys that stuff high-cardinality data (request IDs, timestamps, user-provided values, etc.) into one or more labels of one or more metrics.

A cardinality explosion causes the following problems:

* Increases Prometheus memory consumption, which can eventually lead to be OOM-killed
* Increases scrape durations
* Querying becomes effectively impossible

Monitoring and preventing cardinality explosion is not an easy task. There is a project called [Bomb Squad](https://github.com/open-fresh/bomb-squad), which detects high-cardinality metrics and silences them by rewriting the scrpae config.
However, it's an alpha project and is, in fact, abandoned. 

### Pros

* Easy to set up - just expose some Kyma Prometheus Operator features to end users

### Cons

* Since Prometheus Operator does not support annotation-based discovery of services, its CRD-based API (PodMonitor or ServiceMonitor) has to be exposed to end users. It is very flexible, but requires from end users advanced knowledge of Prometheus
* Since the Prometheus config is under control of Kyma Prometheus Operator, it's not possible to use cardinality prevention alike Bomb Squad
* Hard to guarantee Kyma monitoring stability since custom metrics will be pulled by the same Prometheus instance

## Separate plain Prometheus

Deploy a plain (non-operated) Prometheus server instance. Scraping custom workloads can be enabled by setting the following Pod annotations: 
```yaml
annotations:
  prometheus.io/scheme: "https"
  prometheus.io/scrape: "true"
  prometheus.io/path: $(PROMETHEUS_METRICS_PATH)
  prometheus.io/port: $(PROMETHEUS_METRICS_PORT)
```

Dividing Prometheus into Kyma Prometheus and Custom Worload Prometheus has a lot of advantages. However, we still have to let end users monitor Custom Workload Prometheus's stability and raise alerts if it hits the limits. We also have to suppress high cardinality metrics scraping (e.g. by using something similar to Bomb Squad)

It makes sense to use a shared Grafana that queries both Prometheus instances. It is possible to achieve it without making any changes in the Grafana configuration, just by adding a custom datasource.

### How to test it?

```bash
kubectl apply -f assets/simple-prometheus-setup.yaml # deploy Prometheus server and make it a Kyma Grafana datasource
kubectl apply -f assets/workloads.yaml               # deploy a custom worload that exposes metrics
kubectl apply -f assets/dashboard.yaml               # deploy Grafana dashboard
```

### Pros

* Separating Kyma Prometheus, which is fully under our control, from Custom Worload Prometheus. In this setup, Kyma Prometheus remains unaffected by custom metrics hitting scraping limits, having high-cardinality, etc
* Enabling metric scraping is can easily achieved by setting a set of annotations. In theory, it is possible to wtite a simple mutating webhook server to add those annotations automatically

### Cons

* Limited configuration flexibility (e.g. comparing to ServiceMonitors). According to the [official documentation](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config), it's only possible to set scrape interval and scrape timeout on a per target basis
* Second Prometheus instance needs additional cluster resources

## Separate operated Prometheus (using Kyma Prometheus Operator)

Use Kyma Prometheus Operator to deploy a second Prometheus instance. However, this approach combines disadvantages from both above-mentioned approaches: complex API and hard to prevent cardinality explosion.

## Separate Prometheus Operator

Similar argumentation as above. In addition to that, more cluster resources will be used.

### Conclusions 

Taking all pros and cons into account, Prometheus Operator-based setups do not suit our needs due to the following reasons:

* Unnessesarily complex API
* Enabling Prometheus Operator means exposing all its features, many of those customers do not need (e.g. ThanosRuler)
* Not possible to implement cardinality explosion prevention, since the scrape config is controlled by the operator
* In the long-run we want to drop Kyma Prometheus Operator in favor of plain Prometheus as a more lightweight alternative. It will also decrase chart maintenance costs.

On the other hand, we want to keep Kyma monitoring separate from customer workload monitoring to ensure its stability.
It still makes sense to keep all dashboards in one place, which means opening Kyma Grafana to the users.

It means that the separate plain Prometheus setup is the way to go.

### Questions to answer

* How to integrate alertmanager?
* Investigate what it means to open Kyma dashboard performance-wise (how many dashboards customer can create, etc.)
* Try out the Bomb Squad approach to prevent Prometheus from crashing
