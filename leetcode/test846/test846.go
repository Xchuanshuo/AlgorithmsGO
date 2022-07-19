package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	sort.Ints(hand)
	cntMap := make(map[int]int)
	for _, x := range hand {
		cntMap[x]++
	}
	for _, h := range hand {
		if cntMap[h] == 0 {
			continue
		}
		for i := 0;i < groupSize;i++ {
			next := h + i
			if cntMap[next] == 0 { return false }
			cntMap[next]--
			if cntMap[next] == 0 {
				delete(cntMap, next)
			}
		}
	}
	return true
}

func isNStraightHand1(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	sort.Ints(hand)
	cntMap := make(map[int]int)
	h := &Heap{}
	for _, x := range hand {
		if cntMap[x] == 0 {
			heap.Push(h, x)
		}
		cntMap[x]++
	}
	for h.Len() > 0 {
		last := -1
		tArr := make([]int, 0, groupSize)
		for i := 0;i < groupSize;{
			if h.Len() <= 0 { return false }
			cur := heap.Pop(h).(int)
			if last != -1 && cur-last != 1 {
				return false
			}
			cntMap[cur]--
			if cntMap[cur] > 0 { tArr = append(tArr, cur) }
			last = cur
			i++
		}
		for len(tArr) > 0 {
			t := tArr[0]
			heap.Push(h, t)
			tArr = tArr[1:]
		}
	}
	return true
}

type Heap struct {
	sort.IntSlice
}

func (h Heap) Less(i, j int) bool {
	return h.IntSlice[i] < h.IntSlice[j]
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

func main() {
	hand := []int{1,2,3,6,2,3,4,7,8}
	groupSize := 3
	v := isNStraightHand(hand, groupSize)
	fmt.Println(v)
	m := make(map[int]int)
	_, exist := m[1]
	fmt.Println(exist)
}