# Custom workload metrics

## Shared Prometheus

Kyma comes with a pre-installed Prometheus. Currently, it is only used for monitoring Kyma components and it is not possible for it to scrape custom workload metrics. 
The easiest way to let the customer monitor her worloads would be lifting the restrictions of Kyma Prometheus. However, in this case it's quite hard to guarantee its stability.

One problem that can occur is so called cardinality explosion when the time series count grows because of high-cardinality labels. A good example is metrics exposed by Istio when elastic workloads scale up and down. Another very common use case is bad code deploys that stuff high-cardinality data (request IDs, timestamps, user-provided values, etc.) into one or more labels of one or more metrics.

Monitoring and preventing cardinality explosion is not an easy task. There is a project called [Bomb Squad](https://blog.freshtracks.io/bomb-squad-automatic-detection-and-suppression-of-prometheus-cardinality-explosions-62ca8e02fa32), which detects high-cardinality metrics and silences them by rewriting the scrpae config.

## Separate plain Prometheus

Deploy a plain (non-operated) Prometheus server instance. Scraping custom workloads can be enabled by setting the following Pod annotations: 
```yaml
annotations:
    prometheus.io/scrape: "true"
    prometheus.io/path: $(PROMETHEUS_METRICS_PATH)
    prometheus.io/port: $(PROMETHEUS_METRICS_PORT)
```

Dividing Prometheus into Kyma Prometheus and Custom Worload Prometheus has a lot of advantages.
However, it still makes sense to use a shared Grafana that queries both Prometheus instances. It is possible to achieve it without making any changes in the Grafana configuration, just by adding a custom datasource.

### How to test it?

```bash
kubectl apply -f assets/simple-prometheus-setup.yaml # deploy Prometheus server and make it a Kyma Grafana datasource
kubectl apply -f assets/workloads.yaml               # deploy a custom worload that exposes metrics
kubectl apply -f assets/dashboard.yaml               # deploy Grafana dashboard
```

### Pros

1. Separating Kyma Prometheus, which is fully under our control, from Custom Worload Prometheus. In this setup, Kyma Prometheus remains unaffected by custom metrics hitting scraping limits, having high-cardinality, etc.
2. Enabling metric scraping is can easily achieved by setting a set of annotations. In theory, it is possible to wtite a simple mutating webhook server to add those annotations automatically.

### Cons

1. Limited configuration flexibility (e.g. comparing to ServiceMonitors). According to the [official documentation](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config), it's only possible to set scrape interval and scrape timeout on a per target basis.
2. Second Prometheus instance needs additional cluster resources.
