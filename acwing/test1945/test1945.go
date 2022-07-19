package main

import (
	"fmt"
	"sort"
)
var n int

func main() {
	fmt.Scanln(&n)
	var arr = make([]int, n)
	for i := 0;i < n;i++ {
		var t int
		fmt.Scanln(&t)
		arr[i] = t
	}
	sort.Ints(arr)
	res := 0
	for i := 0;i < n-2;i++ {
		for j := i+1;j < n-1;j++ {
			lv, rv := 2*arr[j]-arr[i], 3*arr[j]-2*arr[i]
			l, r := findL(arr, lv), findR(arr, rv)
			if l == -1 { continue }
			res += r - l + 1
		}
	}
	fmt.Println(res)
}

func findL(arr []int, t int) int {
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] < t {
			l = mid + 1
		} else {
			if mid == 0 || arr[mid-1] < t { return mid }
			r = mid - 1
		}
	}
	return -1
}

func findR(arr []int, t int) int {
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] > t {
			r = mid - 1
		} else {
			if mid == n-1 || arr[mid+1] > t { return mid }
			l = mid + 1
		}
	}
	return -1
}
