package test1298

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	var q = make([]int, 0)
	var hasBox = make(map[int]bool)
	var canOpen = make(map[int]bool)
	var visited = make(map[int]bool)
	for i, b := range status {
		if b == 1 {
			canOpen[i] = true
		}
	}
	for _, b := range initialBoxes {
		hasBox[b] = true
		if canOpen[b] {
			q = append(q, b)
			visited[b] = true
		}
	}
	var res = 0
	for len(q) > 0 {
		var cur = q[0]
		q = q[1:]
		res += candies[cur]
		for _, key := range keys[cur] {
			canOpen[key] = true
			if !visited[key] && hasBox[key] { // 盒子没访问过且有钥匙对应的盒子
				visited[key] = true
				q = append(q, key)
			}
		}
		for _, b := range containedBoxes[cur] {
			hasBox[b] = true
			if !visited[b] && canOpen[b] { // 盒子没访问过且能打开盒子
				visited[b] = true
				q = append(q, b)
			}
		}
	}
	return res
}
