# Motivation
The fluent-bit deployment in kyma gets reconciled as part of the upgrade process. That reconcilation will override any configuration done directly in the cluster.

# Goal
Come up and evaluate a way to have a dynamic in-cluster configuration which will be not resetted by the upgrade process and gets applied automatically (no manual pod restarts required).

# Solutions

## Fluent Bit operator from Kubesphere

https://github.com/kubesphere/fluentbit-operator


## Features

### Dynamic config reloading

> Note that the operator works with kubesphere/fluent-bit, a fork of fluent/fluent-bit. Due to the known issue, the original Fluent Bit doesn't support dynamic configuration. To address that, kubesphere/fluent-bit incorporates a configuration reloader into the original. See kubesphere/fluent-bit documentation for more information.

- [x] Config snippets provided as CRDs: https://github.com/kubesphere/fluentbit-operator#overview
- [x] Dynamic reload
- [ ] Full fluent-bit syntax is supported
- [ ] Debugging a startup problem - how to figure out which config snippet caused a problem at startup
- [ ] Have in mind that the central configuration will be overwritten at any time by reconcilation which should not reset customers config
- [ ] A rollback to the old config might be cool in case of invalid config provided by a customer
- [x] A new config gets picked up dynamically without any further interaction (like pod restart)
- [x] Basic validation to give early feedback on errors
- [-] There is a way to provide the auth details for a backend config in a secure way


## Logging operator from Banzai Cloud
