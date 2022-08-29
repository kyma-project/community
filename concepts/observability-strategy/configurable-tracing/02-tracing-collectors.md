# Comparison of collector technologies

## Requirements


## Alternatives

## Decision
Because the general observability strategy is based on OTLP, support for that as input and output is mandatory. Evaluation of  potential instability of the OpenTelemetry Collector shows that the big contributors are using it in production. Because the trace protocol in OpenTelemetry Collector is stable already, it is just a matter of time till all aspects are marked as stable. Because we plan to bring our own abstraction for supported scenarios, a beta API used underneath is acceptable.
With that, the OpenTelemetry Collector will be used in the further concept.
