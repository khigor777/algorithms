package lib

func BinarySearch(item int, lists []int) int {
	var l = 0
	var h = len(lists) - 1
	for l <= h {
		m := (l + h) / 2
		f := lists[m]
		if item == f {
			return m
		}
		if f > item {
			h = m - 1
		} else {
			l = m + 1
		}
	}
	return 0
}
