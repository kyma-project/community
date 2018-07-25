# Release Notes
Aiming towards continuous delivery, manual intervention has to be reduced as far as possible. Some things, such as writing release notes, cannot be completely automated and depend on humans writing for humans. Following a similar approach from other open source projects, we compile the release notes from the pull requests where the changes are introduced.

## Write a release-note in the PR

All pull requests __must__ include a release-note as part of their description. This way, it is ensured that the author of the pull request makes a conscious decision about the relevance of the change(s) for the end user. If there are user-visible changes, such as new features or bug fixes, then a corresponding note is written.

To meet this requirement, add notes in the release notes block:

    ```release-note
    [Change type]: Your release note here
    ```

The __change type__ must be one of the following:

* __Added__ for new features
* __Changed__ for changes in existing functionality
* __Deprecated__ for soon-to-be removed features
* __Removed__ for now removed features
* __Fixed__ for any bug fixes
* __Security__ in case of vulnerabilities

If the pull request has no user-visible changes and no release note is required, simply write "none" (case insensitive) inside the release note:

    ```release-note
    NONE
    ```
