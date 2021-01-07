# Streams 
In Unix systems we have 2 output stream: `stderr` and `stdout`. The purpose of those streams are:
- stdout, to this stream should be printed conventional output ( for example.: something which can be processed by other tool using piping)
- stderr, to this stream should be printed diagnostic output.
according to [GNU libc documentation](https://www.gnu.org/software/libc/manual/html_node/Standard-Streams.html),
[When to Use STDERR and STDOUT](https://www.jstorimer.com/blogs/workingwithcode/7766119-when-to-use-stderr-instead-of-stdout)
and [Posix](https://pubs.opengroup.org/onlinepubs/9699919799/functions/stderr.html)

In our case we have two options:
- log errors, panics and fatals to stderr, log everything else to stdout 
- log everything to stderr, because it's diagnostic output and filter logs by their level.

## Kubernetes
Kubectl doesn't show the origin of the logs, because it's probably fetching logs directly from pod (need confirmantion).

Kubernetes stores the logs in `/var/log`. In those logs there is the information about log origin, example:
```
/var/log/pods # cat default_stream-test-j6mbg_a0796c8c-bfdc-4f1e-acf3-1dffc1c9c9e5/log/0.log
2020-12-14T15:20:56.1424263Z stderr F "ERROR"
2020-12-14T15:20:57.1438556Z stdout F "TEST"
2020-12-14T15:20:57.1439158Z stdout F "INFO"
```

## Golang
Requirements for Library:
- format the logs in JSON and text
- ability to set timestamp format

Considered Logging libraries for Go:
- Zap
- Zerolog
- apex/log

We shouldn't pick [logurs](https://github.com/sirupsen/logrus), because it's in maintenance mode.

## ZAP
pros:
- configuration of Zap looks very advanced
- library can log in JSON or text format. It has the ability to be extended.
- looks very fast (according to benchmark provided by zap)
- ability to chain loggers (add context)
- api is very intuitive
- configurable level of filtering
- possible to set date format
- available log levels: ERROR, INFO, FATAL, DEBUG, PANIC, WARN, DPANIC(for development),  

## Zerolog

pros:
- according th their benchmarks, the fastest library
- ability to chain loggers (add context)
- library can log in JSON or text format. It has the ability to be extended.
- api is a little different, but it's intuitive.
- possible to set date format
- configurable level of filtering
- available log levels: ERROR, INFO, FATAL, DEBUG, PANIC, WARN, TRACE, no level

cons
- it's not possible to log errors to stderr and other things to stdout, more [info](https://github.com/rs/zerolog/issues/150)

## apex/log
pros:
- ability to chain loggers (add context) using `entry` struct
- library can log in JSON or text format. It has the ability to be extended.
- configurable level of filtering
- available log levels: ERROR, INFO, FATAL, DEBUG, WARN

cons:
- cannot set timestamp format. The default time format looks like this:  `2020-12-22T12:52:39.906885+01:00`
- it's not possible to log errors to stderr and other things to stdout


## Run Examples
run example :
```bash
go run $1.go 1>info.log 2>err.log
```
where '$1' is the name of go main file.

# Summary

I would recommend logging everything to stderr, because:
- zap, zerolog and apex/log and I think other libraries can log everything to stderr 
- it's possible to filter logs by level and it's not needed to filter the logs by stream
- requires less work on current components

For new components, the recommendation is `Zap`.
For already components, we should use what we currently have and unified log messages.
In case of impossibility to fulfil the requirements, the exchange for `Zap`  is recommended.

Keep the logging format consistently it's key for unified logging in Kyma.
