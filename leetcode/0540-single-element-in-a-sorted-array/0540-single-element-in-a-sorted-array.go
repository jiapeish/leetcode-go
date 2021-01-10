package main

func singleNonDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + ((high-low)>>1)
		if mid%2 == 1 {
			mid -= 1
		}

		if nums[mid] == nums[mid+1] {
			low = mid + 2
		} else {
			high = mid
		}
	}
	return nums[low]
}
