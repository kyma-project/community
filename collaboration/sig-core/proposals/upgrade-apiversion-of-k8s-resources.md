# Upgrade API versions of Kubernetes resources

Created on 2019-09-18 by Michal Hudy (@michal-hudy).

## Motivation

API Versions of Kubernetes resources are changing in time and some of them are going to be deprecated. Later, deprecated versions are no longer served by API server. Currently, we have such situation with Kubernetes 1.16 (https://kubernetes.io/blog/2019/07/18/api-deprecations-in-1-16), where some API versions that we are using in Kyma will be no longer served. The problem with an upgrade is that field `apiVersion` is immutable, so the resource cannot be just updated, they needs to be recreated.

## Solution

Solution is not easy as Kubernetes cluster should work without downtime so in ideal world we should create a new resources next to the old one that are deprecated and when new resources are ready, then remove the old one. Unfortunately, world is not ideal and not all resources can be duplicated - ingresses, resources that expects only one instance in the cluster, etc.

In such a case the downtime will be necessary, and we need to the following:
- Create a `pre-upgrade` Helm hook that will be removing old resource definition if there is a new API version available in the cluster
- Update API versions in resources if new API version is available.

Pre-upgrade hook may look as following:

`_helpers.tpl` file:
```yaml
{{/*
Return the appropriate apiVersion for deployment.
*/}}
{{- define "deployment.apiVersion" -}}
{{- if semverCompare "<1.9-0" .Capabilities.KubeVersion.GitVersion -}}
{{- print "apps/v1beta2" -}}
{{- else -}}
{{- print "apps/v1" -}}
{{- end -}}
{{- end -}}
```

`pre-upgrade-hook.yaml` file:
```yaml
---
apiVersion: v1
kind: ServiceAccount
metadata:
  labels:
    app: {{ .Release.Name }}
    release: {{ .Release.Name }}
    heritage: {{ .Release.Service }}
  annotations:
    helm.sh/hook: pre-upgrade
    helm.sh/hook-delete-policy: hook-succeeded
    helm.sh/hook-weight: "0"
  name: {{ .Release.Name }}-api-migration-job
  namespace: {{ .Release.Namespace }}
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  labels:
    app: {{ .Release.Name }}
    release: {{ .Release.Name }}
    heritage: {{ .Release.Service }}
  annotations:
    helm.sh/hook: pre-upgrade
    helm.sh/hook-delete-policy: hook-succeeded
    helm.sh/hook-weight: "1"
  name: {{ .Release.Name }}-api-migration-job
  namespace: {{ .Release.Namespace }}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: {{ .Release.Name }}-migration-job
subjects:
  - kind: ServiceAccount
    name: {{ .Release.Name }}-api-migration-job
    namespace: {{ .Release.Namespace }}
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  labels:
    app: {{ .Release.Name }}
    release: {{ .Release.Name }}
    heritage: {{ .Release.Service }}
  annotations:
    helm.sh/hook: pre-upgrade
    helm.sh/hook-delete-policy: hook-succeeded
    helm.sh/hook-weight: "1"
  name: {{ .Release.Name }}-api-migration-job
  namespace: {{ .Release.Namespace }}
rules:
  - apiGroups: ["*"]
    resources: ["deployments"]
    verbs: ["delete"]
    resourceNames: ["{{ .Release.Name }}-sample-deployment"]
---
apiVersion: batch/v1
kind: Job
metadata:
  labels:
    app: {{ .Release.Name }}
    release: {{ .Release.Name }}
    heritage: {{ .Release.Service }}
  name: {{ .Release.Name }}-api-migration-job
  namespace: {{ .Release.Namespace }}
  annotations:
    helm.sh/hook: pre-upgrade
    helm.sh/hook-delete-policy: hook-succeeded
    helm.sh/hook-weight: "2"
spec:
  template:
    metadata:
      labels:
        app: {{ .Release.Name }}
        release: {{ .Release.Name }}
        heritage: {{ .Release.Service }}
      name: {{ .Release.Name }}-api-migration-job
      namespace: {{ .Release.Namespace }}
      annotations:
        sidecar.istio.io/inject: “false”
    spec:
      serviceAccountName: {{ .Release.Name }}-api-migration-job
      restartPolicy: OnFailure
      containers:
        - name: migration
          image: bitnami/kubectl:1.15
          command:
            - bash
            - -c
            - |
              currentVersion=$(kubectl -n {{ .Release.Namespace }} get deployment {{ .Release.Name }}-sample-ingress -o jsonpath='{.apiVersion}')
              if [[ $currentVersion != "{{ template "deployment.apiVersion" . }}" ]]; then
                kubectl -n {{ .Release.Namespace }} delete deployment {{ .Release.Name }}-sample-ingress --wait=true
                sleep 10
              fi
```

## Action items

* Create pre-upgrade hooks
* Update API versions in resources
