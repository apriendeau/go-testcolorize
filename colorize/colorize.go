package colorize

import (
	"regexp"
	"strings"

	"github.com/jwaldrip/tint"
)

const (
	passing  = tint.LightGreen
	running  = tint.Cyan
	failing  = tint.LightRed
	skipping = tint.Yellow
)

type ColorInfo struct {
	text  string
	color int
	regex string
}

func Color(str string) string {
	colors := []ColorInfo{
		{"--- PASS", passing, ""},
		{"PASS", passing, "^PASS$"},
		{"ok", passing, "^ok"},
		{"--- FAIL", failing, ""},
		{"FAIL", failing, "^FAIL$"},
		{"=== RUN", running, ""},
		{"--- SKIP", skipping, ""},
	}
	for _, c := range colors {
		str = process(str, c)
	}
	return str
}

func process(str string, c ColorInfo) string {
	if c.regex != "" {
		return DyeRegex(str, c.text, c.regex, c.color)
	}
	return Dye(str, c.text, c.color)
}

func DyeRegex(old, value, regex string, color int) string {
	re := regexp.MustCompile(regex)
	if re.MatchString(old) {
		return Dye(old, value, color)
	}
	return old
}
func Dye(old, value string, color int) string {
	return strings.Replace(old, value, tint.Colorize(value, color), 1)
}
