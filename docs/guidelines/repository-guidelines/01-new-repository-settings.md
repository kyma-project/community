---
title: New Repository Settings
label: internal
---

All repositories in `kyma-project` and `kyma-incubator` organizations should be similar in structure, settings, and restrictions. Follow these guidelines to adjust settings of a new repository created in one of these organizations.

> **NOTE:** You have to be an owner of the given organization to create a new repository in it.

## Use the Repository Template

Whenever you create a new repository, use the template from the [`template-repository`](https://github.com/kyma-project/template-repository). This template contains all necessary configuration files for CODEOWNERS, Kyma stale bot, issue and pull request templates, license, and Code of Conduct.

> **NOTE:** Do not mark the **Include all branches** checkbox as it will clone all branches from the `template-repository` to your new repo.

![Template](./assets/template.png)

## Adjust Repository Options

Under the repository name, choose the **Settings** tab. The **Options** view opens as the default one in the left menu.

1. Scroll down to the **Features** section and clear these options:
- Wikis
- Restrict editing to users in teams with push access only
- Projects

![Features](./assets/features.png)

## Set Branch Protection Rules

Define branch protection rules that include enforcing obligatory review and approval of pull requests (PRs), and define which Prow jobs need to pass successfully before merging PR changes into the `main` branch.

To see these settings, go to **Branches** in the left menu, under repository **Settings**:

![Branch protection rules](./assets/branch-protection-rules.png)

In Kyma, the protection rules are defined in the Prow [`config.yaml`](https://github.com/kyma-project/test-infra/blob/main/prow/config.yaml) file handled by a Prow component called [Branch Protector](https://github.com/kyma-project/test-infra/blob/main/docs/prow/prow-architecture.md#branch-protector).

If you add a new repository in:
- `kyma-project`, you do not need to add a new entry to the Prow `config.yaml` file as the branch protection is already defined for [all repositories](https://github.com/kyma-project/test-infra/blob/main/prow/config.yaml#L380) within this organization. The only exception is if you want to specify additional rules that are not handled by Prow.
- `kyma-incubator`, add a new repository entry to the Prow `config.yaml` file, under **branch-protection.orgs.kyma-incubator.repos**.

## Update CLA Assistant Configuration

Ask a [kyma-project owner](https://github.com/orgs/kyma-project/people) to:
- Add the newly created repository to the [Contributor License Agreement](https://cla-assistant.io/) (CLA).
- Add the `kyma-bot` username to be exempt from signing the CLA.

## Enable Markdown Link Check

The `/kyma-project` repositories in GitHub use [Markdown link check](https://github.com/tcort/markdown-link-check) to check their Markdown files for broken links. Configuration and maintenance of the Markdown link check tool is the responsibility of a repository owner.

### Configuration

To configure the Markdown link check, follow these steps:

1. Add or update the `.mlc.config.json` file with the following parameters in the root directory of your repository:

  ```bash
  {
    "replacementPatterns": [
      {
        "_comment": "a replacement rule for all the in-repository references",
        "pattern": "^/",
        "replacement": "{{BASEURL}}/"
      }
    ]
  }
  ```

  See the following examples:
  
  - [`/community/.mlc.config.json`](https://github.com/kyma-project/community/blob/main/.mlc.config.json)
  - [`/telemetry-manager/.mlc.config.json`](https://github.com/kyma-project/telemetry-manager/blob/main/.mlc.config.json)

2. Choose your CI/CD pipeline for the check and set up its workflow. For example, choose GitHub Action and add or update the configuration YAML file(s) to the `/.github/workflows` directory. See the [official GitHub Action - Markdown link check documentation](https://github.com/marketplace/actions/markdown-link-check) for details.

  See the following examples of the GitHub Action configuration for the `/telemetry-manager` and `/kyma` repositories:

  -  `/telemetry-manager`
     - [`markdown-link-check.yml`](https://github.com/kyma-project/telemetry-manager/blob/main/.github/workflows/pr-docu-checks.yml#L23) - checks links in all Markdown files in the repository on every pull request.
  - `/kyma`
     - [`lint-markdown-links-pr.yml`](https://github.com/kyma-project/kyma/blob/main/.github/workflows/lint-markdown-links-pr.yml) - checks links in Markdown files being part of a created pull request.
     - [`lint-markdown-links-daily.yml`](https://github.com/kyma-project/kyma/blob/main/.github/workflows/lint-markdown-links-daily.yml) - checks links in all Markdown files in the repository. This is a periodic, daily check scheduled on the main branch at 5 AM.

## Custom Settings

Track all repository changes that deviate from configuration standard described in the guidelines with an [issue](https://github.tools.sap/kyma/test-infra/issues/new?assignees=&labels=config-change&template=bug_report.md&title=).
