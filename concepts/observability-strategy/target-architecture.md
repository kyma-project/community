# Target picture

This document outlines the targeted picture we want to achieve in the long run. As the adoption of the Open Telemetry ecosystem is not big enough, temporarily solutions will be realized, which don't fit into the final picture. Also, the picture is going to be adjusted over time based on the learnings made on the way.

## The picture

![New Components](./assets/strategy-future-components.drawio.svg)

## Components
### The Manager

By default, the Telemetry module brings the `Telemetry Manager` only, which serves three new CRDs suffixed with `Pipeline` for the end user. The usage of a `Pipeline` instance activates the related components for the related signal type. The manager takes care of the full lifecycle of the setup.

### The Gateways

For every signal type, there will be a central gateway running. That gateway serves an ingestion endpoint in OTLP to which sources/agents can push the signal data. The gateway assures that all data is enriched with relevant resource attributes, and then dispatches the data to the configured backends. Durable buffering can be activated optional (only for logs?).

### The Agents

For every signal type, managed agents will be supported to serve typical setups for collecting data. The agents are all optional and will be activated as soon as a pipeline specifies the desire to leverage them. The agent configuration itself can be influenced only with the static module configuration, not per pipeline.

## Types
### Logs

The logging domain will provide an optional agent for tailing logs from containers running on the Nodes, supporting selection on Namespace name. Furthermore, Kubernetes events can be streamed as logs using a dedicated agent.
The gateway will provide a durable buffering if demanded, which requires to run the gateway as a StatefulSet with a maximum buffer capacity.

### Metrics

For metrics, the core agent will be the annotation-based scraper agent, which will support scraping custom metrics from your applications. On top, Istio proxy and system components can be scraped.
Furthermore, the typical Kubernetes agents will be supported for retrieving host metrics, kubelet metrics, and scraping APIServer metrics. A receiver for metrics about the resources stored at the Kubernetes APIServer will complement the setup.

The metrics gateway will not have any special feature.

### Traces

For traces, there will be no dedicated agent; instead the system components will feed the trace data. However, the gateway setup will be special. To support scalable tail-based sampling and filtering based on full traces, all spans of a trace must be processed by the same instance of the preceding StatefulSet. We achieve this with a two-stage component setup: First a deployment, which then dispatches all spans of a trace to the same instance.
