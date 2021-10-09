# Golang Fundamental Testing

`Example` write basic fundamental testing in Golang

## Installation
```sh
$ go mod download || go get .
```

## Implementation Case

- **01_testing** write testing
- **02_testing** write testing using testify
- **03_testing** write API testing using testify
- **04_testing** write API testing using gin framework, testify and go supertest
- **05_testing** write testing using ginkgo bdd testing framework

## Rule Using Ginkgo

You must set up like this before you run testing using ginkgo, because this is bootstrapping for your test in single file not each every file, if you create new test you must need code like this again with difference name method.

```go
func TestCalculate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculate Suite")
}
```

## Rule Ginkgo Hooks

- **BeforeEach**: run code every test is execute
- **AfterEach**:  run code after every test is execute done
- **BeforeSuite**: run code before all test is execute
- **AfterSuite**:  run code after all test is execute done

More about functionality in ginkgo check [here](https://github.com/onsi/ginkgo/blob/ver2/docs/index.md)

## Command

- ### Testing

```sh
$ make test || go test -v ./...
```

- ### Coverage Testing

```sh
$ make coverage || go test -cover -v ./... || go test -coverprofile=coverage.out ./... || go tool cover -func=coverage.out
```

- ### Ginkgo Testing

```sh
$ make ginkgo || go test -v || ginkgo -v
```

## Testing And Coverage Testing Result

<img src="./images/testing.png"/>
<img src="./images/coverage.png"/>
