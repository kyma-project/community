# Helm chart customizations

## Telemetry

### Fluent Bit

* Standardized priority class name: https://github.com/kyma-project/kyma/blob/main/resources/telemetry/charts/fluent-bit/templates/_pod.tpl#L8
* Standardized image name: https://github.com/kyma-project/kyma/blob/main/resources/telemetry/charts/fluent-bit/templates/_pod.tpl#L39
* Dynamic config sections: https://github.com/kyma-project/kyma/blob/main/resources/telemetry/charts/fluent-bit/templates/configmap.yaml#L12
* Capability check removed: https://github.com/kyma-project/kyma/blob/main/resources/telemetry/charts/fluent-bit/templates/servicemonitor.yaml#L1
* Standardized image name: https://github.com/kyma-project/kyma/blob/main/resources/telemetry/charts/fluent-bit/templates/tests/test-connection.yaml#L12
