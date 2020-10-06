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

>**CAUTION:** If you don't create the Kyma release branch at this point and add a  `post-rel{RELEASE_VERSION_SHORT}-kyma-release-candidate` post-submit job to the `test-infra` master, then pushing anything to the Kyma release branch, creating or rebasing the branch, triggers a new GitHub release.

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

1. Make sure the `RELEASE_VERSION` file includes just a single line, **without the newline character at the end**:  

        ```bash
        echo -n {RELEASE_VERSION} > prow/RELEASE_VERSION
        ```

2. If you had to change the RELEASE_VERSION, create a PR to update it on the release branch.
3. Once this PR is merged you can proceed.

### kyma-project/kyma

#### Create a PR to the release branch

1. Inside the release branch do the following changes.

   1. In `installation/resources/installer.yaml` replace `eu.gcr.io/kyma-project/develop/installer:{image_tag}` with `eu.gcr.io/kyma-project/kyma-installer:{RELEASE_VERSION}`

   2. Find these lines in `tools/kyma-installer/kyma.Dockerfile`:

        ```
        ARG INSTALLER_VERSION="{kyma_operator_version}"
        ARG INSTALLER_DIR={kyma_operator_path}
        FROM $INSTALLER_DIR/kyma-operator:$INSTALLER_VERSION
        ```

      Replace them with:

        ```
        FROM {kyma_operator_path}/kyma-operator:master-{kyma_operator_version}
        ```

   3. In the `resources/core/values.yaml` file, find the `clusterAssetGroupsVersion`

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

#### Execute the tests for the release PR

> **CAUTION:** Never use `/test all` as it might run tests that you do not want to execute.

1. Execute remaining tests. There are dependencies between jobs, so follow the provided order of steps.

    1.  Run `kyma-integration` by adding the `/test pre-rel{RELEASE_VERSION_SHORT}-kyma-integration` comment to the PR.

        > **NOTE:** You don't have to wait until the `pre-rel{RELEASE_VERSION_SHORT}-kyma-integration` job finishes to proceed with further jobs.

    2. Execute the next steps in the following order
        1. Run `/test pre-rel{RELEASE_VERSION_SHORT}-kyma-installer` and wait until it finishes.
        2. Run `/test pre-rel{RELEASE_VERSION_SHORT}-kyma-artifacts` and wait until it finishes.
        3. Run the following tests in parallel and wait for them to finish:

            ```
            /test pre-rel{RELEASE_VERSION_SHORT}-kyma-gke-integration
            /test pre-rel{RELEASE_VERSION_SHORT}-kyma-gke-central-connector
            /test pre-rel{RELEASE_VERSION_SHORT}-kyma-gke-upgrade
            /test pre-rel{RELEASE_VERSION_SHORT}-kyma-gke-compass-integration
            ```

2. If you detect any problems with the release, such as failing tests, wait for the fix that can be either delivered on a PR or cherry-picked to the PR from the `master` branch. Prow triggers the jobs again. Rerun manual jobs as described in .

3. After all checks pass, merge the PR, using the `rebase and merge` option.

   > **CAUTION:** By default, the `rebase and merge` option is disabled. Contact one of the `kyma-project/kyma` repository admins to enable it.

4. Merging the PR to the release branch runs the postsubmit jobs, which:

   * create a GitHub release and trigger documentation update on the official Kyma website
   * trigger provisioning of the cluster from the created release. Use the cluster to test the release candidate.

   > **CAUTION**: The cluster is automatically generated for you, and it is automatically removed after 7 days.

   If you don't have access to the GCP project, post a request in the Slack team channel.

   ```bash
   gcloud container clusters get-credentials gke-{RELEASE_VERSION_DASH} --zone europe-west4-c --project sap-kyma-prow-workloads
   ```

   Follow [these](https://kyma-project.io/docs/#installation-install-kyma-with-your-own-domain-access-the-cluster) instructions to give Kyma teams access to start testing the release candidate.

5. The Github release postsumbit job creates a tag in the `kyma-project/kyma` repository, which triggers the [`post-kyma-release-upgrade`](https://github.com/kyma-project/test-infra/blob/master/prow/jobs/kyma/kyma-release-upgrade.yaml) pipeline. The purpose of this job is to test upgradability between the previous Kyma release, i.e. the latest release that is not a release candidate, and the brand new release published by the release postsubmit job.

    For example, if `1.7.0-rc2` is released, the pipeline will try to upgrade `1.6.0` to `1.7.0-rc2`.

    If you detect any problems with the upgrade, contact the teams responsible for failing components.

    > **CAUTION:** The job assumes no manual migration is involved. If the upgrade process requires any additional actions, the pipeline is likely to fail. In such case, the owners of the components concerned are responsible for running manual tests or modifying the pipeline.

6. Update the `RELEASE_VERSION` file to contain the next patch RC1 version on the release branch. Do it immediately after the release, otherwise, any PR to a release branch overrides the previously published Docker images.

   For example, if the `RELEASE_VERSION` file on the release branch contains `1.4.1`, change it to `1.4.2-rc1`.

7. Validate the `yaml` and changelog files generated under [releases](https://github.com/kyma-project/kyma/releases).

8. Update the release content manually with links to:

   * Instructions on local Kyma installation
   * Instructions on cluster Kyma installation
   * Release notes

   For installation instructions, use the links from the previous release and update the version number in URLs. If contributors want you to change something in the instructions, they would address you directly. Contact technical writers for the link to release notes.

9. Create a new sheet in the [Release Testing](https://docs.google.com/spreadsheets/d/1ty3OciQzgzv0GagTG2Dku9os2AfMupbGNf8QxjHaO88)
   >**NOTE:** Make sure that you are not signed in with your SAP Google Account
   1. Open the **Main** sheet.
   2. Click the **generate new sheet** button.
   3. You will be asked for a GitHub personal access token. This token does not need any additional scopes!

> **NOTE:** After the Kyma release is complete, proceed with [releasing Kyma CLI](/guidelines/releases-guidelines/03-kyma-cli-release-process.md).

## Post-release tasks

1. Ask the Huskies team to upgrade the Kyma Helm chart on AppHub.

        > **NOTE:** Because of a limitation on the AppHub side, only a few people are allowed to create such a PR, which currently includes the members of the Huskies team.

2. Update `prow/RELEASE_VERSION` in the `master` branch of the `test-infra` repository with the name of the next minor release candidate, and merge the pull request to `master`. For example, if the `RELEASE_VERSION` on the `master` branch is set to `1.4.2`, change the version to `1.5.0-rc1`.
