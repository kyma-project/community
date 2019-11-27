---
title: Release strategy
label: internal
---

Read about the release strategy in Kyma.

## Terminology

### Release types

For release versioning, Kyma uses the approach similar to [Semantic Versioning](https://semver.org/). However, while Semantic Versioning focuses on APIs, the release process for Kyma refers to the complete Kyma deliverable. Despite this difference, Kyma follows the same release types:

| Type | Scope |
|-----|---------|
| **MAJOR** | Introduces breaking changes |
| **MINOR** | Introduces a new functionality in a planned schedule |
| **PATCH** **hot.fix** | Introduces a backward-compatible bug fix as an unscheduled release |

### Function freeze

The **function freeze** refers to the point in time when a new release branch is created from the master branch. After the function freeze, no new features are included in the release. This is typically on Monday at 9:00 AM CET/CEST every 4 weeks.

### Roles

These persons actively participate in release preparation and execution:

- **Product Lead** ensures that all capabilities provide a valuable and consistent product when it comes to its features and usability.

- **Release Manager** is responsible for the overall release:
  - Ensures that all persons involved in the release follow the required processes and standards
  - Decides when to create a new release candidate
  - Makes the final call for publishing the release candidate
  - Creates and manages the release schedule and release plan

- **Product Manager** coordinates the process of gathering the input for the release notes.

- **Technical Writer** writes and publishes the release notes as a blog post on the `kyma-project.io` website, based on the input from Product Owners.

- **Developer** ensures that the release branch contains all required features and fixes.

- **Release Master** prepares all mandatory release artifacts and conducts technical release activities described in the release process.

- **Release Publisher** prepares and publishes the social media content based on the release notes.

## Release requirements and acceptance criteria

### Epics and user stories

The functional implementations that are included in the release are documented as GitHub issues. A contributor completes an issue bound by the license agreement and contributing guidelines. All proposed changes must pass the obligatory Continuous Integration checks and require the maintainers' review and approval.

You can close an epic or an issue if the new functionality meets the contributing rules. These include, but are not limited to, the following points:

- Automated unit, integration and end-to-end tests verify this functionality.
- This functionality is documented.

### Test coverage

In Kyma, we follow the test strategy according to which you can verify all functionalities through automated testing that you execute either through the CI or the release pipeline.

The Kyma community must provide the highest possible test coverage to minimize any potential risks and ensure the release is functional and stable.

Consider implementing tests that go beyond verifying the functionality of the system. This includes integration, performance, and end-to-end tests.

### Compliance rules

All Kyma organization members involved in the release process comply with contribution rules defined for Kyma and ensure that all checks built into the continuous integration process pass. These requirements represent the absolute minimal acceptance criteria needed for a release from a compliance perspective.

Additionally, manual pre-release verification must cover these areas:

- the use of the third-party software and components that require licensing
- the use and storage of personal data subject to the GDPR requirements

### Security

Open-source projects, like any other, must ensure secure development. You can provide it by constantly raising the awareness of everyone involved in the project. In Kyma, we strive to make all developers aware of the secure development requirement. For example, we conduct threat modeling workshops during which the participants proactively think about possible security threats in the architecture, infrastructure, and implementation.

The Release Manager in Kyma takes care of formal security validation activities performed before major releases. The results of these activities influence the release decision. Lack of attention to security topics can result in release delay.

## Release schedule

A scheduled release follows the planning cycles for the Kyma development that typically take four weeks. At the beginning of each planning cycle, the Product Lead communicates the specific time plan for a release. After reaching the end of the development cycle, Kyma developers create a release candidate.

All planned releases happen at the end of a planned development cycle. Only the Product Lead and Release Manager can delay or resign from such a release if there are reasons that justify such a decision.

We do not schedule patches but rather provide them on-demand if there is a need to address high priority issues before the next planned release. In this context, we explicitly define high priority issues as those which either affect multiple Kyma installations or result in the financial loss for the users. The Release Manager is involved in the decision to execute a patch release.

### Nightly and weekly builds

It is our goal to have fully automated builds generated on a daily or nightly basis. Such builds can be made available to the community to validate the newest functionality and corrections implemented in development or testing environments, but are not intended for production use.

Once per week, we should update the installation with the newest, weekly build. During the following week, automated tests must run on the cluster to validate the stability and performance of the environment, and ensure that the results remain unchanged for a longer period of time. This also supports the whole release strategy by improving the quality of the product and increasing the public confidence in it.

### Release candidates

After reaching the end of the development cycle, the Kyma developers create a release candidate, published just like any other release in GitHub. Anyone can test it and provide feedback. Depending on whether this is a minor or major release, the period of time for this validation varies:

- two working days for minor releases
- one week for major releases

If during testing a release candidate any **critical issues** are identified, the issues should be documented in GitHub and communicated to both Release Master and Release Manager as soon as possible. The Release Master and Release Manager should be informed about the impact of the issue, time to fix, and any potential implications when fixing it. Not later than when the fix is implemented and merged into the master branch, the Release Manager makes a decision whether or not to cherry-pick the fix to the release branch.

Once testing and validation is complete with no more issues identified, the final release is prepared and executed.

### Planning start

The Product Lead coordinates the product planning, following the planning cycles for the Kyma development.

### Planning end

After completing the planning process, the theme and expected scope of the release is clear. All maintainers know what they have to work on for the release.

At this time, the Release Manager publicly communicates the planned release schedule for Kyma. This communication should include key features or fixes expected to fall within the scope of the release. Possible communication channels include blog posts, social media, and Core SIG meetings.

### Development start

After completing the planning, the Product Lead explicitly hands over the release to the engineering teams during the handover meeting. At this point, the development phase of the new release officially starts.

During the development phase, the Release Manager keeps track of the development and takes early actions required to keep the release on track. When closer to the release date, the Release Manager and the Product Lead discuss the status of the release and decide if there are any additional checkpoints or actions to address.

The Core SIG leaders include the external communication regarding the progress and status of each release to the agenda of the Core SIG bi-weekly meetings.

### Development end

This is the last day of the period planned for the release development.

### Release decision

The Release Manager and the Product Leadership decide to publish a release, based on its status, readiness and compliance of the developed functionality. If any issue is blocking the release, the Release Manager actively discusses the issue, implications and risks with the relevant developer(s), Product Owner and Release Engineer to make a decision to go/no-go.

### Release execution

The release execution is a manual process and is not part of an automated nightly or weekly build. It begins when the Release Engineer creates a release branch with all release artifacts. The process finishes after publishing both the artifacts and the related documentation as a release candidate that can be promoted to a final release.

During the execution of a release, all automated end-to-end test scenarios are validated. If any scenario fails, the release process automatically stops and no release is possible.

The Release Manager notifies the Kyma community on a public Slack channel when a release candidate is available, at which time the community can validate it and provide their feedback. If there are release-related issues identified, the Release Manager decides if any represent blockers for the release.

### Release publishing

The final release is available in the GitHub releases, including the installation instructions. It also includes the complete changelog that lists all pull requests merged into this release.

A Technical Writer publishes the blog post on the public Kyma website to announce the release. The post includes release notes that describe new features, as well as resolved and known issues.

You can learn about the new release from additional notifications published on social media, Slack channels, or announcements made during the Core SIG meetings. Alternatively, you can [subscribe](#subscription-to-releases-subscription-to-releases) to GitHub notifications and receive updates each time a new release is available.

## Release scope

For each of the planning periods described in the release schedule, GitHub epics and issues define and document the release scope with regards to functionality, corrections, and even non-functional requirements. The collection of all documented issues within a release represents the expected scope of that release that the whole organization and all teams define and commit to. The corresponding release in ZenHub identifies all issues and epics that fall within the final release scope.

When planning the release scope, all persons involved in the release must take the GDPR requirements into consideration when processing or storing personal data.

### Version maintenance

There is no guaranteed support for the bug fixes in the previous Kyma versions. The default strategy is to upgrade to the latest version. The Kyma project encourages the community to always upgrade to the latest release.

### Deprecation and backward-compatibility

The 1.0 release and further release versions impose clear expectations regarding the depreciation and backward-compatibility of Kyma versions to ensure some level of stability for the users. This can mean a period of time in which you should not change the provided functionality. This is the practice that other open-source projects also follow.
