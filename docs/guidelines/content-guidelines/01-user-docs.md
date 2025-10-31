# User Documentation

Open-source Kyma and Kyma module user documentation is displayed at [`https://kyma-project.io`](https://kyma-project.io/#/). The website uses [VitePress](https://vitepress.dev/) as a documentation site generator.

The overarching Kyma content is pulled from the `/docs` folder in the `/kyma` repository. Module documentation is pulled from the `/docs/user` folder in respective module repositories.

After initialization, to run, VitePress uses the following files in the `/kyma` repository:

- `config.mjs` for configuration, including configuration of all sidebars 
- `index.md` as the home page
- `.nojekyll` preventing GitHub Pages from ignoring files that begin with an underscore

## Publish a Document from the `/kyma` Repository

To publish a document located in the `/kyma` repository, follow these steps:

1. Create a pull request adding your content to a Markdown file located in the `/docs` folder.
2. Add a new `_sidebar.ts` file including a link to your document, or update an existing `_sidebar.ts` to include it.

## Publish a Document from an Existing Module Repository

To publish a document located in an existing module repository, follow these steps.

1. Create a pull request adding your content to a Markdown file located in the `/docs/user` folder in your module repository.
2. In the `docs/user` folder, add a new `_sidebar.ts` file including a link to your document, or update the existing `_sidebar.ts` to include it.

## Publish a Document from a New Module Repository

To publish a document located in a new module repository, follow the steps from [Publish a document from an existing module](#publish-a-document-from-an-existing-module-repository). Once completed, do the following:

1. In the `/kyma` repository, open the<!-- markdown-link-check-disable-line --> [`/kyma/vitepress/config.mjs`](https://github.com/kyma-project/kyma/blob/main/.vitepress/config.mjs).
2. Add `import {YOUR_MODULE_NAME}Sidebar from '../docs/external-content/{YOUR_MODULE_NAME}/docs/user/_sidebar';` as the next import line.
3. Provide your module details in the `sidebar` section, under `themeConfig`. Use the following pattern:
<!-- markdown-link-check-disable -->
   ```mjs
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

   If you want to add documentation for a new module, you must also add the module repository name do the `deploy.yml` file, under `jobs.copy-docs.strategy.matrix.repository`, and provide the following information under `build-and-deploy.steps`:

   ```yml
   - name: 📥 Download copied docs
     uses: actions/download-artifact@v4
     with:
       name: external-docs-btp-manager
       path: docs/external-content/btp-manager
    ```

<!-- markdown-link-check-enable -->
> **CAUTION:** When you update navigation paths in documentation, make sure you check all `_sidebar.ts` files that may be affected.
