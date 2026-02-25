package wildcard

type matchKey struct {
	s string
	p string
}

var cache = make(map[matchKey]bool)

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
		case p[pi] != '?' && p[pi] != '*' && s[si] != p[pi]:
			return false
		case p[pj] != '?' && p[pj] != '*' && s[sj] != p[pj]:
			return false
		case p[pi] == '?' || s[si] == p[pi]:
			si++
			pi++
		case p[pj] == '?' || s[sj] == p[pj]:
			sj--
			pj--
		case pi+1 <= pj && p[pi] == '*' && p[pi+1] == '*':
			pi++
		case pi <= pj-1 && p[pj-1] == '*' && p[pj] == '*':
			pj--
		case p[pi] == '*' && (pi+1 > pj || (p[pi+1] != '?' && p[pi+1] != s[si])):
			//Can consume current begin character only as wildcard
			si++
		case p[pj] == '*' && (pi > pj-1 || (p[pj-1] != '?' && p[pj-1] != s[sj])):
			//Can consume current end character only as wildcard
			sj--
		case p[pi] == '*':
			// try to represent * in begin as empty substring
			k := matchKey{
				s: s[si : sj+1],
				p: p[pi+1 : pj+1],
			}
			if match, ok := cache[k]; ok {
				return match
			}
			match := isMatch(k.s, k.p)
			cache[k] = match
			if match {
				return true
			}
			si++
		default:
			return false
		}
	}
	return isMatch(s[si:sj+1], p[pi:pj+1])
}
