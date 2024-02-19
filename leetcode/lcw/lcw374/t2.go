package lcw374

import "math"

func minimumAddedCoins(coins []int, target int) int {
	var mp = make(map[int]bool)
	for i := 0; i < 30; i++ {
		if target&(1<<i) != 0 {
			var t = int(math.Pow(2.0, float64(i)))
			mp[t] = true
		}
	}
	for _, c := range coins {
		if mp[c] {
			delete(mp, c)
		}
	}
	return len(mp)
}
