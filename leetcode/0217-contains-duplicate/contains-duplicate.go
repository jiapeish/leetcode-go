package _217_contains_duplicate

func containsDuplicate(nums []int) bool {
	var dict = make(map[int]bool, len(nums))

	for _, num := range nums {
		if dict[num] == true {
			return true
		}
		dict[num] = true
	}
	return false
}