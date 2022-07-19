package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)
	var arr = make([]int, 0)
	for i := 0;i < n;i++ {
		var x, v int
		fmt.Fscan(in, &x, &v)
		arr = append(arr, v)
	}
	res, min := 0, int(1e9) + 1
	for i := len(arr)-1;i >= 0;i-- {
		if arr[i] <= min {
			res++
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	fmt.Println(res)
}

//func main() {
//	var n int
//	fmt.Scanf("%d", &n)
//	var arr = make([][]int, 0)
//	for i := 0;i < n;i++ {
//		var x, v int
//		fmt.Scanf("%d %d\n", &x, &v)
//		arr = append(arr, []int{x, v})
//	}
//	q := []int{arr[0][1]}
//	for i := 1;i < n;i++ {
//		for len(q)>0 && q[len(q)-1]>arr[i][1] {
//			q = q[0:len(q)-1]
//		}
//		q = append(q, arr[i][1])
//	}
//	fmt.Println(len(q))
//}