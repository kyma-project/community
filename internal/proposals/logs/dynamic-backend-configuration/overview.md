# Motivation
Kyma uses Fluent Bit as a log collection solution. Fluent Bit is configured with a plain text file (or multiple files) and does not support dynamic configuration.
It poses the following problems:

1. Static configuration should be provided upon Kyma installation and cannot be changed afterward.
2. It's impossible to separate the Kyma system configuration from the customer one.
3. Every time we change the system configuration in managed Kyma, we have to touch Kyma Environment Broker's code.

# Goal
Come up with a way to dynamically configure Fluent Bit (no manual pod restarting required). It should be possible to apply additional configuration snippets, which are not resettable by the Kyma upgrade process (reconciliation).

# Possible solutions

What immediately comes to mind is Prometheus Operator, which solves very similar problems. 
Prometheus, same as Fluent Bit, is configured with a plain text file (possibly, very long and complex). Configuring Prometheus by hand is not a trivial task.
Prometheus Operator simplifies the configuration by introducing a bunch of CRs, which are then translated into a Prometheus config.
It also makes it possible to dynamically update different parts of a Prometheus config without restarting any pods.

After some investigation, the following open source solutions have been found:

- Fluent Bit Operator from Kubesphere: 
- Logging Operator from Banzai Cloud: https://github.com/banzaicloud/logging-operator

## Fluent Bit operator from Kubesphere

## Logging Operator from Banzai Cloud


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
