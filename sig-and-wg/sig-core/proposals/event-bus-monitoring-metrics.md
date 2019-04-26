# Event Bus Monitoring Metrics Proposal

## Event Bus

- **Throughput**
  
  Number of all events delivered that were successfully received by publish-app per minute

- **event propagation time 99 percentile**

  event propagation time 99 percentile

- **lag**

  number of events in queue

- **subscriptions per {namespace &| source ID &| event type &| event type version &| ready &| endpoint}**

  Number of subscriptions per {namespace &| source ID &| event type &| event type version &| ready &| endpoint}

- **event activations per {namespace &| source ID}**

  Number of event activations per {namespace &| source ID}

- **knative subscriptions per {namespace &| channel name &| channel ready &| subscriber &| ready}**

  Number of subscriptions per {namespace &| channel name &| channel ready &| subscriber &| ready}

- **knative channels per {namespace &| subscriber name &| subscriber URI &| source ID &| event type &| event type version}**

  Number of knative channels per {namespace &| subscriber name &| subscriber URI &| source ID &| event type &| event type version}

- **consumers**

  Number of consumers
  
- **errors per {namespace &| app &| tag &| ...}**

  Number of errors per {namespace &| app &| tag &| ...}

- **middleware availability**

  Health monitor of middleware like NATS Streaming

## Publish App

- **events in**

  Number of received events

- **events in per sec**

  Number of events requests per second

- **events in per {namespace &| source ID &| event type &| event type version}**

  Number of events requests per {namespace &| source ID &| event type &| event type version}

- **Rate of events in per {namespace &| source ID &| event type &| event type version} per sec**

  Rate of events requests per {namespace &| source ID &| event type &| event type version} per second

- **succeeded events per {namespace &| source ID &| event type &| event type version}**

  Number of succeeded events per {namespace &| source ID &| event type &| event type version}

- **succeeded events per {namespace &| source ID &| event type &| event type version} per sec**

  Rate of succeeded events per {namespace &| source ID &| event type &| event type version} per second

- **ignored events per {namespace &| source ID &| event type &| event type version}**

  Number of ignored events per {namespace &| source ID &| event type &| event type version}

- **ignored events per {namespace &| source ID &| event type &| event type version} per sec**

  Rate of ignored events per {namespace &| source ID &| event type &| event type version} per second

- **failed events per {namespace &| source ID &| event type &| event type version}**

  Number of failed events per {namespace &| source ID &| event type &| event type version}

- **failed events per {namespace &| source ID &| event type &| event type version} per sec**

  Rate of failed events per {namespace &| source ID &| event type &| event type version} per second

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