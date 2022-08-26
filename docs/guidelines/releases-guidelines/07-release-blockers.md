---
title: Release blockers
label: internal
---

This document describes best practices for handling `release-blockers`. We would like to improve how `release-blockers` are handled hence we came up with an idea of short guidelines that will be added to each `release-blocker` issue.

### How to use the template

When you label an issue as a `release-blocker` please copy below template and add it as a comment. Don't forget to update the first section: 
> Explain clearly why this is a release blocker (what/who/how is impacted):  `.....`

### Template
---

**This issue was marked as ⛔ `release-blocker`**
Explain clearly why this is a release blocker (what/who/how is impacted):  `.....`

<details>
<summary style="font-size:14px">Please follow below guidelines</summary>
<p>

Assignment:

- Release blockers can't be left unassigned.
- If during SOS or a PO call the SMs or POs discover that it is **unclear who should be assigned**, they should contact @kyma-scrum-masters on #kyma-skr-release - the SMs will then set up a call with team representatives to determine the ownership of the issue and next steps. Regardless of your role, if you discover such an assigned release blocker yourself, you should also report it to the SMs on #kyma-skr-release.

Investigation:

- Split the release blocker into smaller issues if it concerns more teams or more than one bug or error type appears.
- If the most effective way of working on that issue is to have a group call - use #kyma-team to organize it or ask @kyma-scrum-masters to do that.

Time-box:

- Update release-blockers every day with valuable information about progress and planned delivery.
- If you are not able to deliver fix **before branch cut-off** contact Release Managers one business day before release cut-off before noon.
- When issue happens **after branch cut-off** update it by 1PM every day.

Escalation:

- Inform Release Managers about no progress or existing blockers.

Transparency:

- Link any PRs that are related to the release blocker.
- If the release blocker is blocked by other issues, make that clear in a comment with a valid links. Also, link any known issues that are blocked by this release blocker.

</p></details>

---
