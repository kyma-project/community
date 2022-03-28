---
title: Prow workflow
---

This document describes the Prow workflow that we use across all Kyma repositories. This includes basic principles, explanations and commands that interact with [@kyma-bot](https://github.com/kyma-bot).

## Terminology

**Prow** - Prow is a Kubernetes based CI/CD system. Jobs can be triggered by various types of events and report their status to many different services. In addition to job execution, Prow provides GitHub automation in the form of policy enforcement, ChatOps via `/foo` style commands, and automatic PR merging.

## Configuration

## OWNERS and repository ownership

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
