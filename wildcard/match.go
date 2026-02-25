package wildcard

func isMatch(s string, p string) bool {
	// s is empty
	if s == "" {
		for _, ch := range p {
			if ch != '*' {
				return false
			}
		}
		return true
	}

	// s is non-empty
	if p == "" {
		return false
	}

	// s and p are both non-empty
	si, pi := 0, 0
	sj, pj := len(s)-1, len(p)-1
	for si <= sj && pi <= pj {
		switch {
		case p[pi] == '?' || s[si] == p[pi]:
			si++
			pi++
		case p[pi] == '*':
			for pi+1 <= pj && p[pi+1] == '*' {
				pi++
			}
			if pi+1 <= pj && (p[pi+1] == '?' || s[si] == p[pi+1]) {
				// * as empty substring
				if isMatch(s[si:], p[pi+1:pj+1]) {
					return true
				}
			}
			si++
		default:
			return false
		}
	}
	return isMatch(s[si:sj+1], p[pi:pj+1])
}
