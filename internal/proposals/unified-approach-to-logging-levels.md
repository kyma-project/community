# Unified approach to logging levels

It must be possible to easily set a more verbose logging level of each component within the Kyma ecosystem while troubleshooting.

We also have some "external" components (not implemented by the Kyma team) and we need an approach on how to change the logging level in them. For debugging reasons the level could be changed without any code and charts changes just with the in-cluster resource change.

## Default severity levels for the environments

To be consistent, we need to agree on what default level components will log in each environment:

- PROD: `INFO` - this is a minimal severity needed for performing the audit logging,
- STAGE: `INFO` - this is a minimal severity needed for performing the audit logging. It'd also be a good indicator on how the logs will look like on PROD environment,
- DEV: `DEBUG` - this is a lowest possible logging level. It may be helpful for debugging issues while developing.

In the current state we do log everything on every environment. Therefore, changing the logs severity could only decrease amount of the logs on some environments and there is no need to increase amount of needed resources for our components. A potential amount of the resources decreasement could be done while the next iteration of the evaluating Kyma profiles.

## Setting up the log level

Setting up the log level should be easy, do not take much resources and should be easily changeable. The log level will be set as a Helm Chart value. This solution will provide easy possibility to change the log level in every situation.

In the `values.yaml` file of each Kyma component there will be:

```yaml
logLevel: "info"
```

In the `deployment.yaml` or other component container specification there will be an environment variable, for example:

```yaml
    spec:
      containers:
        - env:
          - name: LOG_LEVEL
            value: {{ .Values.global.logLevel }}
          ...
```

The default value for the `logLevel` could be `"info"` so there will be additional resources - overrides - present only on the DEV environment.

In the application there will be an environment variable reader which will parse the value and set the proper logging level in the logger.

While debugging there will be an easy option to change the logging level via kubernetes API, for example:
```bash
kubectl -n compass-system edit deployment compass-runtime-agent
```
Then editting the `LOG_LEVEL` environment variable and hitting `:wq`. The pod will restart with the new logging level value and will be ready to debug.

## Setting up the log level in the external components

> TODO: Brainstorm on it. Currently, we can just filter the logs using for example grep. Damian is on it
