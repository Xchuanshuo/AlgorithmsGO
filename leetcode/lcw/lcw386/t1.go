package lcw386

func isPossibleToSplit(nums []int) bool {
	var cnt = make(map[int]int)
	for _, v := range nums {
		cnt[v]++
		if cnt[v] > 2 {
			return false
		}
	}
	return true
}
