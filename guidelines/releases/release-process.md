# Release Process

This document describes how to create a Kyma release using Prow.

## Preparation

This section only applies to new major and minor versions. If you release a patch, skip the preparation and go to the [**Steps**](#steps) section.

To prepare a release, define new and remove old release jobs.

### Define release jobs

Define release jobs on the `master` branch in the `test-infra` repository. To ensure every job name is unique, prefix it with `pre-rel{versionNumber}`. Remember to provide the version number without any periods. To learn how to define a release job for a component, read the following [document](https://github.com/kyma-project/test-infra/blob/master/docs/prow/release-jobs.md). To find all jobs for the 0.9 release, look for job names with the `pre-rel09` prefix.

1. Navigate to the `test-infra` repository.
1. Define release jobs in the `prow/jobs/test-infra` directory in the following files:
    - `watch-pods.yaml`
1. Define release jobs in the `prow/jobs/kyma` directory in the following files:
   - every `.yaml` in `components` 
   - every `.yaml` in `tests`
   - every `.yaml` in `tools/` **except for**:
     - `tools/docsbuilder/docsbuilder.yaml`
     - `tools/failery/failery.yaml`
   - `kyma-docs.yaml`
   - `kyma-integration.yaml`
   - `kyma-gke-integration.yaml`
   - `kyma-gke-upgrade.yaml`
   - `kyma-artifacts.yaml`
   - `kyma-installer.yaml`
   - `kyma-github-release.yaml`

        > **Note:** Remember to modify the `presets` array. This is an example of the 0.9 release `preset`:

        ```
        - labels:
            preset-target-commit-0.9: "true"
            env:
            - name: RELEASE_TARGET_COMMIT
                value: release-0.9
        ```

1. Ensure that tests for the release jobs exist. Release tests usually iterate through all release versions and run tests for them.
See the `TestBucReleases` test defined in `development/tools/jobs/kyma/binding_usage_controller_test.go` as a reference.

1. Update the `GetAllKymaReleaseBranches()` function
defined in the `development/tools/jobs/tester/tester.go` file under the `test-infra` repository.

1. Define branch protection rules for the release branch in the `prow/config.yaml` file.

    For example, see the release-0.9 definition:

    ```yaml
    release-0.9:
      protect: true
      required_status_checks:
        contexts:
          - pre-rel09-kyma-integration
          - pre-rel09-kyma-gke-integration
          - pre-rel09-kyma-artifacts
          - pre-rel09-kyma-installer
          - pre-rel09-kyma-gke-upgrade
          - pre-rel09-kyma-gke-central-connector
    ```

### Remove previous release jobs

After adding new release jobs, remove the old ones. Remember to leave jobs for three latest releases. For example, during the preparation for the 0.9 release, add `pre-rel09` jobs and remove all `pre-rel06` jobs. Make sure that the only defined jobs are those with `pre-rel07`, `pre-rel08`, and `pre-rel09` prefixes.

## Steps

> **NOTE:** Currently, you are most likely required to push to the `release-x.y` branch more than once. Make sure that a user with **admin** role is present in the `kyma` repository. 

Before a release, make sure there is no mismatch between source code and `.yaml` files. All components are rebuilt from source code, which requires all Helm charts to be up to date.

Follow these steps to create a release:

### kyma-project/test-infra

1. Create a release branch in the `test-infra` repository. The name of this branch should follow the `release-x.y` pattern, such as `release-0.9`.

    ```bash
    git fetch upstream
    git checkout -b $RELEASE_NAME upstream/master
    ```

    >**NOTE:** This point applies only to new major and minor versions.

2. Ensure that the `prow/RELEASE_VERSION` file from the `test-infra` repository on a release branch contains the correct version to be created. The file should contain a release version following the `{A}.{B}.{C}` or `{A}.{B}.{C}-rc{D}` format, where `A`,`B`, `C`, and `D` are numbers. If you define a release candidate version, a pre-release is created.  

    Make sure the `RELEASE_VERSION` file includes just this single line:  
    ```bash
    echo -n {RELEASE_VERSION} > prow/RELEASE_VERSION
    ```

3. Push the branch to the `test-infra` repository.

4. Create a PR to `test-infra/release-x.y`. This triggers the pre-release job for `watch-pods`. 

    > **NOTE:** To trigger the `watch-pods` build without introducing any changes, edit any file within the `test-infra` repository and create a pull request. You don't need to merge it as the job is triggered anyway. After a successful `watch-pods` image build, close the pull request.

5. Update the `RELEASE_VERSION`file with the name of the next minor release candidate and merge the pull request to `master`. For example, if the `RELEASE_VERSION` on the `master` branch is set to `0.9.2`, then change the version to `1.0.0-rc1`.

### kyma-project/kyma

1. Create a release branch in the `kyma` repository.

    Do it only for major and minor releases.
    The name of this branch should follow the `release-x.y` pattern, such as `release-0.9`.

    ```bash
    git fetch upstream
    git checkout -b release-{release_version} upstream/master
    git push upstream release-{release_version}
    ```

    > **NOTE:** If you push anything to the release branch, or if you rebase the branch, a new Github release is created.

2. Create a new branch and do the following changes.
    
    1. Update your PR with the version and the directory of components used in `values.yaml` files.

        Find these values in the files:

        ```yaml
        dir: develop/
        version: {current_version}
        ```

        Replace them with:

        ```yaml
        dir:
        version: {release_version}
        ```

        > **NOTE:** Replace only `develop/` so `develop/tests` becomes `tests/`.

        Every component image is published with a version defined in the `RELEASE_VERSION` file stored in the `test-infra` repository on the given release branch. Test scripts for integration jobs like GKE Integration or GKE Upgrade are also loaded from the `test-infra` release branch.

    1. Check all `yaml` files in the `kyma` repository for references of the following Docker image:

        ```yaml
        image: eu.gcr.io/kyma-project/develop/{IMAGE_NAME}:{SOME_SHA}
        ```

        Change the Docker image to:
        
        ```yaml
        image: eu.gcr.io/kyma-project/{IMAGE_NAME}:{release_version}
        ```

        > **CAUTION**: In `installation/resources/installer.yaml` replace `eu.gcr.io/kyma-project/develop/installer:{image_tag}` with `eu.gcr.io/kyma-project/kyma-installer:{release_version}`
    
    1. In  the `resources/core/values.yaml` file, replace the `clusterDocsTopicsVersion` value with your release branch name. For example, for the 0.9.1 release, find the following section:

        ```yaml
        docs:
            # (...) - truncated
            clusterDocsTopicsVersion: master
        ```

        And replace the `clusterDocsTopicsVersion` value with the following:

        ```yaml
        docs:
            # (...)
            clusterDocsTopicsVersion: release-0.9
        ```
    1. Create a pull request with your changes to the release branch. It triggers all jobs for components.

        ![PullRequest](../../assets/release-PR.png)

3. If any job fails, trigger it again by adding the following comment to the PR:

    ```;
    /test {job_name}
    ```

    > **CAUTION:** Never use `/test all` as it might run tests that you do not want to execute.

4. Wait until all jobs for components and tools finish.
5. Execute remaining tests. The diagram shows you the jobs and dependencies between them.
    ![JobDependencies](../../assets/kyma-rel-jobs.svg)
    1. Run `kyma-integration` by adding the  `/test pre-{release_number}-kyma-integration`  comment to the PR.

        > **NOTE:** You don't have to wait until the `pre-{release_number}-kyma-integration` job finishes to proceed with further jobs.

    1. Run `/test pre-{release_number}-kyma-installer` and wait until it finishes.
    1. Run `/test pre-{release_number}-kyma-artifacts` and wait until it finishes.
    1. Run `/test pre-{release_number}-kyma-gke-integration` and `/test pre-{release_number}-kyma-gke-upgrade`. You can start them in parallel.
    1. Wait for the jobs to finish:
         - `pre-{release_number}-kyma-integration`
         - `pre-{release_number}-kyma-gke-integration`
         - `pre-{release_number}-kyma-gke-upgrade`

6. If you detect any problems with the release, such as failing tests, wait for the fix that can be delivered either on a PR or cherry-picked to the PR from the `master` branch.  
    Prow triggers the jobs again. Rerun manual jobs as described in **step 5**.

7. After all checks pass, merge the PR.
    > **NOTE:** To merge the PR to the release branch, you must receive approvals from all teams.

8. Merging the PR to the release branch runs the postsubmit job that creates a GitHub release. This job also triggers documentation update on the official Kyma website.

9. Update the `RELEASE_VERSION` file to contain the next patch RC1 version on the release branch. Do it immediately after the release, otherwise, any PR to a release branch overrides the previously published Docker images.

    For example, if the `RELEASE_VERSION` file on the release branch contains `0.9.1`, change it to `0.9.2-rc1`.

10.  Validate the `yaml` and changelog files generated under [releases](https://github.com/kyma-project/kyma/releases).

11. Update the release content manually with links to the instruction on how to install the latest Kyma release.  
    Currently, this means to grab the links from the previous release and update the version number in URLs. If contributors want you to change something in the instruction, they would address you directly.

> **NOTE:** All teams should test the release candidate versions. To make the testing easier, provision a publicly available cluster with the release candidate version after performing all steps listed in this document.
