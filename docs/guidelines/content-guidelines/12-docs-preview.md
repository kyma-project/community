# Documentation Preview
You can preview Kyma documentation using GitHub Pages or locally on localhost.

## Preview Documentation Using GitHub Pages

To set up automated documentation previews on GitHub Pages, add a workflow to your repository and configure the Pages settings.

1. In the `.github/workflows/` folder of your repository, create a workflow file and replace the `{MODULE_NAME}` placeholders with your module name. See [deploy-pages.yaml](https://github.com/kyma-project/istio/pull/2028/changes) as an example.

   ```yaml
   name: Deploy VitePress site to Pages
   description: |-
     This workflow creates a local instance of kyma-project.io website and adds the documentation from the branch to this instance.
     To run this workflow you need to
     - Have Github pages enabled in your fork.
     - Have the pages environment set to `docs/*` branches.
     - Have the branch the deployment should take place from in the `docs/*` format. Alternatively it can be manually dispatched.
   on:
     push:
       branches: ["docs/*"]

     workflow_dispatch:

   permissions:
     contents: read
     pages: write
     id-token: write

   concurrency:
     group: "pages"
     cancel-in-progress: false

   jobs:
     # Build job
     build:
       runs-on: ubuntu-latest
       steps:
         - name: Checkout
           uses: actions/checkout@v6
         - name: Setup Pages
           uses: actions/configure-pages@v6
         - name: Clone Kyma project
           shell: bash
           run: |-
             git clone https://github.com/kyma-project/kyma
         - name: Copy external content (documentation of modules)
           shell: bash
           run: |-
             cd kyma/hack
             sed -i 's|mkdir -p "\$TARGET_USER"|mkdir -p "\$TARGET_USER" \&\& rm -d "\$TARGET_USER"|g' ./copy_external_content.sh
             sed -i '/mkdir -p "\$TARGET_ASSETS"/d' ./copy_external_content.sh
             bash ./copy_external_content.sh
         - name: Replace external-content with current branch version
           shell: bash
           run: |-
             rm -rf ./kyma/docs/external-content/{MODULE_NAME}
             mkdir -p ./kyma/docs/external-content/{MODULE_NAME}
             cp docs -r ./kyma/docs/external-content/{MODULE_NAME}/docs
         - name: Install dependencies
           run: |-
             cd kyma
             npm install
         - name: Build with VitePress
           shell: bash
           run: |-
             cd kyma
             sed -i "s|base: '/'|base: '/${{ github.event.repository.name }}/'|" ./.vitepress/config.mjs
             ./move_to_public.sh
             git --no-pager diff
             npx vitepress build
             rm -rf ./docs/public
         - name: Upload artifact
           uses: actions/upload-pages-artifact@v5
           with:
             path: kyma/.vitepress/dist

     # Deployment job
     deploy:
       environment:
         name: github-pages
         url: ${{ steps.deployment.outputs.page_url }}
       runs-on: ubuntu-latest
       needs: build
       steps:
         - name: Deploy to GitHub Pages
           id: deployment
           uses: actions/deploy-pages@v5
   ```

2. In your fork, go to **Settings**.
3. Open the **Pages** section, and choose **GitHub Actions** as build and deployment source.
4. Open the **Environments** section, and choose **github-pages**.
5. In **Deployment branches and tags**, choose **Selected branches and tags**.
6. Choose **Edit** to update the deployment branch rule and enter the `docs/*` name pattern.


When you update documentation, create a branch that starts with the `docs/` pattern (for example, `docs/link-fix`). Every push to a `docs/*` branch now triggers the workflow and deploys the documentation to GitHub Pages at `{YOUR_GH_USERNAME}.github.io/{MODULE_NAME}`.

## Localhost

To preview documentation locally, follow these steps:

1. Go to the [`kyma/hack`](https://github.com/kyma-project/kyma/tree/main/hack), and execute the `copy_external_content.sh` script:
   
   ```bash
   ./copy_external_content.sh
   ```

   This process copies the `docs/user` folder and `docs/assets` folder from the repositories specified in the `sh` file. Everything is copied to the `external-content` folder, and all existing files are overwritten.

> [!NOTE]
> If you want to preview a new module that you added, you must add it to the `REPOS` section in `copy_external_content.sh`.

2. You can run your preview in Development or Production-Like mode.

   - To check if your documentation works locally, use Development mode by executing the following commands from the root folder:

      ```bash
      npm install
      npm run docs:dev
      ```

   - To check if the build is executed properly and the documentation is displayed correctly on the website, run Production-Like mode by executing the following commands from the root folder:

      ```bash
      npm run docs:build
      npm run docs:preview
      ```

> [!NOTE]
> The `npm run docs:build` command copies all the unreferenced assets and non-graphical files (like scripts, documents, etc.) to the `docs/public` directory to include them in the `dist` folder of the project.