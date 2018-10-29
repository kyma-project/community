# Decision 020: Features introduction in Kyma

Created on 2018-10-23 by Ahmed Abdalla(@Abd4llA).

## Decision log

| Name| Description |
|-----------------------|------------------------------------------------------------------------------------|
| Title | Features introduction in Kyma |
| Ultimate decision maker(s) | Kyma Council |
| Due date | 2018-11-3 |
| Input provider(s) | POs, Architects |
| Group(s) affected by the decision | Kyma developers, Kyma users |
| Decision type | Binary |
| Earliest date to revisit the decision | 2019-01-30 |

## Context

Kyma release strategy and branching model defines that master should always be in a
"releasable" state and Kyma follows a cactus branching model with releases
branches. So in principal Kyma is following a trunk based development model with
no long living feature branches and thus it becomes important to have a standard
technique for introducing feature changes in Kyma and controlling when to expose
them or hide them. This decision is about proposing such a convention using Kyma
installer components and feature toggles. 

This way when introducing major changes such as rebasing Kyma on Knative we can
hide all modifications behind a kyma wide feature flag.

[Proposal can be found here](../proposals/feature_introduction.md)

## Decision

TBD

## Status

TBD

## Consequences

Once approved, these are the consequences:

TBD