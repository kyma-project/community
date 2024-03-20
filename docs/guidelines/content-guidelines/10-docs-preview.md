# Documentation Preview

Before you publish new documentation on the Kyma website, you can preview your changes to see if the formatting of the text is correct, images fit well, and links work as expected.

That is possible thanks to the [**docsify preview feature**](https://docsify.js.org/#/quickstart?id=preview-your-site) supported by [docsify](https://docsify.js.org/#/).

## Prerequisites

You have installed docsify-cli.

To install docsify-cli, run `npm i docsify-cli -g`.

## Steps

To preview content on the Kyma website, save your changes and run the local server.

1. Run `docsify serve docs` in the `/kyma` repository.
<!-- markdown-link-check-disable-next-line -->
2. Preview `https://kyma-project.io` in your browser on `http://localhost:3000`.

## Preview Module Documentation

1. In your module repository, create a pull request with documentation changes.
2. In the `/kyma` repository, go to the [`docs/index.html`](https://github.com/kyma-project/kyma/blob/main/docs/index.html) file and change the value of the **alias** parameter for your module. By default, the value points to the raw version of the `/docs` folder on the main branch of your module repository. Change the value to point to the raw version of the respective folder on your branch. Use the following pattern:

  ```html
  alias: {
  '/{MODULE_REPOSITORY_NAME}/(.*)': 'https://raw.githubusercontent.com/{YOUR_GITHUB_USERNAME}/{FORKED_MODULE_REPOSITORY_NAME}/{BRANCH_TO_PREVIEW}/docs/$1',
  }, 
  ```

  For example:
  
  ```html
  alias: {
  '/istio/(.*)': 'https://raw.githubusercontent.com/kymauser/istio/docu-update/docs/$1',
  },
  ```

3. Save your changes.
4. Run `docsify serve docs`.
<!-- markdown-link-check-disable-next-line -->
5. Preview `https://kyma-project.io` in your browser on `http://localhost:3000`.
