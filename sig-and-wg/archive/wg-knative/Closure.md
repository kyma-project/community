# Kyma Knative Integration WG was closed

## Achivements

Kyma 0.8 supports as an alternative to the old implementation approach the Knative framework, leveraging till now its Eventing functionality. 
In the future versions of Kyma, the integration with Knative will be more deeply, making Knative fully available in Kyma environment.  

## Outcomes
- Kyma installation can be used to install both versions of Kyma:
    - Kyma without Knative (the default installation)
    - Kyma based on Knative, currently as an installation feature flag (`--knative`).
- In the Kyma/Knative version, the Event-Bus implementation is based on Knative Eventing 

## Challenges
- The integration of Knative installation into Kyma installation was a complex, evolving task
- The implementation of the Event-Bus was refactored to be based on Knative/Eventing without changing its interfaces. This was possible based on the Knative knowledge acquired by Kyma team. 
- Making the new Event-Bus implementation transparent to the other Kyma components.

## Lessons learned
- A profoundly understanding of Knative implementation was absolutely necessary.
- Cause Knative is evolving rapidly, its integration in Kyma is a continuous process which is time consuming.
- Leverage the Knative functionality in Kyma, if this is feasible for Kyma, it's a viable long-term solution.    
