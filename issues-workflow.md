## Overview

In the Kyma project, we use GitHub Issues for tracking development process, and [ZenHub](https://zenhub.com) to manage the issues on a team and sprint level and to have a clear overview of the work across all Kyma repositories.

This document explains:
- how the issues and pull requests workflows are organized in the Kyma project
- how issues triage is organized
- which tools are used on every stage of the workflow

## Used labels

Our statement is to:
* use default labels provided by GitHub
* introduce new labels only if necessary

### Default labels

The default labels provided by GitHub are as follows:

![](assets/default_labels.png)

### Custom labels

Here are the custom labels introduced by the Kyma team:

* `WIP` indicates that an issue is already in progress.
* `decision` indicates that an issue is related to a decision.
* `priority/critical` indicates the top-priority of a given issue.
* `area/{CAPABILITY_NAME}` indicates which capabilities are related to a given issue. You can assign more than one `area` label an issue.
* `security/{SCORE}` indicates the security issue based on the [CVSSv3](https://www.first.org/cvss/calculator/3.0) security rating.
* `sig/{SIG_NAME}` indicates which [Special interest group (SIG)](./sig-and-wg/README.md) identified the issue and is responsible for further follow-up on the issue.
* `wg/{WG_NAME}` indicates which [Working group (WG)](./sig-and-wg/README.md) identified the issue and is responsible for further follow-up on the issue.

## Issues triage

Here is the flow diagram explaining how issues triage is performed:

![](assets/kyma-triage.svg)

There are five different stages of the triage:

| Stage | Description | Labels |
|--------- |----------|---------|
| Validity | Assess the validity of the issue (whether it is taken for the further triage and proper classification). | invalid, duplicate, wontfix, question |
| Kind | Differentiate whether the related issue is a new feature or a bug. | enhancement, bug |
| Decision | Check if the issue is related to a [decision](governance.md#the-decisions-making-process) | decision |
| Help | Identify issues that do not have high priority and can be taken by the community. | help wanted, good first issue|
| SIG/WG | Clarify which SIG or WG is involved in this issue and is responsible for the further follow-up on the issue. | sig/{SIG_NAME}, wg/{WG_NAME}|
| Security | Specify the [CVSSv3](https://www.first.org/cvss/calculator/3.0) security rating with the support of the security team. | security/{SCORE} |
| Priority | Prioritize issues in the general Kyma backlog to select those which are the most critical and should be taken as first. | priority/critical |
| Area | Clarify which capabilities are involved in a given issue. | area/{CAPABILITY_NAME} |

## Backlog

The Kyma backlog contains issues that went through the triage, are not closed, and have labels added (except for the issues with the `question` label). Backlog prioritization is realized by assigning issues to Kyma milestones (ZenHub Release) and assigning the `priority/critical` label. Critical issues assigned to the current milestone have the highest priority.

>**NOTE**: Issues are taken from the main Kyma backlog by different teams that are responsible for specific areas of Kyma. This is the actual workflow in Kyma which allows distinguishing which team works on a specific issue. This approach allows easy work in team sprints using ZenHub board. We are aware that our teams' names may be cryptic for the external community, but at the moment we are not able to provide any better solution. If you have a better idea, your feedback would be highly appreciated. Sorry for the inconvenience.

### Team sprints
Team Sprint is modeled as a GitHub milestone named with the following pattern: `{TEAM_NAME} Sprint {NUMBER}`. During the planning, a team selects issues from the backlog considering:
- priority (high priority first)
- area (default capability of the team first)
- dependencies (unblock others)

Contributors mark issues with the `WIP` label when they start working on it. When the work is done, the issue is closed.

### Team backlog (optional)
If any team wants to keep team backlog (assign issues they want to work on later), they can assign it to the Github milestone with no due date named with following pattern: `{TEAM_NAME} backlog`.

## ZenHub links

* Install the [ZenHub Chrome plugin](https://chrome.google.com/webstore/detail/zenhub-for-github/ogcgkffhplmphkaahpmffcafajaocjbd)
* Login to [ZenHub](https://www.zenhub.com/)
* Open the [Kyma ZenHub workspace](https://app.zenhub.com/workspace/o/kyma-project/kyma)
