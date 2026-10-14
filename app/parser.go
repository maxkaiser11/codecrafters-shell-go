package main

import "strings"

// parseArgs splits a command line into tokens, handling single quotes,
// double quotes and backslash escapes.
func parseArgs(line string) []string {
	var args []string
	var cur strings.Builder
	var quote rune
	hasToken := false
	escaped := false

	for _, r := range line {
		switch {
		case escaped:
			if quote == '"' && r != '"' && r != '\\' {
				cur.WriteRune('\\')
			}
			cur.WriteRune(r)
			hasToken = true
			escaped = false
		case r == '\\' && (quote == 0 || quote == '"'):
			escaped = true
		case quote == 0 && (r == '\'' || r == '"'):
			quote = r // open quote
			hasToken = true
		case r == quote:
			quote = 0 // close quote
		case r == ' ' && quote == 0:
			if hasToken {
				args = append(args, cur.String())
				cur.Reset()
				hasToken = false
			}
		default:
			cur.WriteRune(r)
			hasToken = true
		}
	}

	if hasToken {
		args = append(args, cur.String())
	}
	return args
}
