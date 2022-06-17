# Concept

## Architecture Overview

- Workloads will be collected by annotations only via scraper, otherwise use OTLP push with a sidecar
- Workloads can push via OTLP
- Scrapers are running as daemonset using a node sharding distribution, they push via OTLP to the central deployment
- Central deployment has only OTLP input and scales dependent on load
- Telemetry Operator configured pieline of central collector
- Monitoring component brings config to push metrics via remotewrite to prometheus
