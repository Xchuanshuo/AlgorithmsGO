package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var a = make([]int, n-1)
	for i := 0; i < n-1; i++ {
		var v int
		fmt.Scanf("%d", &v)
		a[i] = v
	}
	var first = make([]int, 0)
	for i := 1; i < (a[0]+1)/2; i++ {
		first = append(first, i, a[0]-i)
	}
	sort.Ints(first)
	for k := 0; k < len(first); k++ {
		var visited = make(map[int]bool)
		last := first[k]
		visited[last] = true
		var res = []int{last}
		var isFound = true
		for i := 0; i < len(a); i++ {
			cur := a[i] - last
			if cur < 1 || visited[cur] {
				isFound = false
				break
			}
			visited[cur] = true
			res = append(res, cur)
			last = cur
		}
		if isFound {
			for i := 0; i < len(res); i++ {
				fmt.Printf("%d ", res[i])
			}
			break
		}
	}
}
