package _387_first_unique_character_in_a_string

func firstUniqChar(s string) int {
	rec := make(map[rune]int, len(s))

	for _, c := range s {
		rec[c]++
	}

	for idx, c := range s {
		if rec[c] == 1 {
			return idx
		}
	}
	return -1
}

func firstUniqChar2(s string) int {
	rec := make([]int, 26)

	for _, c := range s {
		rec[c-'a']++
	}

	for idx, c := range s {
		if rec[c-'a'] == 1 {
			return idx
		}
	}
	return -1
}