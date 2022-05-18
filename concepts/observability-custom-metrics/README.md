# Custom Workload Metrics

# Simple Prometheus Setup

Deploy a plain (non-operated) Prometheus server instance. Scraping custom workloads can be enabled by setting the following Pod annotations: 
```yaml
annotations:
    prometheus.io/scrape: "true"
    prometheus.io/path: $(PROMETHEUS_METRICS_PATH)
    prometheus.io/port: $(PROMETHEUS_METRICS_PORT)
```

Dividing Prometheus into Kyma Prometheus and Custom Worload Prometheus has a lot of advantages.
However, it still makes sense to use a shared Grafana instance that queries both instances. It is possible to achieve it without making any changes in the Grafana configuration, just adding a custom datasource will suffice.

## How to test it?

```bash
kubectl apply -f assets/simple-prometheus-setup.yaml
kubectl apply -f assets/workloads.yaml
kubectl apply -f assets/dashboard.yaml
```

## Pros

1. Isolating Kyma Prometheus makes it possible to optimize its performance.
2. Easy to enable metric scraping.

## Cons

1. Limited configuration (e.g. comparing to ServiceMonitors).
2. Two Prometheus instances - increased resource consumption.
