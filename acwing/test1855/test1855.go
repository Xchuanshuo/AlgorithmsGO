package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	var a = make([]int, n)
	var mp = make(map[int]int)
	for i := 0;i < n;i++ {
		var v int
		fmt.Scanln(&v)
		a[i] = v
		mp[a[i]] = -1
	}
	max := 0
	for _, v := range a {
		var t = make(map[int]int)
		for k, val := range mp { t[k] = val }
		t[v] = 1
		q := []int{v}
		num := 0
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			num++
			for j := -t[cur];j <= t[cur];j++ {
				if t[cur+j] == -1 {
					q = append(q, cur + j)
					t[cur+j] = t[cur] + 1
				}
			}
		}
		if num > max { max = num }
	}
	fmt.Println(max)
}
