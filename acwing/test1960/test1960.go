package main

import "fmt"

var n, b int

func main() {
	fmt.Scanln(&n, &b)
	var state int
	for i := n-1;i >= 0;i-- {
		var v int
		fmt.Scanln(&v)
		state |= v << i
	}
	var visited = map[int]int{state: 0}
	var res = map[int]int{0: state}
	idx, start := 1, 0
	for true {
		next := getNext(state)
		if v, exist := visited[next];exist {
			start = v
			break
		}
		visited[next] = idx
		res[idx] = next
		state = next
		idx += 1
	}
	if b >= start {
		m := len(visited) - start
		b = (b-start)%m + start
	}
	for i := n-1;i >= 0;i-- {
		fmt.Println(res[b]>>i & 1)
	}
}

func getNext(s int) int {
	t := (s&1)<<(n-1) | (s>>1)
	return s^t
}