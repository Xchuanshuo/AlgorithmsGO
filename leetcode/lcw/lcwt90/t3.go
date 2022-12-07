package main

import (
	"fmt"
	"sort"
)

func destroyTargets(nums []int, space int) int {
	sort.Ints(nums)
	var mp = make(map[int]int)
	var mx = 0
	for _, num := range nums {
		mp[num%space]++
		if mp[num%space] > mx {
			mx = mp[num%space]
		}
	}
	for _, num := range nums {
		if mp[num%space] == mx {
			return num
		}
	}
	return 0
}

func main() {
	var arr = []int{625879766, 235326233, 250224393, 501422042, 683823101, 948619719, 680305710, 733191937, 182186779, 353350082}
	var space = 4
	res := destroyTargets(arr, space)
	fmt.Println(res)
	//fmt.Println(5)
}

//[625879766,235326233,250224393,501422042,683823101,948619719,680305710,733191937,182186779,353350082]
//4
