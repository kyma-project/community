# API Design


```YAML
kind: MetricPipeline # cluster scope
apiVersion: telemetry.kyma-project.io/v1alpha1
metadata:
  name: PrometheusRemoteWrite
spec:
  receiver: # singular, different receivers need different pipelines and exporterssss configs
    type: application # application | system | custom, default is application
    otlp: # maps to central tail pipeline, dealing with the actual application logs
      namespaces: [] # generates the rule for the rewrite_tag assigned to every pipeline
      excludeNamespaces: []
      containers: []
      excludeContainers: []

      podLabels: # generates a "filterprocessor" filter as a first element in the chain selecting logs by the "kubernetes" attributes
        app: icke
      excludePodLabels:
        app: chris
      
    system: {} # maps to systemd based input like kubelet logs; for now, no further spec designed
    custom: # define a custom input, entering unsupported mode
        cloudfoundry:
            rlp_gateway:
            endpoint: "https://log-stream.sys.example.internal"
  processors: # list of processors, order is important
    - add: # maps to "metricstransformprocessor" filter, adds a log attribute
          key: cluster_identifier
          value: icke's cluster
    - remove: # maps to "metricstransformprocessor" filter, removes a log attribute
          key: cluster_identifier
    - include: # maps to "filterprocessor" filter, drops lines where attribute will not match the regexp
        key: tenant
        regexp: icke
    - exclude: # maps to "filterprocessor" filter, drops lines where attribute will not match the regexp
        key: severity
        regexp: debug
    - custom: # entering unsupported mode
        metricstransform:
            transforms:
            - include: host.cpu.usage
              action: insert
              new_name: host.cpu.utilization

  exporter: # only one output, defining no output will fail validation
    prometheusremotewrite: # enables the exporter
      endpoint: "http://monitoring-prometheus.kyma-system.svc.cluster.local:9090/api/v1/write"
      tls:
        insecure: true
    custom: # no filebuffer settings available, entering unsupported mode
      elasticsearch:
        endpoints:
        - "https://localhost:9200"

  variables: # env variables to be used in custom plugins
    - name: myEnv1 # static mapping
      value: myValue1
    - name: myEnv2
      valueFrom:
        secretKeyRef:
          name: my-elastic-credentials
          namespace: default
          key: ES_ENDPOINT
    - name: myEnv3
      valueFrom:
        configMapKeyRef:
          name: my-elastic-credentials
          namespace: default
          key: ES_ENDPOINT

  files: # files to be used in custom plugins
    - name: "labelmap1.json"
      content: |
        {
          "kubernetes": {
                "namespace_name": "namespace",
                "pod_name": "pod"
          },
          "stream": "stream"
        }
    - name: "labelmap2.json"
      contentFrom:
        configMapKeyRef:
          name: my-config
          namespace: default
          key: "label.json"
    - name: "cert.crt"
      contentFrom:
        secretKeyRef:
          name: my-certificates
          namespace: default
          key: "cert.crt"
```