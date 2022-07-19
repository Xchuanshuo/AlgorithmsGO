package test1154

import (
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	strs := strings.Split(date, "-")
	year, _ := strconv.Atoi(strs[0])
	month, _ := strconv.Atoi(strs[1])
	day, _ := strconv.Atoi(strs[2])
	days := [...]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if (year%100!=0 && year%400==0) || year%4 == 0 {
		days[2] += 1
	}
	res := day
	for i := 1;i < month;i++ {
		res += days[i]
	}
	return res
}