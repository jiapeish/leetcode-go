package _209_minimum_size_subarray_sum

func minSubArrayLen(s int, nums []int) int {

	length := len(nums)
	output := length + 1
	sum, left, right := 0, 0, 0

	for right = 0; right < length; right++ {
		sum = sum + nums[right]
		for sum >= s {
			if output > right - left + 1 {
				output = right - left + 1
			}
			sum = sum - nums[left]
			left = left + 1
		}
	}

	if output == length + 1 {
		return 0
	}
	return output

}