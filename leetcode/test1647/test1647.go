package test1647

func minDeletions(s string) int {
	var cnts = make([]int, 26)
	for _, v := range s {
		cnts[v-'a']++
	}
	var set = make(map[int]bool)
	var res = 0
	for i := 0; i < len(cnts); i++ {
		var freq = cnts[i]
		for freq != 0 && set[freq] {
			res++
			freq--
		}
		set[freq] = true
	}
	return res
}
