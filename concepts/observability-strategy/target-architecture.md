# Target picture

The document outlines the targetted picture to be achieved in the long run. As the adoption of the Open Telemetry ecosystem is not big enough, temporarily solutions will be realized which are not fitting into that picture. ALso, the picture gets adjusted over time based on the learnings made on the way.

## The picture

![New Components](./assets/strategy-future-components.drawio.svg)

## Components
### The Manager

By default, the telemetry module brings the `Telemetry Manager` only which is serving three new CRDs suffixed with `Pipeline` for the end-user. The usage of a `Pipeline` instance will activate the related components for the related signal type. The manager takes care of the full lifecycle of the setup.

### The Gateways

For every signal type there will be a central gateway running. That is serving an ingestion endpoint in OTLP where sources/agents can push the signal data to. The gateway will assure that all data gets enriched with relevant resource attributes and then dispatches the data to the configured backends. Durable buffering can be activated optional (only for logs?).

### The Agents

For every signal type, managed agents will be supported to serve typical setups for collecting data. The agents are all optional and will be activated as soon as a pipeline specifies the desire to leverage them.The agent configuration itself can be influenced via the static module configuration only, not per pipeline.

## Types
### Logs

The logging domain will provide an optional agent for tailing logs from containers running on the nodes, supporting selection on namespace name. Furthermore, kubernetes events can be streamed as logs via a dedicated agent.
The gateway will provide a durable buffering if demanded which requires to run the gateway as a StatefulSet with a maximum buffer capacity.

### Metrics

For metrics the core agent will be the annotation based scraper agent which will support scraping custom metrics from your applications. On top, istio proxy and system components can be scraped.
Furthermore, the typical kubernetes agents will be supported for retrieving host metrics, kubelet metrics and scraping APIServer metrics. A receiver for metrics about the resources stored at the kubernetes APIServer will complement the setup.

The metrics gateway will not have any special feature.

### Traces

For traces there will be no dedicated agent, instead the system components will feed the trace data. However, the gateway setup will be special. In order to support a scalable tail based sampling, a two-stage component setup is needed. First a deployment which is then detaching spans belonging to the same trace always to the same instance of the preceding StatefulSet. That will allow filtering based on full traces.
