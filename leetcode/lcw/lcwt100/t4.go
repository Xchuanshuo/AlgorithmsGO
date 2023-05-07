package lcwt100

import "math"

func repairCars(ranks []int, cars int) int64 {
	l, r := int64(0), int64(math.MaxInt64)
	var check = func(total int64) bool {
		var cur = int64(0)
		for _, v := range ranks {
			cur += int64(math.Sqrt(float64(total / int64(v))))
		}
		return cur >= int64(cars)
	}
	for l <= r {
		var mid = l + (r-l)/2
		if !check(mid) {
			l = mid + 1
		} else {
			if mid == 0 || !check(mid-1) {
				return mid
			}
			r = mid - 1
		}
	}
	return 0
}
