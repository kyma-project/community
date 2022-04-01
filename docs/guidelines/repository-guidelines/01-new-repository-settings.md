---
title: New repository settings
label: internal
---

All repositories in `kyma-project` and `kyma-incubator` organizations should be similar in structure, settings, and restrictions. Follow these guidelines to adjust settings of a new repository created in one of these organizations.

> **NOTE:** You have to be an owner of the given organization to create a new repository in it.

## Use the repository template

When you create a new repository use the template `kyma-project/template-repository`. This template contains all necessary configuration files for OWNERS, Kyma stale bot, issue and pull request templates, license and code of conduct.

> **NOTE:** Do not check *Include all branches* option!

![Template](./assets/template.png)

If you are interested in what's inside the template, check the [template repository code](https://github.com/kyma-project/template-repository).

## Adjust repository options

Under the repository name, choose the **Settings** tab. The **Options** view opens as the default one in the left menu.

1. Scroll down to the **Features** section and clear these options:
- Wikis
- Restrict editing to users in teams with push access only
- Projects

![Features](./assets/features.png)

## Set branch protection rules

Define branch protection rules that include enforcing obligatory review and approval of pull requests (PRs), and define which Prow jobs need to pass successfully before merging PR changes into the `main` branch.

To see these settings, go to **Branches** in the left menu, under repository **Settings**:

![Branch protection rules](./assets/branch-protection-rules.png)

In Kyma, the protection rules are defined in the Prow [`config.yaml`](https://github.com/kyma-project/test-infra/blob/main/prow/config.yaml) file generated from rules defined in the [`prow-config.yaml`](https://github.com/kyma-project/test-infra/blob/main/templates/templates/prow-config.yaml) file and handled by a Prow component called [Branch Protector](https://github.com/kyma-project/test-infra/blob/main/docs/prow/prow-architecture.md#branch-protector).

If you add a new repository in:
- `kyma-project`, you do not need to add a new entry to the Prow `config.yaml` file as the branch protection is already defined for [all repositories](https://github.com/kyma-project/test-infra/blob/main/prow/config.yaml#L380) within this organization. The only exception is if you want to specify additional rules that are not handled by Prow.
- `kyma-incubator`, add a new repository entry to the Prow `config.yaml` file, under **branch-protection.orgs.kyma-incubator.repos**. See [an example](https://github.com/kyma-project/test-infra/blob/main/templates/templates/prow-config.yaml)  of such an entry for the `marketplaces` repository.

## Update CLA assistant configuration

Ask a [kyma-project owner](https://github.com/orgs/kyma-project/people) to add the newly created repository to the [Contributor License Agreement](https://cla-assistant.io/) (CLA).

## Add a milv file

If you define any governance-related [Prow job](https://github.com/kyma-project/test-infra/blob/main/prow/jobs/) for the new repository to validate documentation links, you must add a `milv.config.yaml` file at the root of the repository. [See](https://github.com/kyma-project/test-infra/blob/main/milv.config.yaml) an example of the milv file.

## Create labels

Labels are managed by Prow in the [`labels.yaml`](https://github.com/kyma-project/test-infra/blob/main/prow/labels.yaml) configuration file. To add additional labels, simply raise a Pull Request to the test-infra repository with new label definition. Please follow the structure of this file to add a label.

Additionally, you can [define repository exclusive labels](https://help.github.com/en/articles/creating-a-label) for the new repository so you could use them in issues and pull requests. Follow the naming convention and color array used in other repositories such as [`kyma`](https://github.com/kyma-project/kyma/labels).
