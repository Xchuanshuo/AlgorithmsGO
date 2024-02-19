package lcw372

func minimumSteps(s string) int64 {
	var n = len(s)
	var total = int64(0)
	var r = n
	for i := n - 1; i >= 0; i-- {
		if s[i] == '1' {
			total += int64(r - i - 1)
			r--
		}
	}
	return total
}
