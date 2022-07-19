package main

import "fmt"

var isSet = make(map[int]bool)
var posMap = make(map[int]int)
var layer = make([]int, 1)
var n, m, k int

func main() {
	fmt.Scanf("%d %d %d\n", &n, &m, &k)
	var isExist1 = false
	for i := 1; i <= m; i++ {
		var v int
		fmt.Scanf("%d", &v)
		layer = append(layer, v)
		if v == 1 {
			isExist1 = true
		}
	}

	for i := 0; i < k; i++ {
		var c, p int
		fmt.Scanf("%d %d\n", &c, &p)
		posMap[c] = p
		isSet[p] = true
	}
	if isExist1 {
		cal1()
	} else {
		cal2()
	}
	if v, exist := posMap[1]; exist {
		fmt.Println(v)
		return
	}
	for i := 1; i <= n; i++ {
		if !isSet[i] {
			fmt.Println(i)
			break
		}
	}
}

func cal1() {
	for i, j := 1, 1; i <= m && j <= n; i++ {
		if v, exist := posMap[layer[i]]; exist {
			j = v + 1
			continue
		}
		for isSet[j] {
			j++
		}
		posMap[layer[i]] = j
		isSet[j] = true
		j++
	}
}

func cal2() {
	for i, j := m, n; i >= 1 && j >= 1; i-- {
		if v, exist := posMap[layer[i]]; exist {
			j = v - 1
			continue
		}
		for isSet[j] {
			j--
		}
		posMap[layer[i]] = j
		isSet[j] = true
		j--
	}
}
