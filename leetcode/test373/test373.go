package test373

import (
	"container/heap"
)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := &Heap{}
	n1, n2 := len(nums1), len(nums2)
	for i := 0;i < n1;i++ {
		heap.Push(h, []int{nums1[i]+nums2[0], i, 0})
	}
	var res = make([][]int, 0)
	for i := 0;i < k && h.Len()>0;i++ {
		cur := heap.Pop(h).([]int)
		res = append(res, []int{nums1[cur[1]], nums2[cur[2]]})
		if cur[2]+1 < n2 {
			heap.Push(h, []int{nums1[cur[1]]+nums2[cur[2]+1], cur[1], cur[2]+1})
		}
	}
	return res
}

type IntArr [][]int
type Heap struct {
	IntArr
}

func (h Heap) Len() int {
	return len(h.IntArr)
}

func (h Heap) Less(i, j int) bool {
	return h.IntArr[i][0] < h.IntArr[j][0]
}

func (h Heap) Swap(i, j int) {
	h.IntArr[i], h.IntArr[j] = h.IntArr[j], h.IntArr[i]
}

func (h *Heap) Push(x interface{}) {
	h.IntArr = append(h.IntArr, x.([]int))
}

func (h *Heap) Pop() interface{} {
	old := h.IntArr
	n := len(old)
	x := old[n-1]
	h.IntArr = old[0:n-1]
	return x
}