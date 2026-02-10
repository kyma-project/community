## Documentation Templates

We follow a topic-based documentation approach, with one file per topic. Every documentation file has a clearly defined purpose, which is reflected in the title. The content must be able to stand on its own, but you link to other documents as needed.

Here are the content types that we use in Kyma documentation:

### Concept Topics

Concept topics answer "what-is" questions and provide essential background information that users must know.
Conceptual information might explain the nature and components of a product and describe how it fits within a product category. Conceptual information helps readers map their knowledge and understanding to the tasks they need to perform and provides other essential information about a product, process, or system.

Concept topics do not give instructions or include reference information in tables or lists. They may, however, contain links to task topics that explain a related procedure or to reference topics that help define the topic in greater detail.

Use nominal style for the title, for example, "Security" or "Security Concept".

See the [concept topic template](https://github.com/kyma-project/template-repository/blob/main/docs/user/assets/templates/concept.md?plain=1).

### Task Topics

Task topics provide "how-to" instructions that enable users to accomplish a task. Each task topic should tell how to perform a single, specific procedure.

Select a title that describes the task that's accomplished, not the documented software feature. For example, use "Define resource consumption", not "Select a profile". Use the imperative "Select...", rather than the gerund form "Selecting..." or "How to select...".

With regards to structure, it’s nice to have an **introductory paragraph** ("why would I want to do this task?"), **prerequisites** if needed, then the **steps** in a numbered list, and finally the expected **result** that shows the operation was successful.
It's good practice to have 5-9 steps; anything longer can probably be split.

See the [task template](https://github.com/kyma-project/template-repository/blob/main/docs/user/assets/templates/task.md?plain=1).

### Custom Resource Topics

Use this template for a custom resource (CR) document that provides a sample custom resource and describes its fields. Additionally, the document points to the CustomResourceDefinition (CRD) used to create CRs of the given kind.

For the filename, follow the `{RESOURCE_NAME}.md` convention. For the title, use the name of the custom resource written in camel case. For example, "LogPipeline" or "Function". For reference, see [Telemetry Resources](https://github.com/kyma-project/telemetry-manager/blob/main/docs/user/resources/README.md) or [Serverless Resources](https://github.com/kyma-project/serverless/tree/main/docs/user/resources).

Some module teams update the resource documentation manually, while others use tools that generate it from code files. Autogeneration is recommended as it reduces maintenance effort and ensures documentation stays in sync with code. For implementation guidelines, see [Autogenenerate Custom Resource Documentation](./11-autogenerate-crd-docs.md). Regardless of the approach you choose, maintain the basic structure of the file.

Reference topics are typically organized into one or more sections containing a list or table with data that is usually looked up rather than memorized.

You can adjust this tamplate and use it to list other product specifications and parameters, provide essential data, and provide detailed information on subjects such as the commands in a programming language.

See the [custom resource template](https://github.com/kyma-project/template-repository/blob/main/docs/user/assets/templates/custom-resource.md?plain=1).

### Troubleshooting Topics

Troubleshooting topics provide a condition that the reader may want to correct, followed by one or more descriptions of its cause and suggested remedies.

In the document's title, mention the symptom that needs fixing ("Cannot access...") or the error message, for example, `"FAILED" status for created ServiceInstances`. Do not use the cause as a title ("Incompatible version"), because we want to help users who have no idea about the cause and only know something's not working as expected.

It's good practice to use three standard headlines (like “Condition”, “Cause”, "Solution"), each might have just one sentence or more as needed. For solution, use a numbered list if there are multiple steps to follow, and a bullet list or sub-headlines if there are several equally valid solutions.

For all troubleshooting topics, use the [troubleshooting topic template](https://github.com/kyma-project/template-repository/blob/main/docs/user/assets/templates/troubleshooting.md?plain=1).

### Release Notes

Release notes announce what's new in Kyma or in a Kyma module. Module teams generate their release notes automatically in GitHub. See [Automatically generated release notes](https://docs.github.com/en/repositories/releasing-projects-on-github/automatically-generated-release-notes).

For guidelines, see [Release Notes](../../guidelines/content-guidelines/08-release-notes.md).
