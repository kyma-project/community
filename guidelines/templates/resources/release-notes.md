> This template is dedicated to technical writers. Use it to write release notes for Kyma releases. Add them as a blog post under [`website/src/blog-posts/`](https://github.com/kyma-project/website/tree/master/src/blog-posts). Place any related screenshots under the [`assets`](https://github.com/kyma-project/website/tree/master/src/blog-posts/assets) folder. Follow the content-related guidelines and tips for writing [release notes](../../release-notes.md).

<!-- Fill in the required metadata for the blog post to render properly on the "kyma-project.io" website. Remember to remove the code block. -->

```
---
path: "/blog/{link}"
date: "{YYY-MM-DD}"
author: "{Name and surname}, {Role} @Kyma"
tags:
  - release-notes
title: "{Release notes title}"
---
```

<!-- This line adds a button that allows you to download the latest release. Provide the path to the release on GitHub in place of the {path} placeholder and put "Download {version number}" in place of the {content} placeholders. Remember to remove the code block. -->

```
<a class="btn-blog" href="{path}" alt="{content}">{content}</a>
```

> Write an introductory paragraph and present the most important release highlights from all components. List the highlights as bullet points and provide relative links to their corresponding sections.

- [{Feature or fix name}](#relative-link-to-subsection) - {One-sentence description}
- [{Feature or fix name}](#relative-link-to-subsection) - {One-sentence description}
- [{Feature or fix name}](#relative-link-to-subsection) - {One-sentence description}

> For example, write:
> [Application Connector modularization](#modularization) - Components have been moved to separate Helm charts.

> Add the <!-- overview --> comment after this introductory paragraph to separate the excerpt rendered on the main page from the rest of the document. For more details, see [this](https://github.com/kyma-project/website/blob/master/docs/write-blog-posts.md) document.

> Introduce other component features or fixes that are included in the release notes. They should reflect the names of subsections under each component. Add relative links to component sections.

- [Application Connector](#relative-link-to-subsection) - {List of other features and fixes}
- [Console](#relative-link-to-subsection) - {List of other features and fixes}
- [Eventing](#relative-link-to-subsection) - {List of other features and fixes}
- [Logging](#relative-link-to-subsection) - {List of other features and fixes}
- [Monitoring](#relative-link-to-subsection) - {List of other features and fixes}
- [Security](#relative-link-to-subsection) - {List of other features and fixes}
- [Serverless](#relative-link-to-subsection) - {List of other features and fixes}
- [Service Catalog](#relative-link-to-subsection) - {List of other features and fixes}
- [Service Mesh](#relative-link-to-subsection) - {List of other features and fixes}
- [Tracing](#relative-link-to-subsection) - {List of other features and fixes}

> For example, write:
> [Application Connector](#application-connector) - Extended tests, client certificate verification

---

## {Component name}

### {Feature or fix name}

> Write a short paragraph that describes the feature or the fix in details and explains its benefits to the Kyma users. Include screenshots to illustrate the change better.
