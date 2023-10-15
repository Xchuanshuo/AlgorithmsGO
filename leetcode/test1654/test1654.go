package test1654

/**
 * @Description https://leetcode.cn/problems/minimum-jumps-to-reach-home/
 * idea: bfs 上界r分情况讨论，若a>b, 超出x后，继续走a步，实际没法回走了，此时r=b+x
					  若a<b，对于a+b，考虑r=max(f+a+b,x), 超出上界后一定可以
					  通过先走b再走a的方式回到界限内，所以依次作为上界
 **/

func minimumJumps(forbidden []int, a int, b int, x int) int {
	if x == 0 {
		return 0
	}
	var vis = make([]map[int]bool, 2)
	for i := 0; i < len(vis); i++ {
		vis[i] = make(map[int]bool)
	}
	var mf = 0
	for _, f := range forbidden {
		vis[0][f] = true
		vis[1][f] = true
		mf = max(mf, f)
	}
	var limit = max(mf+a+b, b+x)
	var dirs = []int{-b, a}
	var q = make([][]int, 0)
	q = append(q, []int{0, 0})
	vis[0][0], vis[1][0] = true, true
	var level = 0
	for len(q) > 0 {
		var sz = len(q)
		for i := 0; i < sz; i++ {
			var cur = q[0]
			if cur[0] == x {
				return level
			}
			q = q[1:]
			for j, d := range dirs {
				var nxt = cur[0] + d
				if nxt < 0 || vis[j][nxt] || nxt > limit {
					continue
				}
				if j == 0 && cur[1] == 0 {
					continue
				}
				vis[j][nxt] = true
				q = append(q, []int{nxt, j})
			}
		}
		level++
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
