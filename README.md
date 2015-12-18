# go-testcolorize

Colorize your go test output!

![gtc-screenging](https://raw.githubusercontent.com/apriendeau/go-testcolorize/master/img/output.gif)

## Installation

```shell
go get -u github.com/apriendeau/go-testcolorize/cmd/gtc
```

## Example:

You can pipe `go test -v` into gtc like so:

```shell
$ go test -v ./... | gtc
```

or you can use the gtc wrapper `go test` wrapper:

```shell
$ gtc test ./...
```

## Easter Eggs

1.gtc will dye your strings grey if you println if you start them with "//"

```go
println("// some comment")
```

