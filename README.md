# Slf4Go - A simple Logging library

An easy-to-use logging libraty in Go, modeled after Java's Slf4j.

![Version](https://img.shields.io/badge/Version-0.1.0-blue)
![Released](https://img.shields.io/badge/unreleased-green)

[![Author](https://img.shields.io/badge/Author-M.%20Massenzio-green)](https://github.com/massenz)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![OS Debian](https://img.shields.io/badge/OS-Linux-green)

### Copyright & Licensing

**The code is copyright (c) 2022 AlertAvert.com. All rights reserved**<br>
The code is released under the Apache 2.0 License, see `LICENSE` for details.

# Usage

This is a simple logging package, resembling the functionality that is available in typical Java logging libraries (such as Log4j).

For an example usage see `example.go`.

Getting a `logger`:

```go
	var log = logging.NewLog("can.be/anything")
```
the name of the logger will be emitted at each log line, along with the logging level and the source file and line no. of where the logging happened:

```text
2022/06/03 22:48:35 example.go:17: example[INFO] An INFO message
```

Depending on the `logging.LogLevel` for the `Log`, some logs will be emitted, and other may not:

```go
log.Level = logging.INFO
log.Info("An INFO message")
log.Debug("This will NOT be logged")
```

The default level is `INFO`.

There are a couple of "special" convenience logs: a `RootLog` that can be used to quickly log when no special settings are needed, and a `NullLog` which "logs to nowhere" (technically, to `/dev/null`) which can be useful when injected, for example, during tests, where logs would pollute the tests' output.

# Build & Run

Use the package in your program by importing it:

    import "github.com/massenz/slf4go/logging"

The `example` script can be run with:

```shell
$ go build -o bin/log-sample example.go
$ bin/log-sample -trace
```
