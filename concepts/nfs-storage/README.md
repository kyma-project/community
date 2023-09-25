# Requirements

In Kubernetes, the need for ReadWriteMany (RWX) storage solutions arises from a variety of requirements that demand robust, scalable, and highly available shared storage. 

Some requirements collected from our customers:

1. High-Performance of writing small files: Rapidly write numerous small files (e.g., 3,000 files, 20KB each) within 30 seconds.
2. High-Availability: maintain file access during full availability zone outages.
3. Scalability: Scale storage capacity to accommodate over 1TB of data. Storage should expand automatically if needed.
4. Backup and restore: automatic backup and possibility to restore entire volume (disaster recovery) or certain folder
5. Enryption at rest.
6. Secure connection to the cluster

Development and operations related requirements:
1. Easy maintenance and should not require manual intervention for entire lifecycle. 
2. Mature and stable as it involves customer data durability.
3. Easy to protect from accidental (or intentional) damages caused by customer

# Options

## In-cluster Ceph-FS

**Pros:**
- one implementation suitable for all cloud providers
- best performance
- security - not exposed outside of the cluster

**Cons:**
- filesystem service runs in the cluster and can be affected by customer actions (also unintentional)
- complexity much higher than other solutions
- custom backup and restore solution has to be implemented
- lifecycle management can become a burden (upgrades, security patches, etc)
- requires upskilling the team to effectively maintain the solution
- some instability reported for writinglarger files


## Cloud provider NFS storage with dedicated CSI Driver

**Pros:**
- simple implementation: just install and configure CSI driver
- automatic storage instance provisioning (by driver)
- decent performance (depends on cloud provider)
- backups/snaphosts features

**Cons:**
- creating storage instances requires cloud provider credentials present in the cluster. Customers can get access to those credentials (they are cluster admins) and misuse them or even share them without any control
- different implementation for AWS and GCP
- creating storage instances is not always supported (AWS EFS csi driver doesn't create storage) - it has to be implemented in Kyma Control Plane
- [issues](https://cloud.google.com/filestore/docs/create-instance-issues#system_limit_for_internal_resources_has_been_reached_error_when_creating_an_instance) with private network quota for Filestore


## Generic NFS CSI Driver

**Pros:**
- the same driver can be used for AWS EFS and Google Filestore
- decent performance (depends on cloud provider)
- backups/snaphosts features
- security: no credentials exposed to customer, filesystem available only in cluster private network

**Cons:**
- provisioning storage instance and configuration of storage class has to be implemented in the control plane 
- different implementation for cloud providers (storage provisioning part, SKR part can be shared)
 

# NFS from cloud provider

## Use shoot VPC

![](nfs-shoot-vpc.drawio.svg)

## Use own VPC

![](nfs-own-vpc.drawio.svg)