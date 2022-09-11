package test857

import (
	"container/heap"
	"math"
	"sort"
)

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	var n = len(quality)
	var arr = make([]int, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}
	sort.Slice(arr, func(i, j int) bool {
		i1, j1 := arr[i], arr[j]
		return wage[i1]*quality[j1] < wage[j1]*quality[i1]
	})
	var h = hp{make([]int, k)}
	var sum = 0
	for i := 0; i < k; i++ {
		sum += quality[arr[i]]
		h.IntSlice[i] = quality[arr[i]]
	}
	heap.Init(&h)
	var res = float64(sum) * float64(wage[arr[k-1]]) / float64(quality[arr[k-1]])
	for i := k; i < n; i++ {
		var cur = quality[arr[i]]
		if cur < h.IntSlice[0] {
			sum -= h.IntSlice[0] - cur
			h.IntSlice[0] = cur
			heap.Fix(&h, 0)
			res = math.Min(res, float64(sum)*float64(wage[arr[i]])/float64(cur))
		}
	}
	return res
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (hp) Push(interface{})     {}                                       // 由于没有用到，可以什么都不写
func (hp) Pop() (_ interface{}) { return }
