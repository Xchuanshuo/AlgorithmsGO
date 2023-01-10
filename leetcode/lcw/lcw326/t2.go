package lcw326

import (
	"math"
)

var primMap = make(map[int][]int)

func init() {
	var f = func(n int) []int {
		var facs = make([]int, 0)
		var visited = make(map[int]bool)
		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
			if n%i == 0 {
				for n%i == 0 {
					n /= i
					if !visited[i] {
						facs = append(facs, i)
						visited[i] = true
					}
				}
			}
		}
		if n > 1 && !visited[n] {
			facs = append(facs, n)
		}
		return facs
	}
	for i := 2; i <= 1000; i++ {
		primMap[i] = f(i)
	}
}

func distinctPrimeFactors(nums []int) int {
	var visited = make(map[int]bool)
	for _, num := range nums {
		for _, c := range primMap[num] {
			visited[c] = true
		}
	}
	return len(visited)
}
