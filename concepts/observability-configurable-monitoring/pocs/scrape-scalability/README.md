# Scrape Scalability

## Goal

The agent must support scraping metrics in the prometheus pull approach. For that the [prometheusreceiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver) needs to be used.
As the instance keeps state about the scraped data, one target cannot be scraped by interchanging instances (sticky targets we could say).
What will happen if there a magnitude of scrape targets and one instance cannot scrape them all?
Probably only sharding can solve the problem by distributing the targets to scrape to multiple instances.

## Ideas
- Coordinator: Have an instance like the otel-collector-operator which splits the scrape config into distinct sets and assignes it collector instances.
- Self-contained Criterias: Support configs only via service discovery and annotations, then shard the targets via dynamic criterias like namespace or node by the agent itself

## Proposal
Support only dynamic confgiuration of targets via Service Discovery and annotations. Use service discovery attributes as criteria inside every agent to shard the targets over the instances.
Proposal: Use the "node" as criteria element and have one instance per node running. 

Pro:
- No coordinator required
- Scaling of the instances out of the box available by using a daemonset
- If a node crashes anyway all workload is gone, so it is ok that the collector will not scrape in that moment (as it is gone as well)
- The shard size is limited by nature as there is a limited amount of pods which can run on one node

Cons:
- Scale factor might be not enough, there could be a lot of tiny pods running on one big node
- There might be very few pods on a node and the related instance is just idle (however, nodes will scale down at some point)

## PoC

Take a standard kyma cluster and disable the prometheus scraping and enable remotewrite by patching the [prometheus.yaml](./prometheus.yaml).
Then deploy an otel-collector as daemonset using:
```bash
helm upgrade opentelemetry-collector open-telemetry/opentelemetry-collector --version 0.20.0 --install --namespace kyma-system -f otel-collector-values-deployment.yaml
```

That will deploy the otel-collector as a deployment, now try to bring the setup to the limits by deploying workload with custom metrics, being annotated with:
```yaml
prometheus.io/scrape: true
prometheus.io/scheme: https
prometheus.io/port: 9090
```
Figure out what the limits are. Can the limits be pushed by vertical scaling? Are there some borders where resource scale-up will not help?

Afterwards, switch to a daemonset approach by swithcing the values.yaml to otel-collector-values-daemonset.yaml.
That will deploy the collector as daemonset having the scrape config already enriched by sharding based on nodes.

Run the test again. Can that scalability be good enough?
