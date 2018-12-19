> This template is dedicated to technical writers. Use it to write release notes for Kyma releases. Add them as blog posts under [`website/src/blog-posts/`](https://github.com/kyma-project/website/tree/master/src/blog-posts) and any related screenshots under the [`assets`](https://github.com/kyma-project/website/tree/master/src/blog-posts/assets) folder. Follow the content-related guidelines and tips for writing [release notes](../../release-notes.md).

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

<!-- This line adds a button that allows you to download the latest release. Provide the path to the release on GitHub in place of the {path} placeholder and put "Download {version number}" in place of the {content} placeholder. Remember to remove the code block. -->

```
<a class=“btn-blog” href=“{path}“><span>{content}</span></a>
```

> Write an introductory paragraph and present the most important release highlights from all components. List the highlights as bullet points and provide absolute links to their corresponding sections.

- [{Feature or fix name}](#absolute-link-to-subsection) - {One-sentence description}
- [{Feature or fix name}](#absolute-link-to-subsection) - {One-sentence description}
- [{Feature or fix name}](#absolute-link-to-subsection) - {One-sentence description}

> For example, write:
> [Application Connector modularization](#section-link) - Components have been moved to separate Helm charts.

> Introduce other component features or fixes that are included in the release notes. They should reflect the names of subsections under each component. Add absolute links to component sections.

- [Application Connector](#absolute-link-to-subsection) - {List of other features and fixes}
- [Console](#absolute-link-to-subsection) - {List of other features and fixes}
- [Eventing](#absolute-link-to-subsection) - {List of other features and fixes}
- [Logging](#absolute-link-to-subsection) - {List of other features and fixes}
- [Monitoring](#absolute-link-to-subsection) - {List of other features and fixes}
- [Security](#absolute-link-to-subsection) - {List of other features and fixes}
- [Serverless](#absolute-link-to-subsection) - {List of other features and fixes}
- [Service Catalog](#absolute-link-to-subsection) - {List of other features and fixes}
- [Service Mesh](#absolute-link-to-subsection) - {List of other features and fixes}
- [Tracing](#absolute-link-to-subsection) - {List of other features and fixes}

> For example, write:
> [Application Connector](https://kyma-project.io/blog/release-notes-05#application-connector) - Extended tests, client certificate verification

---

## {Component name}

### {Feature or fix name}

> Write a short paragraph that describes the feature or the fix in details. Include screenshots to illustrate the change better.
