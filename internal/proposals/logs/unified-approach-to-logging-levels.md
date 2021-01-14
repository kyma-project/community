# Unified approach to logging levels

We need a consistent way to put the logs into the specific logging severity levels. We also should have a default severity level on each environment such as DEV, STAGE and PROD.
It also must be possible to easily set a more verbose logging level of each component within the Kyma ecosystem while troubleshooting.

We have some "external" components too (not implemented by the Kyma team) and we need an approach on how to change the logging level in them. For debugging reasons the level could be changed without any code and charts changes just with the in-cluster resource change.

## What to put into each severity level

We decided to use the Zap library as the logger. Here is a description of all the needed levels:

- `FATAL` - a component reached a catastrophic error and cannot go further,
- `ERROR` - errors that break something, such as request failures, reading memory failures,
- `WARN` - things that do not break anything but suggest that potential setup should be changed, for example: resources are running low, optional config cannot be fetched, a request failed but will be retried, a connection is lost but will be reacquired,
- `INFO` - requests handling, operating on resources, successes, information about operations in progress, and so on,
- `DEBUG` - options with which the component is deployed, requests' paths, and other information that is useful while making sure that everything works as expected.

No other severity will be used as it is not needed.

## Default severity levels for the environments

To be consistent, we need to agree on what default level components will log in each environment:

- PROD: `WARN` - this is a minimal severity needed for effective debugging of issues,
- STAGE: `WARN` - this is a minimal severity needed for effective debugging of issues. It'd also be a good indicator on how the logs will look like on PROD environment,
- DEV: `DEBUG` - this is a lowest possible logging level. It may be helpful for debugging issues while developing.

In the current state we log everything on every environment. Therefore, changing the logs severity could only decrease amount of the logs on some environments, and there is no need to increase the amount of needed resources for our components. A potential amount of the resources decreasement could be done during the next iteration of evaluating the Kyma profiles.

## Setting up the log level

Setting up the log level should be easy, shouldn't take much resources and should be easily changeable. The log level will be set as a Helm chart value. This solution will provide easy possibility to change the log level in every situation.

In the `values.yaml` file of each Kyma component there will be:

```yaml
log:
  level: "warn"
```

In the `deployment.yaml` or other component container specification there will be an environment variable, for example:

```yaml
    spec:
      containers:
        - env:
          - name: APP_LOG_LEVEL
            value: {{ .Values.global.log.level }}
          ...
```

The default value for the `log.level` could be `"warn"` so there will be additional resources - overrides - present only on the DEV environment.

In the application there will be an environment variable reader which will parse the value and set the proper logging level in the logger.

While debugging there will be an easy option to change the logging level via kubernetes API, for example:
```bash
kubectl -n compass-system edit deployment compass-runtime-agent
```
Then editing the `APP_LOG_LEVEL` environment variable and hitting `:wq`. The pod will restart with the new logging level value and will be ready to debug.

There are components that should not be restarted, so this approach couldn't be used. We should come up with some ideas on how to change their log severity level. A potential solution is to create a configmap, and a listener that would periodically read the value from this resource. [Here is a follow-up issue](https://github.com/kyma-project/community/issues/530).

## Setting up the log level in the external components

Currently, we do have a lot of external components that are set up in many ways. We decided that preparing solutions for all of them is out of scope of this proposal.
