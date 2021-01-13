package main


// solution 1
func findLength1(A []int, B []int) int {
	n1, n2 := len(A), len(B)
	var res int

	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}

	// dp[i][j] == k represents A[i-k:i] == B[j-k:j]
	// and A[i-k-1] != B[j-k-1]
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				res = max(res, dp[i][j])
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// solution 2
func findLength(A []int, B []int) int {


}



















