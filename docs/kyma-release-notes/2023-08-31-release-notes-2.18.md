With some of its offspring having left the parent ship as fully-grown, [independent modules](https://kyma-project.io/#/?id=kyma-modules) and more getting ready to follow suit, [Kyma’s sailing, Kyma’s sailing…](https://www.youtube.com/watch?v=FOt3oQ_k008) and still evolving. To understand its journey towards enhanced functionality, look at the latest updates and fixes.

## Application Connectivity
With this release, all 5XX codes passing through Central Application Gateway are now rewritten to `502`. The original error code is returned in a Target-System-Status header.

## Telemetry
Kyma 2.18 brings the following improvements:

- We’ve [fixed the bug that caused problems scraping the metrics of the Fluent Bit component](https://github.com/kyma-project/kyma/issues/17976)  to third-party vendors.  
- We‘ve added [mTLS support for TracePipeline OTLP outputs](https://github.com/kyma-project/kyma/issues/17995)  
- We‘ve [updated the following components](https://github.com/kyma-project/kyma/pull/18021):  
   - Otel-collector 0.83.0  
   - Fluent Bit 2.1.8


## Eventing
 
### NATS:
The following NATS Images have been updated:  
- [prometheus-nats-exporter `v0.12.0`](https://github.com/nats-io/prometheus-nats-exporter/releases/tag/v0.12.0)  
- [nats-config-reloader `v0.12.0`](https://github.com/nats-io/prometheus-nats-exporter/releases/tag/v0.12.0)

## Serverless

Kyma 2.18 brings more observability into Node.js-based functions.
They are now exposing metrics endpoint containing the following auto-instrumented metrics:  
- histogram for Function execution duration: function_duration_seconds  
- number of calls in total: function_calls_total  
- number of exceptions in total: function_failures_total  
