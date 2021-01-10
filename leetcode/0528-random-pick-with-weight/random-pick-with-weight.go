package main

import (
	"math/rand"
	"sort"
)

type Solution struct {
	counts []int
	total int
}


func Constructor(w []int) Solution {
	total := 0
	counts := make([]int, len(w))
	for i, v := range w {
		total += v
		counts[i] = total
	}

	return Solution{
		counts: counts,
		total:  total,
	}
}


func (this *Solution) PickIndex() int {
	benchmark := rand.Intn(this.total)
	index := sort.Search(len(this.counts), func(i int) bool {
		return this.counts[i] > benchmark
	})
	return index
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */