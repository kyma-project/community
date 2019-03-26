# Kyma Knative Integration WG was closed

Working Group focused on bringing two worlds, Kyma and Knative together was closed (not so) recently. After two months 
of working together we completed what the Group was meant to and may proudly announce that the goal has been achieved. 
The initial scope had to be cut a little bit, though. WG was created with three things in mind:

 * Install Kyma alongside Knative on the same cluster
 * Implement Kyma event-bus using Knative eventing under the hood
 * Implement Kyma serverless with Knative serving instead of Kubeless
 
As scopes of releases 0.6 up to 0.8 emerged to see the daylight the target above had to be shrunk and serverless part 
got removed.

Bad news aside, WG met the expectations and delivered the rest of the scope. Thanks to that from release 0.8 on you can 
have Kyma and Knative on the same Kubernetes cluster and watch both of them flourish. Moreover, in addition to old 
implementation Kyma Event Bus can now run on NATSS streaming provisioner which, by the way, is Kyma contribution to 
Knative.

Was this an easy ride? Not at all. We had to make both frameworks working on one cluster with single istio 
instance. Both Kyma nad Knative make extensive use of istio providing their own customisations. After we got it in the 
air it still wasn't a bed of roses: WG had to reimplement almost all of Event Bus without any other Kyma component 
noticing that change. Change of scope also had negative impact on morale.

Despite the challenges it was pretty instructive period. We had to understand Knative internals and some of us even 
became active members of Knative community. We have learned that even most carefully planned scope may change and we need 
to cope wiht that. And on our own skin we have felt the pain of integration of two rapidly evolving products.

During this two months WG Knative laid some solid foundations for Kyma-Knative integration. Now we look forward how to 
utilise it and make the fusion even more efficient.
