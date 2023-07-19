package test1483

/**
 * @Description https://leetcode.cn/problems/kth-ancestor-of-a-tree-node/
 * idea: dp (树上倍增)，暴力求解时间为O(n*k)，考虑优化。对于每个数k，转换成二进制，可以观察到
		任何数都可以写成多个2^i(i为[0,31])相加，所以可以进行预处理，对于每个节点i存储其第2^j个父节点的值，
		用dp[i][j]表示，dp[i][j] = dp[dp[i][j-1]][j-1], 即求i的第2^j个父节点，可以先求i的2^(j-1)个
		父节点，再继续求该节点的第2^(j-1)个父节点，最终 2^(j-1)+2^(j-1)=2^j
 **/

type TreeAncestor struct {
	dp [][20]int
	n  int
}

func Constructor(n int, parent []int) TreeAncestor {
	var dp = make([][20]int, n)
	for i := 0; i < n; i++ {
		dp[i][0] = parent[i]
		for j := 1; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}
	for j := 1; j < 20; j++ {
		for i := 1; i < n; i++ {
			if dp[i][j-1] != -1 {
				dp[i][j] = dp[dp[i][j-1]][j-1]
			}
		}
	}
	return TreeAncestor{dp, n}
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	var cur = node
	for i := 0; i < 20; i++ {
		if k&(1<<i) != 0 {
			cur = this.dp[cur][i]
		}
		if cur == -1 {
			break
		}
	}
	return cur
}
