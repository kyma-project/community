---
title: {CRD kind}
type: Custom Resource
---

> This document is a ready-to-use template for a custom resource (CR) document that provides a sample custom resource and description of its elements. Additionally, the document points to the Custom Resource Definition (CRD) used to create CRs of the given kind. For reference, see the existing documents for the [Installation](https://github.com/kyma-project/kyma/blob/master/docs/kyma/docs/040-installation-custom-resource.md) and the [Api](https://github.com/kyma-project/kyma/blob/master/docs/api-gateway/docs/011-api-custom-resource.md) CRs.

> **NOTE:** Blockquotes in this document provide instructions. Remove them from the final document.


The {CRD name} Custom Resource Definition (CRD) is a detailed description of the kind of data and the format used to {rest of the CRD description}. To get the up-to-date CRD and show the output in the yaml format, run this command:

```
kubectl get crd {CRD name} -o yaml
```

## Sample custom resource

> In this section, provide an example custom resource created based on the CRD described in the introductory section. Describe the functionality of the CR and highlight all of the optional elements and the way they are utilized.
Provide the custom resource code sample in the ready-to-use format.

This is a sample resource that {description of what the example presents}.

```
apiVersion:
kind:
metadata:
  name:
{another_field}:
```

## Custom resource parameters

This table lists all the possible parameters of a given resource together with their description:


| Parameter   |      Mandatory?      | Description |
|:----------:|:-------------:|:------|
| **metadata.name** |    **YES**   | Specifies the name of the CR. |
| **{another_parameter}** |    **{YES/NO}**   | {Parameter description} |
