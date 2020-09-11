package _205_isomorphic_strings

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ms := make([]int, 256)
	mt := make([]int, 256)
	l := len(s)

	for i := 0; i < l; i++ {
		if ms[s[i]] != mt[t[i]] {
			return false
		}
		ms[s[i]] = i + 1
		mt[t[i]] = i + 1
	}
	return true
}
