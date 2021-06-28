# Test execution organization using labels

Created 2020-03-16 by Korbinian Stoemmer (@k15r)

## Motivation

Kyma developers write a lot of tests. These tests have different purposes and thus require different conditions to run successfully. In order to run those tests under the expected conditions on our CI Infrastructure, multiple shell scripts have been implemented.
These shell scripts are tailored to the exact implementation of the test at the time of writing. This makes it really hard to reimplement a test, or to add a new test that requires the same setup (conditions) as another test.

For example, currently the backup and restore tests are executed the following way on our CI:
1. Install Kyma.
2. Install `TestDefinition` with parameter indicating before backup test.
3. Find all `TestDefinitions` with specific app label.
4. Create `ClusterTestSuite` referencing the found `TestDefinitions` and wait for success.
5. Take backup.
6. Spin up new cluster and restore.
7. Install `TestDefinition` with parameter indicating after restore test.
8. Create `ClusterTestSuite` referencing the found `TestDefinitions` and wait for success.


For upgrade tests the situation is as follows:

1. Install old Kyma
2. Install upgrade test `helm-chart`. This creates a pod that does all the preparation work (compared to the backup test, this one does not use a pre-upgrade `TestDefinition`).
3. Upgrade Kyma.
4. Find **all** `TestDefinitions` in cluster.
5. Create `ClusterTestSuite` referencing the found `TestDefinitions` and wait for success.

Backup/Restore and upgrade tests implement their own framework for executing multiple tests. Event though both tests should be very similar and support reusing test cases, it is very hard to do so. This leads to tests being implemented multiple times.

## Goal

* It should be easy to reuse tests for multiple scenarios.
* It should be easy to implement tests without having to use yet another test runner. Octopus should be the only external requirement.
* It should be possible to execute an unlimited number of test cases per scenario.
* There should be one approach to run those tests from within the shell scripts. This reduces code in the shell scripts to scenario setup code.

## Suggested solution

Octopus allows to specify label queries for `ClusterTestSuites`. A defined set of labels should indicate which `TestDefinition` to execute for which scenario.

### The labels

Each `TestDefinition` should be labeled with one or more of these labels:

* `kyma-project.io/test.before-backup=true`
* `kyma-project.io/test.after-restore=true`
* `kyma-project.io/test.before-upgrade=true`
* `kyma-project.io/test.after-upgrade=true`
* `kyma-project.io/test.integration=true`

>**NOTE:** This list can be extended at will.

Shell scripts install `ClusterTestSuites` with corresponding label queries. For example:

```yaml
apiVersion: testing.kyma-project.io/v1alpha1
kind: ClusterTestSuite
metadata:
  name: before-backup
spec:
  concurrency: 1
  count: 1
  selectors:
    matchLabels:
    - kyma-project.io/test.before-backup: true
```


## Implementation

>**NOTE:** Octopus does not require any changes.

### Required Changes

* Kyma CLI:
  * List available scenarios.
  * Support creation of `ClusterTestSuite` based on scenario(s).
* Existing Tests:
  * Apply labels to all `TestDefinitions`, including the ones that are not part of the basic Kyma installation (e.g. the `tests` folder).
* CI-Scripts:
  * Remove custom test execution logic and replace it with a `kyma-cli` call.
* Upgrade and Backup/Restore tests:
  * Get rid off old upgrade test framework/backup/restore test framework by refactoring the implemented tests into their own test definitions.


### Optional Changes

* Move additional Helm charts from the `tests` folder into the default installation (the `resources` folder). For example, introduce an additional/optional `additional-tests` component.
