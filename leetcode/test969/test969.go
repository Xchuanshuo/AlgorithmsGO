package test969

var res []int

func pancakeSort(arr []int) []int {
	res = make([]int, 0)
	sort(arr, len(arr))
	return res
}

func sort(arr []int, n int) {
	if n == 1 {
		return
	}
	max, mIdx := 0, 0
	for i := 0; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
			mIdx = i
		}
	}
	reverse(arr, 0, mIdx)
	res = append(res, mIdx+1)
	reverse(arr, 0, n-1)
	res = append(res, n)

	sort(arr, n-1)
}

func reverse(arr []int, l, r int) {
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}
