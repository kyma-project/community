# Improvement of log messages usability

> TODO: Ask TWs for a language review

## Motivation

We want to provide a common, unified solution to logging across Kyma components, which will enable users and maintainers to easily track specific events happening in Kyma's ecosystem. The solution should be independent of the platform (AWS, Azure, GCP) and easy to introduce in existing components

Users, developers, and maintainers want to have an easy way to track events happening in Kyma, problems that occur, and debug faulty components. This will make understanding of Kyma's ecosystem easier and more transparent

## Solution regarding the log messages

Every mission-critical component (or a critical part of the component) must produce log entries in order to make the problem diagnosis process quick and effective

Key assumptions for the log messages:

- Transparency - every entry must be as short as possible but containing all essential information about the process
- Tracing - entries sequence must be connected so there is an easy way to filter logs concerning a specific operation
- GDPR-compliance - entries must not contain any sensitive data (password, personal data in the meaning of GDPR)
- Agile - log messages and their context should be refactored. If the log was useful while debugging the incident, make other also that useful. If there was something missing in the log, report it and refactor
- No "responsibility overlapping" - Because there are other tools to help in diagnosis, logs, in general, should not contain the information which can be taken from that tools (i.e there's tracing which can be used for timings)
    > TODO: Ask Artur and Przemek for the review on the potential responsibility overlaping

Additionally, we need an approach to the "external" components used in Kyma (the ones not implemented in the Kyma team). We should have an idea on how to approach this problem and present the potential solution

## Log format

To unify the logs, therefore to make debugging process much easier and to make logs parsing also super easy, I'd like to propose a single log format so every service, job and all the components could follow:

- timestamp - ISO 8601 Date and time with timezone. For example: "2020-12-08T10:33:45+00:00"
- level - logging level. For example: "ERROR"
- message - information with the presented format (notice that there no additional data that could be duplicated in the context structure):
    - error and fatal message: Past tense started with for example "Failed to...", after that the error wrapped with some meaningful context but without additional "failed to" or "error occurred". For example: "Failed to provision runtime: while fetching release: while validating relese: release does not contain installer yaml"
    - info message: Present Continuous tense for the things that are about to be done, for example "Starting processing..." or past tense for the things that are finished, like "Finished successfully!"
    - notice and warning message: A short explanation on what happened and what this can cause. For example: "Tiller configuration not found in the release artifacts. Proceeding to the Helm 3 installation..." or "Connection is not yet established. Retrying in 5 minutes..."
- context - structure of a text contextual information such as operation (for example: "starting workers"), handler/ resolver (for example: "ProvisionRuntime"), controller, resource namespaced name (for example: "production/application1", operation ID, instance ID, operation stage and so on. User must be able to filter the logs that are needed so all the info provided here must be useful and be an unique minimal set for every operation. User must be able to find the needed resource in some store so provide here a name instead of an ID if it's easier to use later
- traceid - 16-byte numeric value as Base16-encoded string. It'll be passed through a header, so the user can filter all the logs regarding the whole business operation in the whole system
- spanid - 16-byte numeric value as Base16-encoded string. It'll be randomly generated for each request handling so the user can filter the component logs for a specific operation handling
> traceid and spanid are required in the log to be compliant with the [OpenTelemetry standards](https://github.com/open-telemetry/oteps/pull/114/files)
> TODO: Ask Andreas for the review on the OpenTelemetry standards

### Log format examples

#### Key-Value Pairs
```text
2020-12-15T07:26:45+00:00 WARNING Tiller configuration not found in the release artifacts. Proceeding to the Helm 3 installation... context.resolver=ProvisionRuntime context.operationID=92d5d8fd-cbdc-4b7a-9bc3-2b2eccfcb109 context.stage=InstallReleaseArtifacts context.shootName=c-3a38b3a context.runtimeId=19eb9335-6c13-4d40-8504-3cd07b18c12f traceid=0354af75138b12921 spanid=14c902d73a
```

#### JSON
```json
{
  "timestamp": "2020-12-15T07:26:45+00:00",
  "level": "WARNING",
  "message": "Tiller configuration not found in the release artifacts. Proceeding to the Helm 3 installation...",
  "context": {
    "resolver": "ProvisionRuntime",
    "operationID": "92d5d8fd-cbdc-4b7a-9bc3-2b2eccfcb109",
    "stage": "InstallReleaseArtifacts",
    "shootName": "c-3a38b3a",
    "runtimeID": "19eb9335-6c13-4d40-8504-3cd07b18c12f"
  },
  "traceid": "0354af75138b12921",
  "spanid": "14c902d73a"
}
```

## Log streams and what to put inside of them

> Damian will take care of it. He's working on his own proposal so we can merge them together later

- stdout: NOTICE and WARN levels
- stderr: ERROR and FATAL levels

TRACE, DEBUG and INFO levels should not be logged on production clusters. All the important context should be provided in the NOTICE and higher levels so there is no need to have a tone of similar INFO logs. On the other environments these three levels could be enabled for the debugging reason and put into the stdout

In case of gathering metrics from the logs, INFO level logs could also be put into the stdout

## What should **not** be logged

We should be careful about the information we put in logs. Every log should be easily connectable with the issue or the error but there should be no internal or confidential data provided in any log entry. Here is what should not be logged:

> TODO: Consult it with the security experts

- authorization data such as passwords, usernames, certificates and tokens. Note that sometimes basic auth could be used in the url such as `http://username:password@example.com/` and other tokens like `http://www.example.com/api/v1/users/1?access_token=123` so you need to be extra careful in these cases
- personal identifiable information such as tenant names, sub account IDs, global account IDs and so on

## **Ideas** for external components

> Suleyman will take care of it. FluentBit seems to be a promising solution

# Workspace - brainstorming

> TODO: Transform these ideas into the solutions
- If there is a quick solution for an error or simple suggestion for a warning, add it to the message so the operator could use it while debugging
- If operation will be retried, say when
- JSON format should be enabled only on clusters where the monitoring is enabled. Text format is more readable for debugging purposes
- In case of debug level add the error stack tracing
- Don't put redundant data into the logs. If something faild 10 times, say that it failes 10 times instead of printing the same error 10 times
