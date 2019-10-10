# Upgrade API versions of Kubernetes resources

Created on 2019-09-18 by Michal Hudy (@michal-hudy).

## Motivation

API Versions of Kubernetes resources are changing in time and some of them are going to be deprecated. Later, deprecated versions are no longer served by API server. Currently, we have such situation with Kubernetes 1.16 (https://kubernetes.io/blog/2019/07/18/api-deprecations-in-1-16), where some API versions that we are using in Kyma will be no longer served. For now, Kyma is not compatible with Kubernetes 1.16.

## Test Environment

I have tested proposed solution with:
 - kubectl v1.16.1
 - helm v2.14.3

## Solution

It looks like solution is quite ease, but it is important to use newest tools. First,  we need to update `apiVersion` field in charts and apply such changes to the cluster, later Kubernetes version can by upgraded.
For having Kyma compatible with multiple Kubernetes versions, we should dynamically detect `apiVersion` in chart, and this can be done with templates. 

Here is an example of template for Deployments resource:

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

Thanks to that, we will migrate to the new `apiVersion` if needed without cluster downtime.
Generally, we can consider build-in such functionality in Kyma Operator, then we will have determination of versions in one place.

## Action items

* Create helper with apiVersions
* Use detected version in charts
