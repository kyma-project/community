# DR 019: Changelog generation with pull request labels

Created on 2018-09-13 by Paweł Kosiec (@pkosiec).

## Decision log

| Name | Description |
|-----------------------|------------------------------------------------------------------------------------|
| Title | Changelog generation with pull request labels |
| Ultimate decision maker(s) | Kyma Council |
| Due date | 2018-09-13 |
| Input provider(s) | Jose Cortina, Łukasz Górnicki, Piotr Bochyński, Marek Nawa, Rakesh Garimella, Paweł Kosiec |
| Group(s) affected by the decision | All Kyma members |
| Decision type | Choice |
| Earliest date to revisit the decision | 01.12.2018 |

## Context

Maintaining changelog is an essential part of developing an open source project. The community wants maximum transparency. As changelog should be easy to read, there is a need to group changes. In the same time, creating changelog should require as little manual work as possible. In addition of changelog

There were two different proposals for creating changelog. In both of them the changelog would be generated automatically in CI pipeline. The only difference was the way to describe changes and categorize them. In first one, there was a special `release` block in pull request description. The second proposal was about using pull request titles written in imperative mood and grouping changes with pull request labels.

The main advantage of second approach is having single place to define changes in pull request: the pull request title, which is usually prefilled with commit message, if there is just only one commit. Selecting type of change is easy - contributor expands the list of available PR labels and chooses right category. Another advantage is that the approach enforces user to make a single type of change in a pull request. There are also many tools available for generating changelog from pull requests labels.

## Decision

The decision is to use pull request titles for describing changes and reuse existing **`area/`** labels for pull request categorization. Changes will be grouped in changelog by area.
Changelog will be generated as a part of CI pipeline using tool [**Lerna Changelog**](https://github.com/lerna/lerna-changelog) in the form of `CHANGELOG.md` file and GitHub release description.

## Status

Accepted on 2018-08-17.

## Consequences

When creating a new pull request, every team member has to:
- write pull request title in imperative mood
- use **`area/`** labels to categorize change made in the pull request

