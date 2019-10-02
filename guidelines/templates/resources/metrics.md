---
title: {Service/Controller name}
type: Metrics
---

>**NOTE:** Blockquotes in this document provide instructions. Remove them from the final document.
>
>This document is a ready-to-use template for the **Metrics** document type that describes service or controller metrics. Follow the `14-{00}-{service-name}.md` convention to name the document.

This table shows the {Service/Controller name} custom metrics, their types, and descriptions.

| Name | Type | Description |
|------|------|-------------|
| `{metric_name}` | {metric_type} | {Metric description} |

> If there are some default metrics, such as Prometheus metrics for [Go applications](https://prometheus.io/docs/guides/go-application/), exposed for the service or controller in addition to the custom ones, add this information to the document. For example, write:

Apart from the custom metrics, the {Service/Controller name} also exposes default Prometheus metrics for [Go applications](https://prometheus.io/docs/guides/go-application/).

> If you describe metrics for a service or a controller that does not expose any custom metrics, write only about the default ones. For example, write:

Metrics for the {Service/Controller name} include:

- {default metric name or description}.
- {another default metric name or description}.

> Provide a reference to the Monitoring documentation in Kyma.

See the [Monitoring](/components/monitoring) documentation to learn more about monitoring and metrics in Kyma.

> For reference, see the existing **Metrics** documents for the [Asset Store](https://kyma-project.io/docs/1.6/components/asset-store/#metrics-metrics).
