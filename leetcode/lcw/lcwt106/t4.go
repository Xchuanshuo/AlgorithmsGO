package lcwt106

import "sort"

/**
 * @Description https://leetcode.cn/problems/find-a-good-subset-of-the-matrix/
 * idea: 1.存在全为0的行，肯定满足条件，直接返回结果
		2.不存在 要满足题意
		 假设取偶数行，则需要满足选中的行，每一列最多有不超过半数的1，若存在则取其中的两行照样满足题意；
         同理选取奇数行，也只能有不超过半数的1，去掉1行也满足题意，去掉后同偶数行一样。
		 所以最少选取两行，看是否存在互补，存在则找到了一个答案。由于列最长为5，直接枚举所有状态即可
 **/

func goodSubsetofBinaryMatrix(g [][]int) []int {
	m, n := len(g), len(g[0])
	var mp = make(map[int]int)
	for i := 0; i < m; i++ {
		var cur = 0
		for j := 0; j < n; j++ {
			cur |= g[i][j] << j
		}
		mp[cur] = i
	}
	if v, exist := mp[0]; exist {
		return []int{v}
	}
	for k1, v1 := range mp {
		for k2, v2 := range mp {
			if k1&k2 == 0 {
				var res = []int{v1, v2}
				sort.Ints(res)
				return res
			}
		}
	}
	return []int{}
}
