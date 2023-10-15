package test2212

/**
 * @Description https://leetcode.cn/problems/maximum-points-in-an-archery-competition
 * idea: 贪心+状态压缩. 枚举所有选择情况s, 对于每个位置分配刚好超过alice一支箭的数量, 统计最大分数即可
 **/

func maximumBobPoints(total int, alice []int) []int {
	var n = len(alice)
	var l = 1 << n
	state, score := 0, 0
	for s := 0; s < l; s++ {
		var cur = total
		var curScore = 0
		for i := 0; i < n; i++ {
			if s&(1<<i) != 0 && cur > alice[i] {
				cur -= alice[i] + 1
				curScore += i
			}
			if curScore > score {
				score = curScore
				state = s
			}
		}
	}
	var res = make([]int, n)
	for i := 0; i < n; i++ {
		if state&(1<<i) != 0 && total > alice[i] {
			total -= alice[i] + 1
			res[i] = alice[i] + 1
		}
	}
	res[n-1] += total
	return res
}
