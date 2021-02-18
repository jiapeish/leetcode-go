package main

func countArrangement(n int) int {
	count := 0
	visited := make([]bool, n+1)
	var calculate func(num int, pos int, visited []bool)

	calculate = func(num int, pos int, visited []bool) {
		if pos > num {
			count++
		}

		for i := 1; i <= num; i++ {
			if !visited[i] && (pos % i == 0 || i % pos == 0) {
				visited[i] = true
				calculate(num, pos+1, visited)
				visited[i] = false
			}
		}
	}

	calculate(n, 1, visited)
	return count
}

