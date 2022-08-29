# Comparison of collector technologies

## Requirements


## Alternatives

## Decision
As the general observability strategy is based on OTLP, support for that as input and output is mandatory. Even considering the potential instability of the OpenTelemetry Collector, it shows that the big contributors are using it in production already. As the trace protocol in OpenTelemetry Collector is stable already, it is just a matter of time till all aspects are marked as stable. Because we plan to bring our own abstraction for supported scenarios, a beta API used underneath is acceptable.
With that, the OpenTelemetry Collector will be used in the further concept.
