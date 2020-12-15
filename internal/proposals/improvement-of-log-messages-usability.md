# Improvement of log messages usability

## Motivation

We want to provide a common, unified solution to logging across Kyma components, which will enable users and maintainers to easily track specific events happening in Kyma's ecosystem. The solution should be independent of the platform (AWS, Azure, GCP) and easy to introduce in existing components

Users, developers, and maintainers want to have an easy way to track events happening in Kyma, problems that occur, and debug faulty components. This will make understanding of Kyma's ecosystem easier and more transparent

## Solution regarding the log messages

Every mission-critical component (or a critical part of the component) must produce log entries in order to make the problem diagnosis process quick and effective

Key facts about log entries:

- Transparency - Entry must be as short as possible but containing all essential information about the process
- Log traces - Must be easily possible to connect the following log entries with a sequence (in the scope of a whole business process). Log entries should contain some unique trace identifier. The trace identifier is useful only when it is known to the operator - can be easily obtained by him and then used to search the logs
- GDPR-compliance - Must not contain any sensitive data (password, personal data in the meaning of GDPR)
- No "responsibility overlapping" - Because there are other tools to help in diagnosis, logs, in general, should not contain the information which can be taken from that tools (i.e there's tracing which can be used for timings)
    > TODO: Ask Artur about the opinion on the log format. Ask also about the known tools used for such resposibilities

Additionally, we need an approach to the "external" components used in Kyma (the ones not implemented in the Kyma team). We should have an idea on how to approach this problem and present the example solution

## Log format

- timestamp - ISO 8601 Date and time with timezone. For example: "2020-12-08T10:33:45+00:00"
- level - logging level. For example: "ERROR"
- message - information with the presented format (notice that there no additional data that could be duplicated in the context structure):
    - error and fatal message: Past tense started with for example "Failed to...", after that the error wrapped with some meaningful context but without additional "failed to" or "error occurred". For example: "Failed to provision runtime: while fetching release: while validating relese: release does not contain installer yaml"
    - info message: Present Continuous tense for the things that are about to be done, for example "Starting processing..." or past tense for the things that are finished, like "Finished successfully!"
    - notice and warning message: A short explanation on what happened and what this can cause. For example: "Tiller configuration not found in the release artifacts. Proceeding to the Helm 3 installation..." or "Connection is not yet established. Retrying in 5 minutes..."
- context - structure of a contextual information such as operation (for example: "starting workers"), handler/ resolver (for example: "ProvisionRuntime"), controller, resource namespaced name, operation ID, instance ID, operation stage and so on. User must be able to filter the logs that are needed so all the info provided here must be useful and be an unique minimal set for every operation. User must be able to find the needed resource in some store so provide here a name instead of an ID if it's easier to use later
- traceID - 16-byte numeric value as Base16-encoded string. It'll be randomly generated for each request handling
    > TODO: Check if trace ID is necessary and how it should be implemented in our case. Suleyman will check it 

### Log format examples

#### Key-Value Pairs
```text
2020-12-15T07:26:45+00:00 WARNING Tiller configuration not found in the release artifacts. Proceeding to the Helm 3 installation... context.resolver=ProvisionRuntime context.operationID=92d5d8fd-cbdc-4b7a-9bc3-2b2eccfcb109 context.stage=InstallReleaseArtifacts context.shootName=c-3a38b3a context.runtimeId=19eb9335-6c13-4d40-8504-3cd07b18c12f traceID=0354af75138b12921
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
  "traceID": "0354af75138b12921"
}
```

## Log streams and what to put inside of them

> TODO: Check if we can distinquish the output streams in Kubernetes. Damian will check it

- stdout: NOTICE and WARN levels
- stderr: ERROR and FATAL levels

TRACE, DEBUG and INFO levels should not be logged on production clusters. All the important context should be provided in the NOTICE and higher levels so there is no need to have a tone of simmilar INFO logs. On the other environments these three levels could be enabled for the debugging reason and put into the stdout

In case of gathering metrics from the logs, INFO level logs could also be put into the stdout

## What should **not** be logged

> TODO: Consult it with the security experts

- authorization data such as passwords, usernames, certificates and tokens
- personal identifiable information such as tenant names

## Ideas for external components

> TODO: Brainstorm on it. Suleyman suggested that it could be done via adapter sidecars and FluentD. Check these

# Workspace - brainstorming

- I assume that error level logs are presented when an action from a support is needed and warning level when the action is not needed yet but may be in the future. Therefore there should be a place where the suggestion on what to do as a support (or user that want to repair it on their own). Only if it is a known issue and there is probably one way to fix it
- If operation will be retried, say when
- Logs shouldn't rely on the previous logs because context may change. All the needed info should be in a single log entry.
- In case of debug level add the error stack tracing
- Errors could have a specific ID so their metrics could be easily tracked (where this ID will be created and stored?)
- **Refactor the logs! If you took part in some incident debugging process, note down which logs were especially useful and what was missing**
- Don't put redundant data into the logs. If something faild 10 times, say that it failes 10 times instead of printing the same error 10 times
- Use context whereever it's possible. Add resource ID, tenant ID, operation ID. Everything YOU CAN FILTER with some base (ID that you cannot use is useless)
- Logs should contain in their context the operation which was requested
- Here is a document from Telemetry - https://github.com/open-telemetry/oteps/blob/a86472bac9a9695da438472aebc72d904a25b9e5/text/logs/0114-log-correlation.md
