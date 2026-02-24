package wildcard

import "strings"

func isMatch(s string, p string) bool {
	// s is empty
	if s == "" {
		p = strings.Trim(p, "*")
		return p == ""
	}

	// s is non-empty
	if p == "" {
		return false
	}

	// s and p are both non-empty
	switch p[0] {
	case '?':
		return isMatch(s[1:], p[1:])
	case '*':
		for len(p) > 1 && p[1] == '*' {
			p = p[1:]
		}
		return isMatch(s, p[1:]) || isMatch(s[1:], p)
	default:
		if s[0] == p[0] {
			return isMatch(s[1:], p[1:])
		}
		return false
	}
}
