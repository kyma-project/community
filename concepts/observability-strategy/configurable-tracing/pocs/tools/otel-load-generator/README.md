# Opentelemetry Load Generator

The application currently generate only trace data, each generated trace contains 1 root span and 10 child spans. The child spans generated with additional 40 attributes each attribute has a 64 byte long random value.
The generated test data will be sent over `gRPC` protocol with batches, each batch contains `512` spans.

## Usage

Load generator has two start parameter:
- -t: The URL of target collector
- -c: This parameter configure concurrent test data generator

```bash
/loadgenerator -t telemetry-collector:4317 -c 10
```
