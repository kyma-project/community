# Improvement of log messages usability


The goal of this proposal is to provide a common, unified solution to logging across Kyma components, which will allow users and maintainers to easily track specific events happening in the Kyma ecosystem. The solution should be independent of the platform (AWS, Azure, GCP) and easy to introduce into the existing components.

Users, developers, and maintainers want to have an easy way to track events and issues that occur in Kyma, and to easily debug faulty components. This will make understanding of the Kyma ecosystem easier and more transparent.

## Solution 

Every component should produce log entries in order to make the issue diagnosis process quick and effective.

Key assumptions for log messages:

- Transparency - every entry must be concise and contain all essential information about the process.
- Tracing - entries sequence must be connected so that it's easy to filter logs concerning a specific operation.
- GDPR-compliance - entries must not contain any sensitive data (password, personal data in the meaning of GDPR).
- No "responsibility overlapping" - as there are other diagnostic tools, logs, in general, should not contain information that can be taken from these tools (for example, there's tracing which can be used for timings).

Additional assumptions:

- If there is one quick solution for an occurring issue, add it to the message on the DEBUG level so the operator could use it while debugging.
- If the operation will be retried, specify when.
- JSON format should be enabled only on environments where monitoring is enabled. Text format is more readable for debugging purposes so make it configurable.
- In case of debug level, add the error stack tracing.
- Don't put redundant data into the logs. If something failed 10 times, log that it failed 10 times instead of printing the same error 10 times.

## Log structure

To unify the logs, which will make the debugging process and logs parsing much easier, I'd like to propose a single log format that every service, job, and all the components should follow:

- **timestamp** - RFC3339 format of date, time, and timezone. For example: "2012-12-12T07:20:50.52Z".
- **level** - logging level. For example: "ERROR"
- **message** - human-readable information with the specified format (notice that there is no data duplication in the context structure):
    - error and fatal message: past tense started with, for example, `Failed to...`, after that the error wrapped with some meaningful context. For example: `Failed to provision runtime: while fetching release: while validating release: release does not contain installer yaml`.
    - info message: present continuous tense for the things that are about to be done, for example, `Starting processing...`, or past tense for the things that are finished, such as `Finished successfully!`.
    - notice and warning message: a short explanation of what happened and what this can cause. For example: `Tiller configuration not found in the release artifacts. Proceeding to the Helm 3 installation...` or `Connection is not yet established. Retrying in 5 minutes...`.
- **context** - structure of the contextual information, such as operation (for example: `starting workers`), handler/ resolver (for example: `ProvisionRuntime`), controller, resource-namespaced name (for example: `production/application1`), operation ID, instance ID, operation stage, and so on. Users must be able to filter the logs so all the info provided here must be a useful and unique minimal set for every operation. Users must be able to find the needed resource in some store so provide here a name instead of an ID if it's easier to use later.
- **traceid** - 16-byte numeric value as a base16-encoded string. It'll be passed through a header, so the user can filter all the logs regarding the whole business operation in the whole system.
- **spanid** - 16-byte numeric value as a base16-encoded string. It'll be randomly generated for each request handling so the user can filter the component logs for a specific operation handling.
>**NOTE:** The **traceid** and **spanid** fields are required in the logs to be compliant with the [OpenTelemetry standards](https://github.com/open-telemetry/oteps/pull/114/files). The standard is still in the development phase, so we should keep an eye on it and change it accordingly.

## Log format

Log format should be configured and changeable. I'd like to propose the **log.format** Helm chart key which could have two possible values: `json` or `text`. For example:

```yaml
log:
  format: "json"
```

In the `deployment.yaml` file or other component's container specification, there will be an environment variable, for example:

```yaml
    spec:
      containers:
        - env:
          - name: APP_LOG_FORMAT
            value: {{ .Values.global.log.format }}
          ...
```

### Key-value pairs example
```text
2012-12-12T07:20:50.52Z WARNING Tiller configuration not found in the release artifacts. Proceeding to the Helm 3 installation... {"context":{"resolver":"ProvisionRuntime","operationID":"92d5d8fd-cbdc-4b7a-9bc3-2b2eccfcb109","stage":"InstallReleaseArtifacts","shootName":"c-3a38b3a","runtimeID":"19eb9335-6c13-4d40-8504-3cd07b18c12f"},"traceid":"0354af75138b12921","spanid":"14c902d73a"}
```

### JSON example
```json
{
  "timestamp": "2012-12-12T07:20:50.52Z",
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

## What should not be logged

We should be careful about the information we put in logs. Every log should be easily connectable with the issue or the error but there should be no internal or confidential data provided in any log entry. Here is the list of items that should not be logged:

- Authorization data, such as passwords, usernames, certificates, and tokens. Note that sometimes basic auth or other tokens can be used in the URL, for example: `http://username:password@example.com/` or  `http://www.example.com/api/v1/users/1?access_token=123` so you need to be extra careful in these cases.
- Person-identifiable information, such as tenant names. If the person-identifiable information is required in the audit log, it should be encrypted.
- User input to avoid log injection. If user input is required in the logs, it should be parsed and validated.
