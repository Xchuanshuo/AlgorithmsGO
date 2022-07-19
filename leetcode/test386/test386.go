package test386

var res []int

func lexicalOrder(n int) []int {
	res = make([]int, 0)
	for i := 1; i <= 9; i++ {
		dfs(i, n)
	}
	return res
}

func dfs(cur int, n int) {
	if cur > n {
		return
	}
	res = append(res, cur)
	for i := 0; i < 10; i++ {
		dfs(cur*10+i, n)
	}
}
