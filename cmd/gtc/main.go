package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/apriendeau/go-testcolorize"
	"github.com/spf13/pflag"
)

var (
	Input        = os.Stdin
	Help    bool = false
	Verbose bool = false
)

const Usage = `gtc - Pipe go test or use the wrapper for some colorful tests

Usage:
	gtc test <args>
	go test -v . |& gtc`

func main() {
	fset := pflag.NewFlagSet("gtc", pflag.ExitOnError)
	fset.BoolVarP(&Help, "help", "h", false, "show help message")
	fset.BoolVarP(&Verbose, "verbose", "v", true, "verbose output")
	fset.Parse(os.Args)

	show, err := fset.GetBool("help")
	if err != nil {
		log.Fatal(err)
	}
	if show {
		fmt.Println(Usage)
		os.Exit(84)
	}

	args := os.Args
	if len(args) >= 2 {
		switch args[1] {
		case "test":
			Test(args[1:])
			return
		}
	}
	code, err := Scan(Input)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(code)
}

func Scan(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)
	exit := 0
	for scanner.Scan() {
		txt := scanner.Text()
		str, err := testcolorize.Color(txt)
		if err != nil {
			if err == testcolorize.ErrFailExitCode {
				exit = 1
			} else {
				return 1, err
			}
		}
		fmt.Println(str)
	}
	if err := scanner.Err(); err != nil {
		return 1, err
	}
	return exit, nil

}
