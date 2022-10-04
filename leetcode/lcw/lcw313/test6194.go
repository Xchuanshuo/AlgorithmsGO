package main

import (
	"math/bits"
)

func minimizeXor(num1 int, num2 int) int {
	var cnt = bits.OnesCount(uint(num2))
	var temp = make([]int, 32)
	for i := 0; i < 32; i++ {
		if (1<<i)&num1 > 0 {
			temp[i] = 1
		}
	}
	var clone = make([]int, 32)
	copy(clone, temp)
	var res = 0
	for i := 31; i >= 0; i-- {
		res = res * 2
		if cnt > 0 && temp[i] == 1 {
			res = res + 1
		}
		if temp[i] == 1 && cnt > 0 {
			temp[i] = 0
			cnt--
		}
	}
	if cnt == 0 {
		return res
	}
	res = 0
	for i := 0; i < 32; i++ {
		if clone[i] == 0 && cnt > 0 {
			clone[i] = 1
			cnt--
		}
	}
	for i := 31; i >= 0; i++ {
		res = res*2 + clone[i]
	}
	return res
}

//func main() {
//	// 24
//	num1, num2 := 25, 72
//	fmt.Println(minimizeXor(num1, num2))
//}
