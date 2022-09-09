package main

import (
	"fmt"
	"math"
)

func preimageSizeFZF(k int) int {
	l, r := findL(k), findR(k)
	if l >= r {
		return 0
	}
	return r - l + 1
}

func findR(k int) int {
	l, r := 0, math.MaxInt64
	for l <= r {
		mid := l + (r-l)/2
		if getCnt(mid) > k {
			r = mid - 1
		} else {
			if mid+1 == math.MaxInt64 || getCnt(mid+1) > k {
				return mid
			}
			l = mid + 1
		}
	}
	return 0
}

func findL(k int) int {
	l, r := 0, math.MaxInt64
	for l <= r {
		mid := l + (r-l)/2
		if getCnt(mid) < k {
			l = mid + 1
		} else {
			if mid == 0 || getCnt(mid-1) < k {
				return mid
			}
			r = mid - 1
		}
	}
	return 0
}

func getCnt(n int) int {
	var res = 0
	t, mul := n, 5
	for t > 0 {
		res += t / mul
		t /= mul
	}
	return res
}

func main() {
	var k = 92
	var res = preimageSizeFZF(k)
	fmt.Println(res)
}
