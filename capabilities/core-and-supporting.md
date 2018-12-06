---
displayName: Core and Supporting
epicsLabels:
  - area/core-and-supporting
  - area/community
---

## Scope

The Core and Supporting capability provides functionality for delivery of the content and its visual representation. For us, the content is not only regular documentation but also specifications and contextual help.
Due to the nature of the content and the number of different areas, it sits in, the Core and Supporting capability provides also many generic tools that do not only support content but also other aspects of the product.

In other words, if some content must be displayed in a given UI, the capability cares also about the rest of the UI of a given business functionality and its backend.

## Vision

- Content is written once and reusable in different contexts in an efficient way: documentation portal, inline help or UI applications. In other words, we provide a headless CMS that is an abstraction on top or more generic ObjectStore solution that allows you to store any static content like for example client side applications or any other objects. This is possible because of :
  - Kubernetes native way of content delivery with support for distributed content sourcing and modularity (content is delivered only if documented component is enabled)
  - Generic reusable UI components for rendering documentation and specification that are reusable in any context, for example Service Catalog view to display documentation for ServiceClasses and their instances or in Applications view to display the documentation of connected applications.
  - Provide UI support for rendering specifications like for example Swagger for REST API, EDM for OData or AsyncApi for any kind of asynchronous communication.
  - Backend that enables reuse of content and specifications` details in any UI context
- Enable support not only for out-of-the-box rendering or content in the Console UI but also make it easy for Kyma user to generate a standalone documentation portal for their services.
- Support for easy content development like a preview before publishing or templates integration
- Support automated content validation, like links, grammar, consistency and accordance with specification
- Assure content is not only available for existing Kyma users but also for new users and the contributors



