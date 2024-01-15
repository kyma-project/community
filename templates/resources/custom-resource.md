---
title: {CRD kind}
---

<!-- Use this template for a custom resource (CR) document that provides a sample custom resource and description of its elements. Additionally, the document points to the CustomResourceDefinition (CRD) used to create CRs of the given kind.

For the filename, follow the `{COMPONENT/AREA}-{NUMBER}-{TITLE}.md` convention.

For reference, see the existing [CR documents](https://kyma-project.io/#/05-technical-reference/00-custom-resources/README).-->

The `{CRD name}` CustomResourceDefinition (CRD) is a detailed description of the kind of data and the format used to {provide the CRD description}. To get the up-to-date CRD and show the output in the `yaml` format, run this command:

```bash
kubectl get crd {CRD name} -o yaml
```

## Sample Custom Resource

<!-- In this section, provide an example custom resource created based on the CRD described in the introductory section. Describe the functionality of the CR and highlight all of the optional elements and the way they are utilized.
Provide the custom resource code sample in a ready-to-use format. -->

This is a sample resource that {provide a description of what the example presents}.

```yaml
apiVersion:
kind:
metadata:
  name:
{another_field}:
```

## Custom Resource Parameters

This table lists all the possible parameters of a given resource together with their descriptions:

| Parameter   | Required |  Description |
|-------------|:---------:|--------------|
| **metadata.name** | Yes | Specifies the name of the CR. |
| **{another_parameter}** | {Yes/No} | {Parameter description} |

## Related Resources and Components

These are the resources related to this CR:

| Custom resource |   Description |
|-----------------|---------------|
| {Related CRD kind} |  {Briefly describe the relation between the resources}. |

These components use this CR:

| Component   |   Description |
|-------------|---------------|
| {Component name} |  {Briefly describe the relation between the CR and the given component}. |
