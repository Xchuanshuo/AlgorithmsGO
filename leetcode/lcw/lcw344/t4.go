package main

import "fmt"

func minIncrements(n int, cost []int) int {
	var res = 0
	var dfs func(int) int
	dfs = func(i int) int {
		if i >= n {
			return 0
		}
		left, right := dfs(i*2+1), dfs(i*2+2)
		res += 2*max(left, right) - (left + right)
		return cost[i] + max(left, right)
	}
	dfs(0)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	n := 15
	cost := []int{764, 1460, 2664, 764, 2725, 4556, 5305, 8829, 5064, 5929, 7660, 6321, 4830, 7055, 3761}
	res := minIncrements(n, cost)
	fmt.Println(res)
}
