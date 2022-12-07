package main

import "fmt"

func makeIntegerBeautiful(n int64, target int) int64 {
	if getV(n) <= target {
		return 0
	}
	var t = n
	var offset int64 = 10
	for offset < n*10 {
		var d = offset - t%offset
		t += d
		if getV(t) <= target {
			return d
		}
		t -= d
		offset *= 10
	}
	return 0
}

func getV(n int64) int {
	var v int64 = 0
	for n > 0 {
		v += n % 10
		n /= 10
	}
	return int(v)
}

func main() {
	var res = makeIntegerBeautiful(8, 2)
	fmt.Println(res)
}
