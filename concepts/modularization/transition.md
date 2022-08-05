# Transition from reconciler to kyma-operator

With almost 20 modules and several teams involved it is not feasible to coordinate the transition to the new modular architecture in the single release. The plan is to switch modules to the new architecture gradually, avoiding big bang releases and keeping existing clusters stable. 

## Phase 1 - local development

You can start developing module operator with the provided guide, tools and libraries. Following the guide you will generate the operator scaffold, copy your helm chart inside and use the [`manifest` package](https://github.com/kyma-project/manifest-operator#use-manifest-packages-in-your-own-operator) to install it in your reconciliation logic. You can add configuration options to the spec of your operator CRD and use it as overrides for helm charts or in your custom logic (you still can use custom code in the operator).
You can test operator on its own by applying operator CRD and creating the instance (you can install operator on your test cluster or even run it locally). When your operator works you can build the image and publish it (using the recommended pipeline template) and create a `ModuleTemplate` that contains a default custom resource used to install the module and operator deployment. You create one module template per release channel you want to support. Having `ModuleTemplate` you can start testing your module with `kyma-operator` . To do that you need to install kyma-operator and its custom resources in your test cluster, create a `Kyma` resource and add your module in the spec. `Kyma-operator` will install your operator CRD, will deploy your operator and start module installation by creating default custom resource. When your operator changes the status.state field it should be propagated to the `Kyma` resource status. You can install `kyma-operator` and enable/disable modules using `kyma CLI` or `kyma Dashboard`.

## Phase 2 - first module managed by `kyma-operator` integrated with KEB

### Phase 2a - new clusters only

`kyma-operator` is deployed in the control-plane and integrated with Kyma Environment Broker (KEB). Additional plan (preview) is created in KEB that creates Kyma clusters with kyma-operator only (no call to reconciler). This plan is the first playground to test KEB integration without affecting existing clusters and flows. The plan is tested with all environments (dev, stage, prod) to ensure we have all major integration issues resolved. The first module (frontrunner) can be submitted by the team and installed in the cluster. The minimal submission process is described and fully automated (manual steps are allowed only before submission). The minimal submission process looks like:
- team verifies on their own if new module (or version) is product standard compliant 
- team creates pull request to the control-plane that contains new/updated ModuleTemplate (by doing it team declares product standard compliance and functional correctness - can be checkbox in PR template)
- automated test verifies if the module operator can be installed/upgraded without issues for all release channels (with default settings) - it is not functional test or integration test 
- automated test verifies if all images installed by operator are declared in the module template, signed by approved CA, vulnerability free
- control-plane maintainer (codeowner) approves the PR which is automatically merged and rolled out


### Phase 2b - move first module (frontrunner) from reconciler to kyma-operator

Regular plans use reconciler and kyma-operator in parallel (clusters are managed by both components, but their lists of modules don't overlap). Frontrunner module is disabled in the reconciler (ignored / replaced by dummy reconciler). `Kyma` resources containing the first module are populated for all existing clusters using script or KEB functionality. In each cluster `kyma-operator` installs frontrunner module operator with its CRD (configuration specification), and creates default configuration resource for the module. From that point in time cluster admin (or any user with sufficient roles) can edit the configuration resource to change module settings or can remove the module completely by deleting it from the list in the `Kyma` resource.

## Phase 3 - all modules are managed by `kyma-operator` 

Reconciler is shutdown and KEB plans integrate only with `kyma-operator` (preview plan is also gone). All teams have submitted their modules using full version of submission process - the minimal submission process introduced in phase 2 extended to cover other aspects like support, microdelivery, etc. Other SAP teams can also submit their modules following improved technical guides and submission process description. 

