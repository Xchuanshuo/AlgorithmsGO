package test1598

func minOperations(logs []string) int {
	var cnt = 0
	for _, log := range logs {
		if log == "./" {
			continue
		}
		if log == "../" {
			cnt--
		} else {
			cnt++
		}
		if cnt < 0 {
			cnt = 0
		}
	}
	return cnt
}
