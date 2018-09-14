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
| Earliest date to revisit the decision | 2018-12-01 |

## Context

Maintaining changelog is an essential part of developing an open source project. It is essential as the community wants to keep the maximum transparency. As changelog should be easy to read, there is a need to group changes. At the same time, creating changelog should require as little manual work as possible.

There were two different proposals for creating changelog. In both of them the changelog is be generated automatically in the CI pipeline. The only difference is the way to describe changes and categorize them. In the first one, there is a special `release` block in the pull request description. The second proposal is about using the imperative mood in pull request titles and grouping changes with the pull request labels.

The main advantage of the second approach is having one place to define changes in a pull request. This place is the pull request title, usually prefilled with the commit message, in case there is only one commit. Selecting type of change is easy - contributor expands the list of available PR labels and chooses right category. Another advantage is that the approach enforces user to make a single type of change in a pull request. There are also many tools available for generating changelog from pull requests labels.

## Decision

The decision is to use pull request titles to describe changes, and to reuse the existing **`area/`** labels for pull request categorization. Changes will be grouped in the changelog by the area.
The changelog will be generated as a part of the CI pipeline using the [**Lerna Changelog**](https://github.com/lerna/lerna-changelog) tool, in the form of `CHANGELOG.md` file and the GitHub release description.

## Status

Accepted on 2018-08-17.

## Consequences

When creating a new pull request, every team member must:
- use the imperative mood for the pull request title
- use the **`area/`** labels to categorize changes
All pull requests without the labels should be rejected during the review process.
