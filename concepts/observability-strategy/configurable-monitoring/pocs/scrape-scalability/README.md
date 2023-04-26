# Scrape Scalability

## Goal

The collector must support scraping metrics in the Prometheus pull approach. Thus, the [Prometheus receiver](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver) must be used.
Because the instance keeps state about the scraped data, one target cannot be scraped by interchanging instances (sticky targets, we could say).
What will happen if there is a large amount of scrape targets and one instance cannot scrape them all?
Probably only sharding can solve the problem by distributing the targets to scrape to multiple instances.

## Ideas
- Coordinator: Have an instance like the otel-collector-operator which splits the scrape config into distinct sets and assignes it collector instances.
- Self-contained criteria: Support configs only by service discovery and annotations, and the collector itself is responsible for sharding the targets based on dynamic criterias like Namespace or Node.

## Proposal
Support only dynamic configuration of targets using service discovery and annotations. Use service discovery attributes as criteria inside every collector to shard the targets over the instances.
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

## Summary

The DaemonSet approach based on sharding by Node assignment is a very simple approach of solving the problem. It allows to scale out in a limited way as there can be situations were particular instances can reach limits dependent on the Node size.
A better approach will be to use a StatefulSet where the scrape pool is sharded in a deterministic way using the instance ID as selector. Then a scale-out can happen dynamically. However, that approach is much more complex to realize, because StatefulSets are harder to manage and probably need a config per instance. For now, the DaemonSet approach should be sufficient. We can switch to the StatefulSet approach if we see that it is actually needed.
