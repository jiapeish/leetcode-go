package _204_count_primes

func countPrimes(n int) int {
	var isComposite = make(map[int]bool, n)
	var count = 0

	for i := 2; i < n; i++ {
		if isComposite[i] == false {
			count++
			for j := 2; j*i < n; j++ {
				isComposite[j*i] = true
			}
		}
	}
	return count
}