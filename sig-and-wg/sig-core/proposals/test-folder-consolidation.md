# Test folder consolidation

## Introduction
The [tests](https://github.com/kyma-project/kyma/tree/master/tests) folder of the kyma-project contains all tests executed against a running kyma cluster to assure integrity and functional correctness of the cluster when all modules are installed together.
In kyma, that kind of tests are called _acceptance tests_.
All subfolders in the _tests_ directory define one test suite, usually focussing on one component.

## Problem
Every subfolder uses a different naming style, like _kubeless_ and _ui-api-layer-acceptance-tests_.

With that:
- Overall it is not easy to decide which convention to follow when creating a new subfolder.
- Not always it is possible to derive the related component (there is no e2e test yet, they are all focussed on components).
- It is confusing for readers

## Goal

Propose a name pattern which places the valuable information, the differentiator of the folder, into the focus.

## Consideration

**kind of test** All tests are of the same style currently, named as _acceptance tests_. As long as there is no different kind of tests available, the information of the kind of test is not valuable as part of the subfolder name

**test prefix/suffix** all subfolders contain test suites. That's why they are located under the _tests_ folder. Having a prefix or suffix in the folder name is redundant and does not bring more value to the reader.

**one component per folder** as long as a test suite is focussing on a specific component under test, the test suite should be named accordingly. With that the component and also the maintainers can be identified more quickly.
There is one special folder called _acceptance_ currently containing test suites for 3 different components, but sharing some few common scripts. As the tendency is to have a dedicated folder and docker image per component, that folder should be converted into three dedicated component folders with own docker images. Common scripts can be still shared on the root level or a dedicated _common_ folder.

**docker images** Every subfolder is producing a docker image using the folder name as image name. If no prefix/suffix is used, the name will collide with the image of the actual component. To have a differentiation here, a subfolder for the image could be introduced

## Proposal

The subfolder specific to a component should be named like the component, that's it.

No prefix, no suffix, no _acceptance_, no multi-components

The resulting docker image of a subfolder should be located in the subfolder _tests_

**Example**: A component _event-bus_ has its acceptance tests in folder _tests/event-bus_ and produces a docker image _XX/tests/event-bus:0.5.1_

Real e2e scenarios (like kubeles--integration) should be bundled into one subfolder _e2e_. Here we should have one test project which will execute all e2e tests organized by scenarios in different packages.

## Concrete Actions

- Have a readme in the _tests_ folder explaining the naming convention
- Link the readme in the developer guide
- Migrate _kubeless-integration_ to _e2e_ project containing e2e scenarios on a per package level resulting in one test suite to be executed.
- All folders will stay in the root _tests_ folder but will be renamed according to the following table:

| old folder | new folder | action required |
|------------|------------|-----------------|
| acceptance | dex | yes |
| _ | service-catalog | yes |
| _ | application-controller | yes |
|api-controller-acceptance-tests|api-controller| yes |
|application-operator-tests|application-operator| yes |
|application-registry-tests|application-registry|yes|
|connector-service-tests|connector-service| yes |
|event-bus|event-bus| - |
|gateway-tests|gateway|yes|
|knative-serving-acceptance|knative-serving|yes|
|kubeless-integration|e2e|yes|
|kubeless|kubeless|-|
|logging|logging|-|
|test-namespace-controller|namespace-controller|yes|
|test-logging-monitoring|monitoring|yes|
|ui-api-layer-acceptance-tests|ui-api-layer|yes|
