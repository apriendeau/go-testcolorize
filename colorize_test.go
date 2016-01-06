package testcolorize_test

import (
	"strings"
	"testing"

	"github.com/apriendeau/go-testcolorize"
	"github.com/jwaldrip/tint"
)

func TestColorize(t *testing.T) {
	var strs = []struct {
		color   int
		name    string
		text    string
		errored bool
	}{
		{text: "--- SKIP", color: tint.Yellow, name: "yellow", errored: false},
		{text: "=== RUN", color: tint.Cyan, name: "cyan", errored: false},
		{text: "ok ", color: tint.LightGreen, name: "light green", errored: false},
		{text: "PASS", color: tint.LightGreen, name: "light green", errored: false},
		{text: "\tError:", color: tint.LightRed, name: "light red", errored: true},
		{text: "FAIL", color: tint.LightRed, name: "light red", errored: true},
		{text: "--- PASS", color: tint.LightGreen, name: "light green", errored: false},
		{text: "--- FAIL", color: tint.LightRed, name: "light red", errored: true},
		{text: "testing.go:12345:", color: tint.Magenta, name: "magenta", errored: false},
		{text: "boom.go:15:", color: tint.Magenta, name: "magenta", errored: false},
		{text: "\tLocation:\tserver.go:1234", color: tint.Magenta, name: "magenta", errored: false},
		{text: "testing.go:1234567890:", color: tint.Magenta, name: "magenta", errored: false},
		{text: "testing.go:1:", color: tint.Magenta, name: "magenta", errored: false},
		{text: "exit status 1", color: tint.LightRed, name: "light red", errored: true},
		{text: "exit status 484839393", color: tint.LightRed, name: "light red", errored: true},
		{text: "exit status 939393", color: tint.LightRed, name: "light red", errored: true},
		{text: "exit status 9123812983", color: tint.LightRed, name: "light red", errored: true},
		{text: "testing: warning: no tests to run", color: tint.Yellow, name: "yellow", errored: false},
		{text: "[no test files]", color: tint.Yellow, name: "yellow", errored: false},
		{text: "// some comment", color: tint.LightGrey, name: "light grey", errored: false},
		{text: "// another comment to be safe", color: tint.LightGrey, name: "light grey", errored: false},
	}
	for _, str := range strs {
		msg, err := testcolorize.Color(str.text)
		if str.errored {
			if err == nil {
				t.Errorf("%s should have had an error", str.text)
			}
			if err != testcolorize.ErrFailExitCode {
				t.Errorf("%s should have had the error: %s", str.text, testcolorize.ErrFailExitCode.Error())
			}
		}
		sample := testcolorize.Dye(str.text, str.text, str.color)
		if msg != sample {
			t.Errorf("%s is not colored %s", msg, str.name)
		}
	}
}

func color(old, value string, color int) string {
	return strings.Replace(old, value, tint.Colorize(value, color), 1)
}
