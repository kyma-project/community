# Container Image Name Consolidation

## Introduction
To run Kyma on a Kubernetes cluster, a lot of different container images are generated. These images are referenced in templates used by Helm charts that are installed by the Kyma Installer.

## Problem
Currently, the correspondence between the image name and the image source code is not always clear.

As a result:
- It's not easy to find the source code for a particular image.
- It's not easy to find the latest image that was built using the source code.

## Goal

Propose a name pattern for container images that helps users quickly find images from source code and vice versa.

## Conventions
According to the Go convention, the `cmd` folder contains the source code for the main executables. Each executable should be put in a separate subfolder.

See the following example:
```shell
component-name
└- cmd
   ├- executable_a
   |  └- main.go
   └- executable_b
      └- main.go
```

> The `cmd` directory is not always required, but it is good practice to use it. If a component's source code can be used to create only a single binary, it is also ok to put its `main` function in a file called `main.go` on the root-level of the `components` folder.

Tests are stored in the `test` folder following [this](https://github.com/kyma-project/community/blob/main/collaboration/sig-core/proposals/test-folder-consolidation.md) proposal.

## Proposal

The path to each binary is unique. Significant information contained in the path of a binary should be used as the base for the image name that contains the binary.

### Components

The naming pattern could be as presented in the following table:

| Path | Image name | Description |
| ---- | ---------- | ---- |
| `components/<component-name>/cmd/<component-binary>` | `<component-name>-<component-binary>` | Applies if multiple images can be built using the component folder |
| `components/<component-name>/` | `<component-name>` | applies if only a single images can be built using the component-folder |

See the following examples:

| Path | Image name | Description |
| ---- | ---------- | ---- |
| components/event-bus/cmd/event-publish-service | event-bus-event-publish-service | Multiple Docker images |
| components/apiserver-proxy/cmd/proxy | apiserver-proxy | Single Docker image |

### Tests

The naming pattern could be as presented in the following table:

| Path | Image name |
| ---- | ---------- |
| `tests/<test-type>/<test-name>` | `<test-name>-<test-type>-tests` |
| `tests/<component-name>` | `<component-name>-tests` |

See the following examples:

| Path | Image name |
| ---- | ---------- |
| `tests/end-to-end/upgrade` | `upgrade-end-to-end-test` |
| `tests/knative-serving` | `knative-serving-test` |


## Actions

- Update the developer's guide to include the naming convention.
- Update all Makefiles to use the new naming convention.
- Update all Helm charts to reference the new images.

> Components and tests not listed below already adhere to the patterns described above.

### Components
Currently, all images follow the proposed pattern and exclusions.

The following table assumes all tests will be renamed according to the naming convention proposed [here](https://github.com/kyma-project/community/blob/main/collaboration/sig-core/proposals/test-folder-consolidation.md)

### Tests
| Path | Old image name | New image name |
| ---- | -------------- | -------------- |
| tests/end-to-end/backup-restore/ | backup-test | backup-restore-end-to-end-tests |
| tests/end-to-end/external-solution/ | e2e-external-integration-test | external-solution-end-to-end-tests |
| tests/end-to-end/kubeless/ | kubeless-integration | kubeless-end-to-end-tests |
| tests/end-to-end/upgrade/ | e2e-upgrade-test | upgrade-end-to-end-tests |
| tests/event-bus/ | event-bus | event-bus-tests |
| tests/integration/dex/ | dex-integration-tests | dex-integration-tests |
| tests/knative-build/ | knative-build-acceptance-tests | knative-build-tests |
| tests/knative-serving/ | knative-serving-acceptance-tests | knative-serving-tests |
| tests/service-catalog/ | service-catalog-acceptance-tests | service-catalog-tests |


## Optional actions

- Automatically generate the image name in a Makefile.
