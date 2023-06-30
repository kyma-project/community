 
## Telemetry
- Moved Loki LogPipeline of logging component to telemetry module, so that it gets installed if logging component is available but the component will not fail on the absence of the telemetry module (https://github.com/kyma-project/kyma/issues/17549)
- Update of components 
  - Otel-collector 0.79.0 (https://github.com/kyma-project/kyma/pull/17629)
  - Fluent bit 2.1.4 (https://github.com/kyma-project/kyma/pull/17658)
- Improved custom loki example (https://github.com/kyma-project/examples/pull/243)
- Support multiple TracePipelines

Added second replica of telemetry-trace-collector to improve availablility

## Service Mesh
### Istio upgraded to 1.18.0
In this release, we upgraded Istio from 1.17.3 to 1.18.0. For more details on the changes, read the official [Istio 1.18.0 release notes](https://istio.io/latest/news/releases/1.18.x/announcing-1.18/upgrade-notes/). 

## API Gateway
### APIRule v1alpha1end of life
APIRule v1alpa1 is deprecated since kyma 2.5. End of life is planned for Kyma 2.19. Migrate your APIRules to v1beta1.

## Eventing
Update nats-server to 2.9.18


## Serverless
### Deprecation of Node.js 16 Serverless runtime
Because of the [planned EOL for Node.js 16](https://github.com/nodejs/release#release-schedule) we are planning to remove Node.js 16 from the list of the supported runtimes.

For now, we recommend that you don’t use Node.js 16 as a runtime for your new Functions and re-configure all your existing Node.js 16 Functions to run on the latest available Node.js runtime.

See this [blog post](https://blogs.sap.com/2022/03/09/changing-the-function-runtime-version-of-a-running-function/) to learn how to update existing Functions.

### Node.js 14 runtime removed
Node.js 14 has reached the end of its life. Therefore, followed by the depreciation of Node.js 14 Serverless runtime, we decided to finally remove it from the list of the available Function runtimes.

Your Node.js 14 Functions' workloads will continue to run, but you will not be able to edit them without changing the `runtime` field. Sooner or later, you must change the spec of your existing Node.js 14-based Functions and change the `runtime` field to either `nodejs16` or `nodejs18`.

For more information about the Node.js 14 deprecation, see the [Kyma 2.12 release notes](https://github.com/kyma-project/kyma/releases/tag/2.12.0).
