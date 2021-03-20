# Learning Go [![Build Status][actions-img]][actions]

Learning Go using book **"Introducing Go"**, written by Caleb Doxsey, and other
online materials.

Module | Description
:--- | :---
`core` | The core features from the Go SDK.
`elasticsearch` | Elasticsearch Go client.
`http` | Using HTTP client in Go.
`json` | JSON marshalling and unmarshalling.
`temporal` | Writing invincible workflows with <https://temporal.io/>.

## Installation

Install Go on macOS:

```sh
> brew install go
```

Check Go version:

```sh
> go version
```

## Run

Run the main program:

```sh
> go run main.go
Hello, world!
```

## Testing

Run tests for the current directory:

```sh
> go test
```

Run tests for all directories (current one and the sub-directories):

```sh
> go test ./...
```

Run test with the verbose mode (`-v`):

```sh
> go test -v ./...
```

## Code Style

The source code is formatted automatically using
[gofmt](https://golang.org/cmd/gofmt/) tool:

```sh
> go fmt
```

## References

- Dan Buch, "travis-ci-examples/go-example", _GitHub_, 2019.
  <https://github.com/travis-ci-examples/go-example>
- Go, "Package testing", _Golang_, 2019.
  <https://golang.org/pkg/testing/>
- Andrew Gerrand, "go fmt your code", _Golang_, 2013.
  <https://blog.golang.org/go-fmt-your-code>
- Mark McGranaghan, "Go by Example", _gobyexample_, 2019.
  <https://gobyexample.com/>

[actions]: https://github.com/mincong-h/learning-go/actions
[actions-img]: https://github.com/mincong-h/learning-go/workflows/Actions/badge.svg
