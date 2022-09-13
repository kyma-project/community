# API Design


```YAML
kind: MetricPipeline # cluster scope
apiVersion: telemetry.kyma-project.io/v1alpha1
metadata:
  name: PrometheusRemoteWrite
spec:
  receiver: # singular, different receivers need different pipelines and exporters configs
    type: otlp # otlp  | custom, default is otlp
    otlp: # maps to central otlp receiver, dealing with the actual application logs
      namespaces: [] # generates the rule for the rewrite_tag assigned to every pipeline
      excludeNamespaces: []
      containers: []
      excludeContainers: [] # allows to exclude istio proxies

      podLabels: # generates a "filterprocessor" filter as a first element in the chain selecting logs by the "kubernetes" attributes
        app: icke
      excludePodLabels:
        app: chris

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

  exporter: # only one output, defining no output will fail validation
    prometheusremotewrite:
      endpoint: "https://my-cortex:7900/api/v1/push"
      external_labels:
        label_name1: label_value1
        label_name2: label_value2
      tls:
          insecure: false
          insecureSkipVerify: true
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
      headers:
          - name: x-token
            value: "value1"
            valueFrom:
              secretKeyRef:
                  name: my-config
                  namespace: default
                  key: "myToken"
```