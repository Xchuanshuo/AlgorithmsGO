package lcw360

func minimumPossibleSum(n int, target int) int64 {
	var total = int64(n * (n + 1) / 2)
	if n+(n-1) < target {
		return total
	}
	var del = make(map[int]bool)
	var r = n
	for i := 1; i < target; i++ {
		if 2*i >= target {
			break
		}
		if del[i] {
			continue
		}
		r++
		del[target-i] = true
	}
	var res = int64(0)
	for i := 1; i <= r; i++ {
		if !del[i] {
			res += int64(i)
		}
	}
	return res
}
