package _389_find_the_difference

func findTheDifference(s string, t string) byte {

	rec := make([]int, 26)

	for i := range s {
		rec[s[i] - 'a']++
		rec[t[i] - 'a']--
	}
	rec[t[len(t)-1] - 'a']--

	var idx, count int

	for idx, count = range rec {
		if count == -1 {
			break
		}
	}
	return byte('a' + idx)
}