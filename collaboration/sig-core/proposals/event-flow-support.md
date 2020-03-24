# Involved CR examples:

## Knative Broker

```yaml
apiVersion: eventing.knative.dev/v1alpha1
kind: Broker
metadata:
  annotations:
    eventing.knative.dev/creator: system:serviceaccount:knative-eventing:eventing-controller
    eventing.knative.dev/lastModifier: system:serviceaccount:knative-eventing:eventing-controller
  creationTimestamp: "2020-03-23T09:58:39Z"
  generation: 1
  labels:
    eventing.knative.dev/namespaceInjected: "true"
  name: default
  namespace: default
  resourceVersion: "410145"
  selfLink: /apis/eventing.knative.dev/v1alpha1/namespaces/default/brokers/default
  uid: 4dfe19e9-37fe-4df5-b79d-2170d34ada4d
spec:
  channelTemplateSpec:
    apiVersion: messaging.knative.dev/v1alpha1
    kind: NatssChannel
status:
  address:
    hostname: default-broker.default.svc.cluster.local
    url: http://default-broker.default.svc.cluster.local
  conditions:
  - lastTransitionTime: "2020-03-23T09:58:40Z"
    status: "True"
    type: Addressable
  - lastTransitionTime: "2020-03-24T09:38:22Z"
    status: "True"
    type: FilterReady
  - lastTransitionTime: "2020-03-23T09:59:02Z"
    status: "True"
    type: IngressReady
  - lastTransitionTime: "2020-03-24T09:38:22Z"
    status: "True"
    type: Ready
  - lastTransitionTime: "2020-03-23T09:58:39Z"
    status: "True"
    type: TriggerChannelReady
  observedGeneration: 1
  triggerChannel:
    apiVersion: messaging.knative.dev/v1alpha1
    kind: NatssChannel
    name: default-kne-trigger
    namespace: default
```


## HTTPSource

```yaml
apiVersion: sources.kyma-project.io/v1alpha1
kind: HTTPSource
metadata:
  creationTimestamp: "2020-03-23T15:24:52Z"
  generation: 1
  name: nachtmaar-test-2
  namespace: kyma-integration
  resourceVersion: "571271"
  selfLink: /apis/sources.kyma-project.io/v1alpha1/namespaces/kyma-integration/httpsources/nachtmaar-test-2
  uid: 1655b2fa-7ecb-4bbd-b572-325c8e196641
spec:
  source: nachtmaar-test-2
status:
  SinkURI: http://nachtmaar-test-2-kn-channel.kyma-integration.svc.cluster.local
  conditions:
  - lastTransitionTime: "2020-03-24T19:15:06Z"
    status: "True"
    type: Deployed
  - lastTransitionTime: "2020-03-24T19:14:51Z"
    severity: Info
    status: "True"
    type: PolicyCreated
  - lastTransitionTime: "2020-03-24T19:15:06Z"
    status: "True"
    type: Ready
  - lastTransitionTime: "2020-03-24T19:14:50Z"
    status: "True"
    type: SinkProvided
```

## ServiceInstance

```yaml
apiVersion: servicecatalog.k8s.io/v1beta1
kind: ServiceInstance
metadata:
  labels:
    servicecatalog.k8s.io/spec.serviceClassRef.name: 71e1f179f0a4ce451d12f91ce744909cbc32fe49b3924d788cc204f5
    servicecatalog.k8s.io/spec.servicePlanRef.name: e3ad3d8cf472eeecc4d471ed4a551c5190b0fec58513bcfef63f96a3
  name: es-all-events-07ee5-bumpy-communication
  namespace: default
  resourceVersion: "413220"
  selfLink: /apis/servicecatalog.k8s.io/v1beta1/namespaces/default/serviceinstances/es-all-events-07ee5-bumpy-communication
  uid: 6bd565af-6172-456f-878f-e7e69598430c
spec:
  externalID: 021b46aa-59fa-4963-b5e3-e79ed73fc7b6
  parameters: {}
  serviceClassExternalName: es-all-events-07ee5
  serviceClassRef:
    name: 78b4757c-472e-497d-abc7-b9811cef71b1
  servicePlanExternalName: default
  servicePlanRef:
    name: 78b4757c-472e-497d-abc7-b9811cef71b1-plan
  updateRequests: 0
  userInfo:
    groups:
    - system:serviceaccounts
    - system:serviceaccounts:kyma-system
    - system:authenticated
    uid: 14a76bc8-9317-4c3c-ae74-9e318ff882b5
    username: system:serviceaccount:kyma-system:service-catalog-controller-manager
status:
  asyncOpInProgress: true
  conditions:
  - lastTransitionTime: "2020-03-23T10:54:47Z"
    message: 'Error polling last operation: Status: 410; ErrorMessage: <nil>; Description:
      <nil>; ResponseError: <nil>'
    reason: ErrorPollingLastOperation
    status: "False"
    type: Ready
  currentOperation: Provision
  deprovisionStatus: Required
  inProgressProperties:
    clusterServicePlanExternalID: ""
    clusterServicePlanExternalName: ""
    servicePlanExternalID: 78b4757c-472e-497d-abc7-b9811cef71b1-plan
    servicePlanExternalName: default
    userInfo:
      groups:
      - system:serviceaccounts
      - system:serviceaccounts:kyma-system
      - system:authenticated
      uid: 14a76bc8-9317-4c3c-ae74-9e318ff882b5
      username: system:serviceaccount:kyma-system:service-catalog-controller-manager
  lastConditionState: ErrorPollingLastOperation
  lastOperation: 01E43GFPHKFXCWN90K1F3PREG8
  observedGeneration: 2
  operationStartTime: "2020-03-23T10:54:47Z"
  orphanMitigationInProgress: false
  provisionStatus: ""
  reconciledGeneration: 0
  userSpecifiedClassName: ServiceClass/es-all-events-07ee5
  userSpecifiedPlanName: default
```