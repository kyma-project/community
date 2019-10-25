# Upgrade API versions of Kubernetes resources

Created on 2019-09-18 by Michal Hudy (@michal-hudy).

## Motivation

API Versions of Kubernetes resources are changing in time and some of them are going to be deprecated. Later, deprecated versions are no longer served by the API server. Currently, we have such a situation with Kubernetes 1.16 (https://kubernetes.io/blog/2019/07/18/api-deprecations-in-1-16), where some API versions that we are using in Kyma will be no longer served. For now, Kyma is not compatible with Kubernetes 1.16.

## Test Environment

I have tested proposed solution with:
 - kubectl v1.16.1
 - helm v2.14.3

## Solution

It looks like a solution is quite easy, but it is important to use the newest tools. First,  we need to update `apiVersion` field in charts and apply such changes to the cluster. The next step is changing the version in code, to be sure that our applications are asking and operating on the latest `apiVersion`. Later Kubernetes version can be upgraded.

## Action items

* Update `apiVersions` to the newest one
* Update code to operate on the newest `apiVersions`
* Pipelines are using newest version of `kubectl` or at least the same as Kubernetes version
