package _349_intersection_of_two_arrays

func intersection(nums1 []int, nums2 []int) []int {
	inter := make([]int, 0, len(nums1))
	m1 := dedup(nums1)
	m2 := dedup(nums2)

	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}

	for k := range m1 {
		if m2[k] {
			inter = append(inter, k)
		}
	}

	return inter
}

func dedup(nums []int) map[int]bool {
	res := make(map[int]bool, len(nums))

	for i := range nums {
		res[nums[i]] = true
	}

	return res
}