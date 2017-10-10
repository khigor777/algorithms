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


func SameBinarySearch(si int, in []int) (index int) {

	low :=0
	high := len(in) - 1

	for low <= high {
		mid := (low + high) / 2
		cur := in[mid]
		if cur == si{
			return mid
		}
		if cur > si {
			high = mid - 1
		}else {
			low = mid + 1
		}

	}
	return 0
}