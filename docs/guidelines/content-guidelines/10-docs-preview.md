# Documentation Preview

Before you publish new documentation on the Kyma website, you can preview your changes to see if the formatting of the text is correct, images fit well, and links work as expected. To use the preview, follow these steps:

1. Go to the [`kyma/hack`](https://github.com/kyma-project/kyma/tree/main/hack), and execute the `copy_external_content.sh` script:
   
   ```bash
   ./copy_external_content.sh
   ```

   This process copies the `docs/user` folder and `docs/assets` folder from the repositories specified in the `sh` file. Everything is copied to the `external-content` folder, and all existing files are overwritten.

> [!NOTE]
> If you want to preview a new module that you added, you must add it to the `REPOS` section in `copy_external_content.sh`.

2. You can run your preview in Development or Production-Like mode.

   - To check if your documentation works locally, use Development mode by executing the following commands:

   ```bash
   npm install
   npm run docs:dev
   ```

   - To check if the build is executed properly and the documentation is displayed correctly on the website, run Production-Like mode by executing the following commands:

   ```bash
   npm run docs:build
   npm run docs:preview
   ```

> [!NOTE]
> The `npm run docs:build` command copies all the unreferenced assets and non-graphical files (like scripts, documents, etc.) to the `docs/public` directory to include them in the `dist` folder of the project.
