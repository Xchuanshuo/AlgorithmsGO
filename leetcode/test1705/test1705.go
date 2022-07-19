package test1705

import "container/heap"

func eatenApples(apples []int, days []int) int {
	i, res := 0, 0
	h := &Heap{}
	for ; i < len(apples); i++ {
		cnt, day := apples[i], i+days[i]
		if cnt > 0 {
			heap.Push(h, []int{cnt, day})
		}
		flag := false
		for h.Len() > 0 {
			cur := heap.Pop(h).([]int)
			if cur[1] > i {
				res++
				cur[0]--
				flag = true
				if cur[0] > 0 { heap.Push(h, cur) }
			}
			if flag { break }
		}
	}
	for h.Len() > 0 {
		cur := heap.Pop(h).([]int)
		if cur[1] > i {
			i++
			res++
			cur[0]--
			if cur[0] > 0 { heap.Push(h, cur) }
		}
	}
	return res
}

type intArr [][]int

type Heap struct {
	intArr
}

func (h Heap) Len() int { return len(h.intArr) }

func (h Heap) Swap(i, j int) {
	h.intArr[i], h.intArr[j] = h.intArr[j], h.intArr[i]
}

func (h Heap) Less(i, j int) bool {
	return h.intArr[i][1] < h.intArr[j][1]
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
