package test658

func findClosestElements(arr []int, k int, x int) []int {
	var n = len(arr)
	l, r := 0, n-1
	for r-l >= k {
		if abs(arr[l]-x) > abs(arr[r]-x) {
			l++
		} else {
			r--
		}
	}
	return arr[l : r+1]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
