# Performance Evalation of JetStream (Kyma: Production Profile)

## Table of Contents
- Test Setup
- Test Scenario 1: Without any server restarts/crash
- Test Scenario 2: NATS Servers deleted during test
- Test Scenario 3: NATS Servers scaled down to 0 and back to 3 during test
- Test Scenario 4: Eventing-controller Pod deleted during test

## Test Setup
* Testing tool: [K6](https://k6.io/)
* Kyma CLI version: `2.2.0`
* Kyma: 
  * Version: main [[commit](https://github.com/kyma-project/kyma/commit/6eb300b0a159fa968763382fbfe918b8cb52f057) and [commit (includes bug fix)](https://github.com/kyma-project/kyma/commit/f8a0c28a43e9eebf192514acc61614300f9909a1)] 
  * Production Profile
  * JetStream with File Storage
* K8s cluster:
  * Kubernetes v1.21.10
  * Gardener cluster [Nodes: 3(min) to 6(max)]
  * GCP machine type: `n1-standard-4`

* Kyma deploy command:
  ```
  kyma deploy --source=main -p production --value global.jetstream.enabled=true --value global.jetstream.storage=file
  ```

## Test Scenario 1: Without any server restarts/crash

### Run ID: 7/6/2022T13:23 [Simple NATS with JetStream Disabled] (Duration: 10m, Event Rate: 150rps)

> **NOTE:** This is the only test run with JetStream disabled. All the other tests were done with JetStream enabled.

![](assets/NATS_07_06_22-10-150_1.png "")

### Run ID: 1/6/2022T11:58 (Duration: 10m, Event Rate: 150rps)
![](assets/01_06_22-10-150_1.png "")

### Run ID: 1/6/2022T12:17 (Duration: 10m, Event Rate: 150rps)
![](assets/01_06_22-10-150_2.png "")

### Run ID: 1/6/2022T12:31 (Duration: 10m, Event Rate: 150rps)
![](assets/01_06_22-10-150_3.png "")

### Run ID: 1/6/2022T12:45 (Duration: 10m, Event Rate: 200rps)
![](assets/01_06_22-10-200_1.png "")

### Run ID: 1/6/2022T13:2 (Duration: 10m, Event Rate: 200rps)
![](assets/01_06_22-10-200_2.png "")

### Run ID: 2/6/2022T7:58 (Duration: 10m, Event Rate: 250rps)
![](assets/02_06_22-10-250_1.png "")

### Run ID: 2/6/2022T8:28 (Duration: 10m, Event Rate: 250rps)
![](assets/02_06_22-10-250_2.png "")

### Run ID: 2/6/2022T10:56 (Duration: 10m, Event Rate: 1000rps)
![](assets/02_06_22-10-1000_1.png "")

---

## Test Scenario 2: NATS Servers deleted during test

> **NOTE:** Deleted (using `kubectl delete`) all thres Pods of NATS at once after 4 minutes.

```
kubectl delete po -n kyma-system eventing-nats-0
kubectl delete po -n kyma-system eventing-nats-1
kubectl delete po -n kyma-system eventing-nats-2
```

### Run ID: 2/6/2022T13:20 (Duration: 10m, Event Rate: 150rps)

```
-> Total Events Sent      by **Test Sender** : 77,625 (+ 5514 Failed = 83,139)
-> Total Events Received  by **Stream**      : 77,625
-> Total Events Processed by **Consumer**    : 77,625
-> Total Events Received  by **Sink**        : 77,629 (Means that 4 events were duplicates)
```

![](assets/crash1_1.png "")
![](assets/crash1_2.png "")
![](assets/crash1_4.png "")
![](assets/crash1_3.png "")

---

## Test Scenario 3: NATS Servers scaled down to 0 and back to 3 during test

> **NOTE:** Scaled down NATS statefulset to 0 after 4 minutes.

```
kubectl scale statefulset eventing-nats -n kyma-system --replicas 0
kubectl scale statefulset eventing-nats -n kyma-system --replicas 3
```

### Run ID: 2/6/2022T13:54 (Duration: 10m, Event Rate: 150rps)

**State before test run:**
- Stream:  
  - LastSeq# 77,625
- Consumer:  
  - Ack Floor: Stream sequence# 77,625

**State after test run:**
- Stream: 
  - LastSeq# 149,259
- Consumer:  
  - Ack Floor: Stream sequence# 149,259
  - Redelivered Messages: 0
  - Unprocessed Messages: 0

```
-> Total Events Sent      by **Test Sender** : 71,634 (+ 11,113 Failed = 82,747)
-> Total Events Received  by **Stream**      : 71,634 (i.e. 149,259 - 77,625)
-> Total Events Processed by **Consumer**    : 71,634 (i.e. 149,259 - 77,625)
-> Total Events Received  by **Sink**        : 71,637 (Means that 3 events were duplicates)
```

![](assets/crash2_1.png "")
![](assets/crash2_2.png "")
![](assets/crash2_4.png "")
![](assets/crash2_3.png "")

---
---
## Test Scenario 4: Eventing-controller Pod deleted during test

### Run ID: 2/6/2022T14:17 (Duration: 10m, Event Rate: 150rps)
> **NOTE:** Deleted (using `kubectl delete`) the Pod of eventing-controller after 4 minutes.

**State before test run:**
- Stream:  
  - LastSeq# 149,259
- Consumer:  
  - Ack Floor: Stream sequence# 149,259

**State after test run:**
- Stream:  
  - LastSeq# 237,514
- Consumer:  
  - Ack Floor: Stream sequence# 237,514
  - Redelivered Messages: 0
  - Unprocessed Messages: 0

```
-> Total Events Sent      by **Test Sender** : 88,255
-> Total Events Received  by **Stream**      : 88,255 (i.e. 237,514 - 149,259) 
-> Total Events Processed by **Consumer**    : 88,255 (i.e. 237,514 - 149,259)
-> Total Events Received  by **Sink**        : 88,255
```

![](assets/crash3_1.png "")
![](assets/crash3_2.png "")
![](assets/crash3_4.png "")
![](assets/crash3_3.png "")

---

## Finding

Eventing controller reaches the CPU limit (500m) on 100 events/sec. Same result for NATS with JetStream disabled.
