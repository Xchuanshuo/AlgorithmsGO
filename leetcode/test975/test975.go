package test975

import "github.com/emirpasic/gods/trees/redblacktree"

// 1.奇数跳: i -> j, min(a[j])>=a[i]
// 2.偶数跳: i -> j, max(a[j])<=a[i]
// dp[i][0]表示从跳奇数步到i的数量，dp[i][1]表示偶数步到i的数量
func oddEvenJumps(arr []int) int {
	var n = len(arr)
	tree := redblacktree.NewWithIntComparator()
	tree.Put(arr[n-1], n-1)
	var nxts = make([][]int, n)
	for i := n - 2; i >= 0; i-- {
		var cur = []int{-1, -1}
		if odd, ok := tree.Ceiling(arr[i]); ok {
			cur[0] = odd.Value.(int)
		}
		if even, ok := tree.Floor(arr[i]); ok {
			cur[1] = even.Value.(int)
		}
		nxts[i] = cur
		tree.Put(arr[i], i)
	}
	var dp = make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
		dp[i][1] = 1
	}
	for i := 0; i < n-1; i++ {
		var nxt = nxts[i]
		if nxt[0] != -1 {
			dp[nxt[0]][0] += dp[i][1]
		}
		if nxt[1] != -1 {
			dp[nxt[1]][1] += dp[i][0]
		}
	}
	return dp[n-1][0] + dp[n-1][1]
}
