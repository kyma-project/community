---
title: {Service/Controller name}
type: Metrics
---

>**NOTE:** Blockquotes in this document provide instructions. Remove them from the final document.
>
>This document is a ready-to-use template for the **Metrics** document type that describes service or controller metrics. Follow the `11-{00}-{service/controller-name}.md` convention to name the document.

This table shows the {Service name} custom metrics, their types, and descriptions.

| Name | Type | Description |
|------|------|-------------|
| `{metric_name}` | {metric_type} | {Metric description} |

> If there are some default metrics, such as Prometheus metrics for [Go applications](https://prometheus.io/docs/guides/go-application/), exposed for the service or controller in addition to the custom ones, add this information to the document. For example, write:

Apart from the custom metrics, the {Service name} also exposes default Prometheus metrics for [Go applications](https://prometheus.io/docs/guides/go-application/).

> If you describe metrics for a controller that does not expose any custom metrics, write only about the default ones. For example, write:

Metrics for the {Controller name} include:

- {default metric name or description}.
- {another default metric name or description}.

> Provide guidelines how to list all metrics for the described service or controller.

To see a complete list of metrics, run this command:

```bash
kubectl -n {namespace_name} port-forward svc/{service/controller_name} {port}
```

To check the metrics, open a new terminal window and run:

```bash
curl http://localhost:{port}/{endpoint}
```

> **TIP:** To use these commands, you must have a running Kyma cluster and kubectl installed. If you cannot access port `{port}`, redirect the metrics to another one. For example, run `kubectl -n {namepace_name} port-forward svc/{service/controller_name} {another_port}:{port}` and update the port in the localhost address.

> Provide a reference to the Monitoring documentation in Kyma.

See the [Monitoring](/components/monitoring) documentation to learn more about monitoring and metrics in Kyma.

> For reference, see the existing **Metrics** documents for the [Rafter](https://kyma-project.io/docs/1.9/components/rafter/#metrics-metrics).
