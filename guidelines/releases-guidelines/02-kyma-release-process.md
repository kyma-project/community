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

## Preparation

> **NOTE:** This section applies only to new major and minor versions. If you release a patch, skip the preparation and go to the [**Steps**](#kyma-release-process-kyma-release-process-steps) section.

To prepare a release:

1. Create a release branch in the `kyma` repository. The name of this branch should follow the `release-x.y` pattern, such as `release-1.4`.

   ```bash
    git fetch upstream
    git checkout -b release-{release_version} upstream/master
    git push upstream release-{release_version}
    ```

   > **NOTE:** If you don't create the Kyma release branch at this point and add a  `post-rel{release_version}-kyma-release-candidate` post-submit job to the `test-infra` master, then pushing anything to the Kyma release branch, creating or rebasing the branch, triggers a new GitHub release.

2. [Define new release jobs](#kyma-release-process-kyma-release-process-preparation-define-new-release-jobs) in the `test-infra` repository.
3. [Remove old release jobs](#kyma-release-process-kyma-release-process-preparation-remove-old-release-jobs) in the `test-infra` repository.

### Define new release jobs

1. Navigate to the `test-infra` repository.
2. Open `templates/config.yaml`
3. Change `global.releases` to contain new release. Also remove oldest release on the list.
4. Change `global.nextRelease` to contain version for future release. 
5. Run `go run development/tools/cmd/rendertemplates/main.go --config templates/config.yaml` in the root of the repository to generate jobs.
6. Run `go test development/tools/jobs/...` in the root of the repository. If anything is red fix it using these guidelines: 
  * For release tests using `GetKymaReleasesSince` with a release that is no longer supported change it to use `GetAllKymaReleases`.
  * For release tests using `GetKymaReleasesUntil` with a release that is no longer supported simply remove that part of the test.
7. When tests are green commit all jobs and you just got yourself new release jobs.

## Steps

Before completing these guidelines, make sure there is no mismatch between the source code and `.yaml` files. All components are rebuilt from the source code, which requires all Helm charts to be up to date.

Follow these steps to release another Kyma version.

### kyma-project/test-infra

1. Create a release branch in the `test-infra` repository. The name of this branch should follow the `release-x.y` pattern, such as `release-1.4`.

   > **NOTE:** This step applies only to new major and minor versions.

   ```bash
   git fetch upstream
   git checkout -b {RELEASE_NAME} upstream/master
   ```

2. Ensure that the `prow/RELEASE_VERSION` file from the `test-infra` repository on a release branch contains the correct version to be created. The file should contain a release version following the `{A}.{B}.{C}` or `{A}.{B}.{C}-rc{D}` format, where `A`,`B`, `C`, and `D` are numbers. If you define a release candidate version, a pre-release is created.  

   Make sure the `RELEASE_VERSION` file includes just this single line:  

   ```bash
   echo -n {RELEASE_VERSION} > prow/RELEASE_VERSION
   ```

3. Push the branch to the `test-infra` repository.

4. Create a PR to `test-infra/release-x.y`. This triggers the pre-release job for `watch-pods`. 

   > **NOTE:** To trigger the `watch-pods` build without introducing any changes, edit any file within the `test-infra` repository and create a pull request. You don't need to merge it as the job is triggered anyway. After a successful `watch-pods` image build, close the pull request.

5. Update the `RELEASE_VERSION` file with the name of the next minor release candidate and merge the pull request to `master`. For example, if the `RELEASE_VERSION` on the `master` branch is set to `1.4.2`, then change the version to `1.5.0-rc1`.

### kyma-project/kyma

1. Inside the release branch do the following changes.

   i. Update your PR with the version and the directory of components used in `values.yaml` files.

   Find these values in the files:

   ```yaml
   dir: develop/
   version: {current_version}
   ```
  
   ```yaml
   dir: pr/
   version: {current_version}
   ```

   Replace them with:

   ```yaml
   dir:
   version: {release_version}
   ```

   > **NOTE:** Replace only `develop/` so `develop/tests` becomes `tests/`.

   Every component image is published with a version defined in the `RELEASE_VERSION` file stored in the `test-infra` repository on the given release branch. Test scripts for integration jobs like GKE Integration or GKE Upgrade are also loaded from the `test-infra` release branch.

   For example, for the first release candidate of 1.4.0, the release version will be `1.4.0-rc1` and the `yaml` files should be modified as follows:

   ``` yaml
   dir:
   version: 1.4.0-rc1
   ```

   > **CAUTION**: Do **not** update the version of components whose `dir` section does not contain `develop`, as is the case with Console-related components. Also do not change octopus version in `kyma/resources/testing/values.yaml` even though the directory is `develop` .

   ii. Check all `yaml` files in the `kyma` repository for references of the following Docker image:

   ```yaml
   image: eu.gcr.io/kyma-project/develop/{IMAGE_NAME}:{SOME_SHA}
   ```

   Change the Docker image to:

   ```yaml
   image: eu.gcr.io/kyma-project/{IMAGE_NAME}:{release_version}
   ```

   > **CAUTION**: In `installation/resources/installer.yaml` replace `eu.gcr.io/kyma-project/develop/installer:{image_tag}` with `eu.gcr.io/kyma-project/kyma-installer:{release_version}`

   iii. In the `resources/core/values.yaml` file, replace the the `clusterDocsTopicsVersion` value with your release branch name. For example, for the 1.4 release, find the following section:

   ```yaml
   docs:
   # (...) - truncated
   clusterDocsTopicsVersion: master
   ```

   And replace the `clusterDocsTopicsVersion` value with the following:

   ```yaml
   docs:
   # (...)
   clusterDocsTopicsVersion: release-1.4
   ```
  
   iv. Ensure that in the `resources/compass/values.yaml` there are no `PR-XXX` values. All image versions should be in a form of commit hashes.

   v. Create a pull request with your changes to the release branch. It triggers all jobs for components.

   ![PullRequest](./assets/release-PR.png)

2. If any job fails, trigger it again by adding the following comment to the PR:

   ```bash
   /test {job_name}
   ```

   > **CAUTION:** Never use `/test all` as it might run tests that you do not want to execute.

3. Wait until all jobs for components and tools finish.
4. Execute remaining tests. The diagram shows you the jobs and dependencies between them.

   ![JobDependencies](assets/kyma-rel-jobs.svg)

   i.  Run `kyma-integration` by adding the  `/test pre-{release_number}-kyma-integration`  comment to the PR.

    > **NOTE:** You don't have to wait until the `pre-{release_number}-kyma-integration` job finishes to proceed with further jobs.

   ii. Run `/test pre-{release_number}-kyma-installer` and wait until it finishes.

   iii. Run `/test pre-{release_number}-kyma-artifacts` and wait until it finishes.
   iv. Run the following tests in parallel:

     ```bash
     /test pre-rel14-kyma-gke-integration
     /test pre-rel14-kyma-gke-upgrade
     /test pre-rel14-kyma-gke-backup
     ```

   iv. Wait for the jobs to finish:
   * `pre-{release_number}-kyma-integration`
   * `pre-{release_number}-kyma-gke-integration`
   * `pre-{release_number}-kyma-gke-upgrade`
   * `pre-{release_number}-kyma-gke-backup`

5. If you detect any problems with the release, such as failing tests, wait for the fix that can be delivered either on a PR or cherry-picked to the PR from the `master` branch. Prow triggers the jobs again. Rerun manual jobs as described in **step 4**.

6. After all checks pass, merge the PR, using the `rebase and merge` option. To merge the PR to the release branch, you must receive approvals from all teams.

   > **CAUTION:** By default, the `rebase and merge` option is disabled. Contact one of the `kyma-project/kyma` repository admins to enable it.

7. Merging the PR to the release branch runs the postsubmit jobs, which:

   * create a GitHub release and trigger documentation update on the official Kyma website
   * trigger provisioning of the cluster from the created release. The cluster name contains the release version with a period `.` replaced by a dash `-`. For example: `gke-release-1-4-0-rc1`. Use the cluster to test the release candidate.

   > **CAUTION**: The cluster is automatically generated for you, so you don't need to create it. The release cluster, the IP Addresses, and the DNS records must be deleted manually after tests on the RC2 cluster are done.

   If you don't have access to the GCP project, post a request in the Slack team channel.

   ```bash
   gcloud container clusters get-credentials gke-release-1-4-0-rc2 --zone europe-west4-c --project sap-hybris-sf-playground
   ```

   Follow [these](https://kyma-project.io/docs/#installation-use-your-own-domain-access-the-cluster) instructions to give Kyma teams access to start testing the release candidate.

8. Update the `RELEASE_VERSION` file to contain the next patch RC1 version on the release branch. Do it immediately after the release, otherwise, any PR to a release branch overrides the previously published Docker images.

   For example, if the `RELEASE_VERSION` file on the release branch contains `1.4.1`, change it to `1.4.2-rc1`.

9. Validate the `yaml` and changelog files generated under [releases](https://github.com/kyma-project/kyma/releases).

10. Update the release content manually with links to the instruction on how to install the latest Kyma release. Currently, this means to grab the links from the previous release and update the version number in URLs. If contributors want you to change something in the instruction, they would address you directly.

11. Create a spreadsheet with all open issues labeled as `test-missing`. Every team assigned to an issue must cover the outstanding test with manual verification on every release candidate. After the test is finished successfully, the responsible team must mark it as completed in the spreadsheet. Every issue identified during testing must be reported. To make the testing easier, provision a publicly available cluster with the release candidate version after you complete all steps listed in this document.

12. Notify Team Breaking Pixels that the release is available for integration with Faros.

> **NOTE:** After the Kyma release is complete, proceed with [releasing Kyma CLI](/release/#kyma-cli-release-process-kyma-cli-release-process).
