package test1765


var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	q := make([][]int, 0)
	for i := 0;i < m;i++ {
		for j := 0;j < n;j++ {
			if isWater[i][j] == 1 {
				isWater[i][j] = 0
				q = append(q, []int{i, j, 0})
			} else {
				isWater[i][j] = -1
			}
		}
	}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for i := 0;i < 4;i++ {
			nx, ny := cur[0]+dirs[i][0], cur[1]+dirs[i][1]
			if nx<0 || nx>=m || ny<0 || ny>=n || isWater[nx][ny]!=-1 {
				continue
			}
			isWater[nx][ny] = cur[2] + 1
			q = append(q, []int{nx, ny, isWater[nx][ny]})
		}
	}
	return isWater
}