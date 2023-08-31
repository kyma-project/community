---
title: Documentation preview
---

Before you publish new documentation on the Kyma website, you can preview your changes to see if the formatting of the text is correct, images fit well, and links work as expected.

That is possible thanks to the [**docsify preview feature**](https://docsify.js.org/#/quickstart?id=preview-your-site) supported by [Docsify](https://docsify.js.org/#/).

## Steps

To preview the Kyma webiste, save you changes and run the local server.

1. Run `docsify serve docs` in the `/kyma` repository.
2. Preview `https://kyma-project.io` in your browser on http://localhost:3000.

## Preview module documentation

You can also use the preview for module documentation.

1. In your module repository, create a pull request with documentation changes.
2. In the `/kyma` repository, go to the [`docs/index.html`](https://github.com/kyma-project/kyma/blob/main/docs/index.html) file and change the value of the **alias** parameter for your module. By default the value points to the `/docs/user` folder on the main branch of your module repository. Change the value to point to the respective folder in your pull request.
3. Save your changes.
4. Run `docsify serve docs`.
5. Preview `https://kyma-project.io` in your browser on http://localhost:3000.
