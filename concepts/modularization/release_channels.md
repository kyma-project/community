---
Release channels in Kyma
---

## Modular Kyma 

Kyma gives you the freedom to choose from many of its modules and thus adjust the runtime to your specific needs. Releasing modules in three channels allows for testing and experimenting with new functionalities. The modular approach boosts Kyma’s extensibility and scalability. 

## Kyma module

A Kyma module is developed based on a Kyma component by a developing team called a Kyma Module Provider (KMP). Modules are released using a `ModuleTemplate` in one or more of the three release channels: Regular, Fast, and Beta.  For more details, see the [module documentation](https://github.com/kyma-project/kyma/tree/main/modules).

## Release channels

Sets of Kyma modules are deployed in three release channels representing different maturity levels of the modules. You can use modules from one or more release channels in your Kyma cluster, but each module can only be consumed from one channel and can occur only once in a release channel. For example, you can mix different modules from the Beta and Fast Channels in your development cluster, but you cannot deploy the same module in the regular and fast versions in one cluster.

### The Regular Channel

Only stable Kyma modules, verified in the Fast Channel, are released here. These modules are recommended for user production environments. 

### The Fast Channel

This channel offers a minimum period of two weeks for the preview of all new features and changes affecting your environments. After that, the modules are promoted to the Regular Channel. You are welcome to test the new features and provide feedback. However, it is recommended to use Fast releases in non-productive environments.

### The Beta Channel

As defined in the SAP Business Technology Supplemental Terms and Conditions, any Beta Functionality is offered solely for testing purposes. After a maximum period of six months, a KMP evaluates a Beta module and decides on its further usability and development.
Modules available in the Beta Channel are compliant with SAP corporate requirements and SAP Beta Functionality requirements. Still, due to their experimental nature, they may have an unstable API, and be discontinued at any time or changed without keeping their backward compatibility. Therefore, Beta releases are not intended for use in production environments.
