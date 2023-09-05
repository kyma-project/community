# User documentation

Open-source Kyma and Kyma module user documentation is displayed on `https://kyma-project.io`. The website uses [docsify](https://docsify.js.org/#/) as a documentation site generator.

The overarching Kyma content is pulled from the `/docs` folder in the `/kyma` repository. Module documentation is pulled from the `/docs/user` folder in respective module repositories.

After initialization, to run, docsify uses the following files located in the `/docs` folder in the `/kyma` repository:

- `index.html` for configuration
- `README.md` as the home page
- `.nojekyll` preventing GitHub Pages from ignoring files that begin with an underscore
- `_sidebar.md` listing all Markdown documents to be displayed and ruling navigation between those documents on the website

> **NOTE:** If you have a flat directory structure, with all Markdown files on the same level, you need only a single `_sidebar.md` file listing the Markdown documents to be displayed. However, if there are subfolders, you need one `_sidebar.md` per folder, for the navigation between the documents to work properly.

## Publish a document from the `/kyma` repository

To publish a document located in the `/kyma` repository, follow these steps:

1. Create a pull request adding your content to a Markdown file located in the `/docs` folder.
2. Add a new `_sidebar.md` file including a link to your document, or update an existing `_siderbar.md` to include it.

## Publish a document from an existing module repositiory

To publish a document located in an existing module repositiory, follow these steps.

1. Create a pull request adding your content to a Markdown file located in the `/docs/user` folder in your module repository.
2. In the `docs/user` folder, add a new `_sidebar.md` file including a link to your document, or update the existing `_siderbar.md` to include it.

## Publish a document from a new module repostory

To publish a document located in a new module repositiory, follow the steps from [Publish a document from an existing module](#publish-a-document-from-an-existing-module-repositiory). Once completed, do the following:

1. In the `/kyma` repository, open the [`/kyma/docs/_sidebar.md`](https://github.com/kyma-project/kyma/blob/main/docs/_sidebar.md) file and under **Modules**, add a link to your module documentation main page.
2. Update [`index.html`](https://github.com/kyma-project/kyma/blob/main/docs/index.html) by adding a new **alias** entry in the **<body>** part. Use the following pattern:

    ```html
    alias: {
    '/{MODULE_REPOSITORY_NAME}/(.*)': 'https://raw.githubusercontent.com/{GITHUB_PROJECT_NAME}/{MODULE_REPOSITORY_NAME}/main/docs/$1',
    }, 
    ```

    For example:

    ```html
    alias: {
    '/serverless-manager/(.*)': 'https://raw.githubusercontent.com/kyma-project/serverless-manager/main/docs/$1',
    }, 
    ```

> **CAUTION** When you update navigation paths in documentation, make sure you check all `_sidebar.md` files that may be affected.
