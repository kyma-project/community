---
title: Kyma release process
label: internal
---

This document describes how to create a Kyma release. Start from defining release jobs as described in the [**Preparation**](#preparation) section. Then proceed to the [**Steps**](#steps).

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

> **NOTE:** This section applies only to new major and minor versions. If you release a patch, skip the preparation and go to the [**Steps**](#steps) section.

### kyma-project/kyma

#### Perform initial checks

Check if the `main` branch contains any PR-images:

   ```bash
   git grep -e 'version:\s.*[Pp][Rr]-.*' -e 'image:.*:[Pp][Rr]-.*' -e 'tag:\s.*[Pp][Rr]-.*' --before-context=2  resources tests
   ```

   Ask the teams for fixes if this command returns any output.

#### Create a release branch

>**NOTE:** a release branch needs to be created per new major / minor version. Patch releases and release candidates do not have a dedicated release branch.

Create a release branch in the `kyma` repository. The name of this branch should follow the `release-{major}.{minor}` pattern, such as `release-1.4`.

    ```bash
    git fetch upstream
    git checkout --no-track -b release-{RELEASE} upstream/main
    git push -u upstream release-{RELEASE}
    ```

## Steps

Follow these steps to release another Kyma version. Execute these steps for every patch release or release candidate.

#### Create a PR to the release branch

1. Change the `RELEASE_VERSION`. Make sure the `VERSION` file includes just a single line, **without the newline character at the end**:  

    ```bash
    echo -n {RELEASE_VERSION} > VERSION
    ```

2. Create a pull request with your changes to the release branch.

   ![PullRequest](./assets/release-PR.png)

3. If `pre-release-pr-image-guard` fails, ask the owners to change PR-XXX images of the components to the `main` version.
4. If the checks are green, merge the PR and proceed to the next step.

#### Development process towards the release
   > **NOTE:** Every developer who is introducing changes to the specific version can perform steps 1-4.

1. Create a feature-branch based on the given `release-{RELEASE}` branch you want to extend. Add your changes and create a Pull Request.
   Usually, a fix is done first on the `main` branch and then backported into the `release-{RELEASE}` branch. Make sure that not only the PR images are updated but also the changes in the code are included in the `main` and `release-{RELEASE}` branch.

2. Once you create a Pull Request to the release branch, the set of checks is triggered.
   These jobs run in the same way as jobs that run on every Pull Request to the `main` branch.
   If you create a Pull Request that contains changes to the components, the component-building job is triggered.
   If you make any changes in the charts, the integration tests are triggered.

3. If you detect any problems with your PR, fix the issues until your checks pass.

4. After all checks pass, merge your PR to the release branch. Merging the PR triggers the post-submit integration tests automatically.
   The jobs' status will be visible on the Kyma [TestGrid](https://testgrid.k8s.io/kyma_integration) in the corresponding dashboard tab.

5. If there's a need for additional changes in the release branch during the development process, open a new PR to the release branch.
   Repeat steps 1-4 for this PR.

#### Create a release

1. Once the preparation for the release is finished, trigger the [Release Kyma](https://github.com/kyma-project/kyma/actions/workflows/github-release.yaml) GitHub action. 
   Choose the branch that corresponds to the release that you want to trigger. The exact release version is taken from the `VERSION` file.
   When you click the **Run workflow** button, the release process waits for the approval from reviewers.
   <!-- markdown-link-check-disable-next-line -->
   The reviewers list is defined in the ["release" Github Environment](https://github.com/kyma-project/kyma/settings/environments).
   After it is approved, the following will happen:
   * GitHub release is triggered.
   * Documentation update on the official Kyma website is triggered.
   * New release cluster is created for the given Kyma `RELEASE_VERSION`.
     If you don't have access to the GCP project, post a request in the Slack team channel.
     > **CAUTION**: The cluster is automatically generated for you, and it is automatically removed after 7 days.
<!-- markdown-link-check-disable-next-line -->
2. The Github release post-submit job creates a release in the `kyma-project/kyma` repository, which triggers the [`post-rel{RELEASE_VERSION_SHORT}-kyma-release-upgrade`](https://github.com/kyma-project/test-infra/blob/main/prow/jobs/kyma/kyma-release-upgrade.yaml) pipeline. The purpose of this job is to test upgradability between the latest Kyma release that is not a release candidate and the brand new release published by the release post-submit job.
    For example, if `1.7.0-rc2` is released, the pipeline will try to upgrade `1.6.0` to `1.7.0-rc2`.

    If you detect any problems with the upgrade, contact the teams responsible for failing components.

    > **CAUTION:** The job assumes no manual migration is involved. If the upgrade process requires any additional actions, the pipeline is likely to fail. In such case, the owners of the components concerned are responsible for running manual tests or modifying the pipeline.

3. Validate the `yaml` and changelog files generated under [releases](https://github.com/kyma-project/kyma/releases).

4. Update the release content manually with links to:

   * Instructions on local Kyma installation
   * Instructions on cluster Kyma installation
   * Release notes

   For installation instructions, use the links from the previous release and update the version number in URLs. If contributors want you to change something in the instructions, they would address you directly. Contact technical writers for the link to release notes.

> **NOTE:** After the Kyma release is complete, proceed with [releasing Kyma CLI](./03-kyma-cli-release-process.md).
