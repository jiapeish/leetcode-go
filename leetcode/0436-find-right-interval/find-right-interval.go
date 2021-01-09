package main

import (
	"sort"

	"leetcode-go/structures"
)

// solution 1
type Interval = structures.Interval

func findRightInterval1(intervals [][]int) []int {
	if len(intervals) == 0 {
		return []int{}
	}

	intervalsList, res, intervalsMap := []Interval{}, []int{}, map[Interval]int{}
	for i, v := range intervals {
		intervalsList = append(intervalsList, Interval{Start: v[0], End: v[1]})
		intervalsMap[Interval{Start: v[0], End: v[1]}] = i
	}

	structures.Quicksort(intervalsList, 0, len(intervalsList)-1)
	for _, v := range intervals {
		idx := searchFirstGreaterInterval(intervalsList, v[1])
		if idx >= 0 {
			idx = intervalsMap[intervalsList[idx]]
		}
		res = append(res, idx)
	}
	return res
}

func searchFirstGreaterInterval(nums []Interval, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid].Start < target {
			low = mid + 1
		} else {
			if mid == 0 || nums[mid-1].Start < target {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

// solution 2
type interval struct {
	interval []int
	index    int
}

type intervalList []interval

func (in intervalList) Len() int {
	return len(in)
}

func (in intervalList) Less(i, j int) bool {
	return (in[i].interval[0] < in[j].interval[0]) ||
		(in[i].interval[0] == in[j].interval[0] && in[i].interval[1] < in[j].interval[1])
}

func (in intervalList) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	if n == 0 {
		return []int{}
	}

	intervalList := make(intervalList, n)
	for i, v := range intervals {
		intervalList[i] = interval{
			interval: v,
			index:    i,
		}
	}
	sort.Sort(intervalList)

	res := make([]int, n)
	for i := 0; i < n; i++ {
		index := sort.Search(n, func(j int) bool {
			return intervalList[j].interval[0] >= intervalList[i].interval[1]
		})

		if index == n {
			res[intervalList[i].index] = -1
		} else {
			res[intervalList[i].index] = intervalList[index].index
		}
	}

	return res
}
