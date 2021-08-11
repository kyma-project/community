# Upgrade API versions of Kubernetes resources

Created on 2019-09-18 by Michal Hudy (@michal-hudy).

## Motivation

API Versions of Kubernetes resources change with time and some of them become deprecated. Once deprecated, the versions are no longer served by an API server. Currently, we have such a situation with [Kubernetes 1.16](https://kubernetes.io/blog/2019/07/18/api-deprecations-in-1-16), where some API versions that we use in Kyma will be no longer served. For now, Kyma is not compatible with Kubernetes 1.16.

## Test Environment

I tested the proposed solution with:
 - kubectl v1.16.1
 - helm v2.14.3

## Solution

The solution is quite easy, but it is important to use the latest tools. First,  we need to update the `apiVersion` field in charts and apply such changes to the cluster. The next step is changing the version in code to be sure that our applications call and operate on the latest `apiVersion`. Only then the Kubernetes version can be upgraded.

## Action items

* Update `apiVersions` to the latest version
* Update code to operate on the latest `apiVersions`
* Run pipelines using the latest version of `kubectl` or at least the same as the Kubernetes version
