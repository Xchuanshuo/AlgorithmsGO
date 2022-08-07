package test944

func minDeletionSize(strs []string) int {
	var n = len(strs[0])
	var res = 0
	for i := 0; i < n; i++ {
		var last byte = 'a'
		for _, s := range strs {
			if s[i] < last {
				res++
				break
			}
			last = s[i]
		}
	}
	return res
}
