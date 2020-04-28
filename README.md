# Bugsnag handler

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/emperror/handler-bugsnag/CI?style=flat-square)](https://github.com/emperror/handler-bugsnag/actions?query=workflow%3ACI)
[![Codecov](https://img.shields.io/codecov/c/github/emperror/handler-bugsnag?style=flat-square)](https://codecov.io/gh/emperror/handler-bugsnag)
[![Go Report Card](https://goreportcard.com/badge/emperror.dev/handler/bugsnag?style=flat-square)](https://goreportcard.com/report/emperror.dev/handler/bugsnag)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.12-61CFDD.svg?style=flat-square)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/emperror.dev/handler/bugsnag)

**Error handler integration for [Bugsnag](https://bugsnag.com).**


## Installation

```bash
go get emperror.dev/handler/bugsnag
```


## Usage

```go
package main

import (
	"emperror.dev/handler/bugsnag"
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
