## Goal

Prove that Istio can be switched to w3c-tracecontext protocol, and verify that Kyma components will support it natively.

## Strategy

At the moment, the only solution to use w3c-tracecontext is with an OpenCensus tracer for Istio. Because Jaeger does not support the OpenCensus protocol, an otel-collector must be introduced additionally, which converts the data to the Zipkin/Jaeger protocol. As the goal anyway is to introduce an otel-collector, this approach is feasible because the central otel-collector could support OpenCensus temporarily as additional receiver.

## Setup

So the PoC used the traditional approach with the meshConfig and a hardcoded proxy setting. A sample application based on Serverless and Eventing was deployed.

1. Have an OSS Kyma 2.6.1 installed.
2. Deploy the otel-collector `kubectl -n kyma-system apply -f otel-collector.yaml`.
3. Edit `kyma/resources/istio/values.yaml` and add the values taken from [istio-changes.yaml](./istio-changes.yaml).
4. Update Istio: `kyma deploy -s local --component istio`.
5. Verify changes are applied: `kubectl -n istio-system get istiooperators.install.istio.io -oyaml`.
6. Deploy sample app.
    ```
    kubectl -n default apply -f https://raw.githubusercontent.com/a-thaler/kyma-function-examples/master/chain/apirule.yaml
    kubectl -n default apply -f https://raw.githubusercontent.com/a-thaler/kyma-function-examples/master/chain/subscription.yaml
    kubectl -n default apply -f https://raw.githubusercontent.com/a-thaler/kyma-function-examples/master/chain/svc-a.yaml
    kubectl -n default apply -f https://raw.githubusercontent.com/a-thaler/kyma-function-examples/master/chain/svc-b.yaml
    kubectl -n default apply -f https://raw.githubusercontent.com/a-thaler/kyma-function-examples/master/chain/svc-c.yaml
    ```
7. Call `GET demo.<yourClusterDomain>`.
8. Check Jaeger `kubectl port-forward -n kyma-system svc/tracing-jaeger-query 16686:16686`.

The activation of the openCensusAgent as tracer with the [extensionProvider](https://github.com/istio/istio.io/blob/release-1.22/content/en/docs/tasks/observability/distributed-tracing/opencensusagent/index.md) concept was tried as well. It required to use Istio 1.15 and worked out very well by changing the Istio config to use the following settings:
```yaml
  defaultProviders:
    tracing: ["opencensus"]
  extensionProviders:
    - name: "opencensus"
      opencensus:
          service: "otel-collector.kyma-system.svc.cluster.local"
          port: 55678
          context:
          - W3C_TRACE_CONTEXT
```
These settings enabled the tracer properly with a default sampling rate of 1%. Be aware that defining a single extensionProvider of type `tracing` deactivates the old meshConfig completely, including sampling and "enableTracing". Placing the telemetry resource in the istio-system namespace configured the sampling rate for the whole mesh:
```yaml
apiVersion: telemetry.istio.io/v1alpha1
kind: Telemetry
metadata:
  name: mesh-default
  namespace: istio-system
spec:
  tracing:
  - randomSamplingPercentage: 100.00
```

By providing no default telemetry resource, the user can completely configure the tracing (besides the push URL, which is controlled by the extensionProvider). Already now, this approach looks like the preferable way to configure tracing.

## Result

In the access logs of the envoys, you can see that the traceparent is set for all applications. For the first hop to svc-a, the B3 headers are not set. Later they are set because Serverless supports both protocols and enriches the requests with both headers sets. This behaviour could be disabled after the official switch to w3c-tracecontext.
The trace in Jaeger is complete.
