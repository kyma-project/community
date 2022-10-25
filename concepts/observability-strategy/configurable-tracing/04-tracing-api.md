# API Design


```YAML
kind: TelemetryModule # cluster scope
apiVersion: telemetry.kyma-project.io/v1alpha1
metadata:
  name: myModule
spec:
  tracing:
    enabled: true
    otlp:
      endpoint: otlp-traces.kyma-system:55690
```

```YAML
kind: TracePipeline # cluster scope
apiVersion: telemetry.kyma-project.io/v1alpha1
metadata:
  name: myPipeline
spec:
  input:
    istio: # Would configure Istio to send traces to the collector service
      propagation: w3c # w3c | b3
      enabled: true
      samplingPercentage: 100

  processors: # list of processors, order is important
    probabilisticSampler:
        samplingPercentage: 15.3

  output: # Only one output. Defining no output will fail validation
    otlp:
        protocol: grpc #grpc | http
        endpoint:
          value: "myserver.local:55690"
          valueFrom:
            secretKeyRef:
                name: my-config
                namespace: default
                key: "endpoint"
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
