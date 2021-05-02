package _974_subarray_sums_divisible_by_k


func subarraysDivByK(A []int, K int) int {
	N := len(A)
	P := make([]int, N + 1)
	for i := 0; i < N; i++ {
		P[i + 1] = P[i] + A[i]
	}

	count := make([]int, K)
	for _, x := range P {
		count[(x % K + K) % K]++
	}

	ans := 0
	for _, c := range count {
		ans += c * (c - 1) / 2
	}
	return ans
}