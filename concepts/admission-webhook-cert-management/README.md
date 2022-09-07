# Admission webhook certificate management

This document proposes several ways to manage TLS certificates for custom admission webhooks. 

## Problem

Kubernetes admission controllers act as gatekeepers intercepting API requests and can change the request object or deny its entry to the cluster.
It is possible to extend the compiled-in admission controllers with custom webhooks using the [dynamic admission control mechanism](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/).

<img src="assets/admission-controller-certs.drawio.svg">

## How Kyma webhooks manage cerificates

A handful of Kyma operators are deployed with custom validating admission webhooks (api-gateway, pod-preset, serverless, telemetry, etc.). They use different approaches to manage certificates, which can be categorized into 2 groups.

### <a name="helm"></a>Use Helm built-in crypto functions to generate certificates

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


### <a name="server-code"></a>Generate the CA cert and the server cert and update the webhook configuration in the webhook server code itself

This approach is used by serverless and api-gateway. In this case, the certificates are generated upon the server startup. Here's an [example of this solution](https://github.com/kyma-project/api-gateway/blob/main/internal/webhook/certificates.go).

Issuing a certificate in the webhook server code doesn't have the above-mentioned problem. However, in this case it has to be implemented by every operator. In addition to that, the webhook server will have to be provided with extended permissions to change the corresponding `validatingwebhookconfiguration` (or the `mutatingwebhookconfiguration`). In addition to that, both the CA and the server certificate will be recreated upon each pod restart and will never be rotated if the pod is not restarted.


## How other teams manage webhook certificates

The most prominent operator frameworks advertise `cert-manager` as a possible solution to manage webhook certificate:

- Kubebuilder: https://kubebuilder.io/cronjob-tutorial/cert-manager.html
- Operator SDK: https://developer.ibm.com/tutorials/create-a-conversion-webhook-with-the-operator-sdk/

Let's discribe a few possible solutions.

### <a name="gardener-cert-manager"></a>Gardener cert-manager

https://github.com/gardener/cert-management

Since Kyma-on-Gardener is a very common setup, and Gardener has a pre-installed cert-manager, it seems to be a natural choice. However Gardener `cert-manager` is actually a separate open source project and is very different from the upstrem `cert-manager`. It comes with a different feature set (it even has its own CRDs). From the Github project description:

>In a multi-cluster environment like Gardener, using existing open source projects for certificate management like cert-manager becomes cumbersome. With this project the separation of concerns between multiple clusters is realized more easily. The cert-controller-manager runs in a secured cluster where the issuer secrets are stored. At the same time it watches an untrusted source cluster and can provide certificates for it.

However, according to the project maintainers the Gardener `cert-manager` is built for mainly for the “Let’s Encrypt” use-cases (publicly resolvable domain names). It does not support self-signed issuers, nor it supports caBundle injection. The project maintainers suggest using the upstream `cert-manager` from the community for self-signed/webhook certs.

Moreover, there are other supported non-Gardener Kyma environments (k3d), where the Gardener `cert-manager` is not pre-installed. 
### <a name="upstream-cert-manager"></a>Upstream cert-manager

https://github.com/cert-manager/cert-manager

<img src="assets/admission-controller-certs-with-cert-manager.drawio.svg">

The `cert-manager` is an operator running inside a Kyma cluster. It is a Kyma component itself and is used by other Kyma components that have custom webhook servers. A shared self-signed `Issuer` is also deployed.

Each component that needs a TLS certificate would just deploy a `Certificate` resource and mount the generated secret as files to be used by the server container. Since all Kyma operators are implemented using Kubebuilder, secret rotation will work out of the box: whenever the `cert-manager` rotates a secret, the server will notice the files change and reload the http handler.

This approach has one big downsides - `cert-manager is` deployed using Helm, so the chart has to be maintained (upgrades, security patches, etc.)

### <a name="init-container"></a>Init container

https://www.velotio.com/engineering-blog/managing-tls-certificate-for-kubernetes-admission-webhook

This is slightly modified version of [one ofthe before-mentioned approach](#server-code). Instead of implementing the logic in the webhook server code, it's packaged as a Docker image and run as an init container.
