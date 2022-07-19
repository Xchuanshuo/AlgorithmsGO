package test2006

func countKDifference(nums []int, k int) int {
	var cntMap = make(map[int]int)
	res := 0
	for _, num := range nums {
		v1, v2 := num+k, num-k
		res += cntMap[v1] + cntMap[v2]
		cntMap[num]++
	}
	return res
}
