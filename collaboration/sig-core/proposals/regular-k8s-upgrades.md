# Regular upgrades of Kyma to work on latest Kubernetes

Created on 2019-08-27 by Michal Hudy, Marcin Witalis and Lukasz Gornicki.

## Motivation

We should regularly validate Kyma compatibility with latest K8s version to not create technology dept and adjust and react properly to comming changes.

Examples of the issue:
- some time ago apiVersions of some workloads, like Deployments were set to be deprecated. We could use new versions seamlesly few months ago. Unfortunately it didn't happen and because copy/paste we used many deprecated areas of K8s that causes a lot of issues for Kyma upgrades. 
- we know that we are compatible with 1.12 K8s while the latest version is 1.15. This is not the best situation that accumulates possible upgradability problems
- we do not know if Kyma works on **kind** and latest **minikube**

## Solution

We need to have in place the following:
* CI
  * Weekly CI jobs that runs Kyma installation and upgrade on latest [kind](https://github.com/kubernetes-sigs/kind) (kind is the recommended CI tool for testing K8s)
  * Weekly CI jobs that runs Kyma installation and upgrade on latest minikube
* Every release we need a person responsible to:
  * Update the jobs to point to latest kind, minikube and k8s
  * Review release notes and new features presentation and prepare a summary for all Kyma maintainers on important features and deprecations
  * Review the results of the jobs that 
    * fail and creates GitHub issues to provide fixes. If the same fixes are required in several components, this person should provide recommendation how to provide a fix.
    * are successful and provides updates to documentation about support K8s version
  * Checks AKS and GKE support for K8s and makes sure that in docs and CI we always use latest. Managed offerings are always behind. For example now K8s 1.15 is available but managed offerings support 1.13 as latest, with 1.14 as beta.

## Action items

* Create CI plans. Work on **kind** is delayed by Service Catalog that in Kyma still requires separate etcd. Because 1.16 is not yet released, we can work on pipeline for Kyma 1.7 timeline in October
* Create a clear procedure/guideline how to perform manual work so the whole upgrade flow is simple to go through by any maintainer. 
