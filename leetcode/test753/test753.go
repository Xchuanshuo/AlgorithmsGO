package test753

/**
 * @Description https://leetcode.cn/problems/cracking-the-safe
 * idea: Hierholzer算法求解欧拉路径
		1.n-1位基础上每加1位密码，得出一个新的密码对 => 前n-1位作为节点, 第n位用(0,k)作为边，寻找一条经过所有边的路径
        2.具体实现时，从n-1个0组成的节点出发到n-1个0组成的节点结束。由于1<=n<=4, 可以对10^(n-1)取模，这样寻找下一条
		  路径时 实际是用前n位数的后n-1位作为作为上一个节点，依次从(0,k)值的边继续寻找下一路径
 **/

import "math"

func crackSafe(n int, k int) string {
	var h = int(math.Pow10(n - 1))
	var bytes = make([]byte, 0)
	var dfs func(s int)
	var set = make(map[int]bool)
	dfs = func(s int) {
		for i := 0; i < k; i++ {
			var cur = s*10 + i
			if set[cur] {
				continue
			}
			set[cur] = true
			dfs(cur % h)
			bytes = append(bytes, byte('0'+i))
		}
	}
	dfs(0)
	for i := 0; i < n-1; i++ {
		bytes = append(bytes, '0')
	}
	return string(bytes)
}
