package main

import "sort"

func numFriendRequests(ages []int) int {
	i, j := 0, 0
	sort.Ints(ages)
	res, n := 0, len(ages)
	for _, age := range ages {
		if age < 15 { continue }
		for ;i < n && ages[i] <= age/2 + 7;i++ {}
		for ;j+1 < n && ages[j+1] <= age;j++ {}
		res += j - i
	}
	return res
}

func numFriendRequests1(ages []int) int {
	cnt, sum := make([]int, 121), make([]int, 122)
	for _,v := range ages {
		cnt[v]++
	}
	for i := 1;i <= 120;i++ {
		sum[i] = sum[i-1] + cnt[i]
	}
	res := 0
	for _, age := range ages {
		if age < 15 { continue }
		res += sum[age] - sum[age/2+7]-1
	}
	return res
}