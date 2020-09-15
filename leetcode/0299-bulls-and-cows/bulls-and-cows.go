package _299_bulls_and_cows

import "strconv"

func getHint(secret string, guess string) string {
	bulls, cows := 0, 0
	numbers := make([]int, 10)
	length := len(secret)

	for i := 0; i < length; i++ {
		s := secret[i]
		g := guess[i]
		if s == g {
			bulls++
		} else {
			if numbers[s-'0'] < 0 {
				cows++
			}

			if numbers[g-'0'] > 0 {
				cows++
			}

			numbers[s-'0']++
			numbers[g-'0']--
		}
	}
	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}
