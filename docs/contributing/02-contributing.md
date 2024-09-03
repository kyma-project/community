# Contributing Rules

As a Kyma contributor, you must follow certain guidelines and rules.

## Guidelines

Go to the **Guidelines** section to read about rules and tips for providing [content](../guidelines/content-guidelines) and [code](../guidelines/technical-guidelines) to the Kyma repositories. Also, learn how to create a new [repository](../guidelines/repository-guidelines), and how the [release process](../guidelines/releases-guidelines) looks in Kyma. Make your life easier using various document types [templates](https://github.com/kyma-project/community/tree/main/templates) prepared for those who would like to contribute.

## Documentation Types

Read about [template types](../guidelines/templates/templates-type.md) used in specific Kyma repositories. The templates themselves are collected in the [`templates/resources`](https://github.com/kyma-project/community/tree/main/templates/resources) subfolder in the `community` repository.

Extend the list whenever you define a new template for other document types. Make sure to update one of the tables in the [**Document types templates**](../guidelines/templates/templates-type.md) document after you add new templates to the `templates/resources` subfolder.

## Developer Certificate of Origin (DCO)

Due to legal reasons, contributors will be asked to accept a DCO before they submit the first pull request to this project. SAP uses the [standard DCO text of the Linux Foundation](https://developercertificate.org/).  
This happens automatically during the submission process: The CLA assistant tool will add a comment to the pull request. The contributer must click it to check the DCO and accept it on the following screen. CLA assistant will save this decision for upcoming contributions.

This DCO replaces the previously used CLA ("Contributor License Agreement") as well as the "Corporate Contributor License Agreement" with new terms that are well-known standards and hence easier to approve by legal departments. Contributors who had already accepted the CLA in the past may be asked to accept the new DCO.

## Guideline for AI-Generated Code Contributions to SAP Open-Source Software Projects

As artificial intelligence evolves, AI-generated code is becoming valuable for many software projects, including open-source initiatives. While we recognize the potential benefits of incorporating AI-generated content into our open-source projects, certain requirements must be reflected and adhered to when making contributions.

When using AI-generated code contributions in open-source software (OSS) projects, their usage must align with OSS values and legal requirements. We have established these essential guidelines to help contributors navigate the complexities of using AI tools while maintaining compliance with open-source licenses and the broader [Open Source Definition](https://opensource.org/osd).

AI-generated code or content can be contributed to SAP OSS projects if the following conditions are met:

1. **Compliance with AI Tool Terms and Conditions**: Contributors must ensure that the AI tool's terms and conditions do not impose any restrictions on the tool's output that conflict with the project's open-source license or intellectual property policies. This includes ensuring that the AI-generated content adheres to the Open Source Definition.
2. **Filtering Similar Suggestions**: Contributors must use features provided by AI tools to suppress responses that are similar to third-party materials or flag similarities. We only accept contributions from AI tools with such filtering options. If the AI tool flags any similarities, contributors must review and ensure compliance with the licensing terms of such materials before including them in the project.
3. **Management of Third-Party Materials**: If the AI tool's output includes pre-existing copyrighted materials, including open-source code authored or owned by third parties, contributors must verify that they have the necessary permissions from the original owners. This typically involves ensuring that there is an open-source license or public domain declaration that is compatible with the project's licensing policies. Contributors must also provide appropriate notice and attribution for these third-party materials, along with relevant information about the applicable license terms.
4. **Employer Policies Compliance**: If AI-generated content is contributed in the context of employment, contributors must also adhere to their employer’s policies. This ensures that all contributions are made with proper authorization and respect for relevant corporate guidelines.

## Contribution Rules

If you are a contributor, follow these basic rules:

* The contribution workflow in all Kyma repositories is based on the principles of the [GitHub flow](https://guides.github.com/introduction/flow/). Thus, the `main` branch is the most important one. Avoid working directly on it. When you work on new features or bug fixes, work on separate branches.
* Work on forks of the Kyma repositories.
* You can merge a PR if you receive an approval from at least one code owner from each part of the repository to which you contribute in your PR.

Every contributor commits to the following agreement:

* In every PR, include a description or a reference to a detailed description of the steps that the maintainer goes through to check if a PR works and does not break any other functionality.
* Provide clear and descriptive commit messages.
* Label your PRs.
* Follow the accepted documentation rules and use appropriate templates.
* As the creator of the PR, you are responsible for ensuring that the PR follows the correct review and approval flow.

## Contribution Process

This section explains how you can contribute code or content to any Kyma repository, propose an improvement, or report a bug. The contributing process applies both to the members of the Kyma organization and the external contributors.

### Contribute Code or Content

To contribute code or content to a given Kyma repository, follow these steps:

1. Make sure that the change is valid and approved. If you are an external contributor, **open a GitHub issue** before you make a contribution.
2. Fork the Kyma repository that you want to contribute to.
3. Clone it locally, add a remote upstream repository for the original repository, and set up the `main` branch to track the remote `main` branch from the upstream repository. See the [**Git Workflow**](./03-git-workflow.md) document to learn how to configure your fork.
4. Create a new branch out of the local `main` branch of the forked repository.
5. Commit and push changes to your new branch. Create a clear and descriptive commit message in which you specify what you have changed. See the [**Git workflow**](./03-git-workflow.md) document for commit message guidelines.
6. Create a PR from your branch on the forked repository to the `main` branch of the original, upstream repository. Fill in the PR template according to instructions.
7. Read and accept the Contributor Licence Agreement (CLA).
8. If there are merge conflicts on your PR, squash your commits and rebase the `main` branch.
9. In your PR:
   * Provide a reference to any related GitHub issue.
   * Make sure that the [**Allow edits from maintainers**](https://help.github.com/articles/allowing-changes-to-a-pull-request-branch-created-from-a-fork/) option is selected to allow upstream repository maintainers, and those with the push access to the upstream repository, to commit to your forked branch.
   * Choose at least one `area/{capability}` label from the available list and add it to your PR to categorize changes you made. Labels are required to include your PR in the `CHANGELOG.md` file and classify it accordingly.
10. After you create a PR, relevant CI tests need to complete successfully.
    * If you are a Kyma organization member, all related CI tests run automatically after you create a PR. If a test fails, check the reason by clicking the **Details** button next to the given job on your PR. Make the required changes and the tests rerun. If you want to run a specific test, add the `/test {test-name}` or `/retest {test-name}` comment to your PR. To rerun all failed tests, add the `/retest` comment.
    * If you are an external contributor, contact the repository maintainers specified in the [`CODEOWNERS`](https://github.com/kyma-project/community/blob/main/CODEOWNERS) file to review your PR and add the `/test all` comment to your PR to trigger all tests. A Kyma organization member needs to rerun the tests manually each time you commit new changes to the PR.
11. As a PR owner, you are responsible for ensuring the success of the submission process, including pre- and postsubmit Prow jobs.  You can ask for more help from the Prow jobs owner only after the initial analysis.
    > **NOTE:** For more information on code owners' responsibilities, see the [Kyma working model](../governance/01-governance.md) document.
12. Wait for the Kyma maintainers to review and approve your PR. The maintainers can approve it, request enhancements to your change, or reject it.  
    > **NOTE:** The reviewer must check if all related CI tests have completed successfully before approving the PR.
13. When the maintainers approve your change, merge the PR or wait until the Kyma Bot merges it for you.

### Report an Issue

If you find a bug to report or you want to propose a new feature, go to the GitHub issue tracker of a given repository and create an issue. If you are not certain which repository your bug or feature relates to, raise it on the `Kyma` repository.

> **NOTE:** The repository maintainers handle only well-documented, valid issues that have not been reported yet. Before you create one, check if there are no duplicates. Provide all details and include examples. When you report a bug, list the exact steps necessary to reproduce it.

See the [**Issues workflow**](../governance/03-issues-workflow.md) document for details on issues triage and processing workflow.

> **NOTE:** The community is relentless about Kyma security. To report a sensitive security issue, send an email with details directly to [kyma-security@googlegroups.com](mailto:kyma-security@googlegroups.com) instead of using a public issue tracker.

## Maintenance Rules

Every maintainer reviews each contribution according to the rules listed in this document.

Although it is the responsibility of the owner of the PR to ensure that the maintainers review and approve the PR, maintainers need to coordinate the overall number of unreviewed and unapproved PRs in their queue, and, if required, take appropriate measures to handle them effectively.

To learn more about maintainers' responsibilities and rules for appointing new maintainers, and removing the existing ones, refer to the [**Kyma working model**](../governance/01-governance.md) document.

## Maintainers

To identify the maintainers of particular parts of your repository, see the [`CODEOWNERS`](https://github.com/kyma-project/community/blob/main/CODEOWNERS) file in the root directory of each Kyma repository.
