package _072_edit_distance

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// target: dp[i][j] means from[:i] -> to[:j]

			// case 1: from[:i-1] -> to[:j]
			//dp[i][j] = 1 + dp[i-1][j]

			// case 2: from[:i] -> to[:j-1]
			//dp[i][j] = 1 + dp[i][j-1]

			dp[i][j] = min(1+dp[i-1][j], 1+dp[i][j-1])

			// case 3: from[:i-1] -> to[:j-1]
			//	case 3.1: from[i-1] == to[j-1]
			//	case 3.2: from[i-1] != to[j-1]
			replace := 1
			if word1[i-1] == word2[j-1] {
				replace = 0
			}
			//dp[i][j] = dp[i-1][j-1] + replace

			dp[i][j] = min(dp[i][j], dp[i-1][j-1]+replace)
		}
	}
	return dp[m][n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}