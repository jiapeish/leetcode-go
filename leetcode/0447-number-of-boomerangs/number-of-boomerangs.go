package _447_number_of_boomerangs

func numberOfBoomerangs(points [][]int) int {
	var pairs int

	total := len(points)
	if total < 3 {
		return 0
	}


	for i := 0; i < total; i++ {
		var distsMap = make(map[int]int, total)
		for j := 0; j < total; j++ {
			if j == i {
				continue
			}
			d := dist(points[i], points[j])
			distsMap[d]++
		}

		for _, count := range distsMap {
			pairs += count * (count-1)
		}
	}
	return pairs
}

func dist(a, b []int) int {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	return dx * dx + dy * dy
}