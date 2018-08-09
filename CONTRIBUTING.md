## Overview

The purpose of this document is to give guidance to the Kyma contributors. All the rules are listed and explained here, and all contributors and maintainers must follow them.

## Naming guidelines

To learn how to name objects in this project, see the [naming](/guidelines/naming.md) document.

## Documentation types

The main [`README.md`](/guidelines/templates/README.md) document in the [`templates`](/guidelines/templates/) subfolder in this repository contains an overview of document templates used in specific Kyma repositories. The templates themselves are collected in the [`resources`](/guidelines/templates/resources/) subfolder.

Extend the list whenever you define a new template for other document types. Make sure to update the table in the [`README.md`](/guidelines/templates/README.md) document after you add new templates to the [`resources`](/guidelines/templates/resources/) subfolder.

## Agreements and licenses

Read the subsections to learn the details of the agreements to submit and licences to comply with as a Kyma contributor.

### Individual contributor license agreement

As a Kyma contributor, you must accept the Kyma project's licenses and submit the
[Individual Contributor License Agreement](https://gist.github.com/CLAassistant/bd1ea8ec8aa0357414e8) before you contribute code or content to any Kyma repository. Kyma maintainers will not accept contributions made without such a consent. This applies to all contributors, including those contributing on behalf of a company. If you agree to the content of the Agreement, click on the link posted by the CLA assistant as a comment to the pull request. Click it to review the CLA and accept it on the next screen if you agree to it. The CLA assistant will save your decision for the upcoming contributions and will notify you if there is any change to the CLA in the meantime.

### Corporate contributor license agreement

Employees of a company who contribute code need to submit one company agreement in addition to the individual agreement above. This is mainly for the protection of the contributing employees.

An authorized company representative needs to download, fill in, and print
the [Corporate Contributor License Agreement](https://github.com/kyma-project/community/blob/master/docs/cla/SAP%20Corporate%20Contributor%20License%20Agreement%20(5-26-15).pdf) form. Scan it and send it to [info@kyma-project.io](mailto:info@kyma-project.io). The form contains a list of employees who are authorized to contribute on behalf of your company. To report any changes on the list, contact [info@kyma-project.io](mailto:info@kyma-project.io).

## Contribution rules

If you are a contributor, follow these basic rules:

* The contribution workflow in all Kyma repositories bases on the principles of the [GitHub Flow](https://guides.github.com/introduction/flow/). Thus, the `master` branch is the most important one. Avoid working directly on it. When you work on new features or bug fixes, work on separate branches.
* Work on forks of Kyma repositories.
* Squash and rebase every pull request before merging.
* You can merge a pull request if you receive an approval from at least one code owner from each part of the repository to which you contribute in your pull request.

Every contributor commits to the following agreement:

* In each pull request, include a description or a reference to a detailed description of the steps that the maintainer goes through to check if a pull request works and does not break any other functionality.
* Provide clear and descriptive commit messages.
* Follow the squash and rebase process.
* Follow the accepted documentation rules and use appropriate templates.
* Choose the proper directory for the content of your pull request, depending on the role of the component, and how it relates to the project structure.
* As the creator of the pull request, you are responsible for ensuring that the pull request follows the correct review and approval flow.
* Make sure that the person who owns the documentation, such as a technical writer, reviews the changes in the documentation before you merge the pull request.

## Contribution process

This section explains how you can contribute code or content to any Kyma repository, propose an improvement, or report a bug. The contributing process applies both to the members of the Kyma organization and the external contributors.

### Contribute code or content

To contribute code or content to a given Kyma repository, follow these steps:

1. Make sure that the change is valid and approved. If you are an external contributor, **open a GitHub issue** before you make a contribution.
2. Fork the Kyma repository that you want to contribute to.
3. Clone it locally, add a remote upstream repository for the original repository, and set up the `master` branch to track the remote `master` branch from the upstream repository. See the [git-workflow](git-workflow.md) document for details on fork configuration.
4. Create a new branch out of the local `master` branch of the forked repository.
5. Commit and push changes on your new branch. Create a clear and descriptive final commit message in which you specify what you have changed.
6. Create a pull request from your branch on the forked repository to the `master` branch of the original, upstream repository. Fill in the pull request template according to instructions.
7. Read and accept the Contributor Licence Agreement (CLA).
8. If there are merge conflicts on your pull request, squash your commits and rebase the `master` branch.
9. On the pull request, select the [**Allow edits from maintainers**](https://help.github.com/articles/allowing-changes-to-a-pull-request-branch-created-from-a-fork/) option to allow upstream repository maintainers and those with the push access to the upstream repository, to commit to your forked branch.
10. If your change relates to any existing GitHub issue, provide a link to it in your pull request.
11. Wait for the Kyma maintainers to review and approve your pull request. The maintainers can approve it, request enhancements to your change, or reject it.

> **NOTE:** The reviewer must check if the related CI job and tests have completed successfully before approving the pull request.

12. When the maintainers approve your change, merge the pull request. If you are an external contributor, contact the repository maintainers specified in the `CODEOWNERS` file to do the merge.

Read the [git-workflow](git-workflow.md) document. It describes the Kyma contribution workflow that relies on forks, branches, rebasing, and squashing. The document also contains guidelines for writing commit messages.

### Report an issue

If you find a bug to report or you want to propose a new feature, go to the GitHub issue tracker of a given repository and create an issue. If you are not certain which repository your bug or feature relates to, raise it on the `kyma` repository.

> **NOTE:** The repository maintainers handle only well-documented, valid issues that have not been reported yet. Before you create one, check if there are no duplicates. Provide all details and include examples. When you report a bug, list the exact steps necessary to reproduce it.

After you report an issue, the maintainers of a given repository either ask for more details required to work with the issue you reported or close the issue if it is not valid or cannot be fixed at the moment.

## Maintenance rules

Every maintainer reviews each contribution according to the rules listed in this document.

Although it is the responsibility of the owner of the pull request to ensure that the maintainers review and approve the pull request, maintainers need to coordinate the overall number of unreviewed and unapproved pull requests in their queue, and, if required, take appropriate measures to handle them effectively.

To learn more about maintainers' responsibilities and rules for appointing new maintainers, and removing the existing ones, refer to the [governance.md](governance.md) document.

## Owners

To identify the owners of particular parts of your repository, see the `CODEOWNERS` file in the root directory.
