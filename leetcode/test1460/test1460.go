package test1460

func canBeEqual(target []int, arr []int) bool {
	var pos = make(map[int]int)
	for _, t := range target {
		pos[t]++
	}
	for _, a := range arr {
		pos[a]--
		if pos[a] < 0 {
			return false
		}
	}
	return true
}
