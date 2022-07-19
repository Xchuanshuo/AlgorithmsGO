package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)
	var hasEdge = make([][]bool, n)
	for i := 0; i < n; i++ {
		hasEdge[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		var f, s int
		fmt.Scanf("%d %d\n", &f, &s)
		f, s = f-1, s-1
		hasEdge[f][s] = true
		hasEdge[s][f] = true
	}
	var bytes = make([]byte, n)
	bytes[0] = '1'
	for i := 1; i < n; i++ {
		var visited = make(map[byte]bool)
		for j := 0; j < i; j++ {
			if hasEdge[j][i] {
				visited[bytes[j]] = true
			}
		}
		for j := 1; j <= 4; j++ {
			if !visited[byte('0'+j)] {
				bytes[i] = byte('0' + j)
				break
			}
		}

	}
	fmt.Println(string(bytes))
}
