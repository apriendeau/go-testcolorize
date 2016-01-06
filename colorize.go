package testcolorize

import (
	"errors"
	"regexp"
	"strings"

	"github.com/jwaldrip/tint"
)

const (
	passing       = tint.LightGreen
	running       = tint.Cyan
	failing       = tint.LightRed
	skipping      = tint.Yellow
	file          = tint.Magenta
	comment       = tint.LightGrey
	FailRegex     = "^(FAIL|\\tError:)"
	FileRegex     = "^.*.go:\\d*:"
	ExitRegex     = "^exit status [1-9][0-9]*"
	FailStr       = "--- FAIL"
	PassRegex     = "^PASS$"
	NoTestsRegex  = "(\\[no test files\\]$|no tests to run)"
	StartOKRegex  = "^ok "
	CommentRegex  = "^//*"
	LocationRegex = "\\tLocation:"
)

var (
	ErrFailExitCode = errors.New("go test failed, exit code 1")
)

type ColorInfo struct {
	text  string
	color int
	regex string
}

func Color(str string) (string, error) {
	colors := []ColorInfo{
		{"--- PASS", passing, ""},
		{"PASS", passing, PassRegex},
		{"ok ", passing, StartOKRegex},
		{FailStr, failing, "--- FAIL"},
		{"FAIL", failing, FailRegex},
		{"=== RUN", running, ""},
		{"--- SKIP", skipping, ""},
		{"", file, FileRegex},
		{"", file, LocationRegex},
		{"", failing, ExitRegex},
		{"", skipping, NoTestsRegex},
		{"", comment, CommentRegex},
	}
	var err error
	var exit error
	for _, c := range colors {
		str, err = process(str, c)
		if err != nil {
			exit = err
		}
	}
	return str, exit
}

func process(str string, c ColorInfo) (string, error) {
	if c.regex != "" {
		return DyeRegex(str, c.text, c.regex, c.color)
	}
	return Dye(str, c.text, c.color), nil
}

func DyeRegex(old, value, regex string, color int) (string, error) {
	re := regexp.MustCompile(regex)
	if re.MatchString(old) {
		switch regex {
		case FailRegex, ExitRegex, FailStr:
			value = re.FindString(old)
			return Dye(old, value, color), ErrFailExitCode
		case FileRegex:
			value = re.FindString(old)
			return Dye(old, value, color), nil

		case NoTestsRegex, StartOKRegex, CommentRegex, LocationRegex:
			return Dye(old, old, color), nil
		default:
			return Dye(old, value, color), nil
		}
	}
	return old, nil
}
func Dye(old, value string, color int) string {
	return strings.Replace(old, value, tint.Colorize(value, color), 1)
}
