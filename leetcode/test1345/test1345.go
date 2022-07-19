package main

import "fmt"

func minJumps(arr []int) int {
	n := len(arr)
	var graph = make(map[int][]int)
	for i := n-1;i >= 0;i--{
		v := arr[i]
		if list, exist := graph[v];!exist {
			graph[v] = []int{i}
		} else {
			list = append(list, i)
			graph[v] = list
		}
	}
	var visited = make(map[int]bool)
	q := [][]int{{0,0}}
	visited[0] = true
	for len(q) > 0 {
		pos, dis := q[0][0], q[0][1]
		q = q[1:]
		if pos == n-1 {
			return dis
		}
		for _, nxt := range graph[arr[pos]] {
			if visited[nxt] { continue }
			visited[nxt] = true
			q = append(q, []int{nxt, dis+1})
		}
		delete(graph, arr[pos])
		if pos+1 < n && !visited[pos+1] {
			q = append(q, []int{pos+1, dis+1})
			visited[pos+1] = true
		}
		if pos-1 >= 0 && !visited[pos-1] {
			q = append(q, []int{pos-1, dis+1})
			visited[pos-1] = true
		}
	}
	return 0
}

func main() {
	//var arr = []int{100,-23,-23,404,100,23,23,23,3,404}
	var arr = []int{6, 1, 9}
	res := minJumps(arr)
	fmt.Println(res)
}