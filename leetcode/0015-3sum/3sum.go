package _015_3sum

import "sort"

func threeSum(nums []int) [][]int {
	var output [][]int
	sort.Ints(nums)
	length := len(nums)

	for i := 0; i < length-2; i++ {
		if (i == 0) || (i > 0 && nums[i] != nums[i-1]) {
			lo := i + 1
			hi := length - 1
			sum := 0 - nums[i]
			for lo < hi {
				if nums[lo] + nums[hi] == sum {
					output = append(output, []int{nums[i], nums[lo], nums[hi]})
					lo, hi = dedup(lo, hi, nums)
				} else if nums[lo] + nums[hi] < sum {
					lo++
				} else {
					hi--
				}
			}
		}
	}

	return output
}

func dedup(lo, hi int, nums []int) (int, int) {
	for lo < hi {
		if nums[lo] == nums[lo+1] {
			lo++
		} else if nums[hi] == nums[hi-1] {
			hi--
		} else {
			lo++
			hi--
			return lo, hi
		}
	}
	return lo, hi
}