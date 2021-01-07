# Improvement of log messages usability

## Motivation

We want to provide a common, unified solution to logging across Kyma components, which will enable users and maintainers to easily track specific events happening in Kyma's ecosystem. The solution should be independent of the platform (AWS, Azure, GCP) and easy to introduce in existing components

Users, developers, and maintainers want to have an easy way to track events happening in Kyma, problems that occur, and debug faulty components. This will make understanding of Kyma's ecosystem easier and more transparent

## Solution regarding the log messages

Every component should produce log entries in order to make the problem diagnosis process quick and effective

Key assumptions for the log messages:

- Transparency - every entry must be as short as possible but containing all essential information about the process
- Tracing - entries sequence must be connected so there is an easy way to filter logs concerning a specific operation
- GDPR-compliance - entries must not contain any sensitive data (password, personal data in the meaning of GDPR)
- No "responsibility overlapping" - Because there are other tools to help in diagnosis, logs, in general, should not contain the information which can be taken from that tools (i.e there's tracing which can be used for timings)

Additional assumptions:

- If there is one quick solution for an occurring issue, add it to the message on the DEBUG level so the operator could use it while debugging
- If operation will be retried, say when
- JSON format should be enabled only on environments where the monitoring is enabled. Text format is more readable for debugging purposes so make it configurable
- In case of debug level add the error stack tracing
- Don't put redundant data into the logs. If something failed 10 times, log that it failed 10 times instead of printing the same error 10 times

## Log format

To unify the logs, therefore to make debugging process much easier and to make logs parsing also super easy, I'd like to propose a single log format so every service, job and all the components could follow:

- timestamp - ISO 8601 (or RFC3339) Date and time with timezone. For example: "2020-12-08T10:33:45+00:00"
- level - logging level. For example: "ERROR"
- message - information with the presented format (notice that there no additional data that could be duplicated in the context structure):
    - error and fatal message: Past tense started with for example "Failed to...", after that the error wrapped with some meaningful context but without additional "failed to" or "error occurred". For example: "Failed to provision runtime: while fetching release: while validating relese: release does not contain installer yaml"
    - info message: Present Continuous tense for the things that are about to be done, for example "Starting processing..." or past tense for the things that are finished, like "Finished successfully!"
    - notice and warning message: A short explanation on what happened and what this can cause. For example: "Tiller configuration not found in the release artifacts. Proceeding to the Helm 3 installation..." or "Connection is not yet established. Retrying in 5 minutes..."
- context - structure of a text contextual information such as operation (for example: "starting workers"), handler/ resolver (for example: "ProvisionRuntime"), controller, resource namespaced name (for example: "production/application1", operation ID, instance ID, operation stage and so on. User must be able to filter the logs that are needed so all the info provided here must be useful and be an unique minimal set for every operation. User must be able to find the needed resource in some store so provide here a name instead of an ID if it's easier to use later
- trace_id - 16-byte numeric value as Base16-encoded string. It'll be passed through a header, so the user can filter all the logs regarding the whole business operation in the whole system
- span_id - 16-byte numeric value as Base16-encoded string. It'll be randomly generated for each request handling so the user can filter the component logs for a specific operation handling
> trace_id and span_id are required in the log to be compliant with the [OpenTelemetry standards](https://github.com/open-telemetry/oteps/pull/114/files)

### Log format examples

#### Key-Value Pairs
```text
2020-12-15T07:26:45+00:00 WARNING Tiller configuration not found in the release artifacts. Proceeding to the Helm 3 installation... context.resolver=ProvisionRuntime context.operationID=92d5d8fd-cbdc-4b7a-9bc3-2b2eccfcb109 context.stage=InstallReleaseArtifacts context.shootName=c-3a38b3a context.runtimeId=19eb9335-6c13-4d40-8504-3cd07b18c12f trace_id=0354af75138b12921 span_id=14c902d73a
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
  "trace_id": "0354af75138b12921",
  "span_id": "14c902d73a"
}
```

## What should **not** be logged

We should be careful about the information we put in logs. Every log should be easily connectable with the issue or the error but there should be no internal or confidential data provided in any log entry. Here is what should not be logged:

- authorization data such as passwords, usernames, certificates and tokens. Note that sometimes basic auth could be used in the url such as `http://username:password@example.com/` and other tokens like `http://www.example.com/api/v1/users/1?access_token=123` so you need to be extra careful in these cases
- personal identifiable information such as tenant names, sub account IDs, global account IDs and so on. If the personal identifiable information is required in the audit log it should be encrypted
- user input to avoid Log Injection. If user input is required in the log, it should be parsed and validated
