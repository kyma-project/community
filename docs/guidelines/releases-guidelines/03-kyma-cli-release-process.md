---
title: Kyma CLI Release Process
label: internal
---

This document describes how to create a Kyma CLI release.

## Release Content

A Kyma CLI release consists of:

* GitHub release with automated changelog generation
* Artifacts, including `cli` binaries and source code archives
* A release tag or branch holding the code

## Steps

1. After Kyma is released, create a release branch in the `cli` repository. The name of this branch must follow the `release-x.y` pattern, such as `release-2.4`.

   >**NOTE:** This step applies only to new major and minor versions.

   ```bash
   git fetch upstream
   git checkout --no-track -b {RELEASE_NAME} upstream/main
   git push upstream {RELEASE_NAME}
   ```

2. Check the `internal/config/config.go` file from the `cli` repository on the release branch:

   * The `DefaultKyma2Version` variable must contain the latest Kyma `2.x.x` version.
   * Run `make docs` to update the CLI help accordingly.

3. Create a PR to `cli/release-x.y` that triggers the presubmit job for `cli`.

4. After merging the PR, create a tag on the release branch that has the value of the new `cli` release version. If you define a release candidate version, a pre-release is created.

   ```bash
   git tag -a {RELEASE_VERSION} -m "Release {RELEASE_VERSION}"
   ```

    Replace {RELEASE_VERSION} with the new `cli` release version, for example, `2.0.0-rc1`, `2.0.0-rc2` or `2.0.0`.

5. Create the new version.
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

6. After the release tarball is available on Github, update the Homebrew formula, either automatically or manually:
    * If no changes to the build command are needed, this can be done automatically by the Homebrew CLI:

       ```bash
       brew bump-formula-pr --strict kyma-cli --url https://github.com/kyma-project/cli/archive/{RELEASE_VERSION}.tar.gz
       ```

      It creates a PR to [Homebrew/homebrew-core](https://github.com/Homebrew/homebrew-core) that contains the version update changes to the kyma-cli formula. 
      If all checks in this PR are passed, the rest is handled by Homebrew maintainers.
      If some tests fail, identify the root cause.
      If it's not due to bugs, simply close your PR and re-run the previous `brew bump-formula-pr` command with `--force` mode.
        * If you got the message `fatal: a branch named 'bump-kyma-cli-{RELEASE_VERSION}' already exists`, delete this local branch.
        * If you got the message `Error: You need to bump this formula manually since the new URL and old URL are both: https://github.com/kyma-project/cli/archive/{RELEASE_VERSION}.tar.gz`, restore the change in you local repo. The local repo is under `/usr/local/Homebrew/Library/Taps/homebrew/homebrew-core`(Intel) or `/opt/homebrew/Library/Taps/homebrew/homebrew-core`(Apple Silicon).
    * Alternatively, create a PR to the [kyma-cli Homebrew formula](https://github.com/Homebrew/homebrew-core/blob/master/Formula/k/kyma-cli.rb).

    When a Homebrew maintainer approves your PR, the formula is updated.

7. Verify a day later that the chocolatey formula has been updated: `https://community.chocolatey.org/packages/kyma-cli`
