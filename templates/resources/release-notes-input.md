> This template is dedicated to Product Owners. Use it to provide input for release notes for your Kyma component. See the template for [release notes](./release-notes.md) to see how your input fits into the whole release notes picture.

## {Component Name}

### {Feature or fix name}

> List other component features or fixes that you want to include in the release notes. Add each feature or fix as a separate heading and describe it in a short paragraph. The paragraph should not only provide details of the feature or fix but also explain its benefits to the Kyma users. Include screenshots to illustrate the changes better.

### Known issues

> List the known issues affecting the component. Add a short description of the issue and a workaround for it, if there is any.
- {Known issue} - {Short description}
    - {Workaround} - {Workaround description}

### Migrations and upgrades

> If you changed the way your component works or removed any of its subcomponents, clearly explain those changes and reasons behind them to the users. If users need to perform any steps to migrate, describe these steps in a separate migration guide and place it under `kyma/docs/migration-guides`.

### Security vulnerabilities fixed

> List the solved security vulnerability issues related to the Kyma project. Provide a short issue description, its calculated risk assessment, and a link to the pull request that solves the issue. You can also include a GitHub link to the issue itself. The calculated risk assessment is provided in each issue of the [`Security Vulnerability`](https://github.com/kyma-project/kyma/issues/new?template=security-vulnerability.md) type created on Github.
- {Short issue description} - {Issue link} - {Calculated risk assessment} - {PR link}

>For example, write:
> Update the K8s version in the Asset Store Controller Manager - [Issue](https://github.com/kyma-project/kyma/issues/5080) - CVSS:3.0/AV:N/AC:H/PR:N/UI:N/S:U/C:H/I:H/A:H **8.1 (High)** - {PR link}
