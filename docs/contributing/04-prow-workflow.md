---
title: Prow workflow
---

This document describes the Prow workflow that we use across all Kyma repositories. This includes basic principles, explanations, and commands that interact with [@kyma-bot](https://github.com/kyma-bot).

Prow is a Kubernetes-based CI/CD system.
In addition to job execution, Prow provides GitHub automation in the form of policy enforcement, ChatOps through the `/foo` style commands, and automatic merging of pull requests (PRs). The Prow bot checks if issues have the required labels assigned, automatically assigns the labels to pull requests,
ensures that pull requests follow the two-person approval flow, automatically merges PRs with approved state, and many more.

Prow defines the ownership of a repository and its directories based on the mandatory OWNERS file and the optional OWNERS_ALIASES file. Read the [OWNERS file](#owners-file) and [OWNERS_ALIASES file](#owners_aliases-file) sections to learn more.

You can find the configuration of the production Prow instance in the [`test-infra/prow`](https://github.com/kyma-project/test-infra/tree/main/prow) directory. If you see any misbehavior in Prow, [create an issue](https://github.com/kyma-project/test-infra/issues/new) in the `test-infra` repository, or contact the repository maintainers.

## OWNERS file

The OWNERS file is the Prow feature that allows you to create the ownership structure across the entire repository.
This file consists of the list of approvers, reviewers, and labels.

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
- Reviewers look for general code quality, correctness, sane software engineering, style, etc.
- Approvers look for holistic acceptance criteria, including dependencies to other features, forwards/backwards compatibility, API and flag definitions, etc.

When a person is in both reviewers and approvers groups, their approving review marks the PR as `approved`, which skips the requirement for a second review and immediately allows the PR to be merged.

You can use multiple OWNERS files across the entire repository to diversify the ownership of specific directories between multiple people.

For more information, refer to the [Kubernetes documentation](https://github.com/kubernetes/community/blob/master/contributors/guide/owners.md) regarding the OWNERS file as well as the [ANNOUNCEMENTS.md](https://github.com/kubernetes/test-infra/blob/master/prow/ANNOUNCEMENTS.md) file in the `kubernetes/test-infra` repository.

## OWNERS_ALIASES file

The OWNERS_ALIASES file contains a list of aliases that can be used in the OWNERS files instead of the GitHub usernames.
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

## Adding or removing people in OWNERS files

If you wish to add yourself as a reviewer or approver of the directory, raise a PR with the updated OWNERS file. The current owners of the directory will then review your PR based on the following requirements:

* You need to be a member of `kyma-project`, or `kyma-incubator` organization.
* You need to actively contribute to the directory you want to be the reviewer for.
* If you want to add yourself as an approver, you need to have a deep understanding of the code structure and how it is working. Thus you need to be a reviewer for some time.

Once you meet all the requirements above, your Pull Request will be approved and corresponding changes will be added to the OWNERS file.

The same formula goes for OWNERS_ALIASES, however you need to ask repository adminitrators for approval, once the specific area members approve your application.

If you want to remove yourself from the OWNERS file, simply open a PR where you remove yourself from the file. If you were an approver then add yourself to the [`emeritus_approvers`](https://github.com/kubernetes/community/blob/master/contributors/guide/owners.md#emeritus) section of the OWNERS.

Additionally add yourself to the [`emeritus.md`](https://github.com/kyma-project/community/blob/main/emeritus.md) file with a new PR. Keep the file structure consistent when you add yourself to it.

## Code review process

Since the review process in Kyma is based on Prow, our flow is the same as described in the official [Kubernetes documentation](https://github.com/kubernetes/community/blob/master/contributors/guide/owners.md#the-code-review-process).

## Commands

Prow provides GitHub automations and ChatOps using `/foo` styled slash-commands. You can use these commands by adding a comment under a GitHub issue or pull request.


This is a basic set of commands you need to know to manage issues and PRs in Kyma:

|Command|Example|Description|Used by|Plugin|
|---|---|---|---|---|
|`/close`|`/close`|Close an issue or PR.|Authors and members of the organization |lifecycle|
|`/reopen`|`/reopen`|Reopen an issue or PR.|Authors and members of the organization |lifecycle|
|`/[remove-]lifecycle [stale/frozen/rotten/active]`| `/lifecycle frozen`, `/remove-lifecycle stale`|Command used to operate on lifecycle labels respected by [Kyma Stale Bot](https://github.com/apps/kyma-stale-bot).|anyone|lifecycle|
|`/[un][remove-]hold [cancel]`|`/hold`, `/unhold`, `/remove-hold` `/hold cancel`|Adds or removes the `do-not-merge/hold` label used to indicate that the PR should not be automatically merged.|anyone|hold|
|`/auto-cc`|`/auto-cc`|Requests review based on the OWNERS files if the reviewers were not assigned.|anyone|blunderbuss|
|`/[un]cc [@username]`|`/cc @Ressetkk`, `/uncc`|Requests review from the specific person. You can also use it to assign yourself as a reviewer.|anyone|assign|
|`/[un]assign [@username]`|`/assign` `/unassign @Ressetkk`|(Un)assigns a person from an issue or PR.|anyone|assign|
|`/[remove-](area,kind,priority,label,language) [name]`|`/area prow`, `/label Epic`, `/kind bug` `/remove-area ci`|Applies or removes a label from one of the recognized types of labels.|Anyone can trigger this command on issues and PRs. `triage/accepted` is a restricted label and can only be added by organization members. Only users that belong to at least one of the configured teams can use the restricted labels.|label|
|`/test (all,test-name)`|`/test all`, `/test pre-test-infra-build`|- Manually starts all automatically triggered test jobs. <br> - Lists all possible jobs when no jobs or invalid jobs are specified.|anyone|trigger|
|`/test ?`|`/test ?`|Lists available test jobs for a trusted PR.|anyone|trigger|
|`/retest`|`/retest`|Rerun test jobs that have failed.|anyone|trigger|

For more information about commands available for the Kyma Prow instance, refer to the [Command Help](https://status.build.kyma-project.io/command-help) page.
