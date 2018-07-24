# Release process

Read this document to learn how to create tags and releases in Kyma component repositories.

## Create a tag

Use tags to mark new functionalities in the repository.
You can create a tag using either the terminal or the GitHub UI.

> **NOTE:** Create the tags manually if you do not have an automated process in place to tag the `master` branch on merge.

### Terminal

Follow these steps to create a tag in the repository using the terminal:

1. Merge the pull request that introduces a new functionality, checkout the `master` branch of the repository and pull the latest changes.
2. If there are already releases on your repository, run `git tag` to list all existing tags and see the latest one.
3. The Kyma organization uses [lightweigh tags](https://git-scm.com/book/en/v2/Git-Basics-Tagging#_lightweight_tags). To create a new tag, run `git tag {tagname}`, where {tagname} stands for the consecutive release version you want the tag to represent. For example, `v2.3.14`.

> **NOTE:** When you create a new tag, follow the [semantic versioning](https://semver.org/) naming schema.

4. The `git tag {tagname}` command creates the tag locally. To push it to the remote repository, run the `git push origin {tagname}` command.

5. The tag is applied to the `master` branch of the given repository.

### GitHub UI

To create a tag in the GitHub UI, you must create and publish a release first.
The **Tag version** box in which you enter the new tag name appears when you start writing a draft of the release.

See the [**Create a release**](#create-a-release) section for more details.

To learn more about the tagging process in general, read the [GitHub documentation](https://git-scm.com/book/en/v2/Git-Basics-Tagging).

## Create a release

Releases base on tags. After you create a tag, you can add a release to it to describe the changes that the given tag introduces.

> **NOTE:** You need to have write access to the repository to create or edit releases, and to view drafts of releases.

Follow these steps to create a release:

1. Go to the `https://github.com/kyma-project/{repositoryname}/releases` page.
2. Select **Draft a new release**.

> **NOTE:** If there are no releases in this repository yet, the **Create a new release** button appears instead.

3. Select the tag version from the drop-down menu or create a new one using the [semantic versioning](https://semver.org/) naming schema. For example, `v2.3.14`.

4. Enter the **Release title**.

5. Provide a description of the release in the **Description** box.

> **NOTE:** It is a good practice to include the content of the well-described pull requests merged before the release into the release notes description. See the [git-workflow](../../git-workflow.md) document for rules on how to write descriptive commit titles and bodies.

6. If your release needs additional files, add them manually or use the drag-and-drop method.

> **NOTE:** There is no limit to the overall size of all combined files in the release. However, a single file must be under 2 GB in size.

7. Select the **This is a pre-release** check box if the release is unstable and not ready for the production yet.
8. Select **Publish release** to publish the release or **Save draft** to work on it later.

> **NOTE:** Before you publish the release, ask another contributor to review it.

To learn more about the release process in general, read the [GitHub documentation](https://help.github.com/categories/releases/).
