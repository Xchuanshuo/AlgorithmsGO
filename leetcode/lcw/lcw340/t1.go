package lcw340

import "math"

func diagonalPrime(nums [][]int) int {
	var m = len(nums[0])
	var res = 0
	for i := 0; i < m; i++ {
		j1, j2 := i, m-i-1
		if nums[i][j1] > res && isPrim(nums[i][j1]) {
			res = nums[i][j1]
		}
		if nums[i][j2] > res && isPrim(nums[i][j2]) {
			res = nums[i][j2]
		}
	}
	return res
}

func isPrim(a int) bool {
	if a == 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}
