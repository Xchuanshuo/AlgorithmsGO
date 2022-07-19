package main

import "fmt"

var n int
var points = make([][]int, 0)

func main() {
	fmt.Scanln(&n)
	minX, minY, maxX, maxY := int(1e9), int(1e9), 0, 0
	for i := 0;i < n;i++ {
		var x, y int
		fmt.Scanln(&x, &y)
		points = append(points, []int{x, y})
		minX = min(minX, x)
		minY = min(minY, y)
		maxX = max(maxX, x)
		maxY = max(maxY, y)
	}
	res := calc(minX)
	res = min(res, calc(minY))
	res = min(res, calc(maxX))
	res = min(res, calc(maxY))
	fmt.Println(res)
}

func calc(except int) int {
	var isExcepted = false
	minX, minY, maxX, maxY := int(1e9), int(1e9), 0, 0
	for i := 0;i < n;i++ {
		x, y := points[i][0], points[i][1]
		if (x == except || y == except) && !isExcepted {
			isExcepted = true
			continue
		}
		minX = min(minX, x)
		minY = min(minY, y)
		maxX = max(maxX, x)
		maxY = max(maxY, y)
	}
	return (maxX - minX) * (maxY - minY)
}

func min(a, b int) int {
	if a < b { return a }
	return b
}

func max(a, b int) int {
	if a > b { return a }
	return b
}