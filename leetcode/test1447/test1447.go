package test1447

import "fmt"

func simplifiedFractions(n int) []string {
	res := make([]string, 0)
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if gcd(i, j) == 1 {
				res = append(res, fmt.Sprintf("%d/%d", i, j))
			}
		}
	}
	return res
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
