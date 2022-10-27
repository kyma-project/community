# Throughput Evalation of NATS Jetstream on different Cloud Providers

## Table of Contents
- Test Setup
- Test Results
- How tests are executed

## Test Setup
* Kyma CLI version: `2.6.2`
* Kyma: 
  * Version: main 
  * Production Profile
  * JetStream with File Storage
  * NATS JetStream Version: `2.9.0` 
* K8s cluster:
  * Kubernetes v1.24.6
  * Gardener cluster [Nodes: 2(min) to 6(max)]
  * Machine Type: `cpu: 2, memory: 16GB`

* Kyma deploy command:
    ```
    kyma deploy --source=${kyma_version} --profile production --value eventing.nats.nats.jetstream.fileStorage.storageClassName=${storage_class_name}
    ```
* Testing tool: [NATS Benchmarking Tool](https://docs.nats.io/using-nats/nats-tools/nats_cli/natsbench)
<br/><br/>

## Test Results
As you see from the table all the NATS Jetstream throughput tests were similar on GCP, Azure, and AWS. 

| Cloud Provider | msgs/sec (NATS Bench) | msgs | storage type | pub/sub | CPU / Memory |
| --- | --- | --- | --- | --- | --- |
| Azure | 4130 - 4313 | 10000 | SSD | 5/5 | 4 / 16 |
| GCP | 2000 - 2500 | 10000 | SSD | 5/5 | 4 / 15 |
| AWS | 3441 - 4289 | 10000 | SSD | 5/5 | 4 / 16 |
|  |  |  |  |  |  |
| Azure | 1630 - 2682 | 10000 | SSD | 1/5 | 4 / 16 |
| GCP | 850  - 1320 | 10000 | SSD | 1/5 | 4 / 15 |
| AWS | 1739 - 2071 | 10000 | SSD | 1/5 | 4 / 16 |
|  |  |  |  |  |  |
| GCP | 2297-2447 | 10000 | HDD | 5/5 | 4 / 16 |
|  |  |  |  |  |  |
| Azure | 940   - 1198 | 10000 | SSD | 5/5 | 2 / 16 |
| Azure | 502  - 584 | 10000 | SSD | 1/5 | 2 / 16 |
<br/>

## How tests are executed
Run the following commands against the installed NATS Jetstream

### Stop EC

```
$ kubectl -n kyma-system scale --replicas 0 deploy eventing-controller
```

### Get into nats-box pod to run benchmarks in k8s cluster:
```
$ kubectl -n kyma-system run -i --rm --tty nats-box --image=natsio/nats-box --restart=Never
```
### Create consumer
Create a durable consumer simulating Kyma NATS Jetstream consumer:
```
$ nats con add sap test-consumer --server=nats://eventing-nats:4222" --deliver-group test-consumer-GROUP --target test-consumer-DELIVERY --filter "kyma.sap.kyma.custom.natsbenchmark.execution.created.v1" --deliver new --ack explicit --replay instant --max-deliver 100 --max-pending 10 --flow-control --heartbeat 1m --wait 30s
```

### Testing with durable consumer
```
$ nats bench "kyma.sap.kyma.custom.natsbenchmark.execution.created.v1" --js --stream sap --consumer test-consumer --push --syncpub --consumerbatch=10 --pub 10 --sub 5 --size 1024 --msgs 300000
```

Note: there is an issue that most of the time the tools gets stuck. Just try multiple times to get the results. If publisher and subscriber throughtput is similar, then it doesn't get suck according to the observation.

### Using nats-top CLI
[nats-top](https://docs.nats.io/using-nats/nats-tools/nats_top) cli can be used to monitor NATS Jetsream server. You can monitor incoming and outging msms/sec, bytes/sec, etc.

### Sample executions
Here are some sample NATS benchmarking executions for GCP. This is just to show the executions look like:

pub: 5, sub: 5
```
nats-box:~# ky-nats bench "kyma.sap.kyma.custom.natsreceiver.order.created.v1" --js --stream sap --consumer 1c40e439fbcfe6b540fbca6f4af84d52 --push --syncpub --consumerbatch=10 --pub 5 --sub 5 --size 1024 --msgs 10000
...
NATS Pub/Sub stats: 2,580 msgs/sec ~ 2.52 MB/sec
 Pub stats: 1,301 msgs/sec ~ 1.27 MB/sec
  [1] 382 msgs/sec ~ 382.60 KB/sec (2000 msgs)
  [2] 309 msgs/sec ~ 309.38 KB/sec (2000 msgs)
  [3] 263 msgs/sec ~ 263.51 KB/sec (2000 msgs)
  [4] 260 msgs/sec ~ 260.50 KB/sec (2000 msgs)
  [5] 260 msgs/sec ~ 260.26 KB/sec (2000 msgs)
  min 260 | avg 294 | max 382 | stddev 47 msgs
 Sub stats: 2,584 msgs/sec ~ 2.52 MB/sec
  [1] 529 msgs/sec ~ 529.82 KB/sec (2000 msgs)
  [2] 531 msgs/sec ~ 531.64 KB/sec (2000 msgs)
  [3] 529 msgs/sec ~ 529.07 KB/sec (2000 msgs)
  [4] 518 msgs/sec ~ 518.35 KB/sec (2000 msgs)
  [5] 518 msgs/sec ~ 518.37 KB/sec (2000 msgs)
  min 518 | avg 525 | max 531 | stddev 5 msgs
```
pub: 1, sub: 5
```
nats-box:~# ky-nats bench "kyma.sap.kyma.custom.natsreceiver.order.created.v1" --js --stream sap --consumer 1c40e439fbcfe6b540fbca6f4af84d52 --push --syncpub --consumerbatch=10 --pub 1 --sub 5 --size 1024 --msgs 10000
...
NATS Pub/Sub stats: 990 msgs/sec ~ 990.46 KB/sec
 Pub stats: 495 msgs/sec ~ 495.58 KB/sec
 Sub stats: 495 msgs/sec ~ 495.55 KB/sec
  [1] 102 msgs/sec ~ 102.70 KB/sec (2000 msgs)
  [2] 101 msgs/sec ~ 101.19 KB/sec (2000 msgs)
  [3] 100 msgs/sec ~ 100.06 KB/sec (2000 msgs)
  [4] 99 msgs/sec ~ 99.27 KB/sec (2000 msgs)
  [5] 99 msgs/sec ~ 99.16 KB/sec (2000 msgs)
  min 99 | avg 100 | max 102 | stddev 1 msgs
```
