---
title: API Gateway release process
label: internal
---

This document describes how to create an API Gateway release.

## Release content

An API Gateway release consists of:

* A GitHub release with automated changelog generation
* Artifacts, including `api-gateway` binary and source code archives
* A release tag and branch holding the code

## Steps

1. Create a release branch in the `api-gateway` repository. The name of this branch must follow the `release-x.y` pattern, such as `release-1.2`.

   >**NOTE:** This step applies only to new major and minor versions.

   ```bash
   git fetch upstream
   git checkout --no-track -b {RELEASE_NAME} upstream/main
   git push upstream {RELEASE_NAME}
   ```

3. Create a PR to `api-gateway/release-x.y` that triggers the presubmit job for `api-gateway`.

5. After merging the PR, create a tag on the release branch that has the value of the new `api-gateway` release version.

   ```bash
   git tag -a {RELEASE_VERSION} -m "Release {RELEASE_VERSION}"
   ```

   Replace {RELEASE_VERSION} with the new `api-gateway` release version, for example, `1.2.0`.

6. Create the new version.
   1. Push the tag to trigger a postsubmit job that creates the GitHub release. Check whether the release is available under [releases](https://github.com/kyma-project/api-gateway/releases).

      ```bash
      git push upstream {RELEASE_VERSION}
      ```

   2. Verify the [Prow Status](https://status.build.kyma-project.io/?repo=kyma-project%2Fapi-gateway&type=postsubmit) of the matching revision ({RELEASE_VERSION}).

   3. If the postsubmit job failed, you can re-trigger it by removing the tag from upstream and pushing it again.

      ```bash
      git push --delete upstream {RELEASE_VERSION}
      git push upstream {RELEASE_VERSION}
      ```

7. After the release tarball is available on Github update Kyma resources.

   Create a PR to Kyma and include details on the features released with the new API Gateway version, description example:

   ```
   **Description**

   Changes proposed in this pull request:

   Bump `api-gateway` component to 1.2.0, which includes:
   - Include support of namespaces on APIRule CRD
   - Refactoring and optimizations
   - [Release changelog and info](https://github.com/kyma-project/api-gateway/releases/tag/1.2.0)

   **Related issue(s)**
   https://github.com/kyma-project/api-gateway/issues/12345
   ```

   In this PR update the release version in the [Chart](https://github.com/kyma-project/kyma/blob/main/resources/api-gateway/Chart.yaml) and the image versions in [Chart values](https://github.com/kyma-project/kyma/blob/main/resources/api-gateway/values.yaml)
