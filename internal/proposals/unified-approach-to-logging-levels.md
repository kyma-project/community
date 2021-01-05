# Unified approach to logging levels

It must be possible to easily set a more verbose logging level of each component within the Kyma ecosystem while troubleshooting.

We also have some "external" components (not implemented by the Kyma team) and we need an approach on how to change the logging level in them. For debugging reasons the level could be changed without any code and charts changes just with the in-cluster resource change.

## Setting up the log level

The log level will be set as a Helm Chart value. This solution will provide easy possibility to change the log level.

In the `values.yaml` file there will be:

```yaml
logLevel: "info"
```

In the `deployment.yaml` or other component container specification there will be for example:

```yaml
    spec:
      containers:
        - env:
          - name: LOG_LEVEL
            value: {{ .Values.global.logLevel }}
          ...
```

> TODO: Consider passing the log level to the Installer so there won't be a massive amount of overrides for every single component in all the environments

## Default severity levels for the environments

- PROD: INFO
> TODO: Brainstorm on how the audit logging is currently performed. Maybe there is no need to log on INFO level on the production and WARN will be enough 
- STAGE: INFO
- DEV: DEBUG

> TODO: Make sure that increased amount of logs (lowered the logging level) won't overload the cluster
> TODO: Align components' resources with the increased consumption caused by the increased amount of logs

## Setting up the log level in the external components

> TODO: Brainstorm on it. Currently, we can just filter the logs using for example grep
