#  Installation Working Group Closure

The group achieved its goal. Several improvements and design changes were discussed and applied in the Kyma installation area. The key outcomes (decisions):
- Kyma installation profiles were introduced (evaluation with limited resource consumption and production for better scalability and availability)
- CLI will be recommended for the Kyma installation task and the Kyma Operator (installer) will be deprecated
- A new installation library (used in Kyma CLI) will replace the current installer code and will be used also to provision managed Kyma
- Parallel installation of modules will be enabled by default
- Better handling of Custom Resource Definitions (CRDs) — not managed by Helm
- Default local setup based on k3s (much faster than Minikube)

The decisions are now in the implementation phase, but some of the ideas are already applied and bring visible benefits. The best example is the new [integration pipeline](https://status.build.kyma-project.io/job-history/kyma-prow-logs/logs/kyma-integration-k3s) that installs Kyma and executes integration tests in about 7 minutes.
