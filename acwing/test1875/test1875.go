package main

import "fmt"

var oddMap = map[int]int{}
var evenMap = map[int]int{}
var a = map[int]int{}
var res = 0
const str = "BESIGOM"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0;i < n;i++ {
		var c,num int
		fmt.Scanf("%c %d\n", &c, &num)
		if num%2 == 0 {
			evenMap[c]++
		} else {
			oddMap[c]++
		}
	}
	dfs(0, 1)
	fmt.Println(res)
}

func dfs(th, x int) {
	if th == 7 {
		if (a['B']+a['I'])%2==0 || a['M']%2==0 || (a['G']+a['O']+a['E']+a['S'])%2==0 {
			res += x
		}
		return
	}
	ch := int(str[th])
	a[ch] = 1
	dfs(th+1, x*oddMap[ch])
	a[ch] = 2
	dfs(th+1, x*evenMap[ch])
}