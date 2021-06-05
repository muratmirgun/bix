
# log

This extremely simple log library is intended for use with micro service
applications, which have only two logging targets: the console while
developing the application, and a log aggregation service in production.

This library was inspired by:

* [Logging Packages in Golang](https://www.client9.com/logging-packages-in-golang/)
* [logrus](https://github.com/sirupsen/logrus)
* [Benchmarking Logging Libraries for Go](https://github.com/imkira/go-loggers-bench)

## Features

* The `fmt` library is not used to minimize memory allocations.
* Logs to `stdout` by default.
* Can be configured to send messages to a UDP log aggregator.

## Usage

By default, `log` will emit messages on `os.Stderr` which is supposed to be
an unbuffered stream. The destination can be changed with: 

```go
log.SetOutput(os.Stdout)
``` 

To prevent console logging from being visible at all, change the destination
to `/dev/null`:

```go
devnull, err := os.OpenFile(os.DevNull, O_WRONLY, 0666)
if err != null {
	panic(err)
}
defer devnull.Close()
log.SetOutput(devnull)
```

To send output to a UDP log aggregator, just set the address and port of the
service as follows:

```go
log.SetServer("10.10.10.10:8080")
```

Now every log message is formatted in JSON and will be sent to the aggregator:

```go
log.Info("The quick brown fox")
// Output: {"time":1554370662469959000,"name":"main","level":"INFO","message":"The quick brown fox"} 
```


