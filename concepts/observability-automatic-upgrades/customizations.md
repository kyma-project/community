# Helm chart customizations

## Telemetry

### Fluent Bit

* Standardized priority class name: https://github.com/kyma-project/kyma/blob/main/resources/telemetry/charts/fluent-bit/templates/_pod.tpl#L8
* Standardized image name: https://github.com/kyma-project/kyma/blob/main/resources/telemetry/charts/fluent-bit/templates/_pod.tpl#L39
* Standardized image name: https://github.com/kyma-project/kyma/blob/main/resources/telemetry/charts/fluent-bit/templates/tests/test-connection.yaml#L12
* Dynamic config sections: https://github.com/kyma-project/kyma/blob/main/resources/telemetry/charts/fluent-bit/templates/configmap.yaml#L12
* Capability check removed: https://github.com/kyma-project/kyma/blob/main/resources/telemetry/charts/fluent-bit/templates/servicemonitor.yaml#L1
* Custom dashboard name: https://github.com/kyma-project/kyma/blob/main/resources/telemetry/charts/fluent-bit/dashboards/fluent-bit.json#L1302
