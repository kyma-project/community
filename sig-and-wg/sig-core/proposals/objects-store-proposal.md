# Objects Store

Created on 2019-01-03 by Lukasz Gornicki (@derberg).

## Status

Proposed on 2019-01-04

## Motivation

1. Have a generic solution for storing in s3 like storage any kind of object: zip file, markdown file, PNG or JS (JavaScript)
2. Have a storage solution not locked to one vendor, like AWS S3 only, or like to Minio on-premise only
3. Have a storage solution that not only stores the file but also exposes it directly to internet

## Use case

1. Storage for documentation and related images
2. Storege for API specifications
3. Storage for static client-side applications

## Solution
  
![](assets/storage.svg)

1. Location of the object is specified in the Object custom resource.
2. ObjectStore controller fetches the object basing on the information given in the custom resource.
3. The controller performs:
    - Mutation of the object by communicating with mutation webhook specified in the custom resource
    - Validation of the object by communicating with mutation webhook specified in the custom resource
    - New file creation, if such file was referenced in the resource definition as a ConfigMap
   If any of above operations failed, controller updates the resource with `ready: False` status  
4. Controller uploads the object to minio to a bucket that name is specified in the custom resource. You need a bucket to upload objects, you create it separately as a Bucket custom resource
5. Controller updates the status of the Object custom resource with information about location of the file

### Bucket custom resource

You might want to use different bucket per solution. This is why you need to be able to specify multiple buckets in the ObjectStore. For example one bucket for documentation and one bucket per Web application.

Another use cases for having multiple Buckets configuration:
- Future extensibility by alowing bucket policy specification per solution
- Minio doesn't support setting a bucket to behave as a static website host. Future controller of the Bucket custom resource will have to handle this additional functionality
- Minio doesn't support setting a CDN for your objects. Future controller of the Bucket custom resource will have to handle this additional functionality. For example controller will be reesponsible for configuring a CloudFront for your bucket on S3

Example resource for first version of the ObjectStore:
```
apiVersion: bucket.objectstore.kyma-project.io/v1alpha1
kind: Bucket
metadata:
  name: my-bucket
  namespace: stage
#ObjectStore 2.0  
#spec:
#  policy: public #or other policies
status:
  ready: False
  reason: BucketCreationFailure
  message: "service unavailable"
```

You reference the Bucket in your Object CR with the following info in the spec:
```
  bucketRef:
    name: my-bucket
```

It must be provided because the Object controller checks the Bucket custom resource status to make sure the bucket exists.

## Object custom resource
