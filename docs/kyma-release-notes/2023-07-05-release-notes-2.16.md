Listen up, my fellow travelers! Your schedules may already be jam-packed with exciting plans for this vacation season, but let us tell you about one more adventure that's too good to miss. Embark on a journey to enchanting Two Sixteen Island, nestled in the very heart of the Kyma archipelago. Once there, get ready to discover the latest and greatest versions of Istio, OTel-collector, Fluent Bit, and NATS server, along with multiple Telemetry enhancements. Read on to uncover the wonders of Kyma 2.16!

## Telemetry
- We have moved the Loki LogPipeline out of the Logging component into the Telemetry module. Thus, it is installed if the Logging component is available, but the component does not fail if the Telemetry module is absent. For more information, see the [PR](https://github.com/kyma-project/kyma/issues/17549).
- The following components have been updated:
  - [OTel-collector to version 0.79.0](https://github.com/kyma-project/kyma/pull/17629)
  - [Fluent Bit to version 2.1.4](https://github.com/kyma-project/kyma/pull/17658)
- We have improved the [custom Loki example](https://github.com/kyma-project/examples/pull/243).
- We have introduced support for multiple TracePipelines.
- To improve availability, we have added a second replica of `telemetry-trace-collector`.

## Service Mesh
### Istio upgraded to version 1.18.0
With this release, we have upgraded Istio from version 1.17.3 to 1.18.0. To learn more about the new version, read the official [Istio 1.18.0 release notes](https://istio.io/latest/news/releases/1.18.x/announcing-1.18/upgrade-notes/). 

## API Gateway
### Removal of APIRule in version `v1alpha1`
APIRule in version `v1alpa1` has been deprecated since Kyma 2.5, and its end of life is planned for Kyma 2.19. Migrate your APIRules to `v1beta1`.

## Eventing
The NATS server has been updated to version 2.9.18.

## Serverless
### Deprecation of Node.js 16 Serverless runtime
Because of the [scheduled EOL for Node.js 16](https://github.com/nodejs/release#release-schedule), we plan to remove Node.js 16 from the list of supported runtimes.

For now, we recommend that you don’t use Node.js 16 as a runtime for your new Functions and re-configure all your existing Node.js 16 Functions to run on the latest available Node.js runtime.

Read this [blog post](https://blogs.sap.com/2022/03/09/changing-the-function-runtime-version-of-a-running-function/) to learn how to update existing Functions.

### Node.js 14 runtime removed
Node.js 14 has reached the end of its life. Therefore, followed by the depreciation of Node.js 14 Serverless runtime, we decided to finally remove it from the list of the available Function runtimes.

Your Node.js 14 Functions' workloads will continue to run, but you will not be able to edit them without changing the `runtime` field. Sooner or later, you must change the **spec** of your existing Node.js 14-based Functions and change the `runtime` field to either `nodejs16` or `nodejs18`.

For more information about the Node.js 14 deprecation, see the [Kyma 2.12 release notes](https://github.com/kyma-project/kyma/releases/tag/2.12.0).
