package main

import (
	"fmt"
	"math"
)

var src = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var isPal = func(v int64) bool {
	var s = fmt.Sprintf("%d", v)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var set = make(map[int64]bool)
	for curV := 1000000; curV <= 10000000; curV++ {
		var tv = int64(math.Sqrt(float64(curV)))
		if tv*tv == int64(curV) && isPal(int64(curV)) && isPal(tv) {
			set[int64(curV)] = true
		}
	}
	fmt.Println(set)
	fmt.Println(math.Sqrt(1002001))
}
