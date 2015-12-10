package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/apriendeau/go-testcolorize"
)

var (
	Input = os.Stdin
)

func main() {
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
