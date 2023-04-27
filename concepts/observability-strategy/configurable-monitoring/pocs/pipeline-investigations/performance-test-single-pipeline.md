# Metric Performance Test

## Setup

The metric performance test uses the [metric-gen](../tools/metric-gen/main.go) app to generate metrics. The application is deployed as a deployment with multiple replicas to simulate increasing replicas.

As test environment, a Kyma cluster installed from the main branch ([6bce47168452](https://github.com/kyma-project/kyma/tree/6bce47168452a87c78c636b4cfc65f3dc9592735)). This was done because at the time of performing the performance test, the MetricPipeline is not released yet. The following steps are performed to set up the cluster:
1. Deploy Kyma from branch 6bce47168452:
```bash
kyma deploy -s 6bce47168452 --value telemetry.operator.controllers.metrics.enabled=true
make install # within telemetry-manager repo
kubectl edit clusterrole telemetry-operator-manager-role # add metricpipeline to cluster role
```

2. Add following roles to the cluster roles:

```yaml
 - apiGroups:
  - telemetry.kyma-project.io
  resources:
  - metricpipelines
  verbs:
  - create
  - delete
  - get
  - list
  - patch
  - update
  - watch
- apiGroups:
  - telemetry.kyma-project.io
  resources:
  - metricpipelines/finalizers
  verbs:
  - update
- apiGroups:
  - telemetry.kyma-project.io
  resources:
  - metricpipelines/status
  verbs:
  - get
  - patch
  - update
```

3. Apply the pipeline:
```bash
kubectl apply -f ./config/samples/telemetry_v1alpha1_metricpipeline.yaml # within telemetry-manager repo
```

## Goal

The goal of this performance test is to derive the following values:

Metric Gateway Pod:

- CPU Limit
- Memory Limit

OTEL Configuration:

- `exporters.otlp.sending_queue.queue_size`
- `processor.batch.send_batch_size`
- `processor.batch.timeout`
- `processor.batch.send_batch_max_size`
- `processor.memory_limiter.check_interval`
- `processor.memory_limiter.limit_percentage`
- `processor.memory_limiter.spike_limit_percentage`

The test uses the image built by [metric-gen](../tools/metric-gen/main.go) which produces around 5.3k metrics per second with a total of 7 attributes per metric.

## Execution

1. Port forward Grafana and a load dedicated dashboard [assets/metric-gateway-grafana-dashboard.json](../assets/metric-gateway-grafana-dashboard.json).
2. Pause reconciliations for metric pipelines.
3. Set values for OTEL configuration in the metric gateway ConfigMap.
4. Restart the metric gateway Deployment.
5. Start the metric-gen tool with the desired number of replicas:

    ```bash
    kubectl create deployment metric-gen --image=<image-name> --replicas=num_replicas
    ```

## Tests with a healthy sink

This test are pushing the metrics to a healthy sink. The following OTEL configurations are the same for all tests:

| Property                                          | Value |
| ------------------------------------------------- | ----- |
| `processor.batch.timeout`                         | 10s   |
| `processor.memory_limiter.check_interval`         | 1s    |
| `processor.memory_limiter.limit_percentage`       | 75    |
| `processor.memory_limiter.spike_limit_percentage` | 10    |

The following table shows the test results with the remaining configurations. The batch size corresponds to the `processor.batch.send_batch_size` and `processor.batch.send_batch_max_size`. The exporter queue size is the current size of the queue and the capacity is the configured `exporters.otlp.sending_queue.queue_size`.

| Test # | Replicas | Request/s | Memory Max/Limit | CPU Max/Limit | Receiver Accepted/Refused | Processor Batch Size | Exporter Queue Size/Capacity |
| ------ | -------- | --------- | ---------------- | ------------- | ------------------------- | -------------------- | ---------------------------- |
| 1      | 1        | 5.3k      | 84MB/1Gi        | 0.69Gi        | 5.34k/0                   | 512                  | 0/1000                       |
| 2      | 2        | 10.6k     | 914MB/1Gi       | 0.64Gi        | 7k/3k                     | 512                  | 717/1000                     |
| 3      | 3        | 12.2k     | 940MB/1Gi       | 0.74Gi        | 7.3k/6k                   | 512                  | 728/1000                     |
| 4      | 1        | 6.3k      | 189 MB/1Gi      | 0.332Gi       | 6.31k/0                   | 1024                 | 0/1000                       |
| 5      | 2        | 10.9k     | 880 MB/1Gi      | 0.650Gi       | 8.5k/2.44k                | 1024                 | 322/1000                     |
| 6      | 3        | 13.6k     | 974 MB/1Gi      | 0.725Gi       | 7.56k/6.15k               | 1024                 | 384/1000                     |




- Test 1 shows that the average requests/s with one replica is around 5.3k. The metric gateway works as expected and the queue size is always 0 as the metrics are directly exported.
- Test 2 shows that with 2 replicas, some metrics are rejected by the receiver. The memory limiter logs that the hard limit is reached, which causes the rejection in the receiver. Also, it is visible that the exporter queue size increases, but is not yet full.
- Test 3 shows that with 3 replicas, around half of the metrics are rejected by the memory limiter/receiver.
- In Test 4, 5 and 6, the processor batch size was increased to 1024. As a result, the throughput increased.
  
## Tests with a unhealthy sink

To simulate an outage, we configured the metric pipeline to send the metrics to a non-existent endpoint. The following OTEL configurations are the same for all tests (same as in the previous tests):

| Property                                          | Value |
| ------------------------------------------------- | ----- |
| `processor.batch.timeout`                         | 10s   |
| `processor.memory_limiter.check_interval`         | 1s    |
| `processor.memory_limiter.limit_percentage`       | 75    |
| `processor.memory_limiter.spike_limit_percentage` | 10    |

The following table shows the test results with some configuration.

| Test # | Replicas | Request/s | Memory Max/Limit | CPU Max/Limit | Receiver Accepted/Refused | Processor Batch Size | Exporter Queue Size/Capacity |
| ------ | -------- | --------- | ---------------- | ------------- | ------------------------- | -------------------- | ---------------------------- |
| 1      | 1        | 5k        | 900MB/1Gi       | 0.35/1Gi      | 30s                       | 1024                 | 599/1000                     |
| 1      | 1        | 5k        | 850MB/1Gi       | 0.35/1Gi      | 30s                       | 512                  | 599/1000                     |
| 1      | 1        | 5k        | 850MB/1Gi       | 0.35/1Gi      | 30s                       | 512                  | 512/512                     |


## With Default limits
The tests were performed with default limits:
 - processor batch size: 8192
 - Queue Length: 5000
### With Healthy sink

| Test # | Replicas | Request/s   | Memory Max/Limit | CPU Max/Limit | Processor Batch Size | Exporter Queue Size/Capacity |
| ------ | -------- | ---------   | ---------------- | ------------- | -------------------- | ---------------------------- |
| 1      | 1        | 6k          | 115 MB/1Gi       | 0.3         | 8192                 | 0/5000                       |
| 1      | 2        | 9.6k        | 1022 MB/1Gi      | 0.54       | 8192                   | 58/5000                      |
| 1      | 2        | 9.6k        | 1722 MB/2Gi      | 0.65       | 8192                   | 118/5000                     |



### With Unhealthy Sink
| Test # | Replicas | Request/s | Memory Max/Limit | CPU Max/Limit | Receiver Accepted/Refused | Processor Batch Size | Exporter Queue Size/Capacity |
| ------ | -------- | --------- | ---------------- | ------------- | ------------------------- | -------------------- | ---------------------------- |
| 1      | 1        | 6k        | 1660MB/2Gi       | 0.38/1        | 2mins                     | 8192                 | 132/5000                     |
| 1      | 1        | 5.4k      | 1630MB/2Gi       | 0.287/1       | 3mins                     | 1024                 | 599/1000                     |
### Final Conclusion

We have following assumptions:
 - Input rate 5.3k
 - No of labels in each metric: 7
 - Current rate of metrics (prometheus_tsdb_head_samples_appended_total): 1.450k metrics/sec 


With a healthy sink, we see the following results:
- A throughput of [9.3k metrics/s](#with-healthy-sink) was reached (Queue: 5000, Batch size: 8192 and Memory: 1Gb)
- A throughput of [8k metrics/s](#tests-with-a-healthy-sink) was reached (Queue: 1024, Batch Size: 1024 and memory 1GB)

With an [unhealthy sink](#tests-with-a-unhealthy-sink), we see the memory limit playing a crucial role (achieved by creating configuration error).
- With 1 GB limit for gateway,  we see the gateway rejects logs after 30s instantly because scrape interval is 30s. The queue is not completely filled.
  - 598/1000 queue filled (Buffer 1024/Queue Size: 1024)
  - 68/5000 queue filled (Buffer 8192/ Queue Size: 5000)

- With [2 Gb limit](#with-unhealthy-sink) for gateway, we see the gateway rejecting logs after 2 minutes


Although we see that a higher throughput can be achieved by having a default buffer size of 8192, we decided to go with a lower value of 1024. The main reason was to be able to support various backends like Dynatrace, which supports a smaller buffer size. Also, we go with a queue size of 512, because we want to keep our setup similar to that with trace collector. The final configuration is the following:

| Property                                          | Value |
| ------------------------------------------------- | ----- |
| `processor.batch.timeout`                         | 10s   |
| `processor.memory_limiter.check_interval`         | 1s    |
| `processor.memory_limiter.limit_percentage`       | 75    |
| `processor.memory_limiter.spike_limit_percentage` | 10    |
| `processor.batch.send_batch_size`                 | 1024  |
| `processor.batch.send_batch_max_size`             | 0     |
| `exporter.otlp.sending_queue.queue-size`          | 512   |