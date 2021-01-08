package structures


type Interval struct {
	Start int
	End int
}

func Interval2IntS(i Interval) []int {
	return []int{i.Start, i.End}
}

func IntS2Interval(s []int) Interval {
	return Interval{
		Start: s[0],
		End:   s[1],
	}
}

func IntervalS2IntSs(is []Interval) [][]int {
	intSs := make([][]int, 0, len(is))
	for _, i := range is {
		intSs = append(intSs, []int{i.Start, i.End})
	}
	return intSs
}

func IntSs2IntervalS(intSs [][]int) []Interval {
	intervalS := make([]Interval, 0, len(intSs))
	for _, intS := range intSs {
		intervalS = append(intervalS, Interval{
			Start: intS[0],
			End:   intS[1],
		})
	}
	return intervalS
}

func Quicksort(a []Interval, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	Quicksort(a, lo, p-1)
	Quicksort(a, p+1, hi)
}

func partition(a []Interval, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j].Start < pivot.Start || (a[j].Start == pivot.Start && a[j].End < pivot.End) {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i+1
}
















