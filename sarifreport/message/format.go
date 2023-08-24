package message

import (
	"regexp"
	"strconv"
)

func format(message string, arguments []string) *string {
	s := regexp.MustCompile("(\\{\\d+}|\\{\\{|}})").ReplaceAllStringFunc(message, makeReplace(arguments))
	return &s
}

func makeReplace(arguments []string) func(s string) string {
	return func(s string) string {
		i := indexOf(s)
		if i >= 0 && len(arguments) > i {
			return arguments[i]
		}

		if s == "{{" {
			return "{"
		}
		if s == "}}" {
			return "}"
		}

		return s
	}
}

func indexOf(s string) int {
	matches := regexp.MustCompile("^\\{(\\d+)}$").FindStringSubmatch(s)
	if len(matches) <= 1 {
		return -1
	}

	result, _ := strconv.Atoi(matches[1])
	return result
}
