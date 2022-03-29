---
title: Prow workflow
---

This document describes the Prow workflow that we use across all Kyma repositories. This includes basic principles, explanations and commands that interact with [@kyma-bot](https://github.com/kyma-bot).

Prow is a Kubernetes based CI/CD system. Jobs can be triggered by various types of events and report their status to many services.
In addition to job execution, Prow provides GitHub automation in the form of policy enforcement, ChatOps via `/foo` style commands, and automatic PR merging.

Prow enforces strict ownership policies based on the OWNERS files that reside on each organisation repository.
The Prow bot user checks if the Issue has required labels assigned, automatically assigns labels to the Pull Requests,
ensures that the Pull Request follows the 2-person approval flow, automatically merges PRs with approved state and many more.

The configuration of production Prow instance can be found under kyma-project/test-infra [prow directory](https://github.com/kyma-project/test-infra/tree/main/prow).

If you see any unusual behaviour in Prow, [raise an Issue](https://github.com/kyma-project/test-infra/issues/new) in the test-infra repository, or contact kyma-project/test-infra maintainers.

## OWNERS

The OWNERS file is the Prow feature that allows you to create an ownership structure across the entire repository.
This file consists of the list of `approvers`, `reviewers`, and `labels`.

The basic structure of this file is as follows:
```yaml
reviewers:
  - githubUser1
  - githubUser2
  - githubUserN 
approvers:
  - githubUser1
  - githubUser2
  - githubUserN
labels: # optional   
  - label1
  - label2
```

Having the same set of reviewers and approvers is considered a bad practice and should be avoided!
Reviewers look for general code quality, correctness, sane software engineering, style, etc.
Approvers look for holistic acceptance criteria, including dependencies with other features, forwards/backwards compatibility, API and flag definitions, etc.
When a person is in both reviewers and approvers groups, their Approving review is also considered as an approve and marks the PR as `approved`. That behaviour skips the requirement for a second review and immediately allows the PR to be merged.

You can use multiple OWNERS files across the entire repository to diversify the ownership of specific directories between multiple people.

For more information, refer to the [Kubernetes documentation](https://github.com/kubernetes/community/blob/master/contributors/guide/owners.md) regarding OWNERS file as well as the [ANNOUNCEMENTS.md](https://github.com/kubernetes/test-infra/blob/master/prow/ANNOUNCEMENTS.md) file on the kubernetes/test-infra repository.

## OWNERS_ALIASES

The OWNERS_ALIASES file contains a list of aliases that can be used in the OWNERS files, instead of the GitHub usernames.
It is useful when the same set of people is responsible for bigger amount of directories in the repo.

The example structure of this file is as follows:
```yaml
aliases:
  owners-group:
    - githubUser1
    - githubUser2
    - githubUserN
```

For more information, refer to the official [Kubernetes documentation](https://github.com/kubernetes/community/blob/master/contributors/guide/owners.md#owners_aliases).

## Code review process

Since in Kyma Prow is responsible for ensuring the review process, our flow is the same as in the official [Kubernetes documentation](https://github.com/kubernetes/community/blob/master/contributors/guide/owners.md#the-code-review-process).

## Commands

Prow provides GitHub automations and ChatOps using `/foo` styled slash-commands. You can use these command by creating a comment under the GitHub Issue or Pull Request.
This flow can be extended by various built-in plugins and can be further extended with external integrations with Prow API that it provides. 
As part of the workflow every developer needs to know a basic set of commands that they can use for managing Issues or Pull Requests.

Prow plugins implement huge set of command everyone, or specific people can use.

Below there is a basic set of commands that need to be known by every person that wants to contribute to Kyma:

|command|example|description|used by|plugin|
|---|---|---|---|---|
|`/close`|`/close`|Close an issue or PR.|authors or members of organisation|lifecycle|
|`/reopen`|`/reopen`|Reopen an issue or PR.|authors or members of organisation|lifecycle|
|`/[remove-]lifecycle [stale/frozen/rotten/active]`| `/lifecycle frozen`, `/remove-lifecycle stale`|Command used to operate on lifecycle labels respected by [kyma-stale-bot](https://github.com/apps/kyma-stale-bot).|anyone|lifecycle|
|`/[un][remove-]hold [cancel]`|`/hold`, `/unhold`, `/remove-hold` `/hold cancel`|Adds or removes the `do-not-merge/hold` Label which is used to indicate that the PR should not be automatically merged.|anyone|hold|
|`/auto-cc`|`/auto-cc`|Request reviews based on the OWNERS files if those were not assigned.|anyone|blunderbuss|
|`/[un]cc [@username]`|`/cc @Ressetkk`, `/uncc`|Request review from the specific person, or yourself.|anyone|assign|
|`/[un]assign [@username]`|`/assign` `/unassign @Ressetkk`|(Un)assign a person from issue or PR.|anyone|assign|
|`/[remove-](area,kind,priority,label,language) [name]`|`/area prow`, `/label Epic`, `/kind bug` `/remove-area ci`|Applies or removes a label from one of the recognized types of labels.|Anyone can trigger this command on issues and PRs. `triage/accepted` can only be added by org members. Restricted labels are only able to be added by teams and users in their configuration.|label|
|`/test (all,test-name)`|`/test all`, `/test pre-test-infra-build`|Manually starts a/all automatically triggered test job(s). Lists all possible job(s) when no jobs/an invalid job are specified.|anyone|trigger|
|`/test ?`|`/test ?`|List available test job(s) for a trusted PR.|anyone|trigger|
|`/retest`|`/retest`|Rerun test jobs that have failed.|anyone|trigger|

For more information about which commands are available for Kyma Prow instance, please refer to the [Command Help](https://status.build.kyma-project.io/command-help) page.
