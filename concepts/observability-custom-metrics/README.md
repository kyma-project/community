# Custom workload metrics

## Simple Prometheus setup

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
2. Enabling metric scraping is can easily achieved by setting a set of annotations.

### Cons

1. Limited configuration flexibility (e.g. comparing to ServiceMonitors). According to the [official documentation](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config), it's only possible to set scrape interval and scrape timeout on a per target basis.
2. Second Prometheus instance needs additional cluster resources.
