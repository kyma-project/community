# Event Bus Monitoring Metrics Proposal

## Event Bus Metrics

|Metric Name |Description |Motivation |
|------------|------------|-----------|
|Throughput |Number of all events delivered that were successfully received by publish-app per minute | Gives insight about the whole system performance|
|event propagation time 99 percentile |event propagation time 99 percentile | Gives more accurate insight about the whole system performance and events delivery|
|lag | number of events in queue | Helps in diagnosing system performance as well as troubleshooting events not delivered cases|
|kyma subscriptions per {namespace &\| source ID &\| event type &\| event type version &\| ready &\| endpoint} |Number of Kyma subscriptions per {namespace &\| source ID &\| event type &\| event type version &\| ready &\| endpoint} | Helps in events not delivered cases and insights about the subscriptions distribution over multiple criteria|
|event activations per {namespace &\| source ID} | Number of event activations per {namespace &\| source ID} | Can help in troubleshooting events not delivered cases and filtering the activated subscriptions based on different criteria|
|knative subscriptions per {namespace &\| channel name &\| channel ready &\| subscriber &\| ready}   |Number of knative subscriptions per {namespace &\| channel name &\| channel ready &\| subscriber &\| ready}  | Comparing this metric to the number of Kyma subscriptions can help in troubleshooting events not delivered cases and system integrity|
|knative channels per {namespace &\| subscriber name &\| subscriber URI &\| source ID &\| event type &\| event type version} |Number of knative channels per {namespace &\| subscriber name &\| subscriber URI &\| source ID &\| event type &\| event type version} | Helps in troubleshooting events not delivered cases and system integrity by filtering the created channels by different criteria and comparing to the correspondent knative subscriptions|
|consumers | Number of consumers | Helps in troubleshooting events not delivered cases and system integrity|
|errors per {namespace &\| app &\| tag &\| ...} | Number of errors per {namespace &\| app &\| tag &\| ...}   | Helps in troubleshooting system failures cases, troubleshooting events not delivered cases and system integrity and fixing system bugs|
|middleware availability | Health monitor of middleware like NATS Streaming | Helps in troubleshooting system failures cases, troubleshooting events not delivered cases and system integrity|

## Publish App Metrics

|Metric Name |Description |Motivation |
|------------|------------|-----------|
|events in |Total Number of received events | Gives insights about system load|
|events in per {namespace &\| source ID &\| event type &\| event type version} |Number of events requests per {namespace &\| source ID &\| event type &\| event type version} | Filtering received events by different criteria combination can help with giving a better insight about the system load as well as troubleshooting events not delivered cases and system integrity|
|succeeded events per {namespace &\| source ID &\| event type &\| event type version} | Number of succeeded events per {namespace &\| source ID &\| event type &\| event type version} | Helps in troubleshooting events not delivered cases and system integrity as well as being a system health indicator|
|ignored events per {namespace &\| source ID &\| event type &\| event type version} | Number of ignored events per {namespace &\| source ID &\| event type &\| event type version} | Helps in troubleshooting events not delivered cases and system integrity as well as being a system health indicator|
|failed events per {namespace &\| source ID &\| event type &\| event type version} | Number of failed events per {namespace &\| source ID &\| event type &\| event type version} | Helps in troubleshooting events not delivered cases and system integrity as well as being a system health indicator|
|latency 99 percentile | Latency 99 percentile of published event request | High latency maybe a reason for not delivered events in some cases and it gives an insight about the system health and performance

## Push App Metrics

|Metric Name |Description |Motivation |
|------------|------------|-----------|
|events in | Number of events push requests | Comparing this metric to the received events by the publish app metric can help in troubleshooting events not delivered cases and system integrity|
|pushed messages | Number of pushed messages to all consumers | Giving insight about knative eventing health by monitoring the rates as well as assuring the system integrity|
|latency 99 percentile to all consumers|Latency 99 percentile of pushed events to all consumers | Gives insight about the knative/events delivery system performance and can help in troubleshooting events not delivered scenarios|
|latency 99 percentile to a single consumer | Latency 99 percentile of pushed message to a single consumer | Assessing the system delivery performance and contributes to judging the whole system performance diagnosis|
|failed pushed messages | Number of failed pushed messages | Help in troubleshooting events not delivered scenarios and system integrity|
| | |