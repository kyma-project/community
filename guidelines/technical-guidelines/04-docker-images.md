---
title: Docker images
---

This document provides guidelines for the Docker image provided in the context of Kyma.

## Naming and structure guidelines

Place images in the Kyma Docker registry located at `eu.gcr.io/kyma-project`. For development and proof of concepts, use the following location: `eu.gcr.io/kyma-project/snapshot`.

All images use the following attributes:

- an image name which is the same as the related project. Do not use prefixes. If the image requires sub-modularization, append it as in "istio-mixer"
- a tag with a semantic version number, like `0.3.2`

Assume an initializer image for the Helm Broker extension. This is the example of the location and the name of the image:

```bash
eu.gcr.io/kyma-project/helm-broker-initializer:0.1.0
```

## Base images

Base all images on an image that is as small as possible in size and dependency. A base image must have a specified version. Do not use the `latest` tag.

An application based on Go should originate from a `scratch` image. If a `scratch` image does not have the specific tooling available, you can use an `alpine` base image having the package catalog updated.
A JavaScript-based application should originate from an `nginx-alpine` base image with an updated package catalog.

## Label images

All images use the `source` label with a link to the GitHub repository containing the sources.

Define labels as in the following example:

```bash
source = git@github.com:kyma-project/examples.git
```

## Third party images

Kyma uses some docker images that originally were not build (and hosted) by us. 
Because of security and reliability reasons we need to copy all external images 
to our own docker registry.
We have two solutions to this problem - third-party-images repo and image-syncer tool.

### Third party repo

If you want to rebuild the image from scratch use the [third-party-images](https://github.com/kyma-incubator/third-party-images) repository.
For every component create a separate directory. You need to provide Dockerfile, Makefile and create prow job for building your images.
See repository content for more information.

### Image syncer

In case you want to "cache" an image from external registry use the [image-syncer
](https://github.com/kyma-project/test-infra/tree/master/development/image-syncer)
tool. 

To copy image to our registry you need to modify the [external-images.yaml](https://github.com/kyma-project/test-infra/blob/master/development/image-syncer/external-images.yaml).
After your change is merged to master branch you can check the new image URL in the logs of the [post-master-test-infra-image-syncer-run](https://status.build.kyma-project.io/job-history/kyma-prow-logs/logs/post-master-test-infra-image-syncer-run) job.

For example: source image `grafana/grafana:7.0.6` will be transformed to `eu.gcr.io/kyma-project/external/grafana/grafana:7.0.6"`.
This URL can then be used in your Helm charts.

## Examples

Go from scratch:

```Dockerfile
FROM scratch
LABEL source=git@github.com:kyma-project/examples.git

ADD main /
CMD ["/main"]
```

Go from alpine:

```Dockerfile
FROM alpine:3.7
RUN apk --no-cache upgrade && apk --no-cache add curl

LABEL source=git@github.com:kyma-project/examples.git

ADD main /
CMD ["/main"]
```

JavaScript from nginx:

```Dockerfile
FROM nginx:1.13-alpine
RUN apk --no-cache upgrade

LABEL source=git@github.com:kyma-project/examples.git

COPY nginx.conf /etc/nginx/nginx.conf
COPY /build var/public

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
```
