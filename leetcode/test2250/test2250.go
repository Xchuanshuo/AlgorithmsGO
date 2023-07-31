package test2250

import "sort"

func countRectangles(rectangles [][]int, points [][]int) []int {
	var arr = make([][]int, 0)
	for _, v := range rectangles {
		arr = append(arr, append(v, -1))
	}
	for i, v := range points {
		arr = append(arr, append(v, i))
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {
			return arr[i][2] > arr[j][2]
		}
		return arr[i][0] < arr[j][0]
	})
	var mp = make(map[int]int)
	var res = make([]int, len(points))
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i][2] < 0 {
			mp[arr[i][1]]++
		} else {
			var cnt = 0
			for j := arr[i][1]; j <= 100; j++ {
				cnt += mp[j]
			}
			res[arr[i][2]] = cnt
		}
	}
	return res
}
