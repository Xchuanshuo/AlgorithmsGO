package lcw337

var dirs = [][]int{{-2, 1}, {-2, -1}, {2, 1}, {2, -1}, {1, 2}, {-1, 2}, {1, -2}, {-1, -2}}
func checkValidGrid(grid [][]int) bool {
	var n = len(grid)
	var cur = []int{-2, -1}
	for i := 0;i <= n*n-1;i++ {
		x, y := cur[0], cur[1]
		var valid = false
		for _, d := range dirs {
			nx, ny := x+d[0], y+d[1]
			if nx<0 || nx>=n || ny<0 || ny>=n {
				continue
			}
			if grid[nx][ny] == i {
				valid = true
				cur = []int{nx, ny}
				break
			}
		}
		if !valid {
			return false
		}
	}
	return true
}
