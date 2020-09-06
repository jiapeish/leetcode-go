package _016_3sum_closest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	output, delta := 0, math.MaxInt32

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			switch {
			case sum < target:
				left++
				if delta > target - sum {
					delta = target - sum
					output = sum
				}
			case sum > target:
				right--
				if delta > sum - target {
					delta = sum - target
					output = sum
				}
			default:
				return sum
			}
		}
	}
	return output
}