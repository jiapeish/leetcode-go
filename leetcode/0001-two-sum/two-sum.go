package _001_two_sum

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, 1)
	for k, v := range nums {
		if id, ok := seen[target-v]; ok {
			return []int{id, k}
		}
		seen[v] = k
	}
	return nil
}