# Proposal: migration to Knative Event Mesh

## Contents

1. [Expectation](#expectation)
2. [Overview](#overview)
3. [First Kyma upgrade : 1.11](#first-kyma-upgrade--111)
     - [I. Missing HttpSource objects](#i-missing-httpsource-objects)
     - [II. Legacy endpoints compatibility](#ii-legacy-endpoints-compatibility)
     - [III. User namespace migration](#iii-user-namespace-migration)
     - [IV. Scale down Event Bus](#iv-scale-down-event-bus)
4. [Second Kyma upgrade : 1.12](#second-kyma-upgrade--112)
     - [I. Purge event-bus component](#i-purge-event-bus-component)

## Expectation

At the end of the Kyma upgrade, the user is left with an event mesh that has feature parity with the previous Kyma
version, without the need for any manual configuration and without message loss.

## Overview

The upgrade happens in two steps over two Kyma releases to avoid creating a major disruption. In
the first step, we ensure Kyma migrated totally to the new Knative Eventing mesh and that `Event Bus` is no longer
functioning. In the second step, we remove leftovers from all other charts and
delete the `event-bus` release in a manual step documented in Kyma upgrade guide.

<!--
described below are orchestrated by the Kyma operator. During a Kyma upgrade, the operator proceeds
component by component, iteratively. The most rational place to hook our migration logic is at the level of the
`cluster-essentials` component, which is the [very first chart](./installation/resources/installer-cr.yaml.tpl#L13-L15)
defined in the `Installation` object, so we can clear the path from there for the actual components' upgrades.
-->

Each of the steps described below is implemented as a Kubernetes `Job` triggered by a `pre-upgrade` Helm hook. A
predictable ordering is enforced using the `helm.sh/hook-weight` annotation.

A failed step restores the state of the cluster as it was when the step started.

## First Kyma upgrade : 1.11

The goal of this step is that the existing `event-bus` objects have their Knative Eventing mesh equivalence ones and
that`Event Bus` is no longer running.

### I. Missing HttpSource objects

#### Changes summary

| Name | Description | Artifact |
|------|-------------|----------|
|HttpSource sync binary| A Go binary which creates missing HttpSources | Docker image|
|`application-connector` `pre-upgrade` hook | A `pre-upgrade` Job which will be executed when Kyma installer upgrades the `application-connector` chart |`application-connector` new version |   

#### Execution logic 

- Create an `HttpSource` for each existing `Application` which has no existing one already.
- Wait, check for failure conditions and fail if any of them are true.
 
#### Failure conditions 
The job should fail when any of those conditions are true so that the release is marked as failed.
 
- If `HttpSource` or `Channel` are not ready after 2 min
- Errors on creating an `HttpSource`
- If it's observed that there is an `HttpSource` adapter but no `Channel`

### II. Legacy endpoints compatibility

#### Changes summary

| Name | Description | Artifact |
|------|-------------|----------|
|`event-service` build| `event-service` providing backward compatibility|Docker image| 
|`application` chart | `application` chart includes the new `event-service` docker image|[application](https://github.com/kyma-project/kyma/tree/master/components/application-operator/charts/application) new chart version| 
|`application-operator` build| A new `application-operator` build that includes the new `application` chart version in previous step|Docker image| 
|`application-connector` chart |`application-connector` new chart version with the updated `application-operator`|[application-connector](https://github.com/kyma-project/kyma/tree/master/resources/application-connector) new chart version|
|`application-connector` `post-upgrade` hook| A `post-upgrade` k8s job that checks if all applications event-services were upgraded correctly and fails if otherwise |`application-connector` new version| 


The new `Event Service` should serve legacy endpoint `/v1/events` and route them to the proper event mesh `HttpSource` adapter.

#### Execution logic

- Once the `application-connector` gets upgraded, the new `application-operator` will be start up.
- Each `Application` has a helm release with the application name (e.g. `commerce`). 
- The new `application` chart includes the new `event-service` deployment.
- When starting, the `application-operator` will compare the `application` chart version in its folder with the version of each release (e.g. `commerce-prod`, `varkes`, etc ) and upgrade it if needed.
- After `application` releases are upgraded the new `event-service` pods will restart and the new ones will start serving the legacy endpoint and routing the events to the HttpSource.
- The `post-upgrade` Job of the `application-connector` chart should wait, check for failure conditions and fail if any of them are true.
 
#### Failure conditions 
The `post-upgrade` job should fail when

- Any of the `event-service` instances are not updated
- Any of the `event-service` instances are failing

#### III. User namespace migration

#### Changes summary

| Name | Description | Artifact |
|------|-------------|----------|
|`ServiceInstance` recreation binary| A Go binary which creates `Trigger` objects for each Kyma `Subscription` objects | Docker image|
|`Trigger` sync binary| A Go binary which recreates `ServiceInstance` objects  | Docker image|
|`core` `pre-upgrade` hooks | Two `pre-upgrade` Jobs which will be executed when Kyma installer upgrades the `core` chart |`core` new version |
   
#### Execution logic

- Recreate each `ServiceInstance` object in its same namespace (check CBS code [here](https://github.com/kyma-project/kyma/blob/master/components/console-backend-service/internal/domain/servicecatalog/serviceinstance_service.go#L239)).
- Find all Kyma `Susbscription` objects and create a matching trigger object.
- Delete each `Susbscription` object after a matching trigger is successfully created.  
- Wait, check for failure conditions and fail if any of them are true.
 
#### Failure conditions 
The `pre-upgrade` job should fail when

- If an `EventActivation` exists and no `Subscription` exists after 2 min. 
- If a kyma `Subscriiption` still exists after 5 min. 

#### IV. Scale down Event Bus

#### Changes summary

| Name | Description | Artifact |
|------|-------------|----------|
|Scale down Event Bus hook | A bashscript `post-upgrade` hook in `core` which will be executed when Kyma installer upgrades the chart |`core` new version |
   
#### Execution logic

- Scale down `event-bus-event-publish-service` deployment to 0
- Scale down `event-bus-subscription-controller` deployment to 0


#### Failure conditions 

- If any pods of `event-bus-event-publish-service` or `event-bus-subscription-controller` still exist. 

## Second Kyma upgrade : 1.12

### I. Purge event-bus component

#### Changes summary

| Name | Description | Artifact |
|------|-------------|----------|
|Remove Event Bus CRDs | A new `cluster-essentials` chart versiion without `eventing-subscription.crd.yaml`| `cluster-essentials` new version |
|Document deleting `event-bus` release | A step documented in Kyma 1.12 migration guide how to helm delete `event-bus` release| Kyma 1.12 migration guide|
|Document deleting `event-bus` crds | A step documented in Kyma 1.12 migration guide how to delete `subscriptions.eventing.kyma-project.io` crd| Kyma 1.12 migration guide|

