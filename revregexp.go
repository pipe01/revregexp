package revregexp

import (
	"fmt"
	"regexp/syntax"
	"strings"
)

func ReversePattern(pattern *syntax.Regexp) (string, error) {
	str := &strings.Builder{}
	if err := Default.Write(pattern, str); err != nil {
		return "", err
	}

	return str.String(), nil
}

func Reverse(pattern string, flags ...syntax.Flags) (string, error) {
	var flag syntax.Flags
	for _, f := range flags {
		flag |= f
	}

	r, err := syntax.Parse(pattern, flag)
	if err != nil {
		return "", fmt.Errorf("failed to parse pattern: %w", err)
	}

	return ReversePattern(r)
}
