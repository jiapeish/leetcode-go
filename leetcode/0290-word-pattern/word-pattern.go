package _290_word_pattern

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	ps := strings.Split(pattern, "")
	ss := strings.Split(s, " ")
	if len(ps) != len(ss) {
		return false
	}
	return isMatch(ps, ss) && isMatch(ss, ps)

}

func isMatch(p []string, s []string) bool {
	if len(p) != len(s) {
		return false
	}

	l := len(p)
	m := make(map[string]string, l)

	for i := 0; i < l; i++ {
		if v, ok := m[p[i]]; ok {
			if v != s[i] {
				return false
			}
		} else {
			m[p[i]] = s[i]
		}
	}
	return true
}