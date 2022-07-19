package test507

import "math"

func checkPerfectNumber(num int) bool {
	if num == 1 { return false }
	r := math.Sqrt(float64(num))
	sum := 1
	for i := 2; i <= int(r); i++{
		if num%i == 0 {
			sum += i + num/i
		}
	}
	return sum == num
}