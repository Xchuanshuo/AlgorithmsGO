package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	times := 0
	h := &Heap{}
	for _, c := range courses {
		if times + c[0] <= c[1] {
			times += c[0]
			heap.Push(h, c[0])
		} else if h.Len() > 0 && c[0] < h.Peek(){
			times += c[0] - heap.Pop(h).(int)
			heap.Push(h, c[0])
		}
	}
	return h.Len()
}


type Heap struct {
	sort.IntSlice
}

func (h Heap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *Heap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *Heap) Pop() interface{} {
	a := h.IntSlice
	x := a[len(a) - 1]
	h.IntSlice = a[0:len(a)-1]
	return x
}

func (h *Heap) Peek() int {
	return h.IntSlice[0]
}

func main() {
	courses := [][]int{{1,2}}
	res := scheduleCourse(courses)
	fmt.Println(res)
}
