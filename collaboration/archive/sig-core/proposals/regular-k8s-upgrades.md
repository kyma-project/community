# Regular upgrades of Kyma to work on latest Kubernetes

Created on 2019-08-27 by Michal Hudy, Marcin Witalis and Lukasz Gornicki.

## Motivation

We should regularly validate Kyma compatibility with the latest K8s version in order not to create technical debt and to adjust and react properly to coming changes.

Issues examples:
- Some time ago API versions of some workloads, like Deployments, were set to be deprecated. We could use new versions seamlessly a few months ago but unfortunately, it didn't happen. Because some data were copy/pasted, we used many deprecated areas of K8s which causes a lot of issues for Kyma upgrades. 
- We are compatible with K8s 1.12 while the latest version is 1.15. This is not the best situation that accumulates possible upgradability problems.
- We do not know if Kyma works on **kind** and the latest **Minikube**.

## Solution

We need the following:
* CI
  * weekly CI jobs that run Kyma installation and upgrade on the latest [kind](https://github.com/kubernetes-sigs/kind), which is a recommended CI tool for testing K8s
  * weekly CI jobs that run Kyma installation and upgrade on the latest Minikube.
* for each release a person responsible for:
  * updating the CI jobs to point to the latest kind, minikube and k8s
  * reviewing k8s release notes and new features presentations, and preparing a summary for all Kyma maintainers on important features and deprecations
  * reviewing the results of the CI jobs and:
    * if any jobs fail, the person should create GitHub issues to provide fixes. If the same fixes are required in several components, this person should provide recommendation on how to provide a fix.
    * if the jobs are successful, the person should provide updates to documentation about the support of a particular K8s version
  * checking AKS and GKE support for K8s and makes sure that in docs and CI we always use the latest. Managed offerings are always behind. For example, now K8s 1.15 is available but managed offerings support 1.13 as the latest, with 1.14 as beta.

## Action items

* Create CI plans. Working on **kind** is delayed because of the Service Catalog that in Kyma still requires separate etcd. Because 1.16 is not released yet, we can work on a pipeline for Kyma 1.7 timeline in October.
* Create a clear procedure/guideline how to perform manual work so the whole upgrade flow is simple to go through by any maintainer. 
