# Decision 020: Feature Gates in Kyma

Created on 2018-09-28 by Ahmed Abdalla(@Abd4llA).

## Decision log

| Name| Description |
|-----------------------|------------------------------------------------------------------------------------|
| Title | Introducing feature gates in Kyma |
| Ultimate decision maker(s) | Kyma Council |
| Due date | 2018-10-10 |
| Input provider(s) | POs, Architects |
| Group(s) affected by the decision | Kyma developers, Kyma users |
| Decision type | Binary |
| Earliest date to revisit the decision | 2019-01-10 |

## Context

Kyma release strategy and branching model defines master should always be in a
"releasable" state and Kyma follows a cactus branching model with releases
branches. So in principal Kyma is following a trunk based development model with
no long living feature branches. A core technique of applying trunk based
development is using feature toggles, and in Kyma case introducing Kyma wide
feature gates is needed and not just per component.

This way when introducing major changes such as rebasing Kyma on Knative we can
hide all modifications behind a kyma wide feature flag.

[Proposal can be found here](../proposals/feature_gates.md)

## Decision

TBD

## Status

TBD

## Consequences

Once approved, these are the consequences:

TBD