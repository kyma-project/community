# Admission webhook certificate management

This document proposes several ways to manage TLS certificates for custom admission webhooks. 

## Problem

Kubernetes admission controllers act as gatekeepers intercepting API requests and can change the request object or deny its entry to the cluster.
It is possible to extend the compiled-in admission controllers with custom webhooks using the [dynamic admission control mechanism](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/).

<img src="assets/admission-controller-certs.drawio.svg">

A handful of Kyma operators are deployed with custom validating admission webhooks (apigateway, serverless, telemetry, etc.).
