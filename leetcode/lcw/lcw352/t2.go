package main

import (
	"fmt"
	"math"
	"sort"
)

var prims = make([]int, 0)

func init() {
	getPrims(int(1e6))
}

func findPrimePairs(n int) [][]int {
	var mp = make(map[int]bool)
	var res = make([][]int, 0)
	for i := 0; i < len(prims); i++ {
		if prims[i] >= n {
			break
		}
		if mp[n-prims[i]] || prims[i]*2 == n {
			res = append(res, []int{n - prims[i], prims[i]})
		}
		mp[prims[i]] = true
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
	return res
}

func getPrims(n int) []int {
	var isPrim = make([]bool, n+1)
	for i := 0; i < len(isPrim); i++ {
		isPrim[i] = true
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if isPrim[i] {
			for j := i * i; j <= n; j += i {
				isPrim[j] = false
			}
		}
	}
	for i := 2; i < len(isPrim); i++ {
		if isPrim[i] {
			prims = append(prims, i)
		}
	}
	return prims
}

func main() {
	fmt.Println(findPrimePairs(10))
}
