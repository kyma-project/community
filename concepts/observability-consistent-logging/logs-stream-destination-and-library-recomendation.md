# Streams
<!-- markdown-link-check-disable -->
In Unix systems, we have 2 output streams: `stderr` and `stdout`. According to [GNU libc documentation](https://www.gnu.org/software/libc/manual/2.36/html_node/Standard-Streams.html)
and [Posix](https://pubs.opengroup.org/onlinepubs/9699919799/functions/stderr.html), the purpose of those streams is as follows:
<!-- markdown-link-check-enable -->
- `stdout` is the one to which conventional output should be printed (for example, something that can be processed by other tools using piping).
- `stderr` is the one to which diagnostic output should be printed.

In our case, we have two options:

1. Log errors, panics, and fatals to `stderr`. Log everything else to stdout.
2. Log everything to `stderr`, because it's the diagnostic output, and filter logs by their level.

## Kubernetes

Kubectl doesn't show the origin of the logs because it's probably fetching logs directly from the Pod (this needs to be confirmed).

Kubernetes stores the logs in `/var/log`. Those logs contain information about log origin, for example:

```
/var/log/pods # cat default_stream-test-j6mbg_a0796c8c-bfdc-4f1e-acf3-1dffc1c9c9e5/log/0.log
2020-12-14T15:20:56.1424263Z stderr F "ERROR"
2020-12-14T15:20:57.1438556Z stdout F "TEST"
2020-12-14T15:20:57.1439158Z stdout F "INFO"
```

## Golang

Requirements for the library:

- Format the logs to JSON and TXT
- Ability to set timestamp format RFC3339
- Ability to suppress logs by level

Considered logging libraries for Go:

- Zap
- Zerolog
- apex/log

We shouldn't pick [logurs](https://github.com/sirupsen/logrus) because it's in the maintenance mode.

## Zap

Pros:

- Configuration of Zap looks very advanced
- Library can save logs in JSON or TXT formats. It has the ability to be extended.
- Looks very fast (according to the benchmark provided by Zap)
- Ability to chain loggers (add context)
- API is very intuitive
- Configurable level of filtering
- Possibility to set the date format
- Available log levels: `ERROR`, `INFO`, `FATAL`, `DEBUG`, `PANIC`, `WARN`, `DPANIC` (for development)
- Support for the `Caller` field

Cons:

-  Method `With`, which accepts unlimited arguments but will print only pairs such as `With("a","b","c")`, will produce only `a:"b"`.

## Zerolog

Pros:

- According to their benchmark, the fastest library
- Ability to chain loggers (add context)
- Library can save logs in JSON or TXT formats. It has the ability to be extended.
- API is a little different, but it's intuitive.
- Has a nice way of adding fields to the log, e.g. `log.Info().Str("a", "b")`.
- Possibility to set the date format
- Configurable level of filtering
- Available log levels: `ERROR`, `INFO`, `FATAL`, `DEBUG`, `PANIC`, `WARN`, `TRACE`, no level
- Support for the `Caller` field

Cons:

- It's not possible to log errors to `stderr` and other things to `stdout`. You can find more info in the [this issue](https://github.com/rs/zerolog/issues/150).

## apex/log

Pros:

- Ability to chain loggers (add context) using the `entry` struct
- Library can save logs in JSON or TXT formats. It has the ability to be extended.
- Configurable level of filtering
- Available log levels: `ERROR`, `INFO`, `FATAL`, `DEBUG`, `WARN`

Cons:

- No `Caller` field
- Cannot set the timestamp format. The default time format looks like this: `2020-12-22T12:52:39.906885+01:00`.
According to Go (`apex/apex_test.go`) and Python's `rfc3339-validator 0.1.2`, this logged time is a valid RFC3339 timestamp.
- It's not possible to log errors to `stderr` and other things to `stdout`.

## Run examples

Run the example in which `$1` is the name of the Go main file:

```bash
go run poc-streams-and-libraries/$1/$1.go 1>info.log 2>err.log
```

## Summary

**The recommended library for logging is Zap.**

Easy to use, flexible in configuration, and with good API.

I would recommend logging everything to `stderr` because:

- Zap, Zerolog, apex/log and, as I think, other libraries can log everything to `stderr`.
- It's possible to filter logs by level and it's not needed to filter the logs by stream.
- It requires less work on current components.

For already components, we should check if the current logger meets the requirements. If yes, we can use it. In case of impossibility to fulfill the requirements, we would need to switch to Zap.

Keep the logging format consistent as it's the key to unified logging in Kyma.
