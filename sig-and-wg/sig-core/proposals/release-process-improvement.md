# Release process improvements

Created on 2019-01-15 by Krystian Cieslik (@crabtree).

## Status

 Proposed on 2019-01-15

## Current situation:

- Components as source code getting versioned independently from kyma resources folder.
- Kyma resources folder copied into installer identifies a kyma version (represented by kyma-installer docker image).
- The release process is taking all the resources and building a kyma version with the lates versions from the components source code.

## Problems:

- The release job consists of two factors - the pipeline definition (described in cont-integration repository) and the pipeline itself defined in Jenkinsfile under the component directory in the consolidated kyma repository. The definition of the pipeline may be removed or renamed as the component evolves. It is hard to determine the particular state of the cont-integration repository from which the release was performed. It may happen that when we remove a component and then remove its job definition the particular bugfix release will no longer be possible. Migration to prow somehow mitigates this problem as the test-infra repository will also have a release branch defined. Unfortunately, jobs configuration is still defined on the master branch of test-infra repository so no job definitions should be deleted until we are sure that we will not need it for a bugfix release.
- The current master branch version is based on components in a specific version which is not necessarily the latest. Docker images set in the charts not necessarily point to the latest builds of the components. When creating a release branch, suddenly the potential release will consist of all components in the latest version. That is a mismatch which is unexpected and will cause problems. Maybe people by purpose have not updated the component version in the charts. The solution would be either we assure that always the latest version is used (which we tried with the dynamic versioning but people have not accepted it) or we do not change the used versions for a release, which is the current proposal.
- Each run of release pipeline requires building of all the components - running unit tests, compiling the source code, building and pushing the docker images. Unfortunately, we do not utilize the caching mechanism offered by the docker on our building nodes (jenkins slaves, prow job containers).
- If we want to create pre-release with a small subset of features for a release, it requires a lot of effort to make master branch stable. Next fixes or features requires cherry-picking into the release branch.
- Before the release, we perform code freeze to ensure our master branch from which we will create final release branch stability.

## Proposed improvements:

In this proposal I identify two release artifacts which are:
- Helm charts of all components which then compose our product
- Basic yaml files for kyma-installer (config yaml referencing the installer image, containing the charts, referencing the components)

The source code of the components is not a release artifact, but it must be possible and easy to track a particular commit in the repository from which the docker image was produced.

We should split the kyma repository into two. First containing charts and installation resources, second the source code (components/tools/tests)

Having such assumptions in mind we can mitigate almost all identified problems. Focusing the release on the charts repository we are sure that what we release is what we already tested. We do not have to freeze the source code repository for a release as the process do not rely on it. All docker images of the components are already released and tested, we do not need to rerun the building process.

### Repositories structure:

#### Charts repository

Following structure represents repository for charts. Under this repository, there will be two main directories: `modules` and `kyma`. In `modules` directory, there are all charts that make up the kyma. In the `kyma` directory, there is Dockerfile that is used to create the kyma-installer and two `*.yaml` files which provide deployment details for the kyma-installer.

```
kyma
|-- Dockerfile
|-- kyma-cluster.yaml
|-- kyma-local.yaml
modules
|-- application-connector
|-- ark
...
|-- service-catalog 
```

#### Components repository

The components repository contains the source code for all kyma components, tests and tools. 

```
components
|-- api-controller
|-- apiserver-proxy
...
tests
|-- acceptance
|-- api-controller-acceptance-tests
...
tools
|-- alpine-net
|-- ark-plugins
```

### Release workflow implications:

#### Creating a pre-release:

Just create a pre-release from a commit in the master branch of the charts repository. I assume, as we are on a pre-release stage, we do not expect to have on master any features that don't belong to the following release. Even if, creating a branch for a release from that point, would also include all of the features available in that point of time.

#### Creating a release:

Create a branch in the charts repository from which then you will create a release. Create a branch because of Github mechanism that requires specifying a branch from which you then want to create a release. Github then tracks that branch in case of later commits and tells that "the release is behind ...". You can use here the master branch but it then may be misleading. Furthermore, if there will be a bugfix needed then the branch is required to create it (bugfix) against the released version, not a master.

#### Bugfixing:

If bugfix is required, we will use our standard workflow, using a forked repository create a branch from a release branch (in the charts repository) and then create PR to merge a fix into the release branch of charts repository. If bugfix is also applicable for master branch introduce the fix using PR (cherrypick that single commit from release branch into master would also be applicable but it will break our workflow).

### Dealing with the components:

We should think about how we want to build components and push them into the registry, for me it may be a totally different process than releasing the kyma. When releasing a component it can be followed with tagging the components repository. Keep in mind that when there is one repository for all components then the tag name must be unique globally in that repository. It may quickly create multiple tags which may not be clear. When building the component the docker image should be labeled with the commit id from which it was built. If we want to be 100% sure that component referenced in helm chart is the one that was built and pushed in a particular time we may consider referencing it by its SHA checksum.

It may happen that after a release we need to do a bugfix in the source code of the component. In such case, we should create a branch from a commit from which the released component was built. If the bug is already fixed on the master branch we can cherry-pick that fixing commit. If the bug still exists we then act following our standard github workflow with PR and review.

In this approach, it is possible that we will have multiple branches for fixes for different components and releases. Such case shows that our quality is something that we should work on. Or we should think about the more granular split of the source code repository (please see Further improvements for consideration -> Modules and components)

## Further improvements for consideration:

### Modules and components

In the first iteration, the concept introduced two repositories, one for the charts and one for the source code. Putting all the source code into one repository may lead to the situation when bug fixing released components is hard to track without polluting the repository with multiple branches.

In kyma we can distinguish main modules like application-connector, service-mesh (istio and istio-patch), logging etc. Having this in mind we can split the source code repository into multiple repositories - each for a single module containing source code of components bounded to that module. Heading to this approach we can gain much cleaner solution in terms of separation in the source code level. Each module has its own dependencies (vendor directory) and there is a more natural separation between module contexts. Also, the problem with release tracing on a component level is much easier. When there is a need to bugfix a component you create a branch on a module level (as each module has its own repository).

### Docker caching

We should investigate the possibility to make use of the docker cache on the building nodes. The migration to prow does not solve the problem when we fetch multiple times the same common layers for our images.

### Helm charts repository

We should consider providing our helm charts for the kyma modules thorugh charts repository https://docs.helm.sh/developing_charts/#chart-repositories
