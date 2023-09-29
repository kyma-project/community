# GitHub Actions

- [GitHub Actions](#github-actions)
  - [Introduction](#introduction)
  - [Restrictions](#restrictions)
  - [Allowed Github Actions](#allowed-github-actions)
  - [GitHub Actions Security Best Practices](#github-actions-security-best-practices)
    - [Security review](#security-review)
    - [Set the minimal scope for credentials](#set-the-minimal-scope-for-credentials)
    - [Do NOT use `pull_request_target` event](#do-not-use-pull_request_target-event)
    - [Treat event context data as untrusted input](#treat-event-context-data-as-untrusted-input)
    - [Use Dependabot](#use-dependabot)
    - [Review Action changes](#review-action-changes)
  - [Additional information](#additional-information)

## Introduction

GitHub Actions are a powerful tool to automate almost every task of the development cycle. The ecosystem of GitHub Actions is growing very fast and one can find a prepared Action for a lot of activities. But like all other usages of Open Source Software, GitHub Actions must be handled carefully because an action is running external code on our code in our GitHub repositories.

## Restrictions

The usage of external GitHub Actions is restricted centrally on the `github.com` organization level. Only Actions in the kyma-project org or approved external Actions are allowed. Actions created by GitHub or [verified creators](https://github.com/marketplace?type=actions&verification=verified_creator) are also allowed.

![Github Actions policy](./assets/gh-actions-policies.png)

GitHub Actions, in particular the GITHUB_TOKEN, have `read` and `write` permissions on repositories but they are not allowed to create and merge pull requests.

![Github Actions permissions](./assets/gh-actions-permissions.png)

## Allowed Github Actions

List of allowed GitHub Actions: [https://github.com/kyma-project/community/docs/contributing/assets/allowed_actions.json](./assets/allowed_actions.json)

To add a Github Action to the list of allowed Actions one has to perform a [security review](#security-review) of the Github Action. After the review was performed one has to create a pull request on the list of allowed Github Actions. Use the template of the json object below to add the relevant information to the `github_actions` array in the file [allowed_actions.json](./assets/allowed_actions.json). The comment entry can be used to add information which can be usefull for others who want to use the same Action.

```json
{
    "name": "full name including Github organization",
    "versions": ["hash digest of the release"],
    "repository": "full link to the repository",
    "marketplace": "full link to the Github Marketplace entry",
    "security_review_performed": true or false,
    "3rd_party_tool": {
        "tool" : "name",
        "pinned_version" : "version tag / hash",
        "repository" : "full link to the repository"},
    "comment" : ""
}
```

By adding the Action to the list and use Github Action one gives an implicit commitment to the best practices listed below.

## GitHub Actions Security Best Practices

Even though the usage of GitHub Actions is restricted, some threats remain. The responsibility to target these threats is on the developer using GitHub Actions.

### Security review

The code of the action has to be reviewed to identify suspicious parts of the code. It should be especially checked if the Action is processing secrets beside the expected usage and what kind of modification the Action does on Github resources.

### Set the minimal scope for credentials

- GITHUB_TOKEN has read and write permissions
- Must be restricted to the least minimum
- Also valid for all other credentials used by the Github Action

### Do NOT use `pull_request_target` event

- Avoid using `pull_request_target` if the workflow doesn't need write repository permissions and doesn't use any repository secrets. They can simply use the `pull_request` trigger instead.
- Assign repository privileges only where needed explicitly through pull_request and workflow_run
  - Handle untrusted code via the `pull_request` trigger so that it is isolated in an unprivileged environment.This workflow started with the `pull_request` trigger should then store any results like code coverage or failed/passed tests in artifacts and exit.
  - The following workflow then starts on `workflow_run` where it is granted write permission to the target repository and access to repository secrets.
  - Find an example in this [blog post](https://securitylab.github.com/research/github-actions-preventing-pwn-requests/)
- If you really need to use the trigger `pull_request_target`, add a condition to the `pull_request_target` to run only if a certain label is assigned the PR, like safe to test that indicates the PR has been vetted by someone with write privileges to the target repository.

### Treat event context data as untrusted input

- Every workflow trigger is provided with a [GitHub context](https://docs.github.com/en/actions/learn-github-actions/contexts#github-context)
  - Contains information about the triggering event
  - This context data has to be treated as untrusted data
- If your Action needs some input data
  - Set the untrusted input value of the expression to an intermediate environment variable

```yaml
- name: print title
  env:
    TITLE: ${{ github.event.issue.title }}
  run: echo "$TITLE"
```

### Use Dependabot

Use Dependabot to regularly update the Github Action as described on the help page [Keeping your actions up to date with Dependabot](https://docs.github.com/en/code-security/dependabot/working-with-dependabot/keeping-your-actions-up-to-date-with-dependabot).

### Review Action changes

Only a small group of people, usually the repository admins, should be able to review and approve Github Actions. The following line has to be added to the CODEOWNERS file of the repository and has to be completed with the users / the team allowed to review and approve PRs regarding Github Actions.

```text
.github/workflows    @user-1 @user-2 @user-x
```

or

```text
.github/workflows    @org/team
```

This entry can be omitted as long as the CODEOWNERS file contains a default owner entry that fits the needs described above.

## Additional information

The information on this page was collected from other resources to give comprehensive guidelines to work with Github Actions. Further information on Github Actions Security can be found on the following pages:

- [GitGuardian: A blog article on how to secure GitHub Actions](https://blog.gitguardian.com/github-actions-security-cheat-sheet/)
- [Security hardening for GitHub Actions](https://docs.github.com/en/actions/security-guides/security-hardening-for-github-actions#using-third-party-actions)
- [Keeping your GitHub Actions and workflows secure Part 1: Preventing pwn requests](https://securitylab.github.com/research/github-actions-preventing-pwn-requests/)
- [Keeping your GitHub Actions and workflows secure Part 2: Untrusted input](https://securitylab.github.com/research/github-actions-untrusted-input/)
- [Keeping your GitHub Actions and workflows secure Part 3: How to trust your building blocks](https://securitylab.github.com/research/github-actions-building-blocks/)
