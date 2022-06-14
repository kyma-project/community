# Revised API

## Motivation for the revision

The basic idea outlined in [the concept for the custom operator](05-custom-fluentbit-operator.md) and [the concept for the servicecatalog integration](06-servicecatalog-integration.md) was to have two layers:
1. A general layer supporting native Fluent Bit config snippets. Out of that, an overall Fluent Bit config is generated, validated, and injected to the daemonset.
2. A second layer on top, bringing the convenience of defining pipelines without knowing the underlying Fluent Bit concepts.

The MVP of the overall feature was based on releasing a first version of the first native layer. 
As we analysed various aspects, we discovered several problems and thus gradually adjusted the concepts, but didn't re-evaluate the overall picture. As a result, more and more abstractions and opinions were built into the first layer, making the layers not sufficiently differentiated anymore. The flexibility on the first layer was lost, while the second layer didn't fully fit on top anymore.

## Considerations

### Tag per pipeline mandatory

Different pipelines will need different filters. By default, there is only one tag like `kube.*` resulting from the central tail plugin. Every `LogPipeline` will need a dedicated `rewrite_tag` filter in order to apply some filtering, otherwise there will be indeterministic effects (the order of the filters is important).
-> Every pipeline needs a `rewrite_tag` introducing a custom tag per pipeline. To generate filters later out of a `LogPreset`, the tag name must be known.

### Filesystem buffer mandatory

We realized that using in-memory buffers will not decouple the pipelines well enough. If one output is down, the other output will stop working as well. To solve the problem, there must be an enabled filesystem buffer for each pipeline, including size management assuring that the node filesystem is not flooded.
-> Every pipeline needs a filesystem buffer; every output needs a storage size defaulting and validation.

### Match expressions should be limited

Support of full match expressions in a `LogPipeline` will harm other pipelines as well as the central `tail` pipeline. People copying code snippets might have match expressions using `*`. That must be prevented.
-> Match expressions should be defaulted and validated.

### Unsupported plugins

We need a clear distinction between supported and tested setups, and unsupported setups. Using an HTTP plugin should be supported with any configuration, while a stackdriver plugin is unsupported because we have no chance to cover testing. The additional resource consumption of some plugins and settings is also a relevant factor. We don't want the user to configure anything that creates a new emitter, which increases the memory usage above the amount that we introduce ourselves.
-> Introduce a mechanism to detect usage of unsupported plugins.

### Dedotting support

There is no filter available for opensearch-typical dedotting. That is a mandatory feature and we planned to solve it with a Lua script. On the other hand, the Lua filter should be an unsupported feature. So having the convenience of dedotting support in the `LogPresetBinding` would result in a Lua script in the `LogPipeline`. Both ways will not be optimal.
-> Add dedotting support as a feature to the `LogPipeline`.

### Meaningfulness of LogPreset

It turned out that an OpenSearch integration with an HTTP plugin required for an internal SAP scenario can be much simpler than expected. A real templating is not needed, and probably a plain simple pipeline config will be better suited. Also, regarding presets to provide, we cannot think of other scenarios that can be delivered. The binding to a secret can be done already in the pipeline itself, workload selectors could be added as well.
-> Do we need the LogPreset concept?

### Support of additional inputs
It turned out that there are typical inputs available (mainly the kubelet logs) that users might want to process as well.
-> Add a source selector.

## Revised API

More and more opinion and abstraction was added to the `LogPipeline` concept, and the meaningfulness of the `LogPreset` is questioned. We came up with the following ideas:
- We could try to push down the `LogPipeline` on the scala of abstraction to be more native (accepting that the user might do really stupid things by combining supported plugins), and revise the `LogPreset` concept.
- Alternatively, we push up the `LogPipeline` on the scala of abstraction even more by introducing the convenience of the `LogPreset`, while keeping native support as a customization option, clearly indicating an unsupported scenario.

This proposal favors the second idea by dropping the `LogPreset` fully and keeping only one layer that is focused on an abstraction and covers the actual value for users. A pure native layer per se adds no real value for users, as the user must understand the Fluent Bit concepts and the specifics of the layer, gaining only the lifecycle management aspects of the Fluent Bit instance - however, the lifecycle management does not provide any clearly guaranteed support.

The proposal takes the existing `LogPipeline` and extends it with a more abstracted syntax, still keeping the option to customize single elements. As parsers are a central config element, they are extracted into a dedicated resource `LogParser`.

```YAML
kind: LogPipeline # cluster scope
apiVersion: telemetry.kyma-project.io/v1alpha1
metadata:
  name: OpenSearchHTTP-App
spec:
  input: # singular, different inputs require different pipelines and output configs
    type: application # application | system | custom, default is application
    application: # maps to central tail pipeline, dealing with the actual application logs
      namespaces: [] # generates the rule for the rewrite_tag assigned to every pipeline
      excludeNamespaces: []
      containers: []
      excludeContainers: []

      podLabels: # generates a "grep" filter as a first element in the chain selecting logs by the "kubernetes" attributes
        app: icke
      excludePodLabels:
        app: chris
      
    system: {} # maps to systemd based input like kubelet logs, no further spec for now designed
    custom: | # define a custom input, entering unsupported mode
      Name dummy
      Dummy {"message":"dummy"}
  filters: # list of filters, order is important
    - add: # maps to "modify" filter, adds a log attribute
          key: cluster_identifier
          value: icke's cluster
    - remove: # maps to "modify" filter, removes a log attribute
          key: cluster_identifier
    - include: # maps to "grep" filter, drops lines where attribute will not match the regexp
        key: tenant
        regexp: icke
    - exclude: # maps to "grep" filter, drops lines where attribute will not match the regexp
        key: severity
        regexp: debug
    - custom: | # no "Match" available, entering unsupported mode
        Name    record_modifier
        Record  myKey myValue

  output: #only one output, no output should fail validation
    http: # enables http output
      Dedot: false # new flag resulting in a lua filter as last element in filter chain, default false as typically not used in http
      Host: # next 3 attributes should be configurable by static value and valueFrom incl. secret rotation
        value: "icke.com"
        valueFrom: ...
      HTTP_User:
        value: "icke"
        valueFrom:
          secretKeyRef:
            name: my-elastic-credentials
            namespace: default
            key: ES_USER
          rotatingSecretKeyRef:
            prefix: my-elastic
            namespace: default
            key: ES_USER
          configMapKeyRef:
            name: my-elastic-credentials
            namespace: default
            key: ES_USER
      HTTP_Password:
        value: "icke.com"
        valueFrom: ...
      URI: /customindex/kyma 
    custom: | # no "Match" available, no filebuffer settings available, entering unsupported mode
      Name               es
      Alias              es-output
      Host               ${ES_ENDPOINT} # Defined in Secret
      HTTP_User          ${ES_USER} # Defined in Secret
      HTTP_Password      ${ES_PASSWORD} # Defined in Secret
      LabelMapPath       /files/labelmap.json

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
        rotatingSecretKeyRef:
          prefix: my-elastic
          namespace: default
          key: ES_ENDPOINT
    - name: myEnv4
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

A `LogParser` specifies exactly one parser or a multiline parser, and is activated instantly on the tail plugin (multiline parser) or to be used in an annotation (parser).

```YAML
kind: LogParser # cluster scope
apiVersion: telemetry.kyma-project.io/v1alpha1
metadata:
  name: multiline-custom-regex
spec:
  parser: # Name is rejected because it is generated from the resource name
          # Will be registered as parser to be used in annotations, or in a pipeline using a custom parser filter
      |
        Format regex
        Regex ^(?<INT>[^ ]+) (?<FLOAT>[^ ]+) (?<BOOL>[^ ]+) (?<STRING>.+)$
  multilineParser: # Name is rejected because it is generated from the resource name
          # Will be registered as multilineparser on the tail plugin
      |
        type          regex
        flush_timeout 1000
        rule      "start_state"   "/(Dec \d+ \d+\:\d+\:\d+)(.*)/"  "cont"
        rule      "cont"          "/^\s+at.*/"                     "cont"
```

Example of minimal OpenSearch HTTP Application Log pipeline:
```YAML
kind: LogPipeline
apiVersion: telemetry.kyma-project.io/v1alpha1
metadata:
  name: OpenSearchHTTP-App
spec:
  input: #  optional section because application is default and all logs are processed by default
    application:
      excludeNamespaces: ["kyma-system", "kube-system"]
  filters: # optional section
    - include:
        key: tenant
        regexp: icke
  output: # mandatory section as a whole
    http:
      Dedot: true
      Host:
        valueFrom:
          secretKeyRef:
            namespace: default
            name: my-elastic
            key: HTTP_HOST
      HTTP_User:
        valueFrom:
          secretKeyRef:
            namespace: default
            name: my-elastic
            key: HTTP_HOST
      HTTP_Password:
        valueFrom:
          secretKeyRef:
            namespace: default
            name: my-elastic
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
  input: # mandatory section to include only Istio containers
    application:
      excludeNamespaces: ["kyma-system", "kube-system"]
      containers: ["istio-proxy"]
  filters: # mandatory section to exclude logs that have no protocol
    - include:
        key: protocol
        regexp: ".+"
  output: # mandatory section as a whole
    http:
      Dedot: true
      Host:
        valueFrom:
          secretKeyRef:
            namespace: default
            name: my-elastic
            key: HTTP_HOST
      HTTP_User:
        valueFrom:
          secretKeyRef:
            namespace: default
            name: my-elastic
            key: HTTP_HOST
      HTTP_Password:
        valueFrom:
          secretKeyRef:
            namespace: default
            name: my-elastic
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
  output:
    grafana-loki:
      Url: "http://logging-loki:3100/loki/api/v1/push"
      Labels:
        "job": "telemetry-fluent-bit"
      RemoveKeys: ["kubernetes", "stream"]
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
