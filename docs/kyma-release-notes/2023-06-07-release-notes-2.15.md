Still warm, as the June sun at noon, comes Kyma 2.15. The hot season has already started, so juicy improvements are here to refresh your experience with our product. Sip the strawberry-Telementry, watermelon-Security, raspberry-API Gateway, and peach-Serverless cool punch news as if you were on a Hawaiian beach. Read on to see what we prepared for you!

## Telemetry
### Manager
- In [preparation for turning the telemetry component into a module](https://github.com/kyma-project/telemetry-manager/issues/150), resources have been consolidated. As a result, you must run a cleanup script when you upgrade to Kyma version 2.15. For more details, read the [2.14-2.15 Migration Guide](https://github.com/kyma-project/kyma/blob/main/docs/migration-guide-2.14-2.15.md).
- Handling of webhook certificates (https://github.com/kyma-project/kyma/issues/16626) has been improved.

### Tracing
- Updated OTel Collector to 0.77.0 (https://github.com/kyma-project/kyma/pull/17469).

### Logging
- Updated Fluent Bit to 2.1.2 (https://github.com/kyma-project/kyma/pull/17485).
- Improved security setup of Fluent Bit (https://github.com/kyma-project/kyma/pull/17574)

## Security
As a part of security hardening and following Kyma security team recommendations, ECDHE-RSA-AES256-SHA and ECDHE-RSA-AES128-SHA cipher suites used in the default Kyma gateway become deprecated starting from the 2.15 Kyma version. With Kyma version 2.18, the configurations will be removed and clients dependent on the cipher suites won’t be accepted.

## API Gateway
This Kyma release brings unified timeout for workloads exposed with APIRules. The default timeout for HTTP requests is 180s and it’s defined on the Istio VirtualService level.

## Serverless
### Simplified internal Docker registry setup
With Kyma 2.15 we simplified Serverless configuration for the internal Docker registry. From now on, the images for Function runtime Pods are pulled from the internal Docker registry via NodePort.

With this change, we improve security as the internal Docker registry is no longer exposed outside of the Kubernetes cluster. Additionally, it makes Serverless fully independent from the Istio module in all installation modes.

### Deployment profiles removed
In preparation for an independent installation model, we removed the predefined deployment profiles, namely evaluation, and production, for Serverless. We are shifting from profiled overrides used during module installation towards runtime-configurable resources.
