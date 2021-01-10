package main

func findMin1(nums []int) int {
	n := len(nums)

	i := 1
	for i < n && nums[i-1] < nums[i] {
		i++
	}
	return nums[i%n]
}

func findMin(nums []int) int {
	low, high := 0, len(nums) - 1
	for low < high {
		if nums[low] < nums[high] {
			return nums[low]
		}
		mid := low + ((high-low)>>1)
		if nums[mid] >= nums[low] {
			low = mid + 1
		} else {
			high = mid
		}

	}
	return nums[low]
}

// 0528.
//func (so *Solution528) PickIndex() int {
//	n := rand.Intn(so.prefixSum[len(so.prefixSum)-1]) + 1
//	low, high := 0, len(so.prefixSum)-1
//	for low < high {
//		mid := low + (high-low)>>1
//		if so.prefixSum[mid] == n {
//			return mid
//		} else if so.prefixSum[mid] < n {
//			low = mid + 1
//		} else {
//			high = mid
//		}
//	}
//	return low
//}