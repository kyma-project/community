# Admission webhook certificate management

This document proposes several ways to manage TLS certificates for custom admission webhooks. 

## Problem

Kubernetes admission controllers act as gatekeepers intercepting API requests and can change the request object or deny its entry to the cluster.
It is possible to extend the compiled-in admission controllers with custom webhooks using the [dynamic admission control mechanism](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/).

<img src="assets/admission-controller-certs.drawio.svg">

A handful of Kyma operators are deployed with custom validating admission webhooks (api-gateway, pod-preset, serverless, telemetry, etc.). They use different apporaches to manage certificates, which can be categorized into 2 groups.

1. Use Helm built-in crypto functions to generate certificates (used by pod-preset, telemetry):
```yaml
{{- $ca := genCA "telemetry-validating-webhook-ca" 3650 }}
{{- $cn := printf "%s-webhook" (include "fullname" .) }}
{{- $altName1 := printf "%s.%s" $cn .Release.Namespace }}
{{- $altName2 := printf "%s.%s.svc" $cn .Release.Namespace }}
{{- $cert := genSignedCert $cn nil (list $altName1 $altName2) 3650 $ca }}
apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingWebhookConfiguration
metadata:
  name: validation.webhook.telemetry.kyma-project.io
webhooks:
- clientConfig:
    caBundle: {{ b64enc $ca.Cert }}
...
---
apiVersion: v1
kind: Secret
metadata:
  name: telemetry-webhook-cert
  labels:
    {{- include "operator.labels" . | nindent 4 }}
    {{- toYaml .Values.extraLabels | nindent 4 }}
type: Opaque
data:
  tls.crt: {{ b64enc $cert.Cert }}
  tls.key: {{ b64enc $cert.Key }}
```

2. Generate the CA cert and the server cert and update the webhook configuration in the webhook server code itself (used by serverless and api-gateway). In this case, the certificates are generated upon the server startup.

The above-mentioned approaches have the following disadvantages:

1. The certs are getting updated every time the Helm Chart is rendered (every reconciliation). This updated is not atomic. For example, when the server cert is updated, but the caBundle of the webhook configuration is not yet updated, the webhook is non-working state and all the corresponding API requests will fail. This situation is actually not unkommon and is documented in [this bug](https://github.com/kyma-project/kyma/issues/15142). You can come up with some workarounds: make reconciler deploy resources in a strict predefined order or make sure the webhook chart does not contain the corresponding CRs. However, it does not fix the underlying problem and the bug may pop up again. 
2. TODO
