---
displayName: Console/Microfrontends
epicsLabels:
  - area/console
  - area/luigi
---

## Scope

The Console/Microfrontends capability is all about how a user interacts with kyma UI. It drives the development of Console, a modular and extensible web user interface for managing all aspects of kyma.

## Vision

* Concise User Experience

    * Provide easy and intuitive user interfaces for kyma to support its users in the best possible way
    * Focus on a consistent user experience based on unified [Fiori 3 Fundamentals](https://sap.github.io/fundamental/components/index.html) style guides
    * Enable most common user journeys in the UI so that usage of CLI is not required
    * Don't hide the kubernetes nature from the user but augment it with kyma-specific user guidance

* Extensible & Modular

    * Use [Luigi orchestration framework](https://github.com/kyma-project/luigi) as UI extension mechanism to ease customization
    * Compose user interfaces from modular and highly reusable UI components
    * Ensure consistent and correct usage of microfrontend-hosting

* Fast & Responsive
    
    * Quick loading time for user interfaces
    * Load only the essential data that is needed for rendering user interfaces and nothing more (use GraphQL)
    * Give the user feedback for his actions (use websockets)   

