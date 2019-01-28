# Loki as bundled logging solution

Created on 2019-01-28 by Andreas Thaler (@a-thaler).

## Status

Proposed on 2019-01-28

## Motivation
With Kyma a "batteries-included" like solution for logging should be provided. Such a solution must be very lightweight as it must run in-cluster and should work even for local development. Current approach is to use [OKLog](https://github.com/oklog/oklog), which is fullfilling that main criteria but unfortunately is discontinued and lacks a lot of user experience features.

## Goal
Replace OKLog by a solution which is:
- leightweight in resource consumption (in-cluster local development support)
- easy to operate even when scaling up the cluster
- allow to query recent logs for a pod in a user-friendly way
- maintained and further developed by a community

## Proposal
Replace OKLog with [Loki](https://github.com/grafana/loki) by:
- Adjust the logging helm chart to deploy the Loki service and the promtail daemonSet instead of OKLog service and logspout daemonSetpromtail
- Enable the experimental feature of grafana to enable the new exporer and integration with Loki as data source

## Details
Loki is a new project by Grafana Labs and aims to provide a similar solution for logging, like Prometheus is providing for monitoring.
It uses a similar approach like OKLog by explicitly not doing a full text indexing, but just an indexing on metadata level. Here, it is leveraging the existing labels of Prometheus and Kubernetes and it integrates with Grafana out-of-the-box (experimental feature for now).

Besides OKLog and Loki there was no other alternative covering the described requirements.
OKLog should be replaced with Loki as
- it has an active development community
- is leveraging the existing labels and with that provides a much better search experience
- it provides the better Graphical user interface by being integrated to Grafana

