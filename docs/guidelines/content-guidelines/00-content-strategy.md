# Content Strategy

## Purpose and Audience

This content strategy focuses on the publicly available documentation for the open-source [Kyma project](https://kyma-project.io). The source of the documentation displayed on the website is stored in [GitHub](https://github.com/kyma-project/kyma) and the content is written in [Markdown](https://daringfireball.net/projects/markdown/).

The assumed readers of this guide and contributors to the documentation have some basic knowledge of technical writing.

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

**Target Group**: Software Developers who quickly want to install Kyma modules on k3d.

Contains a guide that covers typical steps you need to perform to get started.

### Modules

**Target Group**:
  
- Software Developers leveraging all the Kyma modules' functionalities.
- Admins/Operators who make sure their Kyma modules are configured as needed.

Under this tab, there are subtabs dedicated to all Kyma modules. Content of the subtabs is pulled from the `docs/user` folders in module repositories. Module documentation covers: general module description, feature scope, "how-to" instructions that enable users to accomplish a task, installation and configuration instructions, backup info, security documentation, troubleshooting guides, and detailed information, such as architecture diagrams, or configuration charts.

### Community Modules

**Target Group**:
- Software Developers who want to extend Kyma's capabilities by installing modules developed by the Kyma community.
- Contributors who develop and maintain community-driven modules.
- Operators who manage community module deployments in production environments.

Contains comprehensive documentation on community-contributed modules, including:
- Step-by-step installation guides for community modules
- Update and upgrade procedures for existing community modules
- Dedicated subtabs for each available community module with detailed usage instructions, API references, and examples

### User Interfaces

**Target Group**:
- Software Developers who want to learn more about interfaces provided by Kyma.

Contains comprehensive documentation on Kyma's user interfaces, including:
- **Kyma Dashboard**: Web-based graphical interface for managing Kyma clusters, including cluster overview, module management, resource visualization, and configuration options
- **Kyma CLI**: Command-line interface documentation covering installation, available commands, configuration options, and automation scripts

### Operation Guides

**Target Group**: Admins/Operators who make sure the Kyma cluster is configured as needed and keeps running in a healthy and secure way.

Includes installation and configuration instructions, backup info, "how-to" instructions that enable users to uninstall or upgrade a module, and troubleshooting guides.

### Glossary

**Target Group**: Anyone who wants to look up terms they’re not familiar with.

Explains basic terms, with a focus on terms specific to Kyma.