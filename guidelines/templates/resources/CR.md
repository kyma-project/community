---
title: {CRD name}
type: Custom Resource
---

> This document is a ready-to-use template for a Custom Resource (CR) document that provides the sample Custom Resource and description of its elements for a given CRD. For details, see the existing CR documents for the [Installer](https://github.com/kyma-project/kyma/blob/master/docs/kyma/docs/040-installation-custom-resource.md) and the [Api-gateway](https://github.com/kyma-project/kyma/blob/master/docs/api-gateway/docs/011-api-custom-resource.md).

> **NOTE:** Blockquotes in this document provide instructions. Remove them from the final document.



The {CRD name} Custom Resource Definition (CRD) is a detailed description of the kind of data and the format used to {rest of the CRD description}. To get the up-to-date CRD and show the output in the yaml format, run this command:

```
kubectl get crd {rest of the command}
```

## Sample Custom Resource

This is a sample CR that {description of the CR functions}.

```
apiVersion:
kind:
metadata:
  name:
{another_field}:
```

This table analyses the elements of the sample CR and the information it contains:

| Field   |      Mandatory?      |  Description |
|:----------:|:-------------:|:------|
| **apiVersion** |  **YES** | Defined basing on the **group** and **version** fields of the CRD **spec** section. |
| **kind** |    **YES**   | Defined basing on the **names: kind** field of the CRD **spec** section. |
| **metadata.name** |    **YES**   | Specifies the name of the CR. |
| **{another_field}** |    **{YES/NO}**   | {Field description} |
