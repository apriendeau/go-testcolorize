package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/apriendeau/go-testcolorize/colorize"
)

var (
	Input = os.Stdin
)

func main() {
	scanner := bufio.NewScanner(Input)

	for scanner.Scan() {
		fmt.Println(colorize.Color(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
