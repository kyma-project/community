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

> Write a staring paragraph and introduce the most important release highlights from all components. List the highlights as bullet points and provide absolute links to given sections.

- [{Feature or fix name}](#section-link) - {One-sentence description}
- [{Feature or fix name}](#section-link) - {One-sentence description}
- [{Feature or fix name}](#section-link) - {One-sentence description}

> For example, write:
> [Application Connector modularization](#section-link) - Components have been moved to separate Helm charts.

> Introduce other component features or fixes that are included in the release notes. They should reflect the names of subsections under each component. Add absolute links to component sections.

- [Application Connector](#section-link) - {List of other features and fixes}
- [Console](#section-link) - {List of other features and fixes}
- [Eventing](#section-link) - {List of other features and fixes}
- [Logging](#section-link) - {List of other features and fixes}
- [Monitoring](#section-link) - {List of other features and fixes}
- [Security](#section-link) - {List of other features and fixes}
- [Serverless](#section-link) - {List of other features and fixes}
- [Service Catalog](#section-link) - {List of other features and fixes}
- [Service Mesh](#section-link) - {List of other features and fixes}
- [Tracing](#section-link) - {List of other features and fixes}

> For example, write:
> [Application Connector](https://kyma-project.io/blog/release-notes-05#application-connector) - Extended tests, client certificate verification

---

## {Component name}

### {Feature or fix name}

> Write a short paragraph that describes the feature or the fix in details. Include screenshots to illustrate the change better.
