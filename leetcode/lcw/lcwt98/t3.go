package lcwt98

import "math"

func minImpossibleOR(nums []int) int {
	var mp = make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = true
	}
	for i := 0; i < 32; i++ {
		var p = int(math.Pow(2.0, float64(i)))
		if !mp[p] {
			return p
		}
	}
	return 1
}