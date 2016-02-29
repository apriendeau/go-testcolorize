package testcolorize

import (
	"errors"
	"regexp"
	"strings"

	"github.com/jwaldrip/tint"
)

const (
	passing  = tint.LightGreen
	running  = tint.Cyan
	failing  = tint.LightRed
	skipping = tint.Yellow
	file     = tint.Magenta
	comment  = tint.LightGrey
	// FailRegex Is the how we determine if it is a failing test
	FailRegex = "^(FAIL|\\tError:)"
	// FileRegex determines if the line contains a .go file
	FileRegex = "^.*.go:\\d*:"
	// ExitRegex determines the exit status
	ExitRegex = "^exit status [1-9][0-9]*"
	// FailStr is the start of a failing line
	FailStr = "--- FAIL"
	// PassRegex looks for PASS at the beginning of a line
	PassRegex = "^PASS$"
	// NoTestsRegex determines if the line contains no tests
	NoTestsRegex = "(\\[no test files\\]$|no tests to run)"
	// StartOKRegex determines if the line starts with "ok"
	StartOKRegex = "^ok "
	// CommentRegex is the regex for determining comments
	CommentRegex = "^//*"
	// LocationRegex determins if there is a location line
	LocationRegex = "\\tLocation:"
)

var (
	// ErrFailExitCode is generic for if the go test returned an exit code of 1
	ErrFailExitCode = errors.New("go test failed, exit code 1")
)

// ColorInfo is the data for the loop so we can create a generic item to iterate through
type ColorInfo struct {
	text  string
	color int
	regex string
}

// Color determines if the string is valid part of go test output
// and returns the ANSI color encoded string
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

// Dye is the generic function that does the string replace with the colored value
func Dye(old, value string, color int) string {
	return strings.Replace(old, value, tint.Colorize(value, color), 1)
}

func process(str string, c ColorInfo) (string, error) {
	if c.regex != "" {
		return dyeRegex(str, c.text, c.regex, c.color)
	}
	return Dye(str, c.text, c.color), nil
}

func dyeRegex(old, value, regex string, color int) (string, error) {
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
