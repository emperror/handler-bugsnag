# Bugsnag handler

[![CircleCI](https://circleci.com/gh/emperror/handler-bugsnag.svg?style=svg)](https://circleci.com/gh/emperror/handler-bugsnag)
[![Go Report Card](https://goreportcard.com/badge/handler.emperror.dev/bugsnag?style=flat-square)](https://goreportcard.com/report/emperror.dev/handler/bugsnag)
[![GolangCI](https://golangci.com/badges/github.com/emperror/handler-bugsnag.svg)](https://golangci.com/r/github.com/emperror/handler-bugsnag)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://pkg.go.dev/emperror.dev/handler/bugsnag)

**Error handler integration for [Bugsnag](https://bugsnag.com).**


## Installation

```bash
go get handler.emperror.dev/bugsnag
```


## Usage

```go
package main

import (
	"handler.emperror.dev/bugsnag"
)

func main() {
	apiKey := "key"

	handler := bugsnag.New(apiKey)
}
```


## Development

When all coding and testing is done, please run the test suite:

``` bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
