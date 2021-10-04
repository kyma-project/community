---
title: Kyma CLI release process
label: internal
---

This document describes how to create a Kyma CLI release.

## Release content

A Kyma CLI release consists of:

* GitHub release with automated changelog generation
* Artifacts, including `cli` binaries and source code archives
* A release tag or branch holding the code

## Steps
   >**NOTE:** Once `cli 2.0.0` is released, `cli 1.x.x` should not be released any more. If a new `kyma 1.24.x` is released after releasing `cli 2.0.0`, then a corresponding `cli` patch version (`2.x.x`) should be released.
1. After Kyma is released, create a release branch in the `cli` repository. The name of this branch must follow the `release-x.y` pattern, such as `release-1.4`.

   >**NOTE:** This step applies only to new major and minor versions.

   ```bash
   git fetch upstream
   git checkout --no-track -b {RELEASE_NAME} upstream/main
   git push upstream {RELEASE_NAME}
   ```
2. Check the `internal/config/config.go` file from the `cli` repository on the release branch:
   - The `DefaultKyma1Version` variable must contain the latest Kyma `1.24.x` version.
   - Starting with Kyma 2.0.0, the `DefaultKyma2Version` variable must contain the latest Kyma `2.x.x` version.


3. Create a PR to `cli/release-x.y` that triggers the presubmit job for `cli`.

5. After merging the PR, create a tag on the release branch that has the value of the new `cli` release version. If you define a release candidate version, a pre-release is created.  

   ```bash
   git tag -a {RELEASE_VERSION} -m "Release {RELEASE_VERSION}"
   ```

    Replace {RELEASE_VERSION} with the new `cli` release version, for example, `2.0.0-rc1`, `2.0.0-rc2` or `2.0.0`.

6. Create the new version.
   1. Push the tag to trigger a postsubmit job that creates the GitHub release. Check whether the release is available under [releases](https://github.com/kyma-project/cli/releases).

      ```bash
      git push upstream {RELEASE_VERSION}
      ```

   2. Verify the [Prow Status](https://status.build.kyma-project.io/?repo=kyma-project%2Fcli&type=postsubmit) of the matching revision ({RELEASE_VERSION}).
   3. If the postsubmit job failed, you can re-trigger it by removing the tag from upstream and pushing it again.

      ```bash
      git push --delete upstream {RELEASE_VERSION}
      git push upstream {RELEASE_VERSION}
      ```
