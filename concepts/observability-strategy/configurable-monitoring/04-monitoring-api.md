# API Design


```YAML
kind: MetricPipeline # cluster scope
apiVersion: telemetry.kyma-project.io/v1alpha1
metadata:
  name: PrometheusRemoteWrite
spec:
  input: # Metrics can be received from different receivers (OTLP and Prometheus). Metrics are selected based on labels
    application:
      namespaces: # maps to filterprocessor
        include: []
        exclude: []
        system: true
      containers:
        include: []
        exclude: []
      system:
        kubernetes:
          enabled: true
          deployments: false
          daemonSets: true
          persistentVolumes: true
          resources: true
        runtime:
          enabled: true
        istio: # Filter by `destination_service_namespace` or `destination_workload_namespace` if namespaces are limited
          controlPlane:
            enabled: true
          proxies:
            enabled: true
    infrastructure: # Metrics that are not related to any Kubernetes workload (for instance, node or control-plane specific)
      apiserver:
        enabled: true
      nodes:
        enabled: true
      network:
        enabled: true
    labels: # Filter based on custom labels
      include:
        - name: deployment
          value: my-app
        - name: landscape
          value: production
      exclude: []

  processors: # list of processors, order is important
    - addLabel: # maps to "metricstransformprocessor" filter, adds an attribute
          match_type: regexp
          metric_names:
                - prefix/.*
                - prefix_.*
          key: cluster_identifier
          value: icke's cluster
    - removeLabel: # maps to "metricstransformprocessor" filter, removes an attribute
          match_type: regexp
          metric_names:
                - prefix/.*
                - prefix_.*
          key: cluster_identifier
    - includeMetrics: # maps to "filterprocessor" filter, drops metrics where attribute will not match the regexp
        match_type: regexp
        metric_names:
              - prefix/.*
              - prefix_.*
        resource_attributes:
          - Key: container.name
            Value: app_container_1
    - excludeMetrics: # maps to "filterprocessor" filter, drops lines where attribute will not match the regexp
        match_type: strict
        metric_names:
          - hello_world
          - hello/world

  output: # Only one output. Defining no output will fail validation
    prometheusremotewrite:
      endpoint:
          value: "https://my-cortex:7900/api/v1/push"
          valueFrom:
            secretKeyRef:
                name: my-config
                namespace: default
                key: "endpoint"
      external_labels:
        label_name1: label_value1
        label_name2: label_value2
      tls:
        insecure: false
        insecureSkipVerify: true
        ca:
          value: dummy
          valueFrom:
              secretKeyRef:
                name: my-config
                namespace: default
                key: "server.crt"
        cert:
            value: dummy
            valueFrom:
                secretKeyRef:
                    name: my-config
                    namespace: default
                    key: "cert.crt"
        key:
            value: dummy
            valueFrom:
                secretKeyRef:
                    name: my-config
                    namespace: default
                    key: "client.key"
    dynatrace:
      endpoint:
        value: "https://ab12345.live.dynatrace.com"
        valueFrom:
          secretKeyRef:
              name: my-config
              namespace: default
              key: "endpoint"
      apiToken:
        value: "insecure-token"
        valueFrom:
          secretKeyRef:
              name: my-config
              namespace: default
              key: "token"
    otlp:
      protocol: grpc #grpc | http
      endpoint: myserver.local:55690
      tls:
          insecure: false
          insecureSkipVerify: true
          ca:
            value: dummy
            valueFrom:
                secretKeyRef:
                  name: my-config
                  namespace: default
                  key: "server.crt"
          cert:
              value: dummy
              valueFrom:
                  secretKeyRef:
                      name: my-config
                      namespace: default
                      key: "cert.crt"
          key:
              value: dummy
              valueFrom:
                  secretKeyRef:
                      name: my-config
                      namespace: default
                      key: "client.key"
      authentication:
        basic:
          username:
            value: "user"
            valueFrom:
              secretKeyRef:
                  name: my-config
                  namespace: default
                  key: "username"
          password:
            value: "my-insecure-password"
            valueFrom:
              secretKeyRef:
                  name: my-config
                  namespace: default
                  key: "password"
        bearerToken:
          value: "insecure-token"
          valueFrom:
            secretKeyRef:
                name: my-config
                namespace: default
                key: "token"
      headers:
          - name: x-token
            value: "value1"
            valueFrom:
              secretKeyRef:
                  name: my-config
                  namespace: default
                  key: "myToken"
```
