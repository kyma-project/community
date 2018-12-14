# Add a new component to the Kyma documentation

When you add a new component to the `kyma/docs` folder, you must ensure that the content also shows on the `https://kyma-project.io/` website and in the Console UI. To do so, update the following files:

- [`resources/core/charts/docs/charts/documentation/templates/docs-job.yaml`](https://github.com/kyma-project/kyma/blob/master/resources/core/charts/docs/charts/documentation/templates/docs-job.yaml)
- [`docs/manifest.yaml`](https://github.com/kyma-project/kyma/blob/master/docs/manifest.yaml)
- [`docs/Jenkinsfile`](https://github.com/kyma-project/kyma/blob/master/docs/Jenkinsfile)
- [`kyma/docs/docs-build.yaml`](https://github.com/kyma-project/kyma/blob/master/docs/docs-build.yaml)

The `docs-job.yaml` file defines Docker images with documentation which are run to upload this documentation to Minio.

The `docs/manifest.yaml` file specifies the order of documentation topics on the website. If you do not include the component in this file, it doesn't show on the UI.

The `docs/Jenkinsfile` file defines documentation topics for which Docker images are built on the CI tool.

The `docs-build.yaml` file is prepared for the Prow migration and aims to replace `docs/Jenkinsfile` after Kyma migrates to the new CI tool. This file specifies all documentation topics needed for building Docker images. Every item must have the `name` field specified, which represents the name of the Docker image for a topic. The final image name is followed by a `-docs` suffix. If the new component directory is equal to the image name, you don’t have to specify it. Otherwise, add an additional `directory` property.

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
