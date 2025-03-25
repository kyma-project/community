# GitHub Actions

<!-- TOC tocDepth:2..3 chapterDepth:2..6 -->

- [Introduction](#introduction)
- [GitHub Actions Security Best Practices](#github-actions-security-best-practices)
  - [Perform Security Review](#perform-security-review)
  - [Set the Minimal Scope for Credentials](#set-the-minimal-scope-for-credentials)
  - [Do NOT Use `pull_request_target` Event](#do-not-use-pull_request_target-event)
  - [Treat Event Context Data as Untrusted Input](#treat-event-context-data-as-untrusted-input)
  - [Regularly update the Github Actions](#regularly-update-the-github-actions)
  - [Review Changes on Workflows](#review-changes-on-workflows)
- [Additional Information](#additional-information)

<!-- /TOC -->

## Introduction

GitHub Actions is a powerful tool for automating almost every task of the development cycle. The ecosystem of GitHub Actions is growing very fast, and one can find a prepared Action for many activities. But like all other usages of open source software, GitHub Actions must be handled carefully because an action runs external code on our code in our GitHub repositories. The following guidelines are intended to help you use GitHub Actions securely.

## GitHub Actions Security Best Practices

### Perform Security Review

The code of the action must be reviewed to identify suspicious parts of the code. Pay special attention to whether the Action is processing secrets besides the expected usage and what kind of modification the Action does on GitHub resources.

### Set the Minimal Scope for Credentials

By default only `read` permissions are granted to the `GITHUB_TOKEN`. The `GITHUB_TOKEN` is an automatically generated secret that lets you make authenticated calls to the GitHub API in your workflow runs. The permissions of the `GITHUB_TOKEN` have to be [limited to the least minimum the `GITHUB_TOKEN` needs in your workflow](<https://docs.github.com/en/actions/security-guides/automatic-token-authentication#modifying-the-permissions-for-the-github_token>).

You must also ensure, that all secrets used during the workflow follow the so-called least minimum principle. This means that these credentials may only have the permissions they need to carry out their task.

### Do NOT Use `pull_request_target` Event

- Avoid using `pull_request_target` if the workflow doesn't need `write` repository permissions and doesn't use any repository secrets. Just use the `pull_request` trigger instead.
- Assign repository privileges only where needed explicitly through `pull_request` and `workflow_run`.
  - Handle untrusted code via the `pull_request` trigger so that it is isolated in an unprivileged environment. This workflow started with the `pull_request` trigger should then store any results like code coverage or failed/passed tests in artifacts and exit.
  - The following workflow then starts on `workflow_run` where it is granted `write` permission to the target repository and access to repository secrets.
  - Find an example in this [blog post](https://securitylab.github.com/research/github-actions-preventing-pwn-requests/).
- If you really need to use the `pull_request_target` trigger, add a condition to the `pull_request_target` to run only if a certain label is assigned the PR, such as `safe-to-test` that indicates the PR has been vetted by someone with `write` privileges to the target repository.

### Treat Event Context Data as Untrusted Input

- Every workflow trigger is provided with a [GitHub context](https://docs.github.com/en/actions/learn-github-actions/contexts#github-context) that contains information about the triggering event. You must treat this context data as untrusted data.
- If your Action needs some input data, set the untrusted input value of the expression to an intermediate environment variable.

```yaml
- name: print title
  env:
    TITLE: ${{ github.event.issue.title }}
  run: echo "$TITLE"
```

### Regularly update the Github Actions

Use [Dependabot](https://docs.github.com/en/code-security/dependabot/working-with-dependabot/keeping-your-actions-up-to-date-with-dependabot) or [Renovate](https://docs.renovatebot.com/) to regularly update the GitHub Actions in use.

### Review Changes on Workflows

Only a small group of people, usually the repository administrators, should be able to review and approve GitHub Actions. Add the list of users/teams allowed to review and approve PRs regarding GitHub Actions to the CODEOWNERS file of your repository.

```text
.github/workflows    @user-1 @user-2 @user-x
```

or

```text
.github/workflows    @org/team
```

This entry can be omitted as long as the CODEOWNERS file contains a default owner entry that fits the needs described above.

## Additional Information

The information on this page was collected from other resources to give comprehensive guidelines to work with GitHub Actions. Further information on GitHub Actions Security can be found on the following pages:

- [GitGuardian: A blog article on how to secure GitHub Actions](https://blog.gitguardian.com/github-actions-security-cheat-sheet/)
- [Security hardening for GitHub Actions](https://docs.github.com/en/actions/security-guides/security-hardening-for-github-actions#using-third-party-actions)
- [Keeping your GitHub Actions and workflows secure Part 1: Preventing pwn requests](https://securitylab.github.com/research/github-actions-preventing-pwn-requests/)
- [Keeping your GitHub Actions and workflows secure Part 2: Untrusted input](https://securitylab.github.com/research/github-actions-untrusted-input/)
- [Keeping your GitHub Actions and workflows secure Part 3: How to trust your building blocks](https://securitylab.github.com/research/github-actions-building-blocks/)
