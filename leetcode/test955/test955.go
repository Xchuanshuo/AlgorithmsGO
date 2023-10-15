package test955

func minDeletionSize(strs []string) int {
	var m = len(strs[0])
	var res = 0
	var larger = make([]map[int]bool, len(strs))
	for i := 0; i < len(larger); i++ {
		larger[i] = make(map[int]bool)
	}
	for i := 0; i < m; i++ {
		var valid = true
		for j := 1; j < len(strs); j++ {
			if strs[j-1][i] <= strs[j][i] {
				continue
			}
			if !larger[j][j-1] {
				valid = false
				break
			}
		}
		if !valid {
			res++
		} else {
			for j := 1; j < len(strs); j++ {
				if strs[j][i] > strs[j-1][i] {
					larger[j][j-1] = true
				}
			}
		}
	}
	return res
}
