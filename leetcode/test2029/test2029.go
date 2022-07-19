package test2029

func stoneGameIX(stones []int) bool {
	cnt := make([]int, 3)
	for _, v := range stones { cnt[v%3]++ }
	if cnt[0]%2 == 0 {
		return cnt[1] > 0 && cnt[2] > 0
	}
	return abs(cnt[1], cnt[2]) > 2
}

func abs(a, b int) int {
	if a > b { return a - b}
	return b - a
}