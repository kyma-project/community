# [PoC] New engine for the Kyma open-source website

## Context

So far, the open-source project Kyma has been using https://kyma-project.io as its documentation website. The content has been pulled from the [`kyma` repository `/docs` folder](https://github.com/kyma-project/kyma/docs). The website's source is located in the [`website` repository](https://github.com/kyma-project/website/).

As the ownership and responsibility for the `/website` have got lost over time, the project has become hard to maintain. The source code is now outdates and relatively messy. Also, the main focus of the UI Kyma development team is on developing and maintaining Kyma Dashboard and not the documentation portal. The team's size and capacity are limited so the choice of tooling needs to be adjusted to these conditions.

Additionally, [Kyma modularization](../modularization/) initiated changes to hosting all Kyma documentation in a central repository. The modular approach assumes that each module repository includes module documentation. As a result, the content needs to be pulled from different repositories which is not an easy [`website`] feature to implement.

The decision on the tooling choice may be influenced by the fact that the Kyma project is also a foundation of SAP BTP, Kyma runtime which is a part of SAP Business Technology Platform (BTP). The SAP BTP, Kyma runtime is a servie and its documentation must be hosted on the SAP [Help Portal](https://help.sap.com/docs/btp/sap-business-technology-platform/kyma-environment?version=Cloud).

## Potential solutions

To address the challenge, we should find and choose tooling to publish open-source Kyma documentation. Considering different options, take into account the following requirements:

- user experience both for readers and contributors/writers
- concise and user-friendly domain
- intuitive linking
- search feature
- right-side navigation displaying the document's headers (to avoid a long left-side navigation node)
- analytics feature

### SAP Help Portal

The SAP BTP, Kyma runtime documentation is hosted on the SAP [Help Portal](https://help.sap.com/docs/btp/sap-business-technology-platform/kyma-environment?version=Cloud). One of the potential solutions is hosting open-source documentation using the same engine. Documentation on the Help Portal is created using the IXIASoft DITA CCMS.

#### Advantages:

- The Help Portal provides the possibility of having different versions of documentation depending on the user type by adding flags and using different containers and defining different projects and outputs.
- Having all Kyma documentation in Help Portal would ensure a unified user experience and only one entry point for Kyma users.
- DITA CCMS output could be re-used in other publishing channels, for exmaple, GitHub.
- Becoming an author requires a license and training which gives more control over the product documentation.
- We might be able have outputs that aren't part of the BTP Core deliverable, but instead maintained and published by a Kyma Doc Lead.
- Analytics and search options enabled.
- Unified UI and URL.
- Intuitive navigation.
- Blog feature enabled.
- What's New feature enabled.

#### Disadvantages:

- DITA CCMS is slow and requires a relatively big effort to add and publish content.
- DITA CCMS is not a developer-native tool and adding content together with code is not possible. Only UA developers can add content which increases Kyma technical writers' workload.
- The review process requires using comments in the Help Portal or in a GitHub issue.
- Publishing documentation is not instant but needs to be triggered manually and requires waiting at least one day for the build.
- Links to external sources are not verified.
- Search results pointing to various sections of the Help Portal.
- The amount of documentation on various SAP products prevents open-source users from easily finding relevant pieces of information

### Docsify

[Docsify](https://docsify.js.org/) is a site generator that turns Markdown files into a website without a build process.

#### Advantages:

- Micro-websites with generated content pulled from every single module repository
- No build process - simplicity
- Module repositories need to include only one additional file with a list of documents representing the order of the module documents displayed on the website.
- UI settings maintained in one central repository
- General Kyma documentation located in the central repository
- Low maintenance effort

#### Disadvantages:

- Long and relatively ugly URL
- No links back to the central page
- Faulty search feature

### GitHub Pages, Jekyll, Hugo, VitePress

#### Advantages:

- nice UI and user experience

#### Disadvantages:

- UI settings are maintated in an `index.html` file in every single module repository
- build pipeline
