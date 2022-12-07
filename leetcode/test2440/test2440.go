package test2440

/**
 * @Description https://leetcode.cn/problems/create-components-with-same-value/
 * idea: 枚举，关键点 无向【树】
		1.要分成价值相同的块，总节点值需可整除块大小，即【块大小是总节点值的因子】
		2.数据范围大小为2*1e4*50=1e6, 总共因子数为【240】(n因子个数近似值直接取对n开立方即可)
        3.该数据范围下，从小到大枚举所有块值的划分，块值越小，块的个数越多，最终可以删除的边也最大
		  能划分当前块值，删除边数则为块个数【cnt - 1】
		4.子问题：dfs删除值为【块大小】的子树即可
		时间复杂度: O(N)*d(sum) d表示sum因子个数
 **/

func componentValue(nums []int, edges [][]int) int {
	mx, sum := len(nums), 0
	for _, num := range nums {
		sum += num
		mx = max(mx, sum)
	}
	var g = make(map[int][]int)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	var dfs func(val, x, f int) int
	dfs = func(total, cur, f int) int {
		var res = nums[cur]
		for _, nxt := range g[cur] {
			if nxt == f {
				continue
			}
			var val = dfs(total, nxt, cur)
			if val == -1 {
				return -1
			}
			res += val
		}
		if res == total { // 表示找到一个值为total的块
			return 0
		}
		if res > total { // 表示无法删除
			return -1
		}
		return res
	}
	for i := sum; i >= 1; i-- {
		if sum%i != 0 {
			continue
		}
		var total = sum / i
		if dfs(total, 0, -1) != -1 {
			return i - 1
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
