# Update component and its chart in one pull request

Created on 2019-04-23 by Adam Szecowka (@aszecowka) and Pawel Kosiec (@pkosiec). 

## Status
Proposed on 2019-04-23.

## Motivation
Currently, when you introduce changes to your component, you need to do at least two separate PRs:

- Changes made to the component
- Bump image version in the chart

Such an approach has many drawbacks:

- Slow development cycle, dealing with two PRs, asking for approvals twice
- A developer can forget to bump the image, in which case the version can be updated only on the release day. Updating version on the release day is 
the worst case scenario as it results in postponing the integration of our components. This is against the Continuous Integration approach.
- The code in the repository does not reflect the code that is actually used.
- Repository history is not transparent. Many commits only bump the component's image version.

## Goal
- You can introduce your changes and bump the image version of the Helm chart in the same pull request.
- The code in the Kyma repository reflects code executed on the Kyma.
When doing a release we rebuild all components. We should not be surprised on the release day that some tests are failing because some component was not updated for 3 weeks
- Kyma repository has a nice commit history and commits that only bumps a component version are rare.
- Run Kyma integration tests with the modified component within the same pull request status checks.


## Proposed solution

The solution assumes using images built on pull requests on the master branch.

Let assume that I am working on PR-1234. When I modify `componentA`, presubmit jobs build Docker images. 
The image has the same tag as the pull request number.
A developer can use this tag as a version in the `values.yaml` file:
```
component_a:
  dir: pr/
  version: PR-1234
```

With this approach, Prow executes component and integration jobs for the same PR. We have to ensure the proper order of jobs
to build components before using them in integration jobs. To achieve that, an additional step is required at the beginning of 
every integration job that waits for all dependent jobs. See [this](#guard-integration-jobs) section for more details.

To read more details, see [Guard integration jobs](#guard-integration-jobs).

In the beginning, the described approach can be optional, which means that a developer can decide whether he updates code and chart in the same PR or not. 
As a next step, we can introduce a job that checks if a version of the chart is updated in the PR. See
[this](#job-enforcing-changes-in-one-pr) section for more information.

In case there are two PRs that change the same component, there will be a merge conflict for the second PR because
it will modify the same line in the `values.yaml` file. Such approach ensures that the `master` branch contains changes from both PRs.

### Guard integration jobs
To postpone the execution of integration jobs, we should add an additional step at the beginning of every integration job.
To decide whether the integration job can be executed, use the checks of a given pull request:
![](./assets/job-status-checks.png)

Most of these checks are sent by Prow and represent statuses of jobs execution.

The Guard workflow looks as follows:
1. Fetch all required checks sent by Prow for a given PR that represent components build. 
2. If any status is marked as failed, the integration job fails. 
3. If all checks are successful, the integration job execution is continued.
4. If waiting for checks takes more than the defined timeout (10min), the integration job fails.
5. If some statuses are in Pending state, sleep for some time (15s) and go to point 1.

Ad Step 1. Guard filters statuses by its names. In Kyma Prow configuration, there is a convention for
job names. For example, every component job name for master branch starts with `pre-master-kyma-components-`.
Ad Step 2, Guard fails fast integration job to reduce the number of provisioned clusters and VMs.
Ad Step 4, Guard defines a timeout for checking jobs statuses. Prow defines a maximum number of concurrently executed jobs. 
There could be an extremely rare situation, that Prow executes only integration jobs that all wait for components jobs that cannot be executed because a maximum
number of concurrent jobs was reached. In such a case, a developer has to trigger an integration job manually. 

Find more information on Guard implementation in [PR#904](https://github.com/kyma-project/test-infra/pull/904).

### Job enforcing changes in one PR
To require updating charts immediately, a new Presubmit job should be defined. Below you can find it's pseudocode:
```
- get component's changed in current PR (ignore markdown files)
- Check if the chart uses the current version. Use the "path-to-referenced-charts" Makefile target.
```
Still, there should be an option to merge a PR without updating a chart. In such a case, a PR should have a special label, for
example the `postpone-integration`, in which case the presubmit job would succeed immediately.   
