# Transition from reconciler to kyma-operator

With almost 20 modules and several teams involved, it is not feasible to coordinate the transition to the new modular architecture in a single release. The plan is to switch modules to the new architecture gradually, avoiding big bang releases and keeping existing clusters stable. 

## Phase 1 - local development

You can start developing a module operator with the provided guide, tools, and libraries. Following the guide, you generate the module and operator scaffold, copy your Helm chart inside, and use the [`manifest` package](https://github.com/kyma-project/manifest-operator#use-manifest-packages-in-your-own-operator) to install it in your reconciliation logic. You can add configuration options to the spec of your operator CRD and use them as overrides for Helm charts or in your custom logic (you still can use custom code in the operator). This way, you can also differentiate several installation profiles and code paths.
To test the operator on its own, apply the operator CRD and start it independently. You can install the operator on your test cluster or even run it locally.  
You can build the operator image and publish it (using the recommended pipeline template), and generate a `ModuleTemplate` that contains a default custom resource used to install the module and operator deployment. For every release channel you want to support, create one `ModuleTemplate`. 

With the `ModuleTemplate`, you can integrate your module with `kyma-operator`: Install kyma-operator and its custom resources in your test cluster, create a `Kyma` resource, and add your module in the spec. `Kyma-operator` installs your operator CRD, deploys your operator, and starts the module installation by creating the default custom resource. When your operator changes the status.state field (for an example, take a look at the [template-operator](https://github.com/kyma-project/kyma-operator/blob/main/samples/template-operator/api/v1alpha1/mockup_types.go#L38-L54)), it should be propagated to the `Kyma` resource status. You can install `kyma-operator` and will be able to enable and disable modules using `Kyma CLI` or `Kyma Dashboard`.

## Phase 2 - first module managed by `kyma-operator` integrated with KEB

### Phase 2a - new clusters only (internal usage)

`kyma-operator` is deployed in the control-plane and integrated with Kyma Environment Broker (KEB). As a first playground to test KEB integration without affecting existing clusters and flows, an additional plan called "preview" is created in KEB. This plan creates Kyma clusters with kyma-operator only and does not call reconciler.
The plan is tested with all environments (dev, stage, prod) to ensure we have all major integration issues resolved. The first module (frontrunner) can be submitted by the team and installed in the cluster. The minimal submission process is described and fully automated (manual steps are allowed only before submission). The minimal submission process looks like:
1. The team verifies on their own if the new module (or version) complies with the SAP Product Standard.
2. The team creates a pull request to the control-plane that contains one or more new or updated ModuleTemplates. With this action, the team declares product standard compliance and functional correctness; for example, with a checkbox in the PR template.
3. An automated test verifies if the module operator can be installed or upgraded without issues for all release channels (with default settings). This is not a functional test or integration test!
4. An automated test verifies if all images installed by operator are declared in the module template, signed by approved CA, and free of vulnerabilities.
5. The control-plane maintainer (code owner) approves the PR, which is automatically merged and rolled out.


### Phase 2b - move first module (frontrunner) from reconciler to kyma-operator (production)

Regular plans use reconciler and kyma-operator in parallel (clusters are managed by both components, but their lists of modules don't overlap). The frontrunner module is disabled in the reconciler (ignored or replaced by a dummy reconciler). `Kyma` resources containing the frontrunner module are populated for all existing clusters using a script or KEB functionality. In each cluster, `kyma-operator` installs the frontrunner module operator with its CRD (configuration specification), and creates a default configuration resource for the module. Module operator installs all needed components (including UI config map for Kyma Dashboard) and updates the status of the configuration custom resource (status.state = Ready). From that point in time, the cluster admin (or any user with sufficient authorization) can edit the configuration resource to change the module settings, or can remove the module completely by deleting it from the list in the `Kyma` resource.

The submission process is extended with all aspects required for production deployment, like support, commercialization, security, micro-delivery, etc.
## Phase 3 - all modules are managed by `kyma-operator` 

All modules are enabled in `kyma-operator`, and modules are added to Kyma resources for existing clusters. Reconciler is shut down, and KEB plans integrate only with `kyma-operator`. The "preview" plan is also gone. All teams have submitted their modules following the full submission process - the minimal submission process introduced in phase 2 extended to cover other aspects like support, microdelivery, etc. Other SAP teams can also submit their modules following improved technical guides and the full submission process. 

