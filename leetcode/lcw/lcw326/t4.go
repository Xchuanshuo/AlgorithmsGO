package lcw326

import (
	"math"
	"sort"
)

var prims = make([]int, 0)

func init() {
	var n = int(1e6)
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
}

func closestPrimes(left int, right int) []int {
	var l = sort.Search(len(prims), func(i int) bool {
		return prims[i] >= left
	})
	var r = sort.Search(len(prims), func(i int) bool {
		return prims[i] > right
	}) - 1
	var res = []int{-1, -1}
	var v = int(1e6)
	for i := l; i < r; i++ {
		if prims[i+1]-prims[i] < v {
			res = []int{prims[i], prims[i+1]}
			v = prims[i+1] - prims[i]
		}
	}
	return res
}

// 获取2..n范围内的质数
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
