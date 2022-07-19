package test473

var mem map[int]bool

func makesquare(matchsticks []int) bool {
	var total = 0
	for _, v := range matchsticks {
		total += v
	}
	if total%4 != 0 {
		return false
	}
	mem = make(map[int]bool)
	var boxes = make([]int, 0)
	for i := 0; i < 4; i++ {
		boxes = append(boxes, total/4)
	}
	return dfs(matchsticks, boxes, 0, 0)
}

func dfs(sticks []int, boxes []int, cur int, state int) bool {
	if v, exist := mem[state]; exist {
		return v
	}
	if boxes[cur] < 0 {
		return false
	}
	if boxes[cur] == 0 {
		cur++
	}
	if cur > 3 {
		return true
	}
	for i, s := range sticks {
		if state&(1<<i) != 0 {
			continue
		}
		state |= 1 << i
		boxes[cur] -= s
		var res = dfs(sticks, boxes, cur, state)
		mem[state] = res
		if res {
			return true
		}
		boxes[cur] += s
		state &= ^(1 << i)
	}
	return false
}
