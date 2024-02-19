package lcw372

func findMinimumOperations(s1 string, s2 string, s3 string) int {
	var l = min(len(s3), min(len(s1), len(s2)))
	var res = -1
	for i := 1; i < l; i++ {
		var p = s1[0:i]
		if p == s2[0:i] && p == s3[0:i] {
			res = len(s1) - i - 1 + len(s2) - i - 1 + len(s3) - i - 1
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
