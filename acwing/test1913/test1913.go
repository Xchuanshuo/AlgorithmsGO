package main

import (
	"fmt"
	"sort"
)

var arr = make([][]int, 0)
var res = 0

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0;i < n;i++ {
		var x, b int
		fmt.Scanf("%d %c\n", &x, &b)
		arr = append(arr, []int{x, b})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	// 1.全部为G
	calc1('G')
	// 2.全部为H
	calc1('H')
	// 3.G和H数量相等
	calc2()
	fmt.Println(res)
}

func calc1(c int) {
	si := 0
	for i := 0;i < len(arr);i++ {
		if arr[i][1] == c {
			res = max(res, arr[i][0] - arr[si][0])
		} else {
			si = i+1
		}
	}
}

func calc2()  {
	var posMap = make(map[int]int)
	posMap[0] = -1
	cnt := 0
	for i := 0;i < len(arr);i++ {
		if arr[i][1] == 'G' {
			cnt++
		} else {
			cnt--
		}
		if v, exist := posMap[cnt];exist {
			si := v + 1
			res = max(res, arr[i][0] - arr[si][0])
		} else {
			posMap[cnt] = i
		}
	}
}

func max(a, b int) int {
	if a > b { return a }
	return b
}
