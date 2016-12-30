# go-testcolorize

[![GoDoc](https://godoc.org/github.com/apriendeau/go-testcolorize?status.svg)](https://godoc.org/github.com/apriendeau/go-testcolorize)
[![CircleCI](https://circleci.com/gh/apriendeau/go-testcolorize.svg?style=svg)](https://circleci.com/gh/apriendeau/go-testcolorize)

Colorize your go test output!

## Installation

```shell
go get -u github.com/apriendeau/go-testcolorize/cmd/gtc
```

## Command Line Example:

The command tool: `gtc test {args}` is a wrapper that just executes `go test {args}` under the hood.

## Easter Eggs

1.gtc will dye your strings grey if you println if you start them with "//"

```go
println("// some comment")
```

