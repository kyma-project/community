# Kyma Installer with embedded kyma charts and tiller

Created on 2018-09-19 by Piotr Bochyński (@pbochynski).

## Status

Proposed on 2018-09-19

## Abstract

Current installation process involves downloading kyma charts from cloud storage. Such solution allows to release kyma installer and kyma charts independently. In the consequence, there is possibility to install different versions (or flavors) of kyma components using single installer binary. It seems to be a nice feature, but in the reality introduces problems:
- Build/Release process is more complex because it has to publish kyma charts in cloud storage.
- It is possible to run installer with kyma charts version not supported (not tested with the installer in given version).
- Upgrade process is more complex (usually installer has to be upgraded and then kyma release URL has to be updated)

## Proposal

Installer will contain kyma charts and scripts in the docker image. The build process has to be adapted to produce installer image for every kyma snapshot/release. It is important to include also version.env file in the image as it is essential part of kyma charts.

Installer deployment will also contain tiller container that will be used to install helm charts. This change should be applied also in helm-broker. When it is done tiller is no longer required as a prerequisite.

The solution should also support kyma modifications and local development. The minimal requirement is to provide manual how to build own kyma installer image based on modified charts.

The benefits of the solution:
- simpler install/update process - just deploy installer in given version
- installer always works with kyma charts tested with installer version
- better security - tiller is used only inside installer pod - no external access.


