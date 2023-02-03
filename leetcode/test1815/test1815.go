package test1815

/**
 * @Description https://leetcode.cn/problems/maximum-number-of-groups-getting-fresh-donuts/
 * idea: 状态压缩 + 记忆化搜索
		 题目本质上是求解, 前缀和尽可能整除size情况下的最优解, groups长度最大为30，直接枚举所有状态会超时
		 优化: 其实只需要关注group与size取模后，不同余数的组数 使其能达到最优解即可
			  1.对于余数为0的组安排在前，所有组都是感到开心的, 直接统计组数即可
			  2.对于余数不为0的组，由于size最大为9，取余后余数只有[1,8], 而groups最大为30，取余后一个组
				使用5bit即可存储，总共状态数为 8*5=40bit, 所以使用一个int64存储状态即可
			  3.枚举所有方案，每个余数的组数计算: state += 1 << (r * 5)，
						   选取一个组后，状态表示为state-(1<<(r*5)),
				计算当前是否是开心的组，需要将r与上一次状态的mod相加后对size取余, 看结果是否为0即可
 **/

func maxHappyGroups(size int, groups []int) int {
	var mem = make(map[int64]int)
	var state int64 = 0
	var res = 0
	for _, g := range groups {
		var r = g % size
		if r == 0 {
			res++
		} else {
			state += 1 << (r * 5)
		}
	}
	var dfs func(sta int64, mod int) int
	dfs = func(sta int64, mod int) int {
		if v, exist := mem[sta]; exist {
			return v
		}
		var add = 0
		if mod == 0 {
			add = 1
		}
		var curRes = 0
		for r := 1; r < size; r++ {
			if (sta>>(r*5))&31 != 0 {
				var cur = dfs(sta-(1<<(r*5)), (mod+r)%size)
				curRes = max(curRes, cur+add)
			}
		}
		mem[sta] = curRes
		return curRes
	}
	res += dfs(state, 0)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
