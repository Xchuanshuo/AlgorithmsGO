package test640

import (
	"strconv"
	"strings"
)

func solveEquation(equation string) string {
	equation = strings.ReplaceAll(equation, "-", "+-")
	var arr = strings.Split(equation, "=")
	var leftArr = strings.Split(arr[0], "+")
	var rightArr = strings.Split(arr[1], "+")
	left, right := 0, 0
	for _, s := range leftArr {
		if s == "x" {
			left += 1
		} else if s == "-x" {
			left -= 1
		} else if strings.Index(s, "x") != -1 {
			v, _ := strconv.Atoi(s[0 : len(s)-1])
			left += v
		} else {
			v, _ := strconv.Atoi(s)
			right -= v
		}
	}
	for _, s := range rightArr {
		if s == "x" {
			left -= 1
		} else if s == "-x" {
			left += 1
		} else if strings.Index(s, "x") != -1 {
			v, _ := strconv.Atoi(s[0 : len(s)-1])
			left -= v
		} else {
			v, _ := strconv.Atoi(s)
			right += v
		}
	}
	if left == 0 {
		if right == 0 {
			return "Infinite solutions"
		} else {
			return "No solution"
		}
	}
	return "x=" + strconv.Itoa(right/left)
}
