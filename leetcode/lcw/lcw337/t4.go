package lcw337

func findSmallestInteger(nums []int, value int) int {
	var mp = make(map[int]int)
	for _, v := range nums {
		if v < 0 {
			var t = (-v+value-1)/value * value
			mp[(v+t)%value]++
		} else {
			mp[v%value]++
		}
	}
	for i := 0;i <= int(1e5);i++ {
		if mp[i] == 0 && mp[i%value] == 0 {
			return i
		}
		if mp[i] != 0 {
			mp[i]--
		} else if mp[i%value] != 0 {
			mp[i%value]--
		}
	}
	return -1
}
