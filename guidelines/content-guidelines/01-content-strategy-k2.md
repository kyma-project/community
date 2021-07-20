---
title: Content Strategy K2.0 
---

## Table of contents

- [Purpose and Target Group](#purpose-and-target-group)
- [Information types](#information-types)
  - [Concept Topics](#concept-topics)
  - [Task (tutorial) topics](#task-(tutorial)-topics)
  - [Reference topics](#reference-topics)
  - [Troubleshooting topics](#troubleshooting-topics)
  - [Release Notes](#release-notes)
- [Graphics](#graphics)
  - [Architecture and Flow Diagrams](#architecture-and-flow-diagrams)
  - [Screenshots](#screenshots)
 - [Documentation Structure and Target Groups](#documentation-structure-and-target-groups)
   - [Overview](#overview)
   - [Getting Started](#getting-started)
   - [Deep Dive/Tutorials](#deep-dive/tutorials)
   - [Operations](#operations)
   - [Technical References](#technical-references)
   - [Glossary](#glossary)
 - [Content Source](#content-source)

## Purpose and target group

This content strategy focuses on the publicly available documentation under <https://kyma-project.io/docs/>.
More documentation may be found in the Kyma Github repositories.

The assumed reader of this guide has some basic knowledge of technical writing. To learn more, read the excellent [Istio document about adding documentation](https://istio.io/latest/docs/releases/contribute/add-content/).


## Information types

We follow a topic-based documentation approach, with one file per topic. Every documentation file has a clearly defined purpose, which is reflected in the title. The content must be able to stand on its own, but links should point to other documents as needed.

*For the structure within topics, I suggest to follow the DITA standard – because it’s pretty simple, it’s well-established, and it’s what is used in SKR documentation. Instead of currently 12-13 content types, we should thus be able to get away with just four (or five) types:*

### Concept topics

Answer "what-is" questions and provide essential background information that users must know. 
You'll find most concept topics in the Overview section, but they can be useful elsewhere too. 

Use nominal style for the title, for example, "Security" or "Security Concept".

*Not sure we need much predefined structure here; it may depend on the content.*

> [- DITA Concept Topics](http://docs.oasis-open.org/dita/dita/v1.3/errata02/os/complete/part3-all-inclusive/archSpec/technicalContent/dita-concept-topic.html)
> Conceptual information might explain the nature and components of a product and describe how it fits into a category of products. Conceptual information helps readers to map their knowledge and understanding to the tasks they need to perform and to provide other essential information about a product, process, or system.

### Task (tutorial) topics

Provide "how-to" instructions that enable users to accomplish a task. Each task topic should tell how to perform a single, specific procedure.

Select a title that describes the task that's accomplished, not the documented software feature - for example, "Define ressource consumption", not "Select a profile". You can use gerund form "Selecting...", imperative "Select...", or "How to select...".

With regards to structure, it’s nice to have an **introductory paragraph** ("why would I want to do this task?"), **prerequisites** if needed, then the **steps**, and finally the expected **result** that shows the operation was successful.
It's good practice to have 5-9 steps; anything longer can probably be split.

*For longer code blocks, I'd like to use an expandable section as described [here](https://gist.github.com/pierrejoubert73/902cc94d79424356a8d20be2b382e1ab) - but not sure whether our website supports this.*


> [- Dita Task Topics](http://docs.oasis-open.org/dita/dita/v1.3/errata02/os/complete/part3-all-inclusive/archSpec/technicalContent/dita-task-topic.html)
> A task information type answers the "How do I?" question by providing precise step-by-step instructions detailing the requirements that must be fulfilled, the actions that must be performed, and the order in which the actions must be performed. The task topic includes sections for describing the context, prerequisites, expected results, and other aspects of a task. 

### Reference topics

Typically organized into one or more sections containing a list or table with data that is typically “looked up” rather than memorized.

Use nominal style for the title, for example, "Configuration Parameters".

*In our case, architecture diagrams could fall into this category – we could also choose to define them as a separate type – TBD.*

> [- DITA Reference Topics](http://docs.oasis-open.org/dita/dita/v1.3/errata02/os/complete/part3-all-inclusive/archSpec/technicalContent/dita-reference-topic.html)
> Reference topics provide quick access to fact-based information. In technical information, reference topics are used to list product specifications and parameters, provide essential data, and provide detailed information on subjects such as the commands in a programming language. 

### Troubleshooting topics

Provide a condition that the reader may want to correct, followed by one or more descriptions of its cause and suggested remedies.

As title, mention the symptom that needs fixing ("Cannot access...") or the error message. Do not use the cause as title ("Incompatible version").

It's good practice to use three standard headlines (like “Condition”, “Cause”, “Remedy”), each might have just one sentence or more as needed. For remedy, use a numbered list if there are multiple steps to follow, and a bullet list if there are several equally valid solutions.

> [- DITA Troubleshooting Topics](http://docs.oasis-open.org/dita/dita/v1.3/errata02/os/complete/part3-all-inclusive/archSpec/technicalContent/dita-troubleshooting-topic.html)
> In its simplest form, troubleshooting information follows this pattern:
>
> 1. A condition or symptom. Usually the condition or symptom is an undesirable state in a system, a product, or a service that a reader wants to correct.
> 2. A cause for the condition or symptom.
> 3. A remedy for the condition or symptom.
>
> The troubleshooting topic provides sections for describing the condition, causes, and remedies needed to restore a system, a product, or a service to normal.
>
> For some conditions there could be more than one cause-remedy pair. The troubleshooting topic accommodates this. Typically, a cause is immediately followed by its remedy. Multiple cause-remedy pairs can provide a series of successive fall-backs for resolving a condition.
>
> Cause and remedy might occur in combinations other than pairs. It is possible to have:
>
> * Multiple causes with the same remedy
> * A single cause with more than one remedy
> * A remedy with no known cause
> * A cause with no known remedy
> The troubleshooting information type also can be used to document alarm clearing strategies.

### Release Notes

Announce what's new in Kyma. 

After an introductory paragraph that outlines the city that's the namesake of the current release, a list briefly presents the new and changed features. Links lead to longer paragraphs that describe the changes in more detail.


## Graphics

### Architecture and Flow Diagrams

*Klaudia is preparing the new diagram style, see <https://github.com/kyma-project/community/issues/542>.*

### Screenshots

Wherever possible, present screenshots as SUI (simplified user interfaces). Basically, this means to blur out all UI elements that aren't essential for the task at hand.

For more information, see:
• <https://www.techsmith.com/blog/simplified-user-interface/>
• <https://www.tcworld.info/e-magazine/technical-writing/simplified-graphics-and-screenshots-in-software-documentation-1102/>

## Documentation Structure and Target Groups

On the Kyma website, we’ll have five main tabs containing multiple documents each, plus a glossary.

### Overview

Target Group: Decision Makers (Tech Leads) and newbies.
Contains a quick overview of the idea behind Kyma, a diagram of the main areas with a brief explanation, and, as needed, longer documents going into the details of each main area.

### Getting Started

Target Group: Software Developers who quickly want to see what they can do with Kyma.
Contains a guide/tutorial that covers typical steps you need to get started.

### Deep Dive/Tutorials

Target Group: Software Developers leveraging all the Kyma functionalities.
Under this tab, there will be subtabs according to main areas (except UI – Busola and CLI are mentioned as needed within the instructions of the respective main area.)

### Operations

Target Group: Admins/Operators who make sure the Kyma cluster is configured as needed and keeps running in a healthy and secure way.
Contains installation and configuration instructions, backup info, troubleshooting guides…

### Technical References

Target Group: Users who want need to look up specific detailed information.
Contains the architecture diagrams, configuration charts, etc,; no explanation of concepts or instructions.

### Glossary

Target Group: Anyone who wants to look up terms they’re not familiar with.
Explains basic terms, with a focus on terms specific to Kyma.

## Content Source

We write the content in [Markdown](https://daringfireball.net/projects/markdown/) and store it in [Git](https://git-scm.com/) repositories.
