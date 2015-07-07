package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jwaldrip/tint"
)

const (
	passing = tint.LightGreen
	running = tint.Cyan
	failing = tint.LightRed
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		str = color(str, "--- PASS", passing)
		str = color(str, "PASS", passing)
		str = color(str, "ok", passing)
		str = color(str, "--- FAIL", failing)
		str = color(str, "FAIL", failing)
		str = color(str, "=== RUN", passing)
		fmt.Println(str)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func color(old, value string, color int) string {
	return strings.Replace(old, value, tint.Colorize(value, color), -1)
}
