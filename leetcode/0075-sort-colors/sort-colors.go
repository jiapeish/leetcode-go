package _075_sort_colors

func sortColors(nums []int)  {
	i, j, k := 0, 0, len(nums)-1

	// nums[:i] stores 0
	// nums[i:j] stores 1
	// nums[j:k] stores 2

	for j <= k {
		switch nums[j] {
		case 0:
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		case 1:
			j++
		case 2:
			nums[j], nums[k] = nums[k], nums[j]
			k--
		}
	}
}