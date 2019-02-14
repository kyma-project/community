# Release Strategy

Read about the release strategy in Kyma:

- [Terminology](#terminology)
  - [Release types](#release-types)
  - [Function freeze](#function-freeze)
  - [Roles](#roles)
- [Release requirements and Acceptance Criteria](#release-requirements-and-acceptance-criteria)
  - [Epics and user stories](#epics-and-user-stories)
  - [Test coverage](#test-coverage)
  - [Compliance rules](#compliance-rules)
  - [Security](#security)
- [Release schedule](#release-schedule)
  - [Nightly and weekly builds](#nightly-and-weekly-builds)
  - [Release candidates](#release-candidates)
  - [Development planning start](#development-planning-start)
  - [Development planning end](#development-planning-end)
  - [Development start](#development-start)
  - [Development end](#development-end)
  - [Release decision](#release-decision)
  - [Release execution](#release-execution)
  - [Release publishing](#release-publishing)
- [Release scope](#release-scope)
  - [Down-porting](#down-porting)
  - [Deprecation and backward-compatibility](#deprecation-and-backward-compatibility)

## Terminology

### Release types

For release versioning, Kyma follows the approach similar to [Semantic Versioning](https://semver.org/). However, while Semantic Versioning focuses on APIs, the release process for Kyma refers to the complete Kyma deliverable. Despite this difference, the release process in Kyma follows the same release types:

- **MAJOR** release to introduce breaking changes
- **MINOR** release to introduce a new functionality in a planned schedule
- **PATCH** release or a **hot.fix** to introduce a backwards-compatible bug fix as an unscheduled release

>**NOTE:** **Major version zero** refers to the rapid product development that is not production-ready yet. That is why, major changes introduced before the 1.0 release, like in versions 0.5 or 0.6, are not represented as increases in the major version number.

### Function freeze

The function freeze refers to the point in time during which you cannot merge new features to the release branch.

### Roles

These persons actively participate in the release preparation and execution process:

- **Head of Capabilities** ensures that all capabilities provide a valuable and consistent product when it comes to its features and usability.

- **Release Manager** is responsible for the overall release:
  - Ensures that all persons involved in the release follow the required processes and standards
  - Decides when to create a new release candidate
  - Makes the final call for publishing the release candidate
  - Creates and manages the release schedule and release plan

- **Product Manager** coordinates the process of gathering the input for the release notes.

- **Technical Writer** writes and publishes the release notes as a blog post on the `kyma-project.io` website, based on the input from Product Owners.

- **Developer** ensures that the release branch contains all required features and fixes.

- **Release Engineer** prepares all mandatory release artifacts and conducts the technical release activities described in the release process.

- **Release Publisher** prepares and publishes the social media content based on the release notes.


## Release requirements and acceptance criteria

### Epics and user stories

The functional implementations that are included in the release are documented as GitHub issues. A given team completes issues according to the Definition of Done (DoD) checklist that is applicable to the product and the team who implements them. Other developers must review your issue and a corresponding Product Owner, Capability Owner, or Quality Owner must accept it.

You can close an epic or an issue if the implemented functionality meets the respective DoD requirements. These include, but are not limited to, the following points:

- The functionality fulfills the acceptance criteria (functional correctness).
- Automated unit and end-to-end tests verify this functionality.
- This functionality is documented.

To accept an issue which does not comply with these requirements, you must receive an exceptional approval from the Product Owner and the Release Manager.

### Test coverage

The Kyma team follows the accepted test strategy according to which you can verify all functionalities through automated testing that you execute either through the CI or the release pipeline.

The Kyma team must provide the highest possible test coverage to minimize any potential risks and ensure the release is functional and stable.

Whenever possible, we must consider implementing tests that go beyond verifying the functionality of the system. This includes integration, performance, and end-to-end tests.

### Compliance rules

All Kyma organization members involved in the release process comply with contribution rules defined for Kyma and ensure that all checks built into the continuous integration process pass. These requirements represent the absolute minimal acceptance criteria needed for a release from a compliance perspective. These rules can change as Kyma approaches the 1.0 release.

Additional, manual pre-release verification needs to be done for the following:
- the use of the third-party software and components that require prior approval
- the use and storage of personal data subject to the GDPR requirements

### Security

Open-source projects, like any other, must ensure secure development. You can provide it by constantly raising the awareness of everyone involved in the project. In Kyma, we strive to make all developers aware of the secure development requirement. For example, we conduct threat modeling workshops where teams proactively think about possible security threats in the architecture, infrastructure, and implementation. What is more, it is important to keep track of the issues identified during or as a result of threat modeling activities.

The Release Manager in Kyma takes care of formal security validation activities performed before major releases. The results of these activities influence the release decision. Lack of attention to security topics can result in the release delay.

## Release schedule
A planned or scheduled release follows the planning cycles for the Kyma development that typically take four weeks. At the beginning of each planning cycle, the Head of Capabilities communicates the specific time plan for a release. After reaching the end of the development cycle, the Kyma developers create a release candidate which in most cases is identical to the release published on GitHub.

All planned releases happen at the end of a planned development cycle. Only the Product Lead and Release Manager can delay or resign from such a release if there are reasons that justify such a decision.

We do not schedule patches but rather execute them on-demand if there is a need to address high priority issues before the next planned release. In this context, we explicitly define high priority issues as those which either affect multiple Kyma installations or result in the financial loss for the users. The Release Manager is involved in the decision to execute a patch release. However, we do not expect to have patch releases before the official Kyma 1.0 version.

### Nightly and weekly builds
It is our goal to have fully automated builds generated on a daily or nightly basis. The daily builds can be made available to the community to validate the newest functionality and corrections implemented, as long as there is a clear disclaimer that these builds are for development purposes and are not intended for the production use.

Once per week, we should update the installation with the newest, weekly build. During the following week, automated tests must run on the cluster to validate the stability and performance of the environment, and ensure that the results remain unchanged during a longer period of time. This also supports the whole release strategy by improving the quality of the product and increasing the public confidence in it.

### Release candidates

As already mentioned, we produce a release candidate by the end of the last day of each four-week planning iteration.

We publish the release candidate as any other release in GitHub. Anyone can test it and provide their feedback. Depending on whether this is a minor or major release, the period of time for this validation varies:
- two working days for minor releases
- one week for major releases

After this time, and ideally with no issues identified, we can promote the release candidate or, depending on the technical implementation, rebuild it and use it as the final release.

### Development planning start
The Head of Capabilities coordinates the product planning, following the planning cycles for the Kyma development.

### Development planning end
After completing the planning process, the theme and expected scope of the release is clear. All development teams know what they have to work on for this release and commit to completing this work.

At this time, the Release Manager publicly communicates the planned release schedule for Kyma. This communication should include key features or fixes expected to fall into the scope of the release. Possible communication channels include blog posts, social media, and the Core SIG meetings.

### Development start
After completing the planning, the Head of Capabilities explicitly hands over the release to the engineering teams during the handover meeting. At this point, the development phase of the new release officially starts.

The Kyma developers add new release features to the `master` branch of the Kyma repositories.

During the development phase, the Release Manager keeps track of the development with the help of the Scrum of Scrums (SoS) meetings held twice a week. This allows the Release Manager to take early actions required to keep the release on track. When closer to the release date, the Release Manager, the Head of Capabilities, and the Lead Product Engineer discuss the status of the release and decide if they are any additional checkpoints or actions to address.

The Core SIG leaders include the external communication regarding the progress and status of each release to the agenda of the Core SIG bi-weekly meetings. The Release Manager, or a person appointed as a stand-in, presents the release status during those meetings.

### Development end
This is the last day of the period planned for the release development.

### Release decision
The Release Manager and the Product Leadership decide to publish the release, based on the status and readiness of the developed functionality. The Release Manager communicates this decision before the release execution starts.

### Release execution
The Release Engineer executes the release at any time of producing the release artifact. This could result in a release candidate or a final release that is not an automated nightly or weekly build. The release execution starts when the process starts and finishes after publishing all artifacts and the related documentation.

During the execution of a release, Kyma developers run all end-to-end test scenarios. If any single scenario fails, the Release Engineer cannot publish the release.

The Release Manager or Kyma developers notify the Kyma community when the release candidate is available. They pass this information to the public on Slack, requesting the community to validate and provide their short-term feedback. The deadline for the feedback is the day before the planned final release. Ideally, the final release is the same as the last release candidate, avoiding any additional changes. If the Kyma developers identify any release-related issues, they asses them and decide together with the Release Manager if they must fix them and produce a new release candidate.

### Release publishing
The final release is available in the GitHub releases, including the installation instructions. It also includes the complete changelog that lists all pull requests merged into this release.

A Technical Writer publishes the blog post on the public Kyma website to announce the release. The post includes release notes prepared based on the input from the Product Owners.

You can learn about the new release from additional notifications published on social media, Slack channels, or announcements made during the Core SIG meetings.

Alternatively, you can [subscribe](./release-subscription.md) to GitHub notifications and receive updates each time a new release is available.

## Release scope
For each of the planning periods described in the release schedule, GitHub epics and issues define and document the release scope with regards to functionality, corrections, and even non-functional requirements. The collection of all documented issues within a release represents the expected scope of that release that the whole organization and all teams define and commit to. The corresponding release in ZenHub identifies all issues and epics that fall under the final release scope.

When planning the release scope, all persons involved in the release must take into consideration the GDPR requirements when processing or storing personal data.

### Down-porting
There is no guaranteed bug fix down-porting for the Kyma project. The default strategy is to upgrade to the latest version. All changes must align with the Kyma upgrade strategy.

The Kyma team does not plan to provide any patch releases before the first 1.0 production release and encourages the community to always upgrade to the latest release.

### Deprecation and backward-compatibility
The 1.0 release and further release versions impose clear expectations regarding the depreciation and backward-compatibility of Kyma versions to ensure some level of stability for the users. This can denote a period of time in which you should not change the provided functionality. This is the practice that other open-source projects also follow.
