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
	i, j := 0, 0
	for i < len(s) && j < len(p) {
		switch {
		case p[j] == '?' || s[i] == p[j]:
			i++
			j++
		case p[j] == '*':
			for j+1 < len(p) && p[j+1] == '*' {
				j++
			}
			if j+1 < len(p) && (p[j+1] == '?' || s[i] == p[j+1]) {
				if isMatch(s[i:], p[j+1:]) {
					return true
				}
			}
			i++
		default:
			return false
		}
	}
	return isMatch(s[i:], p[j:])
}
