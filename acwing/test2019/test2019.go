package main

import (
	"container/heap"
	"fmt"
)

func main() {
	var n, sx, sy int
	fmt.Scanln(&n, &sx, &sy)
	var grid = make([][]int, 1003)
	for i := 0;i < 1003;i++ {
		grid[i] = make([]int, 1003)
	}
	for i := 0;i < n;i++ {
		var x, y int
		fmt.Scanln(&x, &y)
		grid[x][y] = 1
	}
	res := solve(sx, sy, grid)
	fmt.Println(res)
}

var dirs = [][]int{{1, 0}, {0, 1}, {0, -1}, {-1, 0}}

func solve(x int, y int, grid [][]int) int {
	var INF = int(1e9)
	var dist [1003][1003]int
	for i := 0;i < len(dist);i++ {
		for j := 0;j < len(dist[i]);j++{
			dist[i][j] = INF
		}
	}
	h := &Heap{}
	h.grid = grid
	dist[x][y] = 0
	heap.Push(h, []int{x, y})
	for h.Len() > 0 {
		cur := heap.Pop(h).([]int)
		if cur[0] == 0 && cur[1] == 0 {
			return dist[0][0]
		}
		for i := 0;i < len(dirs);i++ {
			nx, ny := cur[0] + dirs[i][0], cur[1] + dirs[i][1]
			if nx<0 || nx>=1002 || ny<0 || ny>=1002 ||
				dist[nx][ny]<=dist[cur[0]][cur[1]] + grid[nx][ny] {
				continue
			}
			dist[nx][ny] = dist[cur[0]][cur[1]] + grid[nx][ny]
			heap.Push(h, []int{nx, ny})
		}
	}
	return -1
}

type intArr [][]int

type Heap struct {
	intArr
	grid [][]int
}

func (h Heap) Len() int { return len(h.intArr) }

func (h Heap) Swap(i, j int) {
	h.intArr[i], h.intArr[j] = h.intArr[j], h.intArr[i]
}

func (h *Heap) Less(i, j int) bool {
	c1 := h.intArr[i]
	c2 := h.intArr[j]
	return h.grid[c1[0]][c1[1]] < h.grid[c2[0]][c2[1]]
}

func (h *Heap) Push(x interface{}) {
	h.intArr = append(h.intArr, x.([]int))
}

func (h *Heap) Pop() interface{} {
	a := h.intArr
	x := a[len(a)-1]
	h.intArr = a[0 : len(a)-1]
	return x
}


//func solve(x int, y int, grid [][]int) int {
//	var INF = int(1e9)
//	var dist [1003][1003]int
//	for i := 0;i < len(dist);i++ {
//		for j := 0;j < len(dist[i]);j++{
//			dist[i][j] = INF
//		}
//	}
//	h := &Heap{}
//	q := make([][]int, 0)
//	dist[x][y] = 0
//	heap.Push(h, []int{x, y})
//	for len(q) > 0 {
//		cur := q[0]
//		if cur[0] == 0 && cur[1] == 0 {
//			return dist[0][0]
//		}
//		q = q[1:]
//		for i := 0;i < len(dirs);i++ {
//			nx, ny := cur[0] + dirs[i][0], cur[1] + dirs[i][1]
//			if nx<0 || nx>=1002 || ny<0 || ny>=1002 ||
//				dist[nx][ny]<=dist[cur[0]][cur[1]] + grid[nx][ny] {
//				continue
//			}
//			dist[nx][ny] = dist[cur[0]][cur[1]] + grid[nx][ny]
//			if grid[nx][ny] == 1 {
//				q = append(q, []int{nx, ny})
//			} else {
//				nq := make([][]int, 0)
//				nq = append(nq, []int{nx, ny})
//				nq = append(nq, q...)
//				q = nq
//			}
//		}
//	}
//	return -1
//}