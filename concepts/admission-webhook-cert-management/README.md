# Admission webhook certificate management

This document proposes several ways to manage TLS certificates for custom admission webhooks. 

## Problem

Kubernetes admission controllers act as gatekeepers intercepting API requests and can change the request object or deny its entry to the cluster.
It is possible to extend the compiled-in admission controllers with custom webhooks using the [dynamic admission control mechanism](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/).

<img src="assets/admission-controller-certs.drawio.svg">

## How Kyma webhooks manage cerificates

A handful of Kyma operators are deployed with custom validating admission webhooks (api-gateway, pod-preset, serverless, telemetry, etc.). They use different apporaches to manage certificates, which can be categorized into 2 groups.

### Use Helm built-in crypto functions to generate certificates

This approach is used by pod-preset, telemetry:
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

This is a very simple solution, but it has a lot of disadvantages.
The certs are getting updated every time the Helm Chart is rendered (every reconciliation). This updated is not atomic. For example, when the server cert is updated, but the caBundle of the webhook configuration is not yet updated, the webhook is in a non-working state and all the corresponding API requests will fail. This situation is actually not that unkommon and is documented in [this bug](https://github.com/kyma-project/kyma/issues/15142). 

You can come up with some workarounds: make reconciler deploy resources in a strict predefined order or make sure the webhook chart does not contain the corresponding CRs. However, it does not fix the underlying problem and the bug may pop up again. 


### Generate the CA cert and the server cert and update the webhook configuration in the webhook server code itself

This approach is used by serverless and api-gateway. In this case, the certificates are generated upon the server startup. Here's an [example of this solution](https://github.com/kyma-project/api-gateway/blob/main/internal/webhook/certificates.go).

Issuing a certificate in the webhook server code doesn't have the above-mentioned problem. However, in this case it has to be implemented by every operator. In addition to that, the webhook server will have to be provided with extended permissions to change the corresponding `validatingwebhookconfiguration` (or the `mutatingwebhookconfiguration`). In addition to that, both the CA and the server certificate will be recreated upon each pod restart and will never be rotated if the pod is not restarted.
