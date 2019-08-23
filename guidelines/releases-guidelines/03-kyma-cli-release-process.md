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

1. After making sure that Kyma is released, create a release branch in the `cli` repository. The name of this branch should follow the `release-x.y` pattern, such as `release-1.4`.

   >**NOTE:** This step applies only to new major and minor versions.

   ```bash
   git fetch upstream
   git checkout -b {RELEASE_NAME} upstream/master
   git push upstream {RELEASE_NAME}
   ```

2. Ensure that the `KYMA_VERSION` variables in the `Makefile` and `.goreleaser.yml` files from the `cli` repository on the release branch contain the latest Kyma version.

3. Create a PR to `cli/release-x.y` which triggers the presubmit job for `cli`.

4. After merging the PR, create a tag on the release branch that has the same version name as Kyma. If you define a release candidate version, a pre-release is created.  

   ```bash
   git tag -a {RELEASE_VERSION} -m "Release {RELEASE_VERSION}"
   git push upstream {RELEASE_VERSION}
   ```

    where {RELEASE_VERSION} could be `1.4.0-rc1`, `1.4.0-rc2` or `1.4.0`.

5. Pushing the tag triggers the postsubmit job that creates the GitHub release. Validate if the release is available under [releases](https://github.com/kyma-project/cli/releases).
