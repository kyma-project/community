> This template is dedicated to technical writers. Use it to write release notes for Kyma releases. Add them as a `YYYY-MM-DD-release-notes-X.X.md` document under [`community/docs/release-notes/`](../../docs/kyma-release-notes/). Follow the content-related guidelines and tips for writing [release notes](../../docs/guidelines/content-guidelines/07-release-notes.md).

> Write an introductory paragraph and present the most important release highlights from all areas.

> Add a **CAUTION** at the beginning of the release notes whenever there are any important migration and/or update steps required for users to perform before migrating to the new Kyma version. Link to a separate migration guide under `kyma/docs/migration-guides` in which you provide these steps, describe the changes and reasons behind them, and list potential benefits for the users.  

> After the introductory paragraph, list the highlights of the release as bullet points and provide relative links to their corresponding sections.

- [API Exposure](#api-exposure) - {List of other features and fixes}
- [Application Connectivity](#application-connectivity) - {List of other features and fixes}
- [CLI](#CLI) - {List of other features and fixes}
- [Console](#console) - {List of other features and fixes}
- [Eventing](#eventing) - {List of other features and fixes}
- [Observability](#observability) - {List of other features and fixes}
- [Security](#security) - {List of other features and fixes}
- [Serverless](#serverless) - {List of other features and fixes}
- [Service Management](#service-management) - {List of other features and fixes}
- [Service Mesh](#service-mesh) - {List of other features and fixes}
- [Known issues](#known-issues) - {List of all known issues}
- [Fixed security vulnerabilities](#fixed-security-vulnerabilities) - {List of all fixed security vulnerabilities}

> For example, write:
> [Application Connectivity](#application-connectivity) - Extended tests, client certificate verification

> Introduce other component features or fixes that are included in the release notes. They should reflect the names of subsections under each component. Add relative links to component sections.

## {Area name}

### {Feature or fix name}

> Write a short paragraph that describes the feature or the fix in details and explains its benefits to the Kyma users. Include screenshots to illustrate the change better.

## Known issues

> Describe any known issues that the users can face, together with the way on how to solve these issues.

### {Area name and a brief issue description}

> Describe related known issues here. Add a link to a GitHub issue for tracking purposes, if applicable.

## Fixed security vulnerabilities

> Describe any solved security vulnerability issues related to the Kyma project. Provide a short issue description, its calculated risk assessment, and a link to the pull request that solves the issue. You can also include a GitHub link to the issue itself. The calculated risk assessment is provided in each issue of the [Security Vulnerability](https://github.com/kyma-project/kyma/issues/new?template=security-vulnerability.md) type created on Github.
### {Area name}

> Describe related security fixes here.
