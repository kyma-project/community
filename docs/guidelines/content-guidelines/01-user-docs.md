# User Documentation

Open-source Kyma and Kyma module user documentation is displayed at [`https://kyma-project.io`](https://kyma-project.io/#/). The website uses [VitePress](https://vitepress.dev/) as a documentation site generator.

The overarching Kyma content is pulled from the `/docs` folder in the `/kyma` repository. Module documentation is pulled from the `/docs/user` folder in respective module repositories.

After initialization, to run, VitePress uses the following files in the `/kyma` repository:

- `config.mjs` for configuration, including configuration of all sidebars 
- `index.md` as the home page
- `deploy.yml` to deploy documentation to the website
- `.nojekyll` preventing GitHub Pages from ignoring files that begin with an underscore

## Publish a Document from the `/kyma` Repository

To publish a document located in the `/kyma` repository, follow these steps:

1. Create a pull request adding your content to a Markdown file located in the `/docs` folder.
2. Go to [`/kyma/.vitepress/config.mjs`](https://github.com/kyma-project/kyma/blob/main/.vitepress/config.mjs#L138) and provide your document details in the **sidebar** element, under **themeConfig**. Use the following pattern:

   ```bash
    {
       text: 'Troubleshooting',
       link: '/04-operation-guides/troubleshooting/README',
    },
    ```

3. Run the [Deploy VitePress site to GitHub Pages](https://github.com/kyma-project/kyma/actions/workflows/deploy.yml) or wait for the CronJob to start it (every day at midnight).
4. Make sure that the `public` folder in the root of the `/kyma` repository is deleted after the build. If not, delete it manually to clean up the environment.

## Publish a Document from an Existing Module Repository

To publish a document located in an existing module repository, follow these steps.

1. Create a pull request adding your content to a Markdown file located in the `/docs/user` folder in your module repository.
2. Update the `_sidebar.ts` file located in the `/docs/user` folder in the module repository, including a link to your document.
3. In the `/kyma` repository, run the [Deploy VitePress site to GitHub Pages](https://github.com/kyma-project/kyma/actions/workflows/deploy.yml) or wait for the CronJob to start it (every day at midnight).
4. Make sure that the `public` folder in the root of the `/kyma` repository is deleted after the build. If not, delete it manually to clean up the environment.

## Publish a Document from a New Module Repository

1. Create a pull request adding your content to a Markdown file(s) located in the `/docs/user` folder in your module repository.
2. Add a `_sidebar.ts` file in the `/docs/user` folder in the module repository, and create a navigation structure for your module documentation. See the following example of the SAP BTP Operator module [`_sidebar.ts`](https://github.com/kyma-project/btp-manager/blob/main/docs/user/_sidebar.ts). 

3. In the `/kyma` repository, open the<!-- markdown-link-check-disable-line --> [`/kyma/vitepress/config.mjs`](https://github.com/kyma-project/kyma/blob/main/.vitepress/config.mjs).
4. Add `import {YOUR_MODULE_NAME}Sidebar from '../docs/external-content/{YOUR_MODULE_NAME}/docs/user/_sidebar';` as the next import line.
5. Provide your module details in the **sidebar** element, under **themeConfig**. Use the following pattern:
<!-- markdown-link-check-disable -->
      ```bash
      {
         text: 'My Module',
         link: '/external-content/my-module/docs/user/README.md',
         collapsed: true,
         items: makeSidebarAbsolutePath(
           myModuleSidebar,
           'my-module',
         ),
      },
      ```

6. In the `/kyma` repository, add the module repository name to the [`deploy.yml`](https://github.com/kyma-project/kyma/blob/main/.github/workflows/deploy.yml) file under the **jobs.copy-docs.strategy.matrix.repository** element. For example:

      ```yaml
      jobs:
        copy-docs:
          strategy:
            matrix:
              repository:
                - btp-manager
                - istio
                - serverless
                - {YOUR_MODULE_REPO}
      ```
<!-- markdown-link-check-enable -->
> **CAUTION:** When you update navigation paths in documentation, make sure you check all `_sidebar.ts` files that may be affected.

7. Run the [Deploy VitePress site to GitHub Pages](https://github.com/kyma-project/kyma/actions/workflows/deploy.yml) or wait for the CronJob to start it (every day at midnight).
8. Make sure that the `public` folder in the root of the `/kyma` repository is deleted after the build. If not, delete it manually to clean up the environment.

## Execute Prettier

Prettier helps keep your code nicely formatted. The project is set up to use it automatically, as long as you're using Visual Studio Code with the Prettier plugin. However, it's configured to run during Git commits. So, if you're not using the integrated Git in VSCode for commits, you'll need to run it manually with this command:

```bash
npx prettier --config ./.prettierrc --ignore-path ./.prettierignore --write '**/*.{ts,tsx,mjs,js,jsx,json,html,css,yaml,md}'
```
