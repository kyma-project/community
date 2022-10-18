# Performance Evalation of Eventing on different Cloud Providers

## Table of Contents
- Test Setup
- Failure Types
- Test Results
- Screenshots of Crashes

## Test Setup
* Kyma CLI version: `2.6.2`
* Kyma: 
  * Production Profile
  * JetStream with File Storage
* K8s cluster:
  * Kubernetes v1.24.3, v1.24.6
  * Gardener cluster [Nodes: 2(min) to 6(max)]

* Kyma deploy command:
    ```
    kyma deploy --source=main --profile production ...
    ```
* Testing tool: [K6](https://k6.io/) (deployed in the same Kyma cluster)

## Failure Types:
Here are the common crash failures occured during the testing:
- F1  - Only one of the NATS Jetstream nodes storage is full
- F2  - All JetStream node storage is full
- F3  - NATS JetStream just doesn’t accept and dispatch events despite storage is not full. This is solely observed on Azure
<br/><br/>

## Tests Results
In this section the test outcomes of the 2 day running test with the [K6](https://k6.io/). As you see the given K6 tool rps (request-per-second) is different at what is observed at EPP. Therefore, the actual value is EPP rps.
As you see GCP and AWS performed similarly, also SSD and HDD difference was not much. HDD cannot be tested for AWS as it doesn't offer such storage type.
### GCP
CPU: 4, Memory: 15GB
| cluster | K6 rps | EPP rps | JS Crashed | version | test duration | Failure type | NATS version |
| --- | --- | --- | --- | --- | --- | --- | --- |
| gcp1 | 500 | 320 | yes | 2.7.0-rc2 | 5 hours | full storage (F1) |  |
| gcp2 | 250 | 232 | no | 2.7.0-rc2 |  |  |  |
| gcp3 | 350 | 315 | yes | 2.7.0-rc2 | 5.5 hours | full storage (F2) |  |
| gcp4 | 300 | 263 | yes | 2.7.0-rc2 | 13.5 hours | full storage (F2) |  |
| gcp5 | 280 | 240 | yes | 2.7.0-rc2 | 31.5 hours |  |  |
| gcp6 | 280 | 265 | yes | main |  | full storage (F2) | 2.9.0 |
| gcphdd1 | 250 | 240 | no | 2.7.0-rc2 |  |  |  |
| gcp-hdd | 300 | 278 | yes | main |  | full storage (F1) | 2.9.0 |
<br/>

### Azure
### Azure Performance Tests

CPU: 4, Memory: 16GB

| cluster | K6 rps | EPP rps | JS crashed | version | test duration | Failure type | NATS version |
| --- | --- | --- | --- | --- | --- | --- | --- |
| az12 | 250 | 88 (250 for 3h) | yes | main | 19.5 hours | F1 | 2.9.0 |
| az13 | 300 | 93 (298 for 3h) |  | main |  |  | 2.9.0 |
| az14 (hdd) | 300 | 70 (290 for 3h) |  |  |  |  | 2.9.0 |

<br/>

CPU: 2, Memory: 16GB

| cluster | K6 rps | EPP rps | JS crashed | version | test duration | Failure type | NATS version |
| --- | --- | --- | --- | --- | --- | --- | --- |
| ssd | 350 | 88 | no | 2.7.0-rc2 |  |  |  |
| ssd | 2400 | 140 | no | 2.7.0-rc2 |  |  |  |
| ssd | 2400 | 84 | yes | main | 11 hours | full storage (F1) | 2.9.0 |
| ssd | 4800 | 84 | yes | main |  | full storage (F2) | 2.9.0 |
| ssd | 6000 | 174 | yes | 2.7.1 | 11.4 hours | full storage (F2) |  |
| hdd | 1200 | 116 | yes | 2.7.0-rc2 | 4h 20 min | JS pod restart |  |
| hdd | 4800 | 104 | yes | main | 6 hours | F3 - appears w/ new NATS version. Started sending events after consumer leader reelection | 2.9.0 |
| hdd | 4800 | 180 | yes | main | 1.15 hours | F3 - appears w/ new NATS version. Started sending events after consumer leader reelection and restart of k6 pods.                      Probably full storage caused this, but nothing in the dashboard. | 2.9.0 |

<br/>

### AWS
CPU: 4, Memory: 16GB
|  | K6 rps | EPP rps | JS Crashed | version | test duration | Failure type  |
| --- | --- | --- | --- | --- | --- | --- |
| aws1 | 250 | 247 | no | 2.7.1 |  |  |
| aws2 | 500 | 365 | yes | 2.7.1 | 5 hours | F1   -  storage is full |
| aws3 | 350 | 307 | yes | 2.7.1 | 17 hours | F1   -   storage is full |

<br/>

## Takeaways

- Mostly storage is getting filled up. This is normal if publisher faster than subscriber. 
    
- Azure is the worst performer
    - And it is not consistent
    - New 2.9.0 NATS version gets disabled after some time (F1)
    
- GCP and AWS perform similarly


## Screenshots of Common Cases

F1 crash case:  
![image](https://user-images.githubusercontent.com/13185122/196392010-190aa8d5-400d-4f3b-9e91-209e1c0151b7.png)

F2 crash case:  
![image](https://user-images.githubusercontent.com/13185122/196392359-27b39535-c335-4cd6-8882-a85a8c9d549c.png)

F3 crash case:  
As you see in this case the storage is not full, but only 166822 out of 727029 messages are dispatched:  
![image](https://user-images.githubusercontent.com/13185122/196392916-e79b0e5b-410c-4550-a08b-36fb30495fe5.png)

![image](https://user-images.githubusercontent.com/13185122/196393505-3bd743b7-6e1f-406c-80df-53e90e8c2bea.png)




