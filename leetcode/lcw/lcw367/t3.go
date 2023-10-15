package lcw367

func findIndices(nums []int, id int, vd int) []int {
	var n = len(nums)
	if n == 1 && id == 0 && vd == 0 {
		return []int{0, 0}
	}
	var lm = make([][2]int, n)
	lm[0][0], lm[0][1] = nums[0], nums[0]
	for i := 1; i < n; i++ {
		lm[i][0] = min(lm[i-1][0], nums[i])
		lm[i][1] = max(lm[i-1][1], nums[i])
		var l = i - id
		if l >= 0 && (abs(nums[i]-lm[l][0]) >= vd || abs(nums[i]-lm[l][1]) >= vd) {
			for j := l; j >= 0; j-- {
				if abs(nums[i]-nums[j]) >= vd {
					return []int{j, i}
				}
			}
		}
	}
	var rm = make([][2]int, n)
	rm[n-1][0], rm[n-1][1] = nums[n-1], nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rm[i][0] = min(rm[i+1][0], nums[i])
		rm[i][1] = max(rm[i+1][1], nums[i])
		var r = i + id
		if r < n && (abs(nums[i]-rm[r][0]) >= vd || abs(nums[i]-rm[r][1]) >= vd) {
			for j := r; j < n; j++ {
				if abs(nums[i]-nums[j]) >= vd {
					return []int{i, j}
				}
			}
			return []int{r, i}
		}
	}
	return []int{-1, -1}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
