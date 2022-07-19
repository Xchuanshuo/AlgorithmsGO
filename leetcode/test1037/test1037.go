package test1037

func isBoomerang(points [][]int) bool {
	dy1, dx1 := float64(points[1][1]-points[0][1]), float64(points[1][0]-points[0][0])
	dy2, dx2 := float64(points[2][1]-points[1][1]), float64(points[2][0]-points[1][0])
	return dy1*dx2 != dx1*dy2
}
