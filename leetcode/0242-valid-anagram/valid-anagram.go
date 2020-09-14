package _242_valid_anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sr := []rune(s)
	tr := []rune(t)
	occur := make(map[rune]int, len(sr))

	for idx := range sr {
		occur[sr[idx]]++
		occur[tr[idx]]--
	}

	for _, v := range occur {
		if v != 0 {
			return false
		}
	}
	return true
}