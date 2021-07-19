# Session Management

Basic session management service using golang

# Table of contents
  * [Usage](#usage)
    + [Get the package](#get-the-package)
    + [Prerequisites](#prerequisites)
    + [Build Application](#build-application)
    + [Start the Application](#start-application)
    + [Run Test](#run-test)
    + [Genrate Mock with counterfeiter](#generate-mock-using-counterfeiter)
    + [APIs](#apis)
    
    
    
## Usage
### Prerequisites
* Download and install Go: [https://golang.org/doc/install](https://golang.org/doc/install)
* Make sure that your $GOROOT and Go versions are the same

### Validate go installation 
```shell script
$ go version
go version go1.14 darwin/amd64
```

### Get the package 
```shell script
$ go get https://github.com/hecomp/session-management
```

### Build Application
```shell script
$ go build ./cmd/main.go
```

### Start Application
```shell script
$ go run ./cmd/main.go
```

### Run Test
```shell script
# install the ginkgo CLI
$ go get -u github.com/onsi/ginkgo/ginkgo 
# fetch the matcher library
$ go get -u github.com/onsi/gomega/...    
# run all the ginkgo test in the directory
$ ginkgo -r
# run ginkgo with coverage (verbose)
$ ginkgo -r -v -cover
# run ginkgo report with HTML output
$ go test ./... -coverprofile=coverage.out
$ go tool cover -html=coverage.out -o coverage.html
# run ginkgo report with HTML output using original commands
$ ginkgo -r -v -cover -coverprofile=coverage.out -outputdir=. # Generates coverage report
$ go tool cover -html=coverage.out  #  Renders output into HTML and opens in default browser
```

### Generate Mock using `counterfeiter`

```go
// +build tools

package main

import (
	_ "github.com/maxbrunsfeld/counterfeiter/v6"
)
```

```go
package foo

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . MyInterface
type MyInterface interface {
	DoThings(string, uint64) (int, error)
}
```

```shell script
$ go generate ./...
Writing `FakeMyInterface` to `foofakes/fake_my_interface.go`... Done
```

### Using Test Doubles In Your Tests

Instantiate fakes`:

```go
import "my-repo/path/to/foo/foofakes"

var fake = &foofakes.FakeMyInterface{}
```

Fakes record the arguments they were called with:

```go
fake.DoThings("stuff", 5)

Expect(fake.DoThingsCallCount()).To(Equal(1))

str, num := fake.DoThingsArgsForCall(0)
Expect(str).To(Equal("stuff"))
Expect(num).To(Equal(uint64(5)))
```

You can stub their return values:

```go
fake.DoThingsReturns(3, errors.New("the-error"))

num, err := fake.DoThings("stuff", 5)
Expect(num).To(Equal(3))
Expect(err).To(Equal(errors.New("the-error")))
```

### APIs
| Endpoint | Method | Route     |
| :--------| :------| :---------|
| Create   | GET    | /create   |
| Destroy  | GET    | /destroy  |
| Extend   | POST   | /extend   |
| List     | POST   | /list     |