# [PoC] New engine for the Kyma open-source website

## Context

So far, the open-source project Kyma has been using https://kyma-project.io as its documentation webiste. The content has been pulled from the [`kyma` repository `/docs` folder](https://github.com/kyma-project/kyma/docs). The webiste's source is located in the [`website` repository](https://github.com/kyma-project/website/).

As the ownership and responsibility for the `/webiste` have got lost over time, the project has become hard to maintain. Also, the main focus of the UI Kyma develpment team is on developing and maintaining Kyma Dashboard and not the documentation portal. The team's size and capacity is limited so the choice of tooling needs to be adjusted to these conditions.

Additionally, [Kyma modularization](../modularization/) initiated changes to hosting all Kyma documentation in a central repository. The modular approach assumes that each module repository includes module documentation. As a result, the content needs to be pulled from different repositories which is not an easy [`website`] feature to implement.

The decision on the tooling choice may be influenced by the fact that the Kyma project is also a foundation of SAP BTP, Kyma runtime which is a part of SAP Business Technology Platform (BTP). The SAP BTP, Kyma runtime documentation is hosted on the SAP [Help Portal](https://help.sap.com/docs/btp/sap-business-technology-platform/kyma-environment?version=Cloud).

## Potential solutions

To address the challenge, we should find and choose tooling to publish open-source Kyma documentation. Considering different options, take into account the following requirements:

- user experience
- concise and user-friendly domain
- intuitive linking
- search feature
- right-side navigation displaying the document's headers (to avoid a really, really long left-side navigation node)
- analytics feature

### SAP Help Portal

The SAP BTP, Kyma runtime documentation is hosted on the SAP [Help Portal](https://help.sap.com/docs/btp/sap-business-technology-platform/kyma-environment?version=Cloud). One of the potential solutions is hosting open-source documentation using the same engine. Documentation on the Help Portal is created using the DITA CCMS.

Potential setup:

- open-source documentation generally available on the `Cloud Production` level ??
- ??
- internal documentation available on the `Internal draft` level ??

#### Advantages:

- The Help Portal provides the possibility of having different versions of documentation depending on the user type by ading flags and using different containers.
- Having all Kyma documentation in Help Portal would enusre a unified user experince and only one entry point for Kyma users.
- Becoming an author requires license and training which gives more control over the product documentation.
- Analytics and search options enabled.
- Unified UI and URL.
- Intuitive navigation.
- Blog feature enabled.
- What's New feature enabled.

#### Disadvantages:

- DITA CCMS is slow and requires relatively big effort to add and publish content.
- DITA CCMS is not a developer-native tool and adding content together with code is not possible. Only UA developers can add content which increases Kyma technical writers' workload.
- The review process requires using comments in the Help Portal or in a GitHub issue.
- Publishing documentation is not instant but needs to be triggered manually and requires waiting at least one day for the build.
- Links to external sources are not verified.
- Search results pointing to various sections of the Help Portal.
- The amount of documentation on various SAP products prevenst open-source users from easily finidng relevant pieces of information

### Docsify

[Docsify](https://docsify.js.org/) is a site generator that turns Markdown files into a website without a build process.

#### Advantages:

- Micro-webistes with generated content pulled from every single module repository
- No build process - simplicity
- Module repositories need to include only one additional file with a list of documents representing the order of the module documents displayed on the webiste.
- UI settings maintaned in one central repository
- General Kyma documentation located in the central repository

#### Disadvantages:

- Long and relatively ugly URL

### GitHub Pages


#### Advantages:


#### Disadvantages:

- UI settings maintated in an `index.html` file in every single module repository

### Jekyll

https://github.com/jekyll/jekyll

#### Advantages:


#### Disadvantages:

### Hugo 

https://github.com/gohugoio/hugo

#### Advantages:


#### Disadvantages: