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

Define branch protection rules that include enforcing obligatory review and approval of pull requests (PRs), and define which GitHub workflows need to pass successfully before merging PR changes into the `main` branch.

To see these settings, go to **Branches** in the left menu, under repository **Settings**:

![Branch protection rules](./assets/branch-protection-rules.png)

## Update CLA Assistant Configuration

Ask a [kyma-project owner](https://github.com/orgs/kyma-project/people) to:

- Add the newly created repository to the [Contributor License Agreement](https://cla-assistant.io/) (CLA).
- Add the `kyma-bot` username to be exempt from signing the CLA.

## Enable Markdown Link Check

The `/kyma-project` repositories in GitHub use [md-check-link](https://github.com/kyma-project/md-check-link) to check their Markdown files for broken links. Configuration and maintenance of the Markdown link check tool in particular repositories is the responsibility of a repository owner.

### Configuration

#### Repository Config

To configure the md-check-link in your repository, choose your CI/CD pipeline for the check and set up its workflow. For example, choose GitHub Action and add a configuration YAML file to the `/.github/workflows` directory. Paste the following content:

```yaml
name: Verify markdown links

on:
  pull_request:
    branches:
      - "main"
      - "release-*"
  workflow_dispatch:

jobs:
  verify-links:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout repo
        uses: actions/checkout@v4
      - name: Install node.js
        uses: actions/setup-node@v4
        with:
          node-version: '20.x'
      - name: Install md-check-link
        run: npm install -g md-check-link
      - name: Verify links
        run: |
          md-check-link -q -n 8 -c https://raw.githubusercontent.com/kyma-project/md-check-link/main/.mlc.config.json ./
```

With that configuration, the md-check-link verifies all `.md` files in your repository on every PR.

#### Central Config

All Kyma modules' repositories that store user-facing documentation must be included in the central configuration of [md-check-link](https://github.com/kyma-project/md-check-link). It allows Technical Writers to use a nightly GitHub Actions workflow, called [Verify markdown links in Kyma project](https://github.com/kyma-project/md-check-link/actions/workflows/check-kyma-links.yml), that detects broken links in all `.md` files of the listed repositories.

If you create a new module repository, you must add it to the following files:

* [`md-check-link/.github/workflows/check-kyma-links.yml`](https://github.com/kyma-project/md-check-link/blob/main/.github/workflows/check-kyma-links.yml)
* [`md-check-link/.mlc.config.json`](https://github.com/kyma-project/md-check-link/blob/main/.mlc.config.json)

## Custom Settings

Track all repository changes that deviate from configuration standard described in the guidelines with an [issue](https://github.tools.sap/kyma/test-infra/issues/new?assignees=&labels=config-change&template=bug_report.md&title=).
