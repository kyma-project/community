---
title: Kyma release process
label: internal
---

This document describes how to create a Kyma release. Start from defining release jobs as described in the [**Preparation**](#kyma-release-process-kyma-release-process-preparation) section. Then proceed to the [**Steps**](#kyma-release-process-kyma-release-process-steps).

A Kyma release includes the following items:

* Docker images for Kyma components
* A GitHub release including release artifacts, such as source code and configuration
* A Git tag
* A release branch

## Definitions

The table below includes placeholders used throughout this document. When executing the commands, replace each of them with a suitable release number or version. 

| Placeholders | Description | Pattern | Example| 
|-------|------------|---------|--------|
| `RELEASE` | Release number| `{major}.{minor}` | `1.13`|
| `RELEASE_VERSION` | Release version | `{major}.{minor}.{patch}` or `{major}.{minor}.{patch}-rc{candidate}` | `1.13.0` or `1.13.0-rc1` |
| `RELEASE_VERSION_SHORT` | Release version without additional characters | `{major}{minor}` | `113`|
| `RELEASE_VERSION_DASH` | Release version with dashes |`{major}-{minor}-{patch}` or `{major}-{minor}-{patch}-rc{candidate}`| `1-13-0` or `1-13-0-rc1`|

## Preparation

> **NOTE:** This section applies only to new major and minor versions. If you release a patch, skip the preparation and go to the [**Steps**](#kyma-release-process-kyma-release-process-steps) section.

### kyma-project/kyma

#### Perform initial checks

Check if the master branch contains any PR-images:
   
   ```bash
   git grep -e 'version:\s.*[Pp][Rr]-.*' -e 'image:.*:[Pp][Rr]-.*' -e 'tag:\s.*[Pp][Rr]-.*' --before-context=2  resources tests
   ```
  
   Ask the teams for fixes if this command returns any output.

#### Create a release branch

>**NOTE:** a release branch needs to be created per new major / minor version. Patch releases and release candidates do not have a dedicated release branch.

Create a release branch in the `kyma` repository. The name of this branch should follow the `release-{major}.{minor}` pattern, such as `release-1.4`.

    ```bash
    git fetch upstream
    git checkout --no-track -b release-{RELEASE} upstream/master
    git push -u upstream release-{RELEASE}
    ```

### kyma-project/test-infra

#### Update the jobs on master branch

1. Create a PR to `master` containing the following changes to create the new job definitions:

    1. Open `templates/config.yaml`
    2. Add the new release to `global.releases`. Remove the oldest release on the list.
    3. Set `global.nextRelease` to the future release version.
    4. Run `make` in the root of the repository to generate jobs and run tests. If any of the tests is marked red, fix it using these guidelines:
      * For release tests using `GetKymaReleasesSince` or `jobsuite.Since` with a release that is no longer supported, change the method to `GetAllKymaReleases` or `jobsuite.AllReleases` respectively.
      * For release tests using `GetKymaReleasesUntil` or `jobsuite.Until` with a release that is no longer supported, remove the part of the test which includes the method.
    5. If tests are green, commit all jobs.

2. Once the PR is merged to master you can proceed.

#### Create a release branch

>**NOTE:** a release branch needs to be created per new major / minor version. Patch releases and release candidates do not have a dedicated release branch. If this branch already exists this step will be skipped.

Create a release branch in the `test-infra` repository

    ```bash
    git fetch upstream
    git checkout --no-track -b release-{RELEASE} upstream/master
    git push -u upstream release-{RELEASE}
    ``` 

## Steps

Follow these steps to release another Kyma version. Execute these steps for every patch release or release candidate.

### kyma-project/test-infra

Ensure that the `prow/RELEASE_VERSION` file from the `test-infra` repository on a release branch contains the correct version to be created. If you define a release candidate version, a pre-release is created.  

1. Make sure the `prow/RELEASE_VERSION` file includes just a single line, **without the newline character at the end**:  

        ```bash
        echo -n {RELEASE_VERSION} > prow/RELEASE_VERSION
        ```

2. If you had to change the RELEASE_VERSION, create a PR to update it on the release branch.
3. Once this PR is merged you can proceed.

### kyma-project/kyma

#### Create a PR to the release branch

1. Inside the release branch do the following changes.

   1. In `installation/resources/installer.yaml` replace `eu.gcr.io/kyma-project/develop/installer:{image_tag}` with `eu.gcr.io/kyma-project/kyma-installer:{RELEASE_VERSION}`

   2. In the `resources/core/values.yaml` file, find `clusterAssetGroupsVersion`.

        ```yaml
        docs:
        # (...) - truncated
        clusterAssetGroupsVersion: master
        ```

      And replace the `clusterAssetGroupsVersion` value with the following:

        ```yaml
        docs:
        # (...)
        clusterAssetGroupsVersion: release-{RELEASE}
        ```

2. Create a pull request with your changes to the release branch.

   ![PullRequest](./assets/release-PR.png)

3. If `pre-release-pr-image-guard` fails, ask the owners to change PR-XXX images of the components to the master version.
4. If the checks are green merge the PR and proceed to the next step.

#### Development process towards the release
   > **NOTE:** Steps 1-4 can be done by every developer that is introducing changes to the specific version.
1. Create a feature-branch based on the given `release-{RELEASE}` branch you want to extend. Add your changes and create Pull Request.
   
2.
Once you create Pull Request to the release branch the set of checks will be triggered.
These jobs run in the same way as on a Pull Request to the `master` branch. 
If you create a Pull Request that contains change in the components the component-building job will be triggered. 
If you make any changes in the charts, the integration tests will be triggered.

3. If you detect any problems with your PR fix the issues until your checks pass.

4. After all checks pass, merge your PR to the release branch. Merging the PR will trigger the post-submit integration tests automatically.
The jobs' status will be visible on the Kyma [TestGrid](https://testgrid.k8s.io/kyma_integration) in the corresponding dashboard tab.
   
5. If during the development process there will be need for additional changes in the release branch, open a new PR to the release branch with changes.
Repeat steps 1-4 for this PR.

#### Creating release
1. Once the release process is finished and release branch is complete create a new tag in the repository that points to your release branch.
> **CAUTION:** Make sure you are working on the most up-to-date `release-{RELEASE}` branch for a given release. 
```shell
git tag -a {RELEASE_VERSION} -m "Release {RELEASE_VERSION}"
git push upstream {RELEASE_VERSION}
```

Tag has to have the same name as in the `RELEASE_VERSION` file. Creating a new tag will trigger the following actions:
   
   * Create a GitHub release and trigger documentation update on the official Kyma website.

   * Create a new release cluster for the given Kyma `RELEASE_VERSION`.
   > **CAUTION**: The cluster is automatically generated for you, and it is automatically removed after 7 days.
   
   If you don't have access to the GCP project, post a request in the Slack team channel.

   ```bash
   gcloud container clusters get-credentials gke-{RELEASE_VERSION_DASH} --zone europe-west4-c --project sap-kyma-prow-workloads
   ```

   Follow [these](https://kyma-project.io/docs/#installation-install-kyma-with-your-own-domain-access-the-cluster) instructions to give Kyma teams access to start testing the release candidate.

5. The Github release postsumbit job creates a release in the `kyma-project/kyma` repository, which triggers the [`post-rel{RELEASE_VERSION_SHORT}kyma-release-upgrade`](https://github.com/kyma-project/test-infra/blob/master/prow/jobs/kyma/kyma-release-upgrade.yaml) pipeline. The purpose of this job is to test upgradability between the previous Kyma release, i.e. the latest release that is not a release candidate, and the brand new release published by the release postsubmit job.

    For example, if `1.7.0-rc2` is released, the pipeline will try to upgrade `1.6.0` to `1.7.0-rc2`.

    If you detect any problems with the upgrade, contact the teams responsible for failing components.

    > **CAUTION:** The job assumes no manual migration is involved. If the upgrade process requires any additional actions, the pipeline is likely to fail. In such case, the owners of the components concerned are responsible for running manual tests or modifying the pipeline.

6. On the release branch, update the `RELEASE_VERSION` file located in the `prow` folder of the [`kyma-project/test-infra`](https://github.com/kyma-project/test-infra) repository. It must contain the next release candidate version. Do it immediately after the release, otherwise, any PR to a release branch overrides the previously published Docker images.

   For example, if the `RELEASE_VERSION` file on the release branch contains `1.4.1`, change it to `1.4.2-rc1`.

7. Validate the `yaml` and changelog files generated under [releases](https://github.com/kyma-project/kyma/releases).

8. Update the release content manually with links to:

   * Instructions on local Kyma installation
   * Instructions on cluster Kyma installation
   * Release notes

   For installation instructions, use the links from the previous release and update the version number in URLs. If contributors want you to change something in the instructions, they would address you directly. Contact technical writers for the link to release notes.

> **NOTE:** After the Kyma release is complete, proceed with [releasing Kyma CLI](/guidelines/releases-guidelines/03-kyma-cli-release-process.md).

## Post-release tasks

1. Ask the Huskies team to upgrade the Kyma Helm chart on AppHub.

        > **NOTE:** Because of a limitation on the AppHub side, only a few people are allowed to create such a PR, which currently includes the members of the Huskies team.

2. Update `prow/RELEASE_VERSION` in the `master` branch of the `test-infra` repository with the name of the next minor release candidate, and merge the pull request to `master`. For example, if the `RELEASE_VERSION` on the `master` branch is set to `1.4.2`, change the version to `1.5.0-rc1`.
