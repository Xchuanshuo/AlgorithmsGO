package main

import (
	"fmt"
	"sort"
)


func main() {
	var n int
	fmt.Scanln(&n)
	var ascArr = make([]string, 0)
	var descArr = make([]string, 0)
	var arr = make([]string, 0)

	for i := 0;i < n;i++ {
		var chs []byte
		fmt.Scanln(&chs)
		arr = append(arr, string(chs))
		sort.Slice(chs, func(i, j int) bool {
			return chs[i] < chs[j]
		})
		str := string(chs)
		rStr := reverse(str)
		ascArr = append(ascArr, str)
		descArr = append(descArr, rStr)
	}
	sort.Strings(ascArr)
	sort.Strings(descArr)
	fmt.Println(ascArr)
	fmt.Println(descArr)
	fmt.Println(arr)
	solve(arr, ascArr, descArr)
}

func solve(src []string, asc []string, desc []string) {
	for i := 0;i < len(src);i++ {
		var chs = []byte(src[i])
		sort.Slice(chs, func(i, j int) bool {
			return chs[i] < chs[j]
		})
		minS := string(chs)
		maxS := reverse(minS)
		res1 := findL1(minS, desc)+1
		res2 := findL2(maxS, asc)
		fmt.Println(res1, res2)
	}
}

func findL2(t string, arr []string) int {
	n := len(arr)
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] <= t {
			l = mid + 1
		} else {
			if mid == l || arr[mid-1] <= t {
				return mid
			}
			r = mid - 1
		}
	}
	return n
}

func findL1(t string, arr []string) int {
	n := len(arr)
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] < t {
			l = mid + 1
		} else {
			if mid == l || arr[mid-1] < t {
				return mid
			}
			r = mid - 1
		}
	}
	return n-1
}

func reverse(str string) string {
	var chs = []byte(str)
	for i,j := 0, len(chs)-1;i < j;i,j = i+1, j-1 {
		chs[i], chs[j] = chs[j], chs[i]
	}
	return string(chs)
}
