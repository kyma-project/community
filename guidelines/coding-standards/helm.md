# Helm Charts Best Practices

This guide covers the best practices for creating Helm charts every Kyma team should employ. 

### Do not use `crd-install` hook

You should never rely on `crd-install` helm hook. This is because Helm does not trigger this hook on upgrade so new CRDs won't be installed when Kyma is upgraded.

There are several alternatives to `crd-install` hook. Here are some of them order by rising implementation effort:
1. Make CRDs part of separate chart which must be installed before chart that requires them.
   
   Cons:
    * It requires yet another chart.
    * CRDs are managed by Helm and have all the limitations tha come with it.
   
   Pros:
    * There is no additional implementation effort needed.
    * CRD is separate file and may be used without other components (for tests for example).
  
2. Register CRD via its controller (if there is one). 
   
   Cons:
    * Requires controller.
    * CRDs are not listed as part of Helm release.
    * CRD is not available as file.
   
   Pros:
    * CRD is managed by component that is logically responsible for it.
    * CRD is not subject ot Helm limitations.
    
2. Create a job triggered on `pre-install` and `pre-upgrade` which registers new CRDs and removes old ones.
   
   Cons:
    * Jobs are troublesome to debug.
    * CRDs are not listed as part of Helm release.
   
   Pros:
    * CRD may still be available as separate file.
    * If there is any migration that needs to be run it can be easily implemented in job.
    * CRD is not subject ot Helm limitations.
    
### Do not move resources between charts

Sometimes you might want to move a resource (ConfigMap, Deployment, CRD, etc.) from one chart to another. How simple it may sound it is very dangerous operation as it causes charts to loose backward compatibility: it won't be possible to upgrade existing installations to new version. 

First time problem manifested itself because CRD `clustermicrofrontends.ui.kya-project.io` in Kyma repository was moved from chart `core` to chart `cluster-essentials` in release 0.7. During upgrade Kyma was not able to apply changes to `cluster-essentials` release, because `clustermicrofrontends.ui.kya-project.io` has already existed in 0.6 cluster, but was part of `core` chart. 

#### How to do it

To avoid compatibility problems you have to rename your resource when moving it from one chart to another. Bear in mind that this is destructive operation: old copy of resource will be removed and new one will be created. 

CRDs are special in this case: all CRD implementations are removed when CRD is deleted. This may cause user to loose his data. Because of that you should never move CRDs between charts. In case you really need to move CRD migration procedure needs to be prepared. Procedure should:
1. Backup all existing implementations. 
2. Remove old CRD.
3. Run upgrade.
4. Restore CRD implementations.

