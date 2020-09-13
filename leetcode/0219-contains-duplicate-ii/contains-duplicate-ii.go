package _219_contains_duplicate_ii

func containsNearbyDuplicate(nums []int, k int) bool {
	if k <= 0 {
		return false
	}

	position := make(map[int]int, len(nums))
	for i, num := range nums {
		if position[num] != 0 && i - (position[num]-1) <= k {
			return true
		}
		position[num] = i + 1
	}
	return false
}