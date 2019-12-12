# Container Image Name Consolidation

## Introduction
In order to run kyma on a kubernetes cluster a lot of different container images are generated. These images are referenced in templates used by the helm charts that are installed by the `kyma-installer`

## Problem
Currently there is not always a clear relationship between an images name and the source code that resulted in the image.

As a result:
- It's not easily possible to find the source code for a particular image
- It's not easily possible to find the latest image that was built using the source code

## Goal

Propose a name pattern for container images that helps users quickly find images from source and vice versa.

##Conventions
**cmd** its is a golang convention to put source code for the main executable in a folder called `cmd`. A subfolder should exist for every executable that shall be created from the project.
e.g.
```shell
component-name
└- cmd
   ├- executable_a
   |  └- main.go
   └- executable_b
      └- main.go
```

**tests** Tests are sorted into the `test`-folder following this proposal [https://github.com/kyma-project/community/blob/master/collaboration/sig-core/proposals/test-folder-consolidation.md]


## Proposal

The path to each binary is unique. Significant information contained in the path of a binary should be used as the base for the image name that contains the binary.

### Components:

#### Pattern:

| path | image name | rule |
| ---- | ---------- | ---- |
| `components/<component-name>/cmd/<component-binary>` | `<component-name>-<component-binary>` | applies if multiple images can be built using the component-folder |
| `components/<component-name>/` | `<component-name>` | applies if only a single images can be built using the component-folder |

#### Example

| path | image name | rule |
| ---- | ---------- | ---- |
| components/event-bus/cmd/event-publish-service | event-bus-event-publish-service | multiple docker images |
| components/apiserver-proxy/cmd/proxy | apiserver-proxy | single docker image |

### Tests:

 #### Patterns:
 
| path | image name |
| ---- | ---------- |
| `tests/<test-type>/<test-name>` | `<test-name>-<test-type>-tests` |
| `tests/<component-name>` | `<component-name>-tests` |

#### Example

| path | image name |
| ---- | ---------- |
| `tests/end-to-end/upgrade` | `upgrade-end-to-end-test` |
| `tests/knative-serving` | `knative-serving-test` |


## Actions

- Update the developer's guide to include the naming convention
- Update all Makefiles to use the new naming convention
- Update all helm charts to reference the new images

> Components and tests not listed below already adhere to the patterns described above

### Components
Currently all images follow the proposed pattern and exclusions.

The following table assumes all tests will be renamed according to the naming convention proposed in [https://github.com/kyma-project/community/blob/master/collaboration/sig-core/proposals/test-folder-consolidation.md]

### Tests
| path | old image name | new image name |
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


## Optional

- automatically generate the image name in the Makefile.
