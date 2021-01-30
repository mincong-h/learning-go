# Learning Go [![Build Status][actions-img]][actions]

Learning Go using book **"Introducing Go"**, written by Caleb Doxsey, and other
online materials.

<p align="center">
  <a href="https://amzn.to/31Mz7E1">
    <img src="img/introducing-go.jpg" width="200" alt="Caleb Doxsey, Introducing Go" />
  </a>
</p>

## Installation

Install Go on macOS:

    $ brew install go

Check Go version:

    $ go version

## Run

Run the main program:

```
$ go run main.go
Hello, world!
```

## Testing

Run test:

    $ go test

Run test with verbose mode:

    $ go test -v

## Code Style

The source code is formatted automatically using
[gofmt](https://golang.org/cmd/gofmt/) tool:

    $ go fmt

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
