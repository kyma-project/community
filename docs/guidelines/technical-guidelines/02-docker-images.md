# Docker Images

This document provides guidelines for the Docker image provided in the context of Kyma.

## Naming and Structure Guidelines

Place images in the Kyma Docker registry located at `eu.gcr.io/kyma-project`. For development and proof of concepts, use the following location: `eu.gcr.io/kyma-project/snapshot`.

All images use the following attributes:

- an image name which is the same as the related project. Do not use prefixes. If the image requires sub-modularization, append it as in "istio-mixer"
- a tag with a semantic version number, like `0.3.2`

Assume an initializer image for the Helm Broker extension. This is the example of the location and the name of the image:

```bash
eu.gcr.io/kyma-project/helm-broker-initializer:0.1.0
```

## Base Images

Base all images on an image that is as small as possible in size and dependency. A base image must have a specified version. Do not use the `latest` tag.

An application based on Go should originate from a `scratch` image. If a `scratch` image does not have the specific tooling available, you can use an `alpine` base image having the package catalog updated.
A JavaScript-based application should originate from an `nginx-alpine` base image with an updated package catalog.

## Label Images

All images use the `source` label with a link to the GitHub repository containing the sources.

Define labels as in the following example:

```bash
source = git@github.com:kyma-project/examples.git
```

## Third-Party Images

Kyma uses some Docker images that originally were not built (and hosted) by us.
For security and reliability reasons, we need to copy all external images to our own Docker registry.
We have two solutions to this problem: the third-party-images repository and the image-syncer tool.

### Third-Party Repository

If you want to rebuild the image from scratch, use the [third-party-images](https://github.com/kyma-incubator/third-party-images) repository.
For every component, create a separate directory. You need to provide a Dockerfile, a Makefile, and create a ProwJob for building your images.
See the repository content for more information.

### Image Syncer

If you want to "cache" an image from an external registry, use the image-syncer tool.

To copy the image to our registry, modify the `external-images.yaml` file in your repository.

For example, the source image `grafana/grafana:7.0.6` will be transformed to `eu.gcr.io/kyma-project/external/grafana/grafana:7.0.6"`.
This URL can then be used in your Helm charts.

## Image Builder Documentation

For building Docker images within Kyma, refer to the [image-builder documentation](https://pages.github.tools.sap/kyma/documentation/kyma-internal/how-to-guides/090-build-oci-images-with-image-builder.html).
Image Builder is designed to streamline the process of creating and publishing Docker images for Kyma components.

### Dockerfile Recommendations

- Cross-Compilation: If you are building non-native architecture images, implement cross-compilation in your Dockerfile, use
  the [Faster Multi-Platform Builds: Dockerfile Cross-Compilation Guide](https://www.docker.com/blog/faster-multi-platform-builds-dockerfile-cross-compilation-guide/)
  as a reference.
- Bind Mounts: To avoid copying source code for compilation, use bind mounts for the `RUN` command in Dockerfiles.
  However, the speed gain was minimal in our tests: We achieved a speedup of less than ~5 seconds.
- Cache Mounts for Go Compiler:
  Rely on a cache backed by a remote repository, because a new agent is allocated for each pipeline execution, making mount-type caching
  ineffective.
  Use mounts a cache type for Go package downloads. The binary compilation cache did not increase speed during tests.

### Example Dockerfile to Build Publicly Available Images

```dockerfile
FROM --platform=$BUILDPLATFORM golang:1.24.2-alpine3.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download -x

ARG TARGETOS TARGETARCH
RUN --mount=target=. cd /app/cmd/image-builder && CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -buildvcs=false -o /image-builder -a -ldflags '-extldflags "-static"' .

FROM scratch

COPY --from=builder /image-builder /image-builder

ENTRYPOINT ["/image-builder"]
```

### Example Dockerfile to Build Restricted Images

```dockerfile
FROM europe-docker.pkg.dev/kyma-project/restricted-dev/sap.com/python-fips:latest
WORKDIR /app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

CMD ["python", "-m", "your_module"]
```
