## Overview

This document defines the ownership policy and the decision-making process within the [Kyma](../../../) organization.

## Scope

All repositories in the Kyma organization must follow the official guidelines, contributing rules, and the governance process to assure quality and consistency.

The Kyma project also includes the [Kyma Incubator](https://github.com/kyma-incubator) organization. It is a place where all new projects start in a more relaxed environment that facilitates their rapid growth. At that stage, they do not have to comply with all rules that govern the Kyma organization. Once the incubating project is ready to become a part of the main Kyma organization, adjust it to all standards.

## Ownership policy

Kyma repositories are owned by code owners who are a group of people with special privileges in the repositories of the [GitHub](../../../) organization. Each repository has a separate `CODEOWNERS` file located at its root. The file specifies persons who have the ability to approve contribution to the part of the repository they own, after a detailed review of the related pull requests (PRs). Although the name suggests only the code ownership, the `CODEOWNERS` file is not only about the code but the content in general. Apart from the developers, you can define any relevant parties as code owners. For example, technical writers are set up as the owners of all `.md` documents in the Kyma repositories and SIG/WG members are the owners of their SIG/WG's folders' content.

### Code owners' responsibilities

With great power comes great responsibility. Code owners not only review and approve PRs but also truly care about their projects.

Every code owner is expected to:

* Contribute high-quality code and content
* Communicate and collaborate with other code owners to improve the ownership process
* Perform a thorough review of incoming PRs and make sure they follow the [contributing rules](CONTRIBUTING.md)
* Approve only those PRs in which the contributor made the requested improvements
* Check if the related CI job has completed successfully before approving the PR
* Make sure that the PR approval flow runs smoothly
* Prioritize issues and manage their processing
* Proactively fix bugs
* Perform maintenance tasks for their projects

### Add or remove a code owner

To suggest a change in the ownership of a given repository part, create a PR with the required changes in the `CODEOWNERS` file in the project's repository. The required number of code owners needs to approve the PR for the changes to take place. Read [here](https://github.com/kyma-project/community/blob/master/guidelines/internal-guidelines/repository-template/template/CODEOWNERS) how to set up and modify owners of the given repository folders and files.

## The decisions-making process

In general, the relevant Special Interest Groups (SIGs) and Working Groups (WGs) make the decisions that affect the project, including its structure, functionalities, components, or work of the project teams. However, the organizational decisions and those that relate to the product strategy belong to the Kyma Council.

SIGs and WGs follow the "lazy consensus" decision-making process which assumes that:

* All SIG/WG members have an equal voice in the decision-making process.
* Silence is consent. By default, lack of objections to a proposed decision means a silent approval.
* Any objections are good opportunities for healthy and constructive discussions.

> **NOTE:** The described approach only concerns the decisions made by SIGs and WGs. It does not affect any Kyma decisions made during daily team activities.

### Process

The SIG/WG decison-making process looks as follows:

![Decision-making process](assets/decision-making-process.png)

1. Every decision starts with creating a PR with the [**Decision Record**](./guidelines/templates/resources/DR.md) (DR) document that provides the decision details, timelines, status, the related context, and associated consequences. The DR can reference a proposal or any other related document.

> **NOTE:** If a few proposals need a single decision, create only one Decision Record and update it with links to all related proposals. 

2. The PR creator sends the table with the decision log details to the related SIG/WG mailing list. If the decision input providers are not part of the mailing list, add them to the email communication.
3. The discussion on the proposed decision can happen through the mailing list, on the related Slack channel, or directly on the PR. You can also add the topic to the agenda of the upcoming SIG/WG meeting and discuss it with the SIG/WG members. Encourage the discussion and bring up any objections early in the process. A person with a different approach to the topic can create a counter proposal.
4. The discussions on the proposed decision can lead to a change in the actual decision log or the proposal, or end up with no changes required. Update the PR with the outcome of the discussion. Otherwise, take actions required to reach a consensus. Those who created the proposal work with those who had the objections to prepare an improved solution or reach a consensus to decline the proposal.
5. If you reach the "lazy consensus" by the decision-making date, update the status of the decision record with either `Accepted` or `Declined`, and merge the PR.
6. If there are still unresolved objections by the decision due date, merge the PR with the `Proposed` status. The ultimate decision makers defined in the decision log table make the final decision.
7. After the final decision, update the decision log to either `Accepted` or `Declined`.

>**NOTE:** Only the ultimate decision makers can request to revisit the decision before the revision date. To do it, raise an issue in the `community` repository with all details and the reason why you want to discuss the decision earlier than agreed.
