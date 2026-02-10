# Kyma Cluster IP Stability

## Overview

This document provides information about the stability of cluster node public IPs in Kyma runtimes and guidelines for implementing IP-based access restrictions between Kyma clusters and external services.

## Background

Customers frequently need to implement access control lists (ACLs) or firewall rules on target services (such as databases, SFTP servers, or other external services) that are consumed from Kyma workloads. This requires stable and predictable IP addresses for egress traffic from Kyma clusters.

## IP Stability Status

### Current Status (Since October 2022)

Since October 2022, **all new Kyma runtimes have NAT Gateway enabled**, which provides stable egress IP addresses across all cloud providers (AWS, GCP, and Azure).

**Key Confirmation**: These egress IPs are **not expected to change throughout the lifetime of the Kyma runtime**.

### Historical Context

Prior to October 2022:
- AWS and GCP clusters had stable IPs by default
- Azure clusters had NAT Gateway disabled by default, resulting in dynamic egress IPs that could change
- This limitation was addressed through [Issue #13553](https://github.com/kyma-project/kyma/issues/13553)

## How to Retrieve Egress IPs

You can retrieve the stable egress IPs for your Kyma cluster using the following script:

```bash
#!/bin/bash

ZONES=$(kubectl get nodes -o jsonpath='{.items[*].metadata.labels.topology\.kubernetes\.io/zone}')

ZONES=$(for i in ${ZONES[@]}; do echo $i; done | sort -u)

```

**Note**: These IPs represent the VPC IPs that contain the Kyma cluster and are used for egress traffic across different availability zones.

## Ingress IP Stability

The ingress IP address for the Istio Ingress Gateway is also stable:

```bash
kubectl get svc -n istio-system istio-ingressgateway
```

The LoadBalancer IP for `istio-ingressgateway` remains stable throughout the cluster lifetime.

## Use Cases

### Common Scenarios

1. **Database Access Control**: Whitelisting Kyma egress IPs in database firewall rules (e.g., Azure Flexible Server PostgreSQL, HANA Cloud)
2. **SFTP Servers**: Configuring IP restrictions on customer or supplier SFTP servers
3. **External APIs**: Implementing IP-based access control for external services
4. **Service Instance Parameters**: Setting `allow_access` CIDR parameters for Kafka, PostgreSQL, and other services deployed in Kyma

### Example: Azure Flexible Server PostgreSQL

When configuring Azure Flexible Server PostgreSQL firewall rules to allow access from Kyma:

1. Retrieve egress IPs using the script mentioned above
2. Add each egress IP to the PostgreSQL firewall rules
3. The IPs will remain stable for the lifetime of the Kyma runtime

## Important Considerations

### For Managed Kyma (SAP BTP)

- Kyma runtimes provisioned via `btp-cli` on SAP Business Technology Platform have stable egress IPs (since October 2022)
- The SAP BTP documentation should reflect this stable IP guarantee
- Refer to [SAP BTP Regions for Kyma Environment](https://help.sap.com/docs/btp/sap-business-technology-platform/regions-for-kyma-environment) for regional information

### For Older Clusters

For Kyma runtimes provisioned **before October 2022**:
- Azure clusters may not have NAT Gateway enabled by default
- Egress IPs may be dynamic and subject to change
- Consider migrating to a newer runtime or contacting support for NAT Gateway enablement

## Recommendations

1. **Always retrieve current egress IPs** using the provided script to ensure you have the correct values
2. **Verify cloud provider and provisioning date** to confirm NAT Gateway is enabled
3. **Test connectivity** after configuring access lists to ensure proper configuration
4. **Document the IPs** used in your firewall rules for operational purposes
5. **Contact support** if you have questions about older clusters or specific configurations

## Related Resources

- [GitHub Issue #13553 - Enable NAT Gateway for Azure clusters](https://github.com/kyma-project/kyma/issues/13553)
- [SAP BTP Regions for Kyma Environment](https://help.sap.com/docs/btp/sap-business-technology-platform/regions-for-kyma-environment)
- [SAP Community Discussion - IP Ranges of NAT Gateways](https://answers.sap.com/questions/13720194/sap-kyma-ip-ranges-of-nat-gateways.html)

## Support

If you have questions about IP stability or need assistance with older clusters, please contact the SRE team or open a support ticket.

---

**Last Updated**: February 2026  
**Contributors**: Based on team discussions and confirmations from Gaurav Abbi, Benjamin Somhegyi, and the SRE team
