package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/apriendeau/go-testcolorize/colorize"
)

var (
	Input = os.Stdin
)

func main() {
	scanner := bufio.NewScanner(Input)
	exitCode := 0
	for scanner.Scan() {
		str, err := colorize.Color(scanner.Text())
		if err != nil {
			if err == colorize.ErrFailExitCode {
				exitCode = 1
			} else {
				log.Fatal(err)
			}
		}
		fmt.Println(str)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	os.Exit(exitCode)
}
