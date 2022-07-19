package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scanln(&n, &k)
	var posMap = make(map[int]int)
	res := -1
	for i := 0;i < n;i++ {
		var cur int
		fmt.Scanln(&cur)
		if _, exist := posMap[cur]; exist {
			dst := i - posMap[cur]
			if dst <= k && cur > res {
				res = cur
			}
		}
		posMap[cur] = i
	}
	fmt.Println(res)
}
