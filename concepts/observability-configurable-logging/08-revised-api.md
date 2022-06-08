# Revised API

## Motivation For Revise

The basic idea outlined in [the concept for the custom operator](05-custom-fluentbit-operator.md) and [the concept for the servicecatalog integration](06-servicecatalog-integration.md) was to have
1. a general layer supporting native fluent-bit config snippets. Out of that an overall fluent-bit config gets generated, validated and injected to the damonset
2. a second layer on top, bringing the convinience of defining pipelines without knowing the underlaying fluent-bit concepts.

The MVP of the overall feature was based on releasing a first version of the 1. native layer. While doing several problems were discovered while doing analysis of several aspects. The These problems adjusted the concepts slightley step by step without re-evaluating the overall picture once again. As a result more and more abstractions/oppinions were build into the first layer, making the layers not differentiated enough anymore. The flexibility on the first layer got lost while the second layer was not fitting anymore fully on top.

## Considerations

### Tag per pipeline mandatory

Different pipelines will require different filters. By default there is only one tag like `kube.*` resulting from the central tail plugin. Every `LogPipeline` will require a dedicated `rewrite_tag` filter in order to apply some filtering, otherwise you will get indeterministic effects (the order of the filters is important).
-> Every pipeline needs a rewrite_tag introducing a custom tag per pipeline. To generate filters later out of a `LogPreset`, the tag name needs to be known

### Filesystem buffer mandatory

We realized that using in-memory buffers will not decouple the pipelines well enough. If one output is down, the other output will stop working as well. It requires to enable a filesystem buffer for pipeline including size management to assure that the node filesystem gets not flooded
-> Every pipeline needs a filesystem buffer, every output needs a storage size defaulting/validation

### Match expressions should be limited

Support of full match expressions in a `LogPipeline` will harm other pipelines as well as the central `tail` pipeline. People copying code snippets might have match expressions using `*`. It should be prevented.
-> Match expressions should be defaulted/validated

### Unsupported Plugins

We need a clear distinction between supported/tested setups and unsupported setups. An HTTP plugin usage should be supported with any configuration while a stackdriver plugin usage is unsupported as we have no chance to cover testing.
-> Introduce a mechanism to detect usage of unsupported plugins

### Dedotting support

There is no filter available for opensearch-typical dedotting. That is a mandatory feature and we planned to solve it via a lua script. On the other hand, the lua filter should be an unsupported feature. So having the convinience of detotting support in the `LogPresetBinding` would result in a lua script in the `LogPipeline`.
-> Add dedotting support as a feature to the `LogPipeline`

### Meaningfulness of LogPreset

It turned out that an OpenSearch integration via a HTTP plugin required for an internal SAP scenario can be much simpler then expected. A real templating is not needed and probably a plain simple pipeline config will be better suited. Also, what presets to provide, we cannot think of other scenarios which can be delivered. The binding to a secret can be done already in the pipeline itself, workload selectors could be added as well
-> do we need the LogPreset concept?

### Support of additional inputs
It turned out that there are typical inputs available which you would like to process as well, mainly the kubelet logs
-> add a source selector

## Revised API

More and more oppinion/abstraction was added to the `LogPipeline` concept and the meaningfullnes of the `LogPreset` is in the room. We could now try to push down the `LogPipeline` to be more native and accepting that the user can do really stupid things in combination of supported plugins and revise the LogPreset concept.
Or we push up even more the `LogPipeline` API by introducing the convinience of the LogPreset while keeping native support as a customization option clearly indicating an unsupported scenario.
This proposal favors the second idea by dropping the LogPreset fully and having only one layer which is already an abstraction. Giving no chance to configure fluent-bit natively. For that you can simply go with a custom installation.

```YAML
kind: LogPipeline
apiVersion: telemetry.kyma-project.io/v1alpha1
metadata:
  name: OpenSearchHTTP-App
spec:
  input:
    application: # maps to central tail pipeline
      namespaces: [] # maps to rewrite_tag rule
      excludeNamespaces: []
      pods: []
      excludePods: []
      containers: []
      excludeContainers: []
      matchLabels:
        app: icke
    system: {} #maps to systemd based input like kubelet logs
    custom: | # entering unsupported mode
      Name dummy
      Dummy {"message":"dummy"}
  filters:
    - multilineParser:
        type: "java"
    - parser:
        type: "json" # support of built-in parsers
    - modify:
        add:
          key: cluster_identifier
          value: ${KUBERNETES_SERVICE_HOST}
    - modify:
        remove:
          key: cluster_identifier
    - select:
        include:
          key:
          regexp:
        exclude:
          key:
          regexp:
    - custom: | # no "Match" available, entering unsupported mode
        Name    record_modifier
        Record  cluster_identifier ${KUBERNETES_SERVICE_HOST}

  output: #only one output
    http:
      Dedot: true
      Host:
        value: "icke.com"
      HTTP_User:
        value: "icke"
        fromSecretKeyRef:
          name: my-elastic-credentials
          namespace: default
          key: ES_ENDPOINT
      HTTP_Password:
        fromSecretKeyRef:
          name: my-elastic-credentials
          namespace: default
          key: ES_ENDPOINT
    custom:
      | # no "Match" available, no filebuffer settings available, entering unsupported mode
      Name               es
      Alias              es-output
      Host               ${ES_ENDPOINT} # Defined in Secret
      HTTP_User          ${ES_USER} # Defined in Secret
      HTTP_Password      ${ES_PASSWORD} # Defined in Secret
      LabelMapPath       /files/labelmap.json

  custom:
    variables:
      - fromSecretRef:
          name: my-elastic-credentials
          namespace: default
      - fromSecretPrefixRef: # secret rotation
          name: my-elastic-credentials
          namespace: default
      - fromSecretKeyRef:
          name: my-elastic-credentials
          namespace: default
          key: ES_ENDPOINT
      - fromConfigMapRef:
          name: my-elastic-config
          namespace: default
      - fromConfigMapKeyRef:
          name: my-elastic-credentials
          namespace: default
          key: ES_ENDPOINT

    files:
      - name: "labelmap.json"
        content: |
          {
            "kubernetes": {
                  "namespace_name": "namespace",
                  "pod_name": "pod"
            },
            "stream": "stream"
          }
      - fromConfigMapRef:
          name: my-loki-labelMap
          namespace: default
      - fromSecretRef:
          name: my-ip-whitelist
          namespace: default
    parsers:
      - custom: |
          Name dummy_test
          Format regex
          Regex ^(?<INT>[^ ]+) (?<FLOAT>[^ ]+) (?<BOOL>[^ ]+) (?<STRING>.+)$
    multilineParsers:
      - custom: |
          # Example from https://docs.fluentbit.io/manual/pipeline/filters/multiline-stacktrace
          name          multiline-custom-regex
          type          regex
          flush_timeout 1000
          rule      "start_state"   "/(Dec \d+ \d+\:\d+\:\d+)(.*)/"  "cont"
          rule      "cont"          "/^\s+at.*/"                     "cont"
```

Example of typical OpenSearch HTTP Application Log pipeline:
```YAML
kind: LogPipeline
apiVersion: telemetry.kyma-project.io/v1alpha1
metadata:
  name: OpenSearchHTTP-App
spec:
  input:
    application:
      excludeNamespaces: ["kyma-system", "kube-system"]
  filters:
    - parser:
        type: "json"
    - modify:
        add:
          key: cluster_identifier
          value: ${KUBERNETES_SERVICE_HOST}
  output:
    http:
      Dedot: true
      Host:
        fromSecretKeyRef:
          namespace: default
          prefix: my-elastic # secret rotation
          key: HTTP_HOST
      HTTP_User:
        fromSecretKeyRef:
          namespace: default
          prefix: my-elastic # secret rotation
          key: HTTP_HOST
      HTTP_Password:
        fromSecretKeyRef:
          namespace: default
          prefix: my-elastic # secret rotation
          key: HTTP_PASSWD
      URI: /customindex/kyma
```

Example of typical OpenSearch HTTP Istio Access Log pipeline:
```YAML
kind: LogPipeline
apiVersion: telemetry.kyma-project.io/v1alpha1
metadata:
  name: OpenSearchHTTP-Istio
spec:
  input:
    application:
      excludeNamespaces: ["kyma-system", "kube-system"]
      containers: ["istio-proxy"]
  filters:
    - parser:
        type: "json"
    - modify:
        add:
          key: cluster_identifier
          value: ${KUBERNETES_SERVICE_HOST}
    - select:
        include:
          key: protocol
          regexp: ".+"
  output:
    http:
      Dedot: true
      Host:
        fromSecretKeyRef:
          namespace: default
          prefix: my-elastic # secret rotation
          key: HTTP_HOST
      HTTP_User:
        fromSecretKeyRef:
          namespace: default
          prefix: my-elastic # secret rotation
          key: HTTP_HOST
      HTTP_Password:
        fromSecretKeyRef:
          namespace: default
          prefix: my-elastic # secret rotation
          key: HTTP_PASSWD
      URI: /customindex/istio-envoy-kyma

```

Example of default Loki pipeline:
```YAML
apiVersion: telemetry.kyma-project.io/v1alpha1
kind: LogPipeline
metadata:
  name: loki
spec:
  filters:
    - parser:
        type: "json"
  output:
    grafana-loki:
      Url: "http://logging-loki:3100/loki/api/v1/push"
      Labels:
        "job": "telemetry-fluent-bit"
      RemoveKeys: ["kubernetes", "stream"]
      LineFormat: json
      LogLevel: warn
      LabelMap:
        "kubernetes":
          "container_name": "container"
          "host": "node"
          "labels":
            "app": "app"
            "app.kubernetes.io/component": "component"
            "app.kubernetes.io/name": "app"
            "serverless.kyma-project.io/function-name": "function"
          "namespace_name": "namespace"
          "pod_name": "pod"
        "stream": "stream"
```

## Pros

- Approach makes one common layer which is clearly oppinionated but has options to opt-out in a controlled way (match expressions will still be enforced)

## Cons
