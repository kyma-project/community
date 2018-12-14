# Add new component to the official Kyma documentation

When you add a new component to the `kyma/docs` folder, you must ensure that the content displays also on the official Kyma documentation website under `https://kyma-project.io/`. To do so, update the following files:

- [`resources/core/charts/docs/charts/documentation/templates/docs-job.yaml`](https://github.com/kyma-project/kyma/blob/master/resources/core/charts/docs/charts/documentation/templates/docs-job.yaml)
- [`resources/core/charts/azure-broker/templates/azure-service-classes-docu.yaml`](https://github.com/kyma-project/kyma/blob/75cd925a611e4c6b822968f03b863eea4a08b5a8/resources/core/charts/azure-broker/templates/azure-service-classes-docu.yaml) (only for Azure classes documentation)
- [`resources/core/charts/docs/charts/documentation/templates/docs-gcp-service-classes.yaml`](https://github.com/kyma-project/kyma/blob/master/resources/core/charts/docs/charts/documentation/templates/docs-gcp-service-classes.yaml) (only for GCP classes documentation)
- [`resources/core/charts/docs/charts/documentation/templates/docs-hb-service-classes.yaml`](https://github.com/kyma-project/kyma/blob/master/resources/core/charts/docs/charts/documentation/templates/docs-hb-service-classes.yaml) (only for Helm Broker classes documentation)
- [`docs/manifest.yaml`](https://github.com/kyma-project/kyma/blob/master/docs/manifest.yaml)
- [`docs/Jenkinsfile`](https://github.com/kyma-project/kyma/blob/master/docs/Jenkinsfile)
- [`kyma/docs/docs-build.yaml`](https://github.com/kyma-project/kyma/blob/master/docs/docs-build.yaml)

The `docs-job.yaml` file, as well as the corresponding files for broker classes documentation, defines Docker images with documentation to be run to upload the documentation to Minio.

The `docs/manifest.yaml` file specifies the order in which the documentation topics display on the website. If the component is not included in this file, it will not display on the UI.

The `docs/Jenkinsfile` file defines document topics for which Docker images are built on the Jenkins CI. This solution will be outdated as soon as the migration to Prow is complete.

The `docs-build.yaml` file is prepared for the Prow migration and aims to replace `docs/Jenkinsfile` after Kyma migrates from Jenkins to the new CI tool. This file specifies all document topics needed for building Docker images. Every item must have the `name` field specified, which represents the name of the Docker image for a topic. The final image name is followed by a `-docs` suffix. If the new component directory is equal to the image name, you don’t have to specify it. Otherwise, add an additional `directory` property.

Example 1:
```
- name: test
```

This example builds the `test-docs` image from the `kyma/docs/test` directory.


Example 2:
```
- name: example-two
  directory: test/example
```

This example builds the `example-two-docs` image from the `kyma/docs/test/example` directory.
