# Test execution organization using labels

Created 2020-03-16 by Korbinian Stoemmer (@k15r)

## Motivation

Kyma developers write a lot of tests. These tests have different intentions and thus require different conditions to run successfully. In order to run those tests under the expected conditions on our CI Infrastructure, multiple shell scripts have been implemented. 
These shell scripts are tailored to the exact implementation of the test at the time of writing. This makes it really hard to reimplement a test, or to add a new test that requires the same setup (conditions) as another test.

Example:

currently the backup and restore tests are executed the following way on our CI:
* install kyma
* install `TestDefinition` with parameter indicating before backup test
* find all `TestDefinitions` with specific app label
* create `ClusterTestSuite` referencing the found `TestDefinitions` and wait for success
* take backup
* spinup new cluster and restore
* install `Testdefinition` with parameter indicating after restore test
* create `ClusterTestSuite` referencing the found `TestDefinitions` and wait for success


For upgrade tests the situation is as follows:

* install old kyma
* install upgrade test `helm-chart`
  * this creates a pod that does all the preparation work (compared to the backup test, this one does not use a pre-upgrade `TestDefinition`)
* upgrade kyma
* find **all** `TestDefinitions` in cluster
* create `ClusterTestSuite` referencing the found `TestDefinitions` and wait for success

Backup/Restore and upgrade tests implement their own framework for executing multiple tests.Event though both tests should be very similar and support reusing test cases it is very hard to do so. This leads to tests that are implemented multiple times.

## Goal

* It should be easy to reuse tests for multiple scenarios.
* It should be easy to implement tests without having to use yet another test runner. Octopus should be the only external requirement.
* It should be possible execute an unlimited number of test-cases per scenario
* There should be one approach to run those tests from within the shell scripts. This reduces code in the shell scripts to scenario setup code.

## Suggested solution

Octopus allows to specify label queries for `ClusterTestSuites`. A defined set of labels should indicate the which `TestDefinition` to execute for which scenario.

###The labels:

each `TestDefinition` should be labeled with one or multiple of these labels:

* `kyma-project.io/test.before-backup=true`
* `kyma-project.io/test.after-restore=true`
* `kyma-project.io/test.before-upgrade=true`
* `kyma-project.io/test.after-upgrade=true`
* `kyma-project.io/test.integration=true`

>*This list can be extended at will.*

Shell scripts install `ClusterTestSuites` with corresponding label queries e.g.:

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

>*Octopus` does not require any changes*

### Required Changes

* Kyma CLI
  * list available scenarios
  * support creation of `ClusterTestSuite` based on scenario(s)
* Existing Tests:
  * apply labels to all `TestDefinitions` including the ones that are not part of a basic Kyma installation (e.g. tests-folder)
* CI-Scripts
  * remove custom test execution logic and replace by kyma-cli call
* Upgrade and Backup/Restore tests:
  * get rid off old upgrade test framework / backup/restore test framework by refactoring the implemented tests into their own test definitions


### Optional Changes

* move additional `Helm-charts` from tests-folder into default installation (resources-folder)
  >*introduce additional/optional component 'additional-tests'*