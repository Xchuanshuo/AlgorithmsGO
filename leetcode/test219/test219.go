package test219

func containsNearbyDuplicate(nums []int, k int) bool {
	var posMap = make(map[int]int)
	for i, num := range nums {
		if v, exist := posMap[num];exist && i-v <= k {
			return true
		}
		posMap[num] = i
	}
	return false
}