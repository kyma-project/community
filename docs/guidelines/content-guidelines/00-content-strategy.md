# Content Strategy

## Purpose and Audience

This content strategy focuses on the publicly available documentation for the open-source [Kyma project](https://kyma-project.io). The source of the documentation displayed on the website is stored in [GitHub](https://github.com/kyma-project/kyma) and [Kyma modules](https://kyma-project.io/#/06-modules/README)' repositories. The content is written in [Markdown](https://daringfireball.net/projects/markdown/).

The assumed readers of this guide and contributors to the documentation have some basic knowledge of technical writing.

## Information Types

We follow a topic-based documentation approach, with one file per topic. Every documentation file has a clearly defined purpose, which is reflected in the title. The content must be able to stand on its own, but you link to other documents as needed.

Here are the content types that we use in Kyma documentation:

### Concept Topics

Concept topics answer "what-is" questions and provide essential background information that users must know.
You'll find most concept topics in the Overview section, but they can be useful elsewhere too.

Use nominal style for the title, for example, "Security" or "Security Concept".

For all concept topics, use the [concept topic template](https://github.com/kyma-project/community/blob/main/templates/resources/concept.md).

### Task Topics

Task topics provide "how-to" instructions that enable users to accomplish a task. Each task topic should tell how to perform a single, specific procedure.

Select a title that describes the task that's accomplished, not the documented software feature. For example, use "Define resource consumption", not "Select a profile". Use the imperative "Select...", rather than the gerund form "Selecting..." or "How to select...".

With regards to structure, it’s nice to have an **introductory paragraph** ("why would I want to do this task?"), **prerequisites** if needed, then the **steps** in a numbered list, and finally the expected **result** that shows the operation was successful.
It's good practice to have 5-9 steps; anything longer can probably be split.

For all step instructions, use the [task template](../../../templates/resources/task.md).

### Reference Topics

Reference topics are typically organized into one or more sections containing a list or table with data that is usually looked up rather than memorized.

Reference topics provide quick access to fact-based information. In technical information, reference topics are used to list product specifications and parameters, provide essential data, and provide detailed information on subjects such as the commands in a programming language.

Use nominal style for the title, for example, "Configuration Parameters".

Use the templates for [architecture documents](https://github.com/kyma-project/community/blob/main/templates/resources/architecture.md), [configuration parameter charts](https://github.com/kyma-project/community/blob/main/templates/resources/configuration.md), and [custom resources](https://github.com/kyma-project/community/blob/main/templates/resources/custom-resource.md).

### Troubleshooting Topics

Troubleshooting topics provide a condition that the reader may want to correct, followed by one or more descriptions of its cause and suggested remedies.

In the title, mention the symptom that needs fixing ("Cannot access...") or the error message. To quote an error message, start and end with `'` to escape `"` (because quotation marks `"` themselves do not display correctly on the website), for example, `title: '"FAILED" status for created ServiceInstances'`. Do not use the cause as the title ("Incompatible version"), because we also want to help users who have no idea about the cause and only know something's not working as expected.

It's good practice to use three standard headlines (like “Condition”, “Cause”, “Remedy”), each might have just one sentence or more as needed. For remedy, use a numbered list if there are multiple steps to follow, and a bullet list or sub-headlines if there are several equally valid solutions.

For all troubleshooting topics, use the [troubleshooting topic template](https://github.com/kyma-project/community/blob/main/templates/resources/troubleshooting.md).

### Release Notes

Release notes announce what's new in Kyma or in a Kyma module. Module teams generate their release notes automatically in GitHub. See [Automatically generated release notes](https://docs.github.com/en/repositories/releasing-projects-on-github/automatically-generated-release-notes).

For guidelines on how to add release notes content, see the [release notes template](../../guidelines/content-guidelines/07-release-notes.md).

## Target Groups

The general assumption is that the readers of the documentation are familiar with the following terms and do not need an explanation of the technical concepts behind them:

- Kubernetes
- Docker and containers

These are the assumed target groups for Kyma documentation:

### Decision Maker

Assesses the software to make sure that it meets the company's needs. Requires the facts – not just marketing spin – before signing on the dotted line. Wants to choose the right solution for the company, and ensure stakeholders back this decision.

### Software Developer

Interested in technical topics. Solid knowledge of programming languages. Experienced in programming and development projects. Expert in the technical or business area. Uses and contributes to community content. Wants to develop, maintain, and enhance software.

### Admin/Operator

Deals with installation, upgrades, system troubleshooting. Wants to support the ongoing operations and evolution of the Kyma implementation.

## Documentation Structure

On the Kyma website, we have five main tabs containing multiple documents each, plus a glossary.

### What Is Kyma

**Target Group**: Decision Makers (Tech Leads) and newbies.

Contains a quick overview of the idea behind Kyma, presents Kyma's strengths, and explains the connection between Kyma and SAP BTP, Kyma runtime.

### Quick Install

**Target Group**: Software Developers who quickly want to see what they can do with Kyma.

Contains a guide that covers typical steps you need to perform to get started.

### Modules

**Target Group**:
  
- Software Developers leveraging all the Kyma modules' functionalities.
- Admins/Operators who make sure their Kyma modules are configured as needed.

Under this tab, there are subtabs dedicated to all Kyma modules. Content of the subtabs is pulled from the `docs/user` folders in module repositories. Module documentation covers: general module description, feature scope, "how-to" instructions that enable users to accomplish a task, installation and configuration instructions, backup info, security documentation, troubleshooting guides, and detailed information, such as architecture diagrams, or configuration charts.

### Operation Guides

**Target Group**: Admins/Operators who make sure the Kyma cluster is configured as needed and keeps running in a healthy and secure way.

Includes installation and configuration instructions, backup info, "how-to" instructions that enable users to uninstall or upgrade a module, and troubleshooting guides.

### Glossary

**Target Group**: Anyone who wants to look up terms they’re not familiar with.

Explains basic terms, with a focus on terms specific to Kyma.
