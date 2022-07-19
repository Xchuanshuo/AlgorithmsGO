package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var n, pre int
	dif := make(map[int]int)
	arr := make([]int, 0)
	pre = 0
	fmt.Scanln(&n)
	for i := 0;i < n;i++ {
		var step,dir string
		fmt.Scanln(&step, &dir)
		v, _ := strconv.Atoi(step)
		neg := 1
		if dir == "L" {neg = -1}
		next := pre + v*neg
		var cur [2]int
		if dir == "L" {
			cur = [2]int{next, pre}
		} else {
			cur = [2]int{pre, next}
		}
		if _, exist := dif[cur[0]];!exist {
			arr = append(arr, cur[0])
		}
		if _, exist := dif[cur[1]];!exist {
			arr = append(arr, cur[1])
		}
		dif[cur[0]]++
		dif[cur[1]]--
		pre = next
	}
	sort.Ints(arr)
	res, sum := 0, dif[arr[0]]
	for i := 1;i < len(arr);i++ {
		sum += dif[arr[i]]
		if sum - dif[arr[i]] > 1 {
			res += arr[i] - arr[i-1]
		}
	}
	fmt.Println(res)
}
