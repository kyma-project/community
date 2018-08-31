---
title: {CRD kind}
type: Custom Resource
---

> This document is a ready-to-use template for a Custom Resource (CR) document that provides a sample Custom Resource and description of its elements. Additionally, the document points to the Custom Resource Definition (CRD) used to create CRs of the given kind. For reference, see the existing documents for the [Installation](https://github.com/kyma-project/kyma/blob/master/docs/kyma/docs/040-installation-custom-resource.md) and the [Api](https://github.com/kyma-project/kyma/blob/master/docs/api-gateway/docs/011-api-custom-resource.md) CRs.

> **NOTE:** Blockquotes in this document provide instructions. Remove them from the final document.


The {CRD name} Custom Resource Definition (CRD) is a detailed description of the kind of data and the format used to {rest of the CRD description}. To get the up-to-date CRD and show the output in the yaml format, run this command:

```
kubectl get crd {CRD name} -o yaml
```

## Sample Custom Resource

> In this section, provide an example Custom Resource created based on the CRD described in the introductory section. Describe the functionality of the CR and highlight all of the optional elements and the way they are utilized.
Provide the Custom Resource code sample in the ready-to-use format.

```
apiVersion:
kind:
metadata:
  name:
{another_field}:
```

This table analyses the elements of the sample CR and the information it contains:

> In the table, describe CR's fields starting with the **metadata.name**.

| Field   |      Mandatory?      |  Description |
|:----------:|:-------------:|:------|
| **metadata.name** |    **YES**   | Specifies the name of the CR. |
| **{another_field}** |    **{YES/NO}**   | {Field description} |
