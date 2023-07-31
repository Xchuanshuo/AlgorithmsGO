package test1981

/**
 * @Description https://leetcode.cn/problems/minimize-the-difference-between-target-and-chosen-elements/
 * idea: 分组背包，统计和为j的背包是否能装满，需要注意的是每次查看新的组时，大小只能由前一组的和转移过来，即每行必须选择一个
		元素与前面的行选择一个元素后装入背包，所以可以使用两个dp数组进行【交替计算】，保证不会有行被跳过
		初始条件，第1行任选一个元素装入背包
 **/

func minimizeTheDifference(mat [][]int, target int) int {
	m, n := len(mat), len(mat[0])
	var mx = 0
	for i := 0; i < m; i++ {
		var cur = 0
		for j := 0; j < n; j++ {
			cur = max(cur, mat[i][j])
		}
		mx += cur
	}
	var dp0 = make([]bool, mx+1)
	for i := 0; i < n; i++ {
		dp0[mat[0][i]] = true
	}
	for i := 1; i < m; i++ {
		var dp1 = make([]bool, mx+1)
		for j := mx; j >= 1; j-- {
			for _, v := range mat[i] {
				if j < v || !dp0[j-v] {
					continue
				}
				dp1[j] = dp0[j-v]
			}
		}
		dp0 = dp1
	}
	var res = abs(mx - target)
	for i := 1; i <= mx; i++ {
		if dp0[i] {
			res = min(res, abs(i-target))
		}
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
