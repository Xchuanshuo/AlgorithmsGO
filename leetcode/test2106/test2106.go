package test2106

// 滑动窗口，l, r维护左右最远能到达的距离，窗口的扩容呈单调性，计算窗口的总数即可
// 1.全部往左走, 最远距离l=s-k ；全部往右走, 最远距离r=s+k
// 2.先左后右 s-l+r-l; 先右后左 r-s + r-l
func maxTotalFruits(f [][]int, s int, k int) int {
	var n = len(f)
	var l = 0
	cur, res := 0, 0
	for r := 0; r < n; r++ {
		if f[r][0]-s > k {
			break
		}
		cur += f[r][1]
		for l <= r && f[r][0]-f[l][0]+min(abs(s-f[l][0]), abs(f[r][0]-s)) > k {
			cur -= f[l][1]
			l++
		}
		res = max(res, cur)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
