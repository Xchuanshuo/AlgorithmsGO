package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var dis = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dis[i] = make([]int, n+1)
	}
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		dis[a][b] = 1
	}
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				dis[i][j] |= dis[i][k] & dis[k][j]
			}
		}
	}
	res := -1
	for i := 1; i <= n; i++ {
		var cnt = 0
		for j := 1; j <= n; j++ {
			if dis[j][i] == 1 {
				cnt++
			}
		}
		if cnt == n-1 {
			res = i
			break
		}
	}

	fmt.Println(res)
}
