package _080_remove_duplicates_from_sorted_array_ii

func removeDuplicates(nums []int) int {
	length := len(nums)

	i := 2

	for j := 2; j < length; j++ {
		if nums[j] != nums[i-2] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
