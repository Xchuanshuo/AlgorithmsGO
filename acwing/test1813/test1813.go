package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	res := make([]int, 26)
	for i := 0;i < n;i++ {
		var a, b string
		fmt.Scanf("%s %s\n", &a, &b)
		cntA, cntB := make([]int, 26), make([]int, 26)
		for _, ch := range a { cntA[ch-'a']++ }
		for _, ch := range b { cntB[ch-'a']++ }
		for j := 0;j < 26;j++ {
			res[j] += max(cntA[j], cntB[j])
		}
	}
	for i := 0;i < 26;i++ {
		fmt.Println(res[i])
	}
}

func max(a, b int) int {
	if a > b { return a }
	return b
}
