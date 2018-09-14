# Generating changelog with pull request labels

Created on 2018-08-06 by Paweł Kosiec (@pkosiec).

## Status

Proposed on 2018-08-06.

## Key assumptions

- Open source community wants as much transparency as possible.
- Changelog should be easy to read
- Changelog should be available for developers with offline access, if they have cloned repository
- Defining changes by contributors should be as easy as possible.

## Proposed solution

In one sentence: Categorizing changes with pull request labels.

### Workflow

- All developers should categorize their changes with PR labels
- All changes with specific changelog-related labels are included in changelog
- Pull requests should have a title written in imperative mood (like commit messages). The title will be included in changelog, along with the author and a link to the particular PR.
- No label on a pull request means that the change won't be included in changelog.

### Result

- The changelog is written to CHANGELOG.md file and put in GitHub release description via [GitHub API](https://developer.github.com/v3/repos/releases/#create-a-release) for a specific git tags
- All changes are grouped by selected labels

Desired output would look like this:
    - [example GitHub releases](https://github.com/lerna/lerna-changelog/releases)
    - [example CHANGELOG.md](https://github.com/lerna/lerna-changelog/blob/master/CHANGELOG.md)

### Labels

For now we don't introduce any new labels. We reuse `area/` labels from [the accepted proposal](https://github.wdf.sap.corp/SAP-CP-Extension-Factory/community/blob/master/sig-and-wg/wg-github-issues-migration/proposals/github-issues-labels-proposal.md). This means that all changes are grouped by area (e.g. `installation`, `security`...).

**Example:**

> ### Installation
> - Sample change 1
> - Sample change 2
>
> ### Security
> - Sample change 1
> - Sample change 2

In future we can introduce nested grouping for changelog (first level: change type, second one: area of the change). Change type will marked with additional prefixed labels. Choosing right prefix will be discussed later.

**Example:**

> ### Bug fixes
> **Installation**
> - Sample change 1
> - Sample change 2
>
> **Security**
> - Sample change 1
> - Sample change 2
>
> ### Features
> **Installation**
> - Sample change 1
> - Sample change 2
>
> **Security**
> - Sample change 1
> - Sample change 2

## Pros

- One place to define general changes in pull request: pull request title, which is usually prefilled with commit message, if there is just only one commit
- Selecting type of change is a no-brainer: contributor expands the list of available PR labels and chooses right category.
- We can filter unlabeled pull requests ([link](https://github.com/kyma-project/kyma/issues?q=is%3Aopen+is%3Apr+no%3Alabel)) and easily add labels to them to generate complete changelog. Also, it's easy to see the full list of PRs to check if they are named correctly (and adjust titles - we can do that even for closed pull requests)
- It enforces user to make a single type of change in a pull request
- There are many tools available for generating changelog from pull requests labels.

## Cons

- More detailed changes in changelog. For example, when there is a bigger feature, which needs multiple PRs, it will be visible as multiple bullet points.
- Pull request authors and reviewers can forget about adding a label to categorize the change. Anyway, even closed pull requests can be edited before generating the changelog. 

## Tools

The changelog can be generated during CI build job. There are a lot of tools for that. 

#### [Lerna Changelog](https://github.com/lerna/lerna-changelog)

- Very simple tool
- Generates changelog with PR links and authors to standard output ([example](https://github.com/lerna/lerna-changelog/releases)) - we can easily pipe it to a CHANGELOG.md file and as GitHub release
- Allows to define custom labels
- Can be run with parameters to generate a full changelog or partial one (for one specific app version)
- Uses GitHub API

#### [GitHub Changelog Generator](https://github.com/github-changelog-generator/github-changelog-generator)

- Generates changelog based on GitHub issues and PRs
- Many configuration options
- No custom labels for PRs are supported: there are only 3 sections in generated changelog: `bug`, `enhancement` and all other changes
- Uses GitHub API
- Written in Ruby - it can be difficult to contribute to it
