package wildcard

import "testing"

func TestIsMatch(t *testing.T) {
	testCases := []struct {
		t    string
		s    string
		p    string
		want bool
	}{
		{
			t:    "both empty",
			s:    "",
			p:    "",
			want: true,
		},
		{
			t:    "double a",
			s:    "aa",
			p:    "a",
			want: false,
		},
		{
			t:    "wildcard a",
			s:    "aa",
			p:    "*",
			want: true,
		},
		{
			t:    "non-match cb",
			s:    "cb",
			p:    "?a",
			want: false,
		},
		{
			t:    "multiple wildcard",
			s:    "xazt",
			p:    "?*az*",
			want: true,
		},
		{
			t:    "wildcard in the beginning",
			s:    "xat",
			p:    "*at",
			want: true,
		},
		{
			t:    "empty pattern non empty string",
			s:    "aa",
			p:    "",
			want: false,
		},
		{
			t:    "multiple wildcard sequentially",
			s:    "aaabbbaabaaaaababaabaaabbabbbbbbbbaabababbabbbaaaaba",
			p:    "a*******b",
			want: false,
		},
		{
			t:    "multiple wildcard empty string",
			s:    "",
			p:    "******",
			want: true,
		},
		{
			t:    "complex multiple wildcard pattern",
			s:    "abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb",
			p:    "**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb",
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.t, func(t *testing.T) {
			actual := isMatch(tc.s, tc.p)
			if actual != tc.want {
				if tc.want {
					t.Errorf("string\n'%s'\nshould match pattern\n'%s'", tc.s, tc.p)
				} else {
					t.Errorf("string\n'%s'\nshould not match pattern\n'%s'", tc.s, tc.p)
				}
			}
		})
	}
}
