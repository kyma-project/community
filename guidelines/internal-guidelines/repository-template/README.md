# Repository Template

## Overview

This repository offers a unified folder, file, and document structure. Use it for every new repository that you create in Kyma. It helps you to ensure that the project is consistent and standardized.

## Usage

The `template` folder contains all elements required for a skeleton repository. However, before you copy the content of the `template` folder into your new repository, read carefully the following paragraph to learn what the purpose of the specific files and documents is and which of them you need to adjust.

The `template` folder consists of:

* [`.github`](./template/.github) - This folder contains pull request and issue templates.

* [`docs`](./template/docs) - In this folder, put the repository-specific documentation only. Store any architectural decisions or documents applicable to all Kyma repositories in the `community` repository.

* [CODE_OF_CONDUCT.md](./template/CODE_OF_CONDUCT.md) - This document is a ready-to-use template which provides a link to the general `CODE_OF_CONDUCT.md` document from the `community` repository. Copy the template into your own repository.

* [CODEOWNERS](./template/CODEOWNERS) - In this document, specify the owners of particular parts of your repository. The owners are automatically added as reviewers when you open a pull request that modifies the code and content they own. If you additionally [modify the settings](https://help.github.com/articles/enabling-required-reviews-for-pull-requests/) of the master branch and select the **Require review from Code Owners** option, their approvals become obligatory to merge the pull request. Configure the `CODEOWNERS` document and adjust your master branch. [This](./template/CODEOWNERS) `CODEOWNERS` document contains instructions on how to do both properly.

* [CONTRIBUTING.md](./template/CONTRIBUTING.md) - This template makes a reference to the [`CONTRIBUTING.md`](https://github.com/kyma-project/community/blob/master/CONTRIBUTING.md) document that contains the general guidance from the `community` repository and describes the rules for contributing to all Kyma repositories. If there is any additional, project-specific information that you want to add to your project's `CONTRIBUTING.md` document, add them under the same sections as in the general [`CONTRIBUTING.md`](https://github.com/kyma-project/community/blob/master/CONTRIBUTING.md) document.

* [LICENSE](./template/LICENSE) - It is an obligatory element of every open-source repository. Copy the template into your repository.

* [NOTICE.md](./template/NOTICE.md) - The document defines the ownership of the copyright in the repository. Copy the template into your repository.

* [README.md](./template/README.md) - This is a template with sections that you fill in according to the provided suggestions. Add any information specific for a development guide in this document. Describe how your project works, how to use it, and how to develop it. Because all sections are optional, remove those that do not apply to your project.  
