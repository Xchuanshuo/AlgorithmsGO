package main

import (
	"sort"
	"strconv"
	"strings"
)

var days = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	var sum = make([]int, len(days)+1)
	for i := 0; i < len(days); i++ {
		sum[i+1] = sum[i] + days[i]
	}
	var arr = make([]T, 0)
	arr = append(arr, T{arriveAlice, 1}, T{leaveAlice, -1}, T{arriveBob, 1}, T{leaveBob, -1})
	sort.Slice(arr, func(i, j int) bool {
		a, b := strings.Split(arr[i].s, "-"), strings.Split(arr[j].s, "-")
		a1, _ := strconv.Atoi(a[0])
		a2, _ := strconv.Atoi(a[1])
		b1, _ := strconv.Atoi(b[0])
		b2, _ := strconv.Atoi(b[1])
		if a1 == b1 {
			if a2 == b2 {
				return arr[i].v > arr[j].v
			}
			return a2 < b2
		}
		return a1 < b1
	})
	var cnt = 0
	var resArr = make([]string, 0)
	for idx, v := range arr {
		cnt += v.v
		if cnt == 2 {
			resArr = append(resArr, v.s)
			resArr = append(resArr, arr[idx+1].s)
			break
		}
	}
	a, b := strings.Split(resArr[0], "-"), strings.Split(resArr[1], "-")
	a1, _ := strconv.Atoi(a[0])
	a2, _ := strconv.Atoi(a[1])
	b1, _ := strconv.Atoi(b[0])
	b2, _ := strconv.Atoi(b[1])
	return sum[b1] + b2 - (sum[a1-1] + a2)
}

type T struct {
	s string
	v int
}
