# DR 021: Remove components proxying requests to the Kubernetes API server

Created on 2021-07-12 by Piotr Bochyński (@pbochynski).

## Decision log

| Name | Description |
|-----------------------|------------------------------------------------------------------------------------|
| Title | Remove components proxying requests to the Kubernetes API server |
| Due date | 2021-07-19 |
| Status | Proposed on 2021-07-12, Accepted on 2021-08-04|
| Decision type | Binary |
| Affected decisions | [DR 007: GraphQL as API facade for UI](https://github.com/kyma-project/community/blob/main/collaboration/sig-core/decisions/dr-007-GraphQL_as_API_facade_for_UI.md), [DR 015: Authorization for GraphQL](https://github.com/kyma-project/community/blob/main/collaboration/sig-core/decisions/dr-015-Authorization_for_GraphQL.md), [DR 008: DEX as an OIDC authenticator](https://github.com/kyma-project/community/blob/main/collaboration/sig-core/decisions/dr-008-Dex_as_an_OIDC_authenticator.md) |

## Context

Proxying the API server calls was initiated in the early days of Kyma in order to introduce unified authentication and authorization model for Kubernetes clusters coming from different vendors (Google Kubernetes Engine, Azure Kubernetes Service, Amazon EKS, etc). Whereas unification has some value for teams that need multi-cloud support, the solution has several drawbacks that should be addressed. In the meantime, the Kyma project picked [Gardener](https://github.com/gardener) as a managed Kubernetes platform to provide multi-cloud capabilities, and since then the additional layer of proxies does not add any value. The main reason for that decision is the simplification of the authorization concept in Kyma. The proxies together with the custom dex connector for UAA are hard to explain and introduce additional security vulnerabilities. This decision is fundamental for the changes coming in the [Kyma 2.0](https://github.com/kyma-project/kyma/issues/11337) release.

## Decision

Remove the following components from Kyma:
- apiserver-proxy
- console-backend
- iam-kubeconfig-service
- permission-controller
- uaa-activator
- dex

To replace the removed components, use the Kubernetes API server and its features for authentication and authorization, i.e.:
- [Authentication with OpenId Connect Tokens](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#openid-connect-tokens)
- [Role-based access control (RBAC) Authorization](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)


## Consequences

In consequence there is no pre-configured auth model for Grafana/Kiali/Jaeger exposure. We decided to have an oauth2-proxy in front of all this UIs having a default configuration which points to a documentation. Here port-forwarding is described as well as a guide on how to configure the proxy with OIDC, see also https://github.com/kyma-project/kyma/issues/11181

The benefits:
- faster installation and lower resource consumption (6 components less),
- reduced complexity of the authentication and authorization flow,
- improved security (e.g. removed service accounts with powerful roles),
- faster development and easier maintenance of the Console UI.

As a result of this decision, the following decisions are invalidated:
- [DR 007: GraphQL as API facade for UI](https://github.com/kyma-project/community/blob/main/collaboration/sig-core/decisions/dr-007-GraphQL_as_API_facade_for_UI.md)
- [DR 015: Authorization for GraphQL](https://github.com/kyma-project/community/blob/main/collaboration/sig-core/decisions/dr-015-Authorization_for_GraphQL.md)
- [DR 008: DEX as an OIDC authenticator](https://github.com/kyma-project/community/blob/main/collaboration/sig-core/decisions/dr-008-Dex_as_an_OIDC_authenticator.md)
