# Helm chart customizations

## Telemetry

### Fluent Bit

* Standardized priority class name: https://github.com/kyma-project/kyma/blob/main/resources/telemetry/charts/fluent-bit/templates/_pod.tpl#L8
* Standardized image name: https://github.com/kyma-project/kyma/blob/main/resources/telemetry/charts/fluent-bit/templates/_pod.tpl#L38
* Standardized image name: https://github.com/kyma-project/kyma/blob/main/resources/telemetry/charts/fluent-bit/templates/tests/test-connection.yaml#L12
* Dynamic config sections: https://github.com/kyma-project/kyma/blob/main/resources/telemetry/charts/fluent-bit/templates/configmap.yaml#L12
* Capability check removed: https://github.com/kyma-project/kyma/blob/main/resources/telemetry/charts/fluent-bit/templates/servicemonitor.yaml#L1
* Custom dashboard name: https://github.com/kyma-project/kyma/blob/main/resources/telemetry/charts/fluent-bit/dashboards/fluent-bit.json#L1302

### Prometheus & Prometheus Operator
#### /monitoring/

| resource | modification description | prometheus-community | kyma-project |
| -------  | ------------------------ | -------------------- | ------------ |
| charts/prometheus-istio/templates/deploy.yaml | Different path for monitoring CRDs | /crds/ | /installation/resources/crds/monitoring/ | 
| charts/prometheus-istio/templates/deploy.yaml | Different names for monitoring CRDS | crd-*.yaml | *.monitoring.coreos.crd.yaml | 
| charts/prometheus-istio/templates/deploy.yaml | priorityClassName considers global priorityClass | [L#36](https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/templates/server/deploy.yaml#L36) | [L#43](https://github.com/kyma-project/kyma/blob/main/resources/monitoring/charts/prometheus-istio/templates/deploy.yaml#L43) | 
| charts/prometheus-istio/templates/deploy.yaml | images are referenced from global values | [L#57](https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/templates/server/deploy.yaml#L57) |  | 
| charts/prometheus-istio/templates/deploy.yaml | images are referenced from global values | [L#87](https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/templates/server/deploy.yaml#L87) |  | 
| charts/prometheus-istio/templates/deploy.yaml | sidecar has containerSecurityContext | n/a |  | 
| charts/prometheus-istio/templates/sts.yaml | priorityClassName considers global priorityClass | [L#36](https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/templates/server/sts.yaml#L36) |  | 
| charts/prometheus-istio/templates/sts.yaml | images are referenced from global values | [L#57](https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/templates/server/sts.yaml#L57) |  | 
| charts/prometheus-istio/templates/sts.yaml | images are referenced from global values | [L#87](https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/templates/server/sts.yaml#L87) |  | 
| charts/prometheus-istio/templates/sts.yaml | main container and sidecar have containerSecurityContext | n/a |  | 
| templates/grafana/dashboards-1.14/cluster-total.yaml | dashboard not editable | [L#43](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/cluster-total.yaml#L43) |  | 
| templates/_helpers.tpl | define imageurl | n/a | [L#207](https://github.com/kyma-project/kyma/blob/main/resources/monitoring/templates/_helpers.tpl#L207) | 
| templates/grafana/dashboards-1.14/etcd.yaml | dashboard not editable | [L#28](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/etcd.yaml#L28) |  | 
| templates/grafana/dashboards-1.14/grafana-overview.yaml | dashboard disabled | n/a |  | 
| templates/grafana/dashboards-1.14/k8s-coredns.yaml | dashboard not editable | [L#34](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-coredns.yaml#L34) |  | 
| templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml | dashboard not editable | [L#29](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml#L29) |  | 
| templates/grafana/dashboards-1.14/k8s-resources-namespace.yaml | dashboard not editable | [L#29](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-namespace.yaml#L29) |  | 
| templates/grafana/dashboards-1.14/k8s-resources-node.yaml | dashboard not editable | [L#29](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-node.yaml#L29) |  | 
| templates/grafana/dashboards-1.14/k8s-resources-pod.yaml | dashboard not editable | [L#29](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-pod.yaml#L29) |  | 
| templates/grafana/dashboards-1.14/k8s-resources-workload.yaml | dashboard not editable | [L#29](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-workload.yaml#L29) |  | 
| templates/grafana/dashboards-1.14/k8s-resources-workloads-namespace.yaml | dashboard not editable | [L#29](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-workloads-namespace.yaml#L29) |  | 
| templates/grafana/dashboards-1.14/namespace-by-pod.yaml | dashboard not editable | [L#43](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/namespace-by-pod.yaml#L43) |  | 
| templates/grafana/dashboards-1.14/namespace-by-workload.yaml | dashboard not editable | [L#43](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/namespace-by-workload.yaml#L43) |  | 
| templates/grafana/dashboards-1.14/pod-total.yaml | dashboard not editable | [L#43](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/pod-total.yaml#L43) |  | 
| templates/grafana/dashboards-1.14/prometheus-remote-write.yaml | dashboard not editable | [L#35](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/prometheus-remote-write.yaml#L35) |  | 
| templates/grafana/dashboards-1.14/workload-total.yaml | dashboard not editable | [L#43](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/workload-total.yaml#L43) |  | 
| templates/grafana/dashboards-1.14/prometheus.yaml | dashboard not editable | [L#29](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/prometheus.yaml#L29) |  | 
| templates/grafana/dashboards-1.14/prometheus.yaml | remove hard-coded monitoring namespace in label_values | [L#1159](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/prometheus.yaml#L1159) |  | 
| templates/prometheus/rules-1.14/kubernetes-system-kubelet.yaml | "alert KubeletPodStartUpLatencyHigh: group expression uses additional ""service"" label" | [L#110](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/prometheus/rules-1.14/kubernetes-system-kubelet.yaml#L110) |  | 
| templates/prometheus/rules-1.14/etcd.yaml | "renamed annotation ""message"" key to ""description""" | [L#30](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/prometheus/rules-1.14/etcd.yaml#L30) |  | 
| templates/prometheus/rules-1.14/etcd.yaml | "renamed annotation ""message"" key to ""description""" | [L#42](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/prometheus/rules-1.14/etcd.yaml#L42) |  | 
| templates/prometheus/rules-1.14/etcd.yaml | "renamed annotation ""message"" key to ""description""" | [L#54](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/prometheus/rules-1.14/etcd.yaml#L54) |  | 
| templates/prometheus/rules-1.14/etcd.yaml | "renamed annotation ""message"" key to ""description""" | [L#66](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/prometheus/rules-1.14/etcd.yaml#L66) |  | 
| templates/prometheus/rules-1.14/etcd.yaml | "renamed annotation ""message"" key to ""description""" | [L#82](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/prometheus/rules-1.14/etcd.yaml#L82) |  | 
| templates/prometheus/rules-1.14/etcd.yaml | "renamed annotation ""message"" key to ""description""" | [L#98](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/prometheus/rules-1.14/etcd.yaml#L98) |  | 
| templates/prometheus/rules-1.14/etcd.yaml | "renamed annotation ""message"" key to ""description""" | [L#112](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/prometheus/rules-1.14/etcd.yaml#L112) |  | 
| templates/prometheus/rules-1.14/etcd.yaml | "renamed annotation ""message"" key to ""description""" | [L#127](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/prometheus/rules-1.14/etcd.yaml#L127) |  | 
| templates/prometheus/rules-1.14/etcd.yaml | "renamed annotation ""message"" key to ""description""" | [L#138](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/prometheus/rules-1.14/etcd.yaml#L138) |  | 
| templates/prometheus/rules-1.14/etcd.yaml | "renamed annotation ""message"" key to ""description""" | [L#152](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/prometheus/rules-1.14/etcd.yaml#L152) |  | 
| templates/prometheus/rules-1.14/etcd.yaml | "renamed annotation ""message"" key to ""description""" | [L#166](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/prometheus/rules-1.14/etcd.yaml#L166) |  | 
| templates/prometheus/rules-1.14/etcd.yaml | "renamed annotation ""message"" key to ""description""" | [L#180](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/prometheus/rules-1.14/etcd.yaml#L180) |  | 
| templates/prometheus/rules-1.14/etcd.yaml | "renamed annotation ""message"" key to ""description""" | [L#194](https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/prometheus/rules-1.14/etcd.yaml#L194) |  | 
| templates/grafana/dashboards-1.14/alertmanager-overview.yaml | enabled flag | n/a |  | 
| templates/alertmanager/alertmanager.yaml | images are referenced from global values | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/alertmanager/alertmanager.yaml#L15 |  | 
| templates/grafana/dashboards-1.14/apiserver.yaml | dashboard disabled | n/a |  | 
| templates/prometheus/prometheus.yaml | images are referenced from global values | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/prometheus/prometheus.yaml#L34 |  | 
| templates/prometheus/prometheus.yaml | priorityClassName considers global priorityClass | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/prometheus/prometheus.yaml#L297 |  | 
| charts/prometheus-istio/Chart.yaml | deleted home, icon, sources, maintainers, engine, type, dependencies |  |  | 
| charts/prometheus-istio/Chart.yaml | renamed: prometheus-istio |  |  | 
| templates/prometheus-operator/deployment.yaml | priorityClassName considers global priorityClass | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/prometheus-operator/deployment.yaml#L31 |  | 
| templates/prometheus-operator/deployment.yaml | images are referenced from global values (operator) | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/prometheus-operator/deployment.yaml#L37 |  | 
| templates/prometheus-operator/deployment.yaml | images are referenced from global values (configmap reloader) | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/prometheus-operator/deployment.yaml#L69 |  | 
| templates/prometheus-operator/deployment.yaml | added liveness/readiness probes | n/a |  | 
| charts/prometheus-istio/values.yaml | serviceAccounts.nodeExporter.create: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/values.yaml#L14 |  | 
| charts/prometheus-istio/values.yaml | serviceAccounts.alertmanager.create: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/values.yaml#L18 |  | 
| charts/prometheus-istio/values.yaml | serviceAccounts.pushgateway.create: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/values.yaml#L22 |  | 
| charts/prometheus-istio/values.yaml | alertmanager.enabled: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/values.yaml#L33 |  | 
| charts/prometheus-istio/values.yaml | kubeStateMetrics.enabled: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/values.yaml#L464 |  | 
| charts/prometheus-istio/values.yaml | nodeExporter.enabled: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/values.yaml#L474 |  | 
| charts/prometheus-istio/values.yaml | server.enableServiceLinks: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/values.yaml#L684 |  | 
| charts/prometheus-istio/values.yaml | server.global.scrape_interval: 30s | https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/values.yaml#L739 |  | 
| charts/prometheus-istio/values.yaml | server.global.evaluation_interval: 30s | https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/values.yaml#L745 |  | 
| charts/prometheus-istio/values.yaml | server.persistentVolume.enabled: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/values.yaml#L895 |  | 
| charts/prometheus-istio/values.yaml | "server.podAnnotations.sidecar.istio.io/inject: ""false""" | n/a |  | 
| charts/prometheus-istio/values.yaml | server.readinessProbeTimeout: 30 | https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/values.yaml#L1021 |  | 
| charts/prometheus-istio/values.yaml | server.livenessProbeTimeout: 30 | https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/values.yaml#L1026 |  | 
| charts/prometheus-istio/values.yaml | server.resources.limits.cpu: 1000m | n/a |  | 
| charts/prometheus-istio/values.yaml | server.resources.limits.memory: 200Mi | n/a |  | 
| charts/prometheus-istio/values.yaml | server.resources.requests.cpu: 40m | n/a |  | 
| charts/prometheus-istio/values.yaml | server.resources.requests.memory: 200Mi | n/a |  | 
| charts/prometheus-istio/values.yaml | server.containerSecurityContext.allowPrivilegeEscalation: false | n/a |  | 
| charts/prometheus-istio/values.yaml | server.containerSecurityContext.privileged: false | n/a |  | 
| charts/prometheus-istio/values.yaml | "server.retention: ""2h""" | https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/values.yaml#L1124 |  | 
| charts/prometheus-istio/values.yaml | pushgateway.enabled: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/values.yaml#L1129 |  | 
| charts/prometheus-istio/values.yaml | serverFiles.recording_rules.yml | https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/values.yaml#L1405 |  | 
| charts/prometheus-istio/values.yaml | extraScrapeConfigs | n/a |  | 
| charts/prometheus-istio/values.yaml | "envoyStats.sampleLimit: ""20000""" | n/a |  | 
| charts/prometheus-istio/values.yaml | "envoyStats.metricKeepRegexp: ""istio_.*""" | n/a |  | 
| charts/prometheus-istio/values.yaml | "envoyStats.labeldropRegex: ""^(grpc_response_status|source_version|destination_version|source_app|destination_app)$""" | n/a |  | 
| charts/prometheus-istio/values.yaml | extraScrapeConfigs | n/a |  | 
| charts/prometheus-istio/values.yaml | grafana.datasource.enabled: false | n/a |  | 
| charts/prometheus-istio/values.yaml | deleted example scrape configs L1426 - L1808 | https://github.com/prometheus-community/helm-charts/blob/main/charts/prometheus/values.yaml#L1426 |  | 
| values.yaml | defaultRules.rules.etcd: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L37 |  | 
| values.yaml | defaultRules.rules.kubeApiserverBurnrate: false | n/a |  | 
| values.yaml | defaultRules.rules.kubeApiserverHistogram: false | n/a |  | 
| values.yaml | defaultRules.rules.kubeApiserverSlos: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L43 |  | 
| values.yaml | defaultRules.disabled.CPUThrottlingHigh: true | n/a |  | 
| values.yaml | additionalPrometheusRules: |- | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L82 |  | 
| values.yaml | global.domainName | n/a |  | 
| values.yaml | global.containerRegistry | n/a |  | 
| values.yaml | global.images | n/a |  | 
| values.yaml | global.istio | n/a |  | 
| values.yaml | global.rbac.pspEnabled: true | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L108 |  | 
| values.yaml | "global.highPriorityClassName: ""kyma-system-priority""" | n/a |  | 
| values.yaml | global.alertTools | n/a |  | 
| values.yaml | alertmanager.config is multiline string | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L164 |  | 
| values.yaml | alertmanager.tplConfig: true | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L190 |  | 
| values.yaml | alertmanager.serviceMonitor.scheme: http | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L381 |  | 
| values.yaml | alertmanager.serviceMonitor.metricRelabelings | n/a |  | 
| values.yaml | alertmanager.alertmanagerSpec.resources | n/a |  | 
| values.yaml | "alertmanager.alertmanagerSpec.podMetadata.annotations.sidecar.istio.io/inject: ""false""" | n/a |  | 
| values.yaml | "grafana.defaultDashboardsTimezone: """"" | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L668 |  | 
| values.yaml | grafana.rbac.pspEnabled: true | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L675 |  | 
| values.yaml | grafana.sidecar.dashboards.searchNamespace: ALL | n/a |  | 
| values.yaml | grafana.sidecar.datasources.searchNamespace: ALL | n/a |  | 
| values.yaml |  |  |  | 
| values.yaml | kubeApiServer.serviceMonitor.metricRelabelings | n/a |  | 
| values.yaml | kubelet.serviceMonitor.cAdvisorMetricRelabelings | n/a |  | 
| values.yaml | kubelet.serviceMonitor.metricRelabelings | n/a |  | 
| values.yaml | kubeControllerManager.enabled: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L999 |  | 
| values.yaml | coreDns.enabled: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L1064 |  | 
| values.yaml | kubeEtcd.enabled: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L1161 |  | 
| values.yaml | kubeScheduler.enabled: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L1228 |  | 
| values.yaml | kubeProxy.enabled: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L1292 |  | 
| values.yaml | kube-state-metrics.prometheus.monitor.metricRelabelings | n/a |  | 
| values.yaml | prometheus-node-exporter.prometheus.monitor.metricRelabelings | n/a |  | 
| values.yaml | prometheus-node-exporter.rbac.pspEnabled: true | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L1453 |  | 
| values.yaml | prometheusOperator.tls.enabled: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L1463 |  | 
| values.yaml | prometheusOperator.admissionWebhooks.enabled: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L1473 |  | 
| values.yaml | "prometheusOperator.admissionWebhooks.patch.podAnnotations.sidecar.istio.io/inject: ""false""" | n/a |  | 
| values.yaml | "prometheusOperator.podAnnotations.sidecar.istio.io/inject: ""false""" | n/a |  | 
| values.yaml | prometheusOperator.serviceMonitor.metricRelabelings | n/a |  | 
| values.yaml | prometheusOperator.resources | n/a |  | 
| values.yaml | prometheusOperator.containerSecurityContext.privileged: false | n/a |  | 
| values.yaml | prometheusOperator.livenessProbe | n/a |  | 
| values.yaml | prometheusOperator.readinessProbe | n/a |  | 
| values.yaml | prometheusOperator.prometheusConfigReloader.resources.requests.cpu: 100m | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L1735 |  | 
| values.yaml | prometheusOperator.prometheusConfigReloader.resources.limits.cpu: 100m | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L1738 |  | 
| values.yaml | prometheus.serviceMonitor.metricRelabelings | n/a |  | 
| values.yaml | prometheus.prometheusSpec.scrapeInterval: 30s | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L2130 |  | 
| values.yaml | prometheus.prometheusSpec.evaluationInterval: 30s | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L2138 |  | 
| values.yaml | prometheus.prometheusSpec.query.maxConcurrency: 100 | n/a |  | 
| values.yaml | prometheus.prometheusSpec.query.timeout: 30s | n/a |  | 
| values.yaml | prometheus.prometheusSpec.ruleSelectorNilUsesHelmValues: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L2257 |  | 
| values.yaml | prometheus.prometheusSpec.serviceMonitorSelectorNilUsesHelmValues: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L2282 |  | 
| values.yaml | prometheus.prometheusSpec.podMonitorSelectorNilUsesHelmValues: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L2305 |  | 
| values.yaml | prometheus.prometheusSpec.probeSelectorNilUsesHelmValues: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L2325 |  | 
| values.yaml | prometheus.prometheusSpec.retention: 1d | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L2343 |  | 
| values.yaml | prometheus.prometheusSpec.retentionSize: 2GB | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L2347 |  | 
| values.yaml | prometheus.prometheusSpec.walCompression: true | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L2351 |  | 
| values.yaml | prometheus.prometheusSpec.podMetadata | n/a |  | 
| values.yaml | prometheus.prometheusSpec.resources | n/a |  | 
| values.yaml | prometheus.prometheusSpec.storageSpec.volumeClaimTemplate | n/a |  | 
| values.yaml | prometheus.prometheusSpec.volumes | n/a |  | 
| values.yaml | prometheus.prometheusSpec.volumeMounts | n/a |  | 
| values.yaml | pushgateway.enabled: false | n/a |  | 
| values.yaml | prometheus-istio.enabled: true | n/a |  | 
| values.yaml | grafana.serviceMonitor.enabled: false | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/values.yaml#L791 |  | 
| values.yaml | grafana.serviceMonitor.selfMonitor: true   (overrides grafana charts' values) | n/a |  | 
| values.yaml | grafana.serviceMonitor.interval: 1m   (overrides grafana charts' values) | n/a |  | 
| values.yaml | grafana.serviceMonitor.scheme: https    (overrides grafana charts' values) | n/a |  | 
| values.yaml | grafana.serviceMonitor.tlsConfig     (overrides grafana charts' values) | n/a |  | 
| values.yaml | grafana.serviceMonitor.metricRelabelings | n/a |  | 
| templates/grafana/servicemonitor.yaml | new resource overriding the upstream one in grafana templates | n/a |  | 
| templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml | Delete device label selector | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml#L2548 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml | Delete device label selector | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml#L2637 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml | Delete device label selector | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml#L2898 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml | Delete device label selector | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml#L2907 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml | Delete device label selector (2 times) | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml#L2916 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml | Delete device label selector | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml#L2925 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml | Delete device label selector | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml#L2934 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml | Delete device label selector (2 times) | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml#L2943 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml | Delete device label selector | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml#L2928 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml | Delete device label selector | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml#L2929 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml | Delete device label selector | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-cluster.yaml#L2930 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-namespace.yaml | Delete device label selector (2 times) | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-namespace.yaml#L2230 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-namespace.yaml | Delete device label selector (2 times) | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-namespace.yaml#L2319 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-namespace.yaml | Delete device label selector | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-namespace.yaml#L2580 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-namespace.yaml | Delete device label selector | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-namespace.yaml#L2589 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-namespace.yaml | Delete device label selector (2 times) | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-namespace.yaml#L2598 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-namespace.yaml | Delete device label selector | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-namespace.yaml#L2607 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-namespace.yaml | Delete device label selector | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-namespace.yaml#L2616 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-namespace.yaml | Delete device label selector (2 times) | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-namespace.yaml#L2625 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-pod.yaml | Delete device label selector | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-pod.yaml#L1668 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-pod.yaml | Delete device label selector | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-pod.yaml#L1676 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-pod.yaml | Delete device label selector | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-pod.yaml#L1765 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-pod.yaml | Delete device label selector | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-pod.yaml#L1773 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-pod.yaml | Delete device label selector | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-pod.yaml#L2225 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-pod.yaml | Delete device label selector | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-pod.yaml#L2234 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-pod.yaml | Delete device label selector (2 times) | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-pod.yaml#L2243 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-pod.yaml | Delete device label selector | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-pod.yaml#L2252 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-pod.yaml | Delete device label selector | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-pod.yaml#L2261 |  | 
| templates/grafana/dashboards-1.14/k8s-resources-pod.yaml | Delete device label selector (2 times) | https://github.com/prometheus-community/helm-charts/blob/main/charts/kube-prometheus-stack/templates/grafana/dashboards-1.14/k8s-resources-pod.yaml#L2270 |  | 