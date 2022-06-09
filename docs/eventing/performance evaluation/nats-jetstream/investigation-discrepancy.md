# Investigation of the discrepancy between sent and received events

## Why was "Events Sent" not the same as "Events Received" in some test runs?

There were some test runs where the "Events Sent" didn't match "Events Received". This document investigates one such suspicious test run.

> **Conclusion:** Once the minReplicas for the receiver (that is, the sink) was changed to six (so that it won't be scaling up during the tests), we haven't seen any job whose "Events Sent" didn't match "Events Received". Therefore, maybe the scaling up of the receiver caused the issue.

### State before test run
- Stream: 
  - LastSeq# 238,410
- Consumer: 
  - Ack Floor: Stream sequence# 238,410

### State after test run
- Stream: 
  - LastSeq# 360,137
- Consumer: 
  - Ack Floor: Stream sequence# 360,137
  - Redelivered Messages: 0
  - Unprocessed Messages: 0

```
* Total Events Received   by **Stream**   : 360,137 - 238,410 = 121,727
* Total Events Processed  by **Consumer** : 360,137 - 238,410 = 121,727
```
---

### Test Run Dashboard:

![](assets/1000_loss1.png "")

![](assets/1000+loss3run.png "")
![](assets/loss_1000_stream.png "")
![](assets/loss_1000_consumer.png "")

