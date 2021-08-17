---
title: Content Strategy K2.0 
---

## Table of contents

- [Purpose and Audience](#purpose-and-audience)
- [Information types](#information-types)
  - [Concept Topics](#concept-topics)
  - [Task (tutorial) topics](#task-tutorial-topics)
  - [Reference topics](#reference-topics)
  - [Troubleshooting topics](#troubleshooting-topics)
  - [Release Notes](#release-notes)
- [Graphics](#graphics)
  - [Architecture and flow diagrams](#architecture-and-flow-diagrams)
  - [Screenshots](#screenshots)
- [Target Groups](#target-groups)
  - [Decision maker](#decision-maker)
  - [Software developer](#software-developer)
  - [Admin/Operations](#adminoperations)
- [Documentation structure](#documentation-structure)
  - [Overview](#overview)
  - [Getting Started](#getting-started)
  - [Tutorials](#tutorials)
  - [Operations](#operations)
  - [Technical References](#technical-references)
  - [Glossary](#glossary)
- [Content Source](#content-source)

## Purpose and audience

This content strategy focuses on the publicly available documentation for the open-source [Kyma project](https://kyma-project.io/docs/).
More documentation may be found in the [Kyma](https://github.com/kyma-project/kyma) GitHub repositories.

The assumed reader of this guide has some basic knowledge of technical writing. To learn more, read the excellent [Istio document about adding documentation](https://istio.io/latest/docs/releases/contribute/add-content/).

## Information types

We follow a topic-based documentation approach, with one file per topic. Every documentation file has a clearly defined purpose, which is reflected in the title. The content must be able to stand on its own, but you may use links to point to other documents as needed.

Here are the content types that we use in Kyma documentation:

### Concept topics

Answer "what-is" questions and provide essential background information that users must know.
You'll find most concept topics in the Overview section, but they can be useful elsewhere too.

Use nominal style for the title, for example, "Security" or "Security Concept".

For all concept topics, use the [concept topic template](../templates/resources/concept.md).
Learn more about [Concept Topics](http://docs.oasis-open.org/dita/dita/v1.3/errata02/os/complete/part3-all-inclusive/archSpec/technicalContent/dita-concept-topic.html).

### Task (tutorial) topics

Provide "how-to" instructions that enable users to accomplish a task. Each task topic should tell how to perform a single, specific procedure.

Select a title that describes the task that's accomplished, not the documented software feature. For example, use "Define resource consumption", not "Select a profile". You can use the gerund form "Selecting...", imperative "Select...", or "How to select...".

With regards to structure, it’s nice to have an **introductory paragraph** ("why would I want to do this task?"), **prerequisites** if needed, then the **steps** in a numbered list, and finally the expected **result** that shows the operation was successful.
It's good practice to have 5-9 steps; anything longer can probably be split.

For all step instructions, use the [task topic template](../templates/resources/task.md) and learn more about [Task Topics](http://docs.oasis-open.org/dita/dita/v1.3/errata02/os/complete/part3-all-inclusive/archSpec/technicalContent/dita-task-topic.html).

### Reference topics

Typically organized into one or more sections containing a list or table with data that is usually “looked up” rather than memorized.

Use nominal style for the title, for example, "Configuration Parameters".

*In our case, architecture diagrams could fall into this category – we could also choose to define them as a separate type – TBD.*

Reference topics provide quick access to fact-based information. In technical information, reference topics are used to list product specifications and parameters, provide essential data, and provide detailed information on subjects such as the commands in a programming language.

Learn more about [Reference Topics](http://docs.oasis-open.org/dita/dita/v1.3/errata02/os/complete/part3-all-inclusive/archSpec/technicalContent/dita-reference-topic.html).

### Troubleshooting topics

Provide a condition that the reader may want to correct, followed by one or more descriptions of its cause and suggested remedies.

In the title, mention the symptom that needs fixing ("Cannot access...") or the error message. To quote an error message, start and end with `'` to escape `"` (because quotation marks `"` themselves do not display correctly on the website), for example, `title: '"FAILED" status for created ServiceInstances'`. Do not use the cause as title ("Incompatible version"), because we also want to help users who have no idea about the cause and only know something's not working as expected.

It's good practice to use three standard headlines (like “Condition”, “Cause”, “Remedy”), each might have just one sentence or more as needed. For remedy, use a numbered list if there are multiple steps to follow, and a bullet list or sub-headlines if there are several equally valid solutions.

For all troubleshooting topics, use the [troubleshooting topic template](../templates/resources/troubleshooting.md).
Learn more about [Troubleshooting Topics](http://docs.oasis-open.org/dita/dita/v1.3/errata02/os/complete/part3-all-inclusive/archSpec/technicalContent/dita-troubleshooting-topic.html).

### Release notes

Announce what's new in Kyma.

After an introductory paragraph that outlines the city that's the namesake of the current release, a list briefly presents the new and changed features. Links lead to longer paragraphs that describe the changes in more detail.

## Graphics

### Architecture and flow diagrams

For information about our diagram style, see [Diagrams](../02-diagrams.md).

### Screenshots

For information about Screenshots, see [Screenshots](../07-diagrams.md).

## Target groups

The general assumption is that the audience is familiar with the following terms and does not require the explanation of technical concepts behind them:

- Kubernetes
- Docker and containers

<!-- Information about generic target groups is based on the SAP Styleguide: https://help.sap.com/viewer/DRAFT/e33c591ae4494a659a3f5f983c9d1161/PROD/en-US/546dfb06e80c4005aabc4795d548fe35.html -->

### Decision maker

Assesses the software to make sure that it meets the company's needs. Requires the facts – not just marketing spin – before signing on the dotted line. Wants to purchase the right solution for the company, and ensure stakeholders back this decision.

### Software developer

Interested in technical topics. Solid knowledge of programming languages. Experienced in programming and development projects. Expert in the technical or business area. Uses and contributes to community content. Wants to develop, maintain, and enhance software.

### Admin/Operations

Deals with installation, upgrades, system troubleshooting. Wants to support the ongoing operations and evolution of the Kyma implementation.

## Documentation structure

On the Kyma website, we have five main tabs containing multiple documents each, plus a glossary.

### Overview

**Target Group**: Decision Makers (Tech Leads) and newbies.
Contains a quick overview of the idea behind Kyma, a diagram of the main areas with a brief explanation, and, as needed, longer documents going into the details of each main area.

### Getting Started

**Target Group**: Software Developers who quickly want to see what they can do with Kyma.
Contains a guide/tutorial that covers typical steps you need to perform to get started.

### Deep Dive/Tutorials

**Target Group**: Software Developers leveraging all the Kyma functionalities.
Under this tab, there are subtabs according to main areas (except UI – user interfaces are mentioned as needed within the instructions of the respective main area). Documents in the subtabs contain "how-to" instructions that enable users to accomplish a task

### Operations

**Target Group**: Admins/Operators who make sure the Kyma cluster is configured as needed and keeps running in a healthy and secure way.
Contains installation and configuration instructions, backup info, troubleshooting guides…

### Technical References

**Target Group**: Users who want to look up specific detailed information.
Contains the architecture diagrams, configuration charts, etc.; no explanation of concepts or instructions.

### Glossary

**Target Group**: Anyone who wants to look up terms they’re not familiar with.
Explains basic terms, with a focus on terms specific to Kyma.

## Content source

We write the content in [Markdown](https://daringfireball.net/projects/markdown/) and store it in [Git](https://git-scm.com/) repositories.
