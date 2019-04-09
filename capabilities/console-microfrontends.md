---
displayName: Console/Microfrontends
epicsLabels:
  - area/console
  - area/luigi
---

## Scope

The Console/Microfrontends capability is all about how a user interacts with kyma UI. It drives the development of Console, a modular and extensible web user interface for managing all aspects of kyma.

## Vision

* User Experience

    * Provide easy and intuitive user interfaces for kyma to support its users in the best possible way
    * Focus on a consistent user experience based on unified [Fiori 3 Fundamentals](https://sap.github.io/fundamental/components/index.html) style guides
    * Enable most common user journeys in the UI so that usage CLI is not required
    * Don't hide the kubernetes nature from the user but augment it with kyma-specific user guidance

* Extensibility & Modularity

    * Use [Luigi orchestration framework](https://github.com/kyma-project/luigi) as UI extension mechanism to ease customization
    * Compose user interfaces from modular and highly reusable UI components
    * Ensure consistent and correct usage of microfrontend-hosting

* Unified API access
    
    * Use unified GraphQL API facade for all user interfaces
    * Leverage websockets for server-push communication   

