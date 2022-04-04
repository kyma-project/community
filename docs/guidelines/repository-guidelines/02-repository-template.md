---
title: Repository template
label: internal
---

The [`template-repository`](https://github.com/kyma-project/template-repository) offers a unified file, document, and folder structure. Use it for every new repository that you create in Kyma. It helps you to ensure that the project is consistent and standardized.

## Usage

The `template-repository` contains all elements required for a skeleton repository. However, before you create a new repository from this template, read carefully the following paragraph to learn what the purpose of the specific files and documents is and which of them you need to adjust.

The `template-repository` consists of:

* [`.github`](https://github.com/kyma-project/template-repository/tree/main/.github) - This folder contains the pull request template, issue templates, and the Stale Bot that monitors inactive issues, marks them as `stale`, and closes them after the specified period of time.

* [`docs`](https://github.com/kyma-project/template-repository/tree/main/docs) - In this folder, put the repository-specific documentation only. Store any architectural decisions or documents applicable to all Kyma repositories in the `community` repository.

* [CODE_OF_CONDUCT.md](https://github.com/kyma-project/template-repository/blob/main/CODE_OF_CONDUCT.md) - This document is a ready-to-use template which provides a link to the general `CODE_OF_CONDUCT.md` document from the `community` repository.

* [OWNERS](https://github.com/kyma-project/template-repository/blob/main/OWNERS) - In this document, specify the owners of particular parts of your repository. The owners are automatically added as reviewers when you open a pull request that modifies the code and content they own.

* [OWNERS_ALIASES](https://github.com/kyma-project/template-repository/blob/main/OWNERS_ALIASES) - This file contains the aliases that group the GitHub usernames. Use it when you want to use group names for granting ownership across the repository.

* [CONTRIBUTING.md](https://github.com/kyma-project/template-repository/blob/main/CONTRIBUTING.md) - This template makes a reference to the [contributing rules](../../contributing/02-contributing.md) that contain the general guidance from the `community` repository and describes the rules for contributing to all Kyma repositories. If there is any additional, project-specific information that you want to add to your project's `CONTRIBUTING.md` document, add them under the same sections as in the general contributing guide.

* [LICENSE](https://github.com/kyma-project/template-repository/blob/main/LICENSE) - It is an obligatory element of every open-source repository. Copy the template into your repository.

* [NOTICE.md](https://github.com/kyma-project/template-repository/blob/main/NOTICE.md) - The document defines the ownership of the copyright in the repository. Copy the template into your repository.

* [README.md](https://github.com/kyma-project/template-repository/blob/main/README.md) - This is a template with sections that you fill in according to the provided suggestions. Add any information specific for a development guide in this document. Describe how your project works, how to use it, and how to develop it. Because all sections are optional, remove those that do not apply to your project.  
