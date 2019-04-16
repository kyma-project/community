# Event Bus Monitoring Metrics Proposal

## Event Bus
- **Throughput**
  
  Number of all events delivered that were successfully received by publish-app per minute

- **event propagation time 99 percentile**

  event propagation time 99 percentile

- **lag**

  number of events in queue

- **subscriptions**

  Number of subscriptions

- **subscriptions per {source ID &| event type}**

  Number of subscriptions per {source ID &| event type}

- **event activations**

  Number of event activations

- **consumers**

  Number of consumers
  
- **errors per {app &| tag &| ...}**

  Number of errors per {app &| tag &| ...}

## Publish App
- **events in**

  Number of received events

- **events in per sec**

  Number of events requests per second

- **events in per {source ID &| event type &| event type version}**

  Number of events requests per {source ID &| event type &| event type version}

- **Rate of events in per {source ID &| event type &| event type version} per sec**

  Rate of events requests per {source ID &| event type &| event type version} per second

- **succeeded events per {source ID &| event type &| event type version}**

  Number of succeeded events per {source ID &| event type &| event type version}

- **succeeded events per {source ID &| event type &| event type version} per sec**

  Rate of succeeded events per {source ID &| event type &| event type version} per second

- **ignored events per {source ID &| event type &| event type version}**

  Number of ignored events per {source ID &| event type &| event type version}

- **ignored events per {source ID &| event type &| event type version} per sec**

  Rate of ignored events per {source ID &| event type &| event type version} per second

- **failed events per {source ID &| event type &| event type version}**

  Number of failed events per {source ID &| event type &| event type version}

- **failed events per {source ID &| event type &| event type version} per sec**

  Rate of failed events per {source ID &| event type &| event type version} per second

- **latency 99 percentile**

  Latency 99 percentile of published event request

**Push App**
-
- **events in**

  Number of events push requests

- **events in per sec**

  Number of events push requests per second

- **pushed messages**

  Number of pushed messages to all consumers

- **pushed messages per sec**

  Number of pushed messages to all consumers per second

- **latency 99 percentile**

  Latency 99 percentile of pushed events to all consumers

- **failed pushed messages**

  Number of failed pushed messages

- **failed pushed messages per sec**

  Number of failed pushed messages per second

- **latency 99 percentile**

  Latency 99 percentile of pushed message to a single consumer