# Add a new component to the Kyma documentation

When you add a new component to the `kyma/docs` folder, you must ensure that the content also shows on the `https://kyma-project.io/` website and in the Console UI. To do so, update the following files:

- `manifest.yaml` specifies the order of documentation topics on the website. If you do not include the component in this file, it doesn't show on the UI.
- `docs-build.yaml` specifies all documentation topics needed for building Docker images. Every item must have the **name** field specified, which represents the name of the Docker image for a topic. The final image name is followed by a `-docs` suffix. If the name of the new component directory matches the name of the image, you don’t have to specify it. Otherwise, add an additional **directory** property.


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

Follow these steps to add a new component to the Kyma documentation:

1. Update the `docs-build.yaml` and the `manifest.yaml` files, and merge your changes to the `master` branch.
2. Update the image version of the docs.
