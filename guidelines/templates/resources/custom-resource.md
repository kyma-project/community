---
title: {CRD kind}
type: Custom Resource
---

> This document is a ready-to-use template for a Custom Resource (CR) document that provides the sample Custom Resource and description of its elements for a given CRD. For details, see the existing CR documents for the [Installer](https://github.com/kyma-project/kyma/blob/master/docs/kyma/docs/040-installation-custom-resource.md) and the [Api-gateway](https://github.com/kyma-project/kyma/blob/master/docs/api-gateway/docs/011-api-custom-resource.md).

> **NOTE:** Blockquotes in this document provide instructions. Remove them from the final document.


The {CRD name} Custom Resource Definition (CRD) is a detailed description of the kind of data and the format used to {rest of the CRD description}. To get the up-to-date CRD and show the output in the yaml format, run this command:

```
kubectl get crd {CRD name} -o yaml
```

## Sample Custom Resource

> In this section, describe a Custom Resource created using the CRD mentioned in the introductory section. Describe its functionality, stressing all the options used. Provide the Custom Resource code sample in the applicable, ready-to-use format.
In case there are more examples relating to a given Custom Resource, provide all of them together with an appropriate description.

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
