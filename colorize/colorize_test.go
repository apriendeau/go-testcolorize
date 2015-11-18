package colorize_test

import (
	"strings"
	"testing"

	"github.com/apriendeau/go-testcolorize/colorize"
	"github.com/jwaldrip/tint"
)

func TestColorize(t *testing.T) {
	var strs = []struct {
		color int
		name  string
		text  string
	}{
		{text: "--- SKIP", color: tint.Yellow, name: "yellow"},
		{text: "=== RUN", color: tint.Cyan, name: "cyan"},
		{text: "ok", color: tint.LightGreen, name: "light green"},
		{text: "PASS", color: tint.LightGreen, name: "light green"},
		{text: "FAIL", color: tint.LightRed, name: "light red"},
		{text: "--- PASS", color: tint.LightGreen, name: "light green"},
		{text: "--- FAIL", color: tint.LightRed, name: "light red"},
	}
	for _, str := range strs {
		msg, err := colorize.Color(str.text)
		if err != nil && err != colorize.ErrFailExitCode {
			t.Errorf(err.Error())
		}
		sample := colorize.Dye(str.text, str.text, str.color)

		if msg != sample {
			t.Errorf("%s is not colored %s", msg, str.name)
		}
	}
}

func color(old, value string, color int) string {
	return strings.Replace(old, value, tint.Colorize(value, color), 1)
}
