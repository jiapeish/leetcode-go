package _852_peak_index_in_a_mountain_array

func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)
	for {
		m := l + (r-l)/2
		switch {
		case arr[m] < arr[m+1]:
			l = m+1
		case arr[m] < arr[m-1]:
			r = m
		default:
			return m
		}
	}
}