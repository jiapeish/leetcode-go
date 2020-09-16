package _350_intersection_of_two_arrays_ii

func intersect(nums1 []int, nums2 []int) []int {
	var inter []int

	m1 := dedup(nums1)
	m2 := dedup(nums2)

	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}

	for num, count := range m1 {
		m1[num] = min(count, m2[num])
	}

	for num, count := range m1 {
		for i := 0; i < count; i++ {
			inter = append(inter, num)
		}
	}

	return inter
}

func dedup(nums []int) map[int]int {
	res := make(map[int]int, len(nums))

	for i := range nums {
		res[nums[i]]++
	}

	return res
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}