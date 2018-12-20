# Central Connector Service

Created on 2018-12-17 by Lukasz Szymik (@lszymik).

## Status

Proposed on 2018-12-17.

## Motivation

The Connector service is responsible for establishing a secure connection between the connected external solutions and Kyma runtime. Such a connection is achieved by providing the client certificate to the connected solution. This certificate is later validated by the Application Connector. The client certificate is used for registering application metadata, such as APIs, and sending events to Kyma.
Currently, the connection between an external solution and Kyma is always point-to-point. 

As the customers work with multiple Kyma clusters, they will benefit from extending the provisioning of Kyma client certificates. A central Connector Service would manage the provisioning of certificates for multiple Kyma clusters and connected clients. Such approach allows the users to control their entire Kyma ecosystem from a single, central point.
Additionally, creating a central Connector Service enables the development of a central Application Registry. The central Application Registry can be used in a similar manner, serving multiple Kyma clusters and their connected external solutions.


## Goal

1. The Connector service handles client certificate provisioning for the connection with the Application Registry.
2. The Connector service handles client certificate provisioning for the connection with the Event Service.
3. The Connector service handles certificate provisioning for Kyma runtime.
4. The Connector service handles certificate rotation.
5. The Connector service returns information about the available cluster endpoints.


## Suggested solution

The Connector Service (CS) is deployed as a central component.

  - The CS is deployed as a global component in implementations with multiple Kyma clusters where one cluster takes the role of a master.
  - The CS exposes a secured connection for requesting client certificates signed with root CA.
  - The CS exposes a secured connection for requesting server certificates signed with root CA and deployed to Kyma runtime.
  - The client certificate enables a trusted connection with the central Kyma cluster where the App Registry is stored.
  - The client certificate enables a trusted connection with the Kyma runtime where the server certificate is delivered.
  - The server certificate enables a trusted connection with central Kyma cluster.
  - For standalone Kyma clusters, the Connector Service is deployed locally and works in the same manner.


### Component Diagram

![Connector Service Component](assets/connector-service-component-diagram.svg)

### Connection Flow Diagram

![Connector Service Flow](assets/connector-service-flow.svg)

#### Connection Flow overview

1. The application requests the client certificate from the central Connector Service. The flow is a standard one-time token flow currently used in Kyma. 

   The client certificate has the following properties:

   - Is signed by the root CA
   - The subject of the client certificate contains the unique ID of the Application and information about the group to which the Application is assigned

2. After the runtime is provisioned, it requests for the intermediate certificate. As a response, it receives a certificate chain consisting of the generated intermediate certificate and the root CA certificate. The intermediate certificate has the following properties:

   - Is signed by the root CA.
   - Contains the information about the runtime name for which it is generated.

3. The Application can access the master Kyma cluster and the Kyma runtime using the single certificate. The identity of the Application and the Kyma clusters is encoded in the certificate subject. It allows the verification of the calling parties.

### Cluster information

The Connector Service exposes the `info` endpoint which returns information about the connected clusters, including the App Registry URL, URL to the Event Service working in the runtime, etc.
A connected Application calls this endpoint periodically and checks the cluster status. 

### Certificate revocation 

The client certificates and intermediate server certificates must be revoked as soon as they are compromised. The list of revoked certificates will be stored in the central Connector Service and synchronized with all Kyma clusters. The list will contain both the client certificates and the server intermediate certificates.

The fingerprint of the compromised certificate will be added to the Application Connector. It will ensure that the connected Application which uses the compromised certificates can no longer perform any calls.

>**NOTE:** The current versions of Istio and the Nginx-controller do not support the revocation list. The further plan to add support for this feature will be provided later.

## Proof of Concept

### Prerequisites:

- Private key (`rootCA.key`) and certificate (`rootCA.crt`) generated as a root CA
- Kyma cluster provisioned with `rootCA.key` and `rootCA.crt` as a CA
- Application `test-application` created

### Steps

1. Generate the client certificate (`client.crt`) for the `test-application` using the Connector Service.
2. Access the cluster using the generated `client.crt`

    ```
    curl  https://gateway.{CLUSTER_DOMAIN}/test-application/v1/metadata/services --cert client.crt --key client.key
    ```

3. Generate the intermediate CA signed with the root CA

    ```
    openssl genrsa -out intermediate.key 4096
    openssl req -new -out intermediate.csr -key intermediate.key -subj /CN="intermediate"
    openssl x509 -req -sha256 -in intermediate.csr -out intermediate.crt -CAkey rootCA.key -CA rootCA.crt -days 1800 -CAcreateserial -CAserial serial
    ```

4.  Create the certificate chain containing `rootCA.crt` and `intermediate.crt`.

    ```
    cat rootCA.crt intermediate.crt > intermediate-chain.crt
    ```

5. Edit the secret containing CA `nginx-auth-ca` to use `intermediate-chain.crt` as `ca.crt` and `intermediate.key` as `ca.key`.

    ```
    export CERT=$(cat intermediate-chain.crt | base64)
    export KEY=$(cat intermediate.key | base64)
    cat <<EOF | kubectl apply -f -
    apiVersion: v1
    data:
      ca.crt: $CERT
      ca.key: $KEY
    kind: Secret
    metadata:
      name: nginx-auth-ca
      namespace: kyma-integration
    type: Opaque
    EOF
    kubectl -n kyma-system delete po -l "app=nginx-ingress"
    ```

6. Wait for the Pod to restart. Access the cluster using the previously generated `client.crt`:

    ```
    curl  https://gateway.{CLUSTER_DOMAIN}/test-application/v1/metadata/services --cert client.crt --key client.key
    ```

7. Revert changes by setting the root CA-signed certificate as the original one and call the cluster using of the intermediate certificate as the client certificate.
