package test1652

func decrypt(code []int, k int) []int {
	var n = len(code)
	var res = make([]int, n)
	if k == 0 {
		return res
	}
	var sums = make([]int, n+1)
	for i, c := range code {
		sums[i+1] = sums[i] + c
	}
	for idx, _ := range code {
		var f = 0
		if k > 0 {
			if idx+1+k > n {
				f = sums[n] - sums[idx+1] + sums[idx+k+1-n]
			} else {
				f = sums[idx+1+k] - sums[idx+1]
			}
		} else {
			if idx+k < 0 {
				f = sums[idx] + sums[n] - sums[n+(idx+k)]
			} else {
				f = sums[idx] - sums[idx+k]
			}
		}
		res[idx] = f
	}
	return res
}
