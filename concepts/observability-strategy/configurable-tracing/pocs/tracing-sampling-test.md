# OpenTelemetry/Istio Sampling Rate Analysis

This analysis aim to find out impact of different sampling rate configuration on istio proxy and overall call chain.

Different scenarios should be compared end of analysis, like istio resource consumption, throughput, and impact of other kubernetes components. 

## Setup
### Opentelemetry
OpenTelemetry version 0.22.1 deployed on a Kyma cluster (version 2.6.0) with a minimal setup,
minimal setup mean :
- Standard deployment of OpenTelemetry Helm Chart version 0.22.1
- As receivers standard Jaeger, Zipkin, and OTLP configured
- As processor only memory limiter processor with standard configuration
- As exporter only log/info and/or NOP exporter to keep impact of OpenTelemetry as low as possible
- No andy additional extensions

### Sampling App Deployment

As sampling app, 3 serverless function deployed with istio sidecar injected and which calling each other in a chain like
**Extern Call -> FunctionA -> FunctionB -> FunctionC** to simulate a trace fully sampled with an external call.

All functions are deployed with NodeJs version 16.

### Call Simulator
As call simulator Gatling version 3.8.3 used.

## Scenario
All scenarios are having same OpenTelemetry deployment and configuration described above.

Same Sampling app used in all scenarios.

Same Gatling call simulator used in all scenarios, Gatling shall call **FunctionA** from extern with at least 5 simultaneous users up to 10 users maximum.
Call is simple URL call of FunctionA with no additional data or http headers to keep any influence of those on trace self.

Call simulation should run 100 minutes long to put enough load on call chain and generate enough metrics to get precise result.
Gatling will generate around 50K call towards to FunctionA, istio proxy of deployed functions will call tracing collectors for each call should be sampled

### Scenario 1
Kyma standard (version 2.6.0) deployment with istio sampling rate configured to **1%** sampling, to observe istio behavior like resource consumption and throughput.

### Scenario 2 
Kyma standard deployment from main branch with istio sampling rate configuration changed to **100%** sampling, to observe istio behavior like resource consumption and throughput.

### Scenario 3
Like **Scenario 2** described above with additional configuration on **Jaeger** and **Zipkin** receivers services without endpoint.

Services **Jaeger** and **Zipkin** in this scenario are exist but pointing no valid endpoint. This scenario should focus additionally impact on kubernetes components.

### Scenario 4

Like Scenario 2 described above, this setup should run without collectors **Jaeger** and **Zipkin** to analyse istio behaviour and impact on kubernetes components like **CoreDNS**.

## Result

Overall analysis focus on istio behavior, resource consumption under different configuration and setup, in **Scenario 3** and **Scenario 4** additionally focus on kubernetes components.

### Scenario 1

Following picture shown call execution summary (from client perspective) of test cluster with istio sampling rate 1% setup. 
Success rate of call execution is 100% and average response time is around **400ms.**

| ![Call execution summary](../assets/istio-1per-call-summary.jpg) |
| :--: |
| Fig. 1 Call execution summary |


Screenshot below show overview istio mesh network, as call execution summary cluster metrics also show 100% of success rate inclusive local call chain.
Latency on average stay around 400ms with a 6.5 operation per second per service.

| ![Istio mesh mverview](../assets/istio-1per-overwiew.jpg) |
| :--: |
| Fig. 2 Istio Mesh Overview |

| ![Istio service detail](../assets/istio-1per-servicedetail.jpg) |
| :--: |
| Fig. 3 Istio Service Overview |

| ![Service network detail 1](../assets/istio-1per-servicedetail-2.jpg) |
| :--: |
| Fig. 4 Service Network Details |

| ![Service network detail 2](../assets/istio-1per-servicedetail-3.jpg) |
| :--: |
| Fig. 5 Service Network Details |

Screenshot below show istio proxy resource consumption during test, bytes transfered during test phase increased from around 10KB/s to 60KB/s in peak time.

Memory consumtion of istio proxy increased from 61MB to 68.7 MB in peak time, on CPU consumption side, before test started was 0.100 in peak time the value increased to around 0.150.

Although resource consumtion during test execution increased, there are no considarable resource consumption observed.

| ![Istio resource consumption](../assets/istio-1per-resource.jpg) |
| :--: |
| Fig. 6 Istio Proxy Resource Consumption |

On Kubernetes CoreDNS service side nothing important to mention here, screenshots below shown an overview.
DNS request per second increase during test execution from 2 packet/sec. to 3.5 packet/sec. also DNS lookups increased. 

DNS Cache overview show cache hits are increased which indicate DNS records request results are mostly coming from cache.

| ![Kubernetes CoreDNS Overview](../assets/istio-1per-coredns-overview.jpg) |
| :--: |
| Fig. 7 Kubernetes CoreDNS Overview |

| ![Kubernetes CoreDNS Cache](../assets/istio-1per-coredns-cache.jpg) |
| :--: |
| Fig. 8 Kubernetes CoreDNS Cache |

#### Summary
Test results shown there is a small impact on istio proxy but none of those are overall considerably.
Also, on Kubernetes side look good, CoreDNS service side is no considerable DNS record request observed and DNS records mostly read from DNS cache.

### Scenario 2

Following picture shown call execution summary (from client perspective) of test cluster with istio sampling rate 100% setup.
Success rate of call execution is 99% and average response time is around **950ms.** 

Unsuccessfully calls mostly result HTTP 503, main cause of not successfully calls are, either istio proxy can't access upstream service (serverless function self) or jaeger tracing collector service,
which result both upstream services (serverless function and jaeger) are can't deal such load but on istio proxy side noting suspicion observed.  


| ![Call execution summary](../assets/istio-100per-call-summary.jpg) |
| :--: |
| Fig. 1 Call execution summary |


Screenshot below show overview istio mesh network, as call execution summary cluster metrics also show 100% of success rate inclusive local call chain.
Latency on average stay over 500ms with a 6.7 operation per second per service.

Global Request Volume increased from 10 to 65.8 operation per second which indicate high traffic on istio side.

| ![Istio mesh mverview](../assets/istio-100per-overwiew.jpg) |
| :--: |
| Fig. 2 Istio Mesh Overview |

| ![Istio service detail](../assets/istio-100per-servicedetail.jpg) |
| :--: |
| Fig. 3 Istio Service Overview |

| ![Service network detail 1](../assets/istio-100per-servicedetail-2.jpg) |
| :--: |
| Fig. 4 Service Network Details |

| ![Service network detail 2](../assets/istio-100per-servicedetail-3.jpg) |
| :--: |
| Fig. 5 Service Network Details |

Screenshot below show istio proxy resource consumption during test, bytes transferred during test phase increased from around 10KB/s to 300KB/s in peak time.

Memory consumption of istio proxy increased from 61MB to 68.7 MB in peak time, on CPU consumption side, before test started was 0.100 in peak time the value increased to around 0.170.

Although resource consumption during test execution increased, there are no considerable resource consumption observed.

| ![Istio resource consumption](../assets/istio-100per-resource.jpg) |
| :--: |
| Fig. 6 Istio Proxy Resource Consumption |

On Kubernetes CoreDNS service side nothing important to mention here, screenshots below shown an overview.
DNS request per second increase during test execution from 2 packet/sec. to 3.3 packet/sec. also DNS lookups increased.

DNS Cache overview show cache hits are increased which indicate DNS records request results are mostly coming from cache.

| ![Kubernetes CoreDNS Overview](../assets/istio-100per-coredns-overview.jpg) |
| :--: |
| Fig. 7 Kubernetes CoreDNS Overview |

| ![Kubernetes CoreDNS Cache](../assets/istio-100per-coredns-cache.jpg) |
| :--: |
| Fig. 8 Kubernetes CoreDNS Cache |

#### Summary
Test results shown there is a small impact on istio proxy but none of those are overall considerably.
On Kubernetes side look good, CoreDNS service side is no considerable DNS record request observed and DNS records mostly read from DNS cache.

Test shown 100% sampling rate, increase istio proxy network traffic, in certain circumstances can cause increased network latency.

### Scenario 3

Following picture shown call execution summary (from client perspective) of test cluster with istio sampling rate 100% setup.
Success rate of call execution is 33% and average response time is around **9 seconds.**

Unsuccessfully calls mostly result HTTP 503, main cause of not successfully calls is, istio proxy can't access upstream service jaeger tracing collector service,
this service has a broken endpoint.


| ![Call execution summary](../assets/istio-100per-broken-svc-call-summary.jpg) |
| :--: |
| Fig. 1 Call execution summary |


Screenshot below show overview istio mesh network, as call execution summary cluster metrics also show 33% of success rate inclusive local call chain.
Latency on average stay over 951ms with a 7 operation per second per service.

Global Request Volume increased from 10 to 72.8 operation per second which indicate high traffic on istio side.

| ![Istio mesh mverview](../assets/istio-100per-broken-svc-overwiew.jpg) |
| :--: |
| Fig. 2 Istio Mesh Overview |

| ![Istio service detail](../assets/istio-100per-broken-svc-servicedetail.jpg) |
| :--: |
| Fig. 3 Istio Service Overview |

| ![Service network detail 1](../assets/istio-100per-broken-svc-servicedetail-2.jpg) |
| :--: |
| Fig. 4 Service Network Details |

| ![Service network detail 2](../assets/istio-100per-broken-svc-servicedetail-3.jpg) |
| :--: |
| Fig. 5 Service Network Details |

Screenshot below show istio proxy resource consumption during test, bytes transferred during test phase increased from around 10KB/s to 180KB/s in peak time.

Memory consumption of istio proxy increased from 61MB to 67 MB in peak time, on CPU consumption side, before test started was 0.100 in peak time the value increased to around 0.140.

Compare to test **Scenario 2** amount of data transferred is almost 50% decreased which show no tracing data could be transferred to the collectors.

Although resource consumption during test execution increased, there are no considerable resource consumption observed.

| ![Istio resource consumption](../assets/istio-100per-broken-svc-resource.jpg) |
| :--: |
| Fig. 6 Istio Proxy Resource Consumption |

On Kubernetes CoreDNS service side, DNS request significantly increased, from 2 paket per second to almost 8 packet per second during the test execution time.
DNS lookup times also increased significantly, compare to test case 100% sampling in **Scenario 2** increased 100%.

DNS cache behavior also changed if we compare with **Scenario 2**, cache hits and misses doubled during test execution.

| ![Kubernetes CoreDNS Overview](../assets/istio-100per-broken-svc-coredns-overview.jpg) |
| :--: |
| Fig. 7 Kubernetes CoreDNS Overview |

| ![Kubernetes CoreDNS Cache](../assets/istio-100per-broken-svc-coredns-cache.jpg) |
| :--: |
| Fig. 8 Kubernetes CoreDNS Cache |

#### Summary
Test results shown there is a small impact on istio proxy resource consumption but network throughput getting bad.
Network latency increased and Kubernetes CoreDNS service has more work compare to **Scenario 2**

### Scenario 4

Following picture shown call execution summary (from client perspective) of test cluster with istio sampling rate 100% setup.
Success rate of call execution is 99% and average response time is around **1 seconds.**

Unsuccessfully calls mostly result HTTP 503, main cause of not successfully calls are, either istio proxy can't access upstream service (serverless function self) or jaeger tracing collector service,
which result both upstream services (serverless function and jaeger) are can't deal such load but on istio proxy side noting suspicion observed.


| ![Call execution summary](../assets/istio-100per-no-svc-call-summary.jpg) |
| :--: |
| Fig. 1 Call execution summary |


Screenshot below show overview istio mesh network, as call execution summary cluster metrics also show 100% of success rate inclusive local call chain.
Latency on average stay over 727ms with a 7.30 operation per second per service.

Global Request Volume increased from 10 to 21.1 operation per second which indicate increased traffic on istio side.

| ![Istio mesh mverview](../assets/istio-100per-no-svc-overwiew.jpg) |
| :--: |
| Fig. 2 Istio Mesh Overview |

| ![Istio service detail](../assets/istio-100per-no-svc-servicedetail.jpg) |
| :--: |
| Fig. 3 Istio Service Overview |

| ![Service network detail 1](../assets/istio-100per-no-svc-servicedetail-2.jpg) |
| :--: |
| Fig. 4 Service Network Details |

| ![Service network detail 2](../assets/istio-100per-no-svc-servicedetail-3.jpg) |
| :--: |
| Fig. 5 Service Network Details |

Screenshot below show istio proxy resource consumption during test, bytes transferred during test phase increased from around 10KB/s to 77KB/s in peak time.

Memory consumption of istio proxy increased from 66MB to 67 MB in peak time, on CPU consumption side, before test started was 0.050 in peak time the value increased to around 0.100.

By test start, memory consumption as well as CPU consumption increased rapidly but decreased in 5 minutes to the normal values.

Although resource consumption during test execution increased, there are no considerable resource consumption observed.

| ![Istio resource consumption](../assets/istio-100per-no-svc-resource.jpg) |
| :--: |
| Fig. 6 Istio Proxy Resource Consumption |

Like **Scenario 3** Kubernetes CoreDNS service side DNS request per second increased significantly, from 2 packet per second to 8 packet per second.
DNS cache misses also doubled in the test execution time, which also explain increased DSN lookup response time, this value increased from 1 second to 4 second during test execution time.

| ![Kubernetes CoreDNS Overview](../assets/istio-100per-no-svc-coredns-overview.jpg) |
| :--: |
| Fig. 7 Kubernetes CoreDNS Overview |

| ![Kubernetes CoreDNS Cache](../assets/istio-100per-no-svc-coredns-cache.jpg) |
| :--: |
| Fig. 8 Kubernetes CoreDNS Cache |

#### Summary
Test results shown there is a small impact on istio proxy but none of those are overall considerably.

CoreDNS service lookup times increased significantly because of unresolved dns entries.

This scenario show bad behavior of CoreDNS service and bad istio proxy network traffic result.
### Peroration

TBD