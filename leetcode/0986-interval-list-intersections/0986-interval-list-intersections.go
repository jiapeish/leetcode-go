package _986_interval_list_intersections

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var ans [][]int
	i := 0
	j := 0

	for i < len(firstList) && j < len(secondList) {
		// lo: the start point of the intersection
		// hi: the end point of the intersection
		lo := max(firstList[i][0], secondList[j][0])
		hi := min(firstList[i][1], secondList[j][1])
		if lo <= hi {
			ans = append(ans, []int{lo, hi})
		}

		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
