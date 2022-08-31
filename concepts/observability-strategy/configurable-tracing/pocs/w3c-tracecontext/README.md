## Goal

Prove that istio can be switched to w3c-tracecontext protocol and verify that Kyma components will support it native

## Strategy

At the moment, the only solution to use w3c-tracecontext already is via an openCensus tracer for istio. As Jaeger does not support the openCensus protocol, an otel-collector needs to be introduced additionally which does the conversion to the zipkin/jaeger protocol. As the goal anyway is to introduce an otel-collector, that approach sournds feasible as the central otel-collector could support openCensus temporarily as additional receiver.

## Setup

The activation of the openCensusAgent as tracer via the [extensionProvider](https://preliminary.istio.io/latest/docs/tasks/observability/distributed-tracing/opencensusagent/) concept was not working proper yet. Here, the sampling rate was not taken into account proper with strange effects. The istio-ingressgateway was not reporting spans, the downstream services received trace headers without sampling decision and still were reporting spans.

So the PoC used the traditional approach via the meshConfig and an hardcoded proxy setting. An sample application based on Serverless and Eventing was deployed.

1. Have an OSS Kyma 2.6.1 installed
2. Deploy the otel-collector `kubectl -n kyma-system apply -f otel-collector.yaml`
3. Edit kyma/resources/istio/values.yaml and add the values taken from [istio-changes.yaml](./istio-changes.yaml)
4. Update istio: `kyma deploy -s local --component istio`
5. Verify changes are applied: `kubectl -n istio-system get istiooperators.install.istio.io -oyaml`
6. Deploy sample app
```
kubectl -n default apply -f https://raw.githubusercontent.com/a-thaler/kyma-function-examples/master/chain/apirule.yaml
kubectl -n default apply -f https://raw.githubusercontent.com/a-thaler/kyma-function-examples/master/chain/subscription.yaml
kubectl -n default apply -f https://raw.githubusercontent.com/a-thaler/kyma-function-examples/master/chain/svc-a.yaml
kubectl -n default apply -f https://raw.githubusercontent.com/a-thaler/kyma-function-examples/master/chain/svc-b.yaml
kubectl -n default apply -f https://raw.githubusercontent.com/a-thaler/kyma-function-examples/master/chain/svc-c.yaml
```
7. Call `GET demo.<yourClusterDomain>`
8. Check Jaeger `kubectl port-forward -n kyma-system svc/tracing-jaeger-query 16686:16686`

## Result

In the access logs of the envoys you can see that the traceparent is set for all applications. For the first hop to svc-a the B3 headers are not set. Later they are set as Serverless is supporting both protocols and is enriching the requests with both headers sets. These could be disabled after the official switch to w3c-tracecontext.
The trace in Jaeger is complete.
