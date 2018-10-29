# Feature Introduction Proposal

## Overview

Using the Kyma installer components and overrides along with Helm values, we can
sktech a simple technique in Kyma to provide changes that are hidden behind a 
feature toggle. This proposal provides the details of such a technique using 
Kyam on Knative as a use case.  

## Use Case: Kyma on Knative

The goal is to have Kyma running on top of knative with Kyma components
leveraging knative features (serve, build and eventing). The transition will be
gradual such that Kyma can run in two modes, legacy mode (i.e without Knative)
and Knative mode. In such mode any component in Kyma can behave as following:

| |Mode: Legacy| Mode: Knative|
|---|-----------|--------------|
|Legacy-only Component| Installed: Yes | Installed: No (component is totally replaced with Knative functionality)|
|Dual Mode Component: Unchanged| Component is installed | Component is unaffected with Knative functionality|
|Dual Mode Component: Modified| Component is installed | Component changes its behavior to work in this mode and needs to be aware it's running in this mode|
|Knative-only Component| Installed: No| This component is only installed in Knative mode|

#### Requirements

* To incorporate knative support as an alpha feature.
* Can start Kyma in legacy mode or knative mode.
* Can upgrade an old cluster to the new version without enabling the knative
  mode.

## Proposed Solution

This is the proposed way to implement the knative use case 

![Proposal
](assets/feature-introduction.png)

## Installer

Currently the installer supports selective installation of components and values
overrides. These can be used to define an installation profile that defines 
which components to be installed and which feature toggles are set.

### Installer Profiles

The `run.sh` script accepts an optional `--profile PROFILE` argument which
specifies which installation profile to use. A profile basically consists of the
following pair of templates underneath 
`installation/resources/profiles/[PROFILE NAME]/` directory:

- `installer-cr.yaml.tmpl`: The installation custom resource file.
- `installer-config.yaml.tmpl`: The installation configuration overrides file.


### Feature Toggles

In the installer profile `installer-config.yaml.tmpl` file, feature toggles can
be globally defined following the convention `global.feature.[TOGGLE NAME]`. 
Component charts templates can check for this feature toggle and define 
a component toggle (e.g. an environment variable or a cmd argument that get set
only when the toggle is enabled). The following snippet from 
`installer-config.yaml.tmp` enables `knative` feature toggle.


```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: features-overrides
  namespace: kyma-installer
  labels:
    installer: overrides
data:
  global.feature.knative: "true" 
```