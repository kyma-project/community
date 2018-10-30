## Overview

Content Strategy is a term that relates to the approach you need to to define before you start content development. You need to know:
* Who you write for
* What to document
* How and where to develop the content
* Where the content will be visible
* What types of content you have
* How the reader will navigate through the content
* What is the review process

## Location and context

One of the main Kyma principles is to care about the developer experience. That is why the Kyma content developers' focus is to provide documentation in the context of the developer, inside the cluster, and in the Console UI. This makes it easy to access and its version corresponds to the given cluster version.

It is important to remember that you need to convince developers to use Kyma before they start to work on it. That is why technical documentation must be exposed to the public without a prerequisite to start a cluster first. Publicly available documentation should not only contain the technical content but also a more overarching explanation and showcases that help to convince developers and the business decision makers to use Kyma.

This reasoning leads to a strategy of having two different locations for the documentation:
- A specific Kyma instance (cluster):
    - Contextual help in the Service Catalog so the Kyma user does not have to search for the documentation of a specific service in the general location.
    - The Docs view in the Console UI with the overarching documentation.
- A publicly available Kyma documentation portal.

## Structure

The decision for the document structure is to take a standard approach with the topic-oriented documentation. Looking at the structure of Kyma, there are two types of topics to differentiate:
* Component
* Task

The content creation starts with the component-oriented structure that is easier to follow without clear customer expectations. In the long-term, once the Kyma content developers create the whole content and know what customers want, they need to assess the task-oriented structure.

To be more precise, now readers need to know at the start what such terms as `Service Catalog` or `Serverless` mean as this is their starting point. In the long-term, they may prefer a topic, such as `Extensibility`, to see documents explaining how to quickly extend the application by provisioning external services through the catalog, and how to quickly extend the application with the lambda.

## Topic types

Every independent Kyma component is a separate documentation topic. The only exceptions from this rule are as follows:
- Kyma is an overall topic treated as a grouping point for the overarching Kyma documentation.
- For the publicly available documentation, there needs to be a separate topic for handling more business and marketing-oriented documentation with showcases.

## Documentation types

There is a set of documents that a given technical topic must include. You can also add additional document types to expose a specific topic better. However, before you create an additional document type, consult it with other Kyma content developers.  

### Obligatory

Each technical topic must have the following document types:

>**NOTE:** The Kyma content developers create templates for a given document type once there are at least two documents to use as a base for such a template.

- [Overview](../templates/resources/overview.md) - Use it to describe the component in general. It serves as an entry point for the topic. Make sure it is short but descriptive.
- [Architecture](../templates/resources/architecture.md) - Use it to describe in detail the architecture of the component. Include a diagram in this document.
- [Getting Started](../templates/resources/getting-started.md) - Use it to provide a clear step-by-step instruction that helps the user to understand a given concept better. The user must be able to go through all the steps of the document and complete them. There is no separate tutorial type. The document does not have to explicitly point out the example used as, at the end, the explicit reference to the example will be in the main content of the guide.
- [Details](../templates/resources/details.md) - Use it to describe more technical details of the component that do not fit into any other document type. Among other things, include a detailed explanation of the application lifecycle that describes how the resource is created and what other resources are created, how it is updated, how it is removed, and what each operation means from the technical point of view.
- **Configuration** - Use it to describe configuration options for a given component. Define the settings that a user can change and the expected outcome of such changes. Include the table structure with the settings in the document.
- [CLI Reference](../templates/resources/cli-reference.md) - Use it to describe the syntax and the use of CLI commands for a given component.
- [Custom Resource](../templates/resources/custom-resource.md) - Use it to document details of Custom Resource Definitions (CRDs) that are part of a given component.
- **API** - Use it to document the exposed external API of components that the Kyma administrators use to integrate them with Kyma.
- **Operational Guide** - Use it to describe how to operate the component. Explain all details needed for the component troubleshooting.

>**NOTE:** The examples are exceptions from the above types. The plan is to expose a table-like view as a reference to all examples and to enhance it with examples that the external Kyma users provide. Another plan is to implement an automated way of filling in the content in such documents.

### Optional

You can add the following document type to the Kyma documentation:
- **Installation** - Use it to describe the installation process. This includes guides for local, cluster, or component installation, as well as documents describing installation scripts.

## The content source

The Kyma content developers write the content in [markdown](https://daringfireball.net/projects/markdown/) and store it in [Git](https://git-scm.com/) repositories.

## Audience

For the documentation that is part of the Kyma cluster, the audience is the Kyma user. The assumption is that a person that gained access to the cluster and can sign in to the Console already knows the basics and knows what to use Kyma for. Therefore, the intended audience are the following technical people that operate the cluster:
- Developers
- Administrators

As for the documentation that is published in the publicly available portal, the audience is much more diverse and it requires much more documentation to understand what Kyma is. Nevertheless, the assumption is that the audience has the basic technical understanding of such terms as containers, cloud, and Kubernetes:
- Developers
- Technical analysts
- Business decision makers
Because of such a diversified audience, the navigation of the Kyma portal needs to clearly separate technical content from the more showcase-based content.

>**NOTE:** When you write a given document type, adjust its voice and tone to the audience that you address. See the **Voice and tone** section in the [guidelines](https://github.com/YaaS/REST_API_Documentation_Guidelines/blob/master/010_About_Style_And_Standards.html.md#voice-and-tone) for more details.

### The assumed reader's knowledge

The assumption is that the audience is familiar with the following terms and does not require the explanation of technical concepts behind them:
- Kubernetes
- Docker and containers

## The main purpose of instructions

One of the main Kyma principles is that CLI is, metaphorically speaking, the first-level citizen. Therefore, the documentation's main focus is to explain concepts and provide step-by-step instructions using CLI commands instead of Console UI screenshots. The only exception from the rule applies to the components that cannot be managed through the CLI, in which case screenshots are essential to explain specific functionalities.

## Quality

A technical writer must review any content produced for Kyma. The review not only checks the language quality of a given document but also verifies its structure, consistency, and compliance with the guidelines.

## Release notes

Release notes are written and by default displayed in the release notes section of the GitHub repository. The owner of the release notes, similarly to any other content, is the team that owns the component that the release notes describe. Technical writers must review and accept all release notes.

All release notes must be visible in one view in the Kyma documentation portal.
