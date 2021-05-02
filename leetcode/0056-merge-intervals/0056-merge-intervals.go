package main

import (
	"fmt"
	"math/rand"
)

func main() {
	data := [][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	}
	ans := merge(data)
	fmt.Printf("ans %v\n", ans)
}

func merge(intervals [][]int) [][]int {
	quicksort(intervals, 0, len(intervals)-1)
	fmt.Printf("sort res %v\n", intervals)

	res := make([][]int, 0, 1)
	res = append(res, intervals[0])
	curr := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= res[curr][1] {
			res[curr][1] = max(res[curr][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
			curr++
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

func quicksort(A [][]int, p, r int) {
	if p > r {
		return
	}
	q := randPartition(A, p, r)
	quicksort(A, p, q-1)
	quicksort(A, q+1, r)
}

func randPartition(A [][]int, p, r int) int {
	i := p + rand.Intn(r-p+1)
	A[i], A[r] = A[r], A[i]
	return partition(A, p, r)
}

func partition(A [][]int, p, r int) int {
	x := A[r]
	i := p - 1
	for j := p; j < r; j++ {
		if A[j][0] <= x[0] {
			i = i + 1
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}

