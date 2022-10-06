package lcwt86

/**
 * @Description
 * idea: 1.维护预算小于等于budge的窗口 2.单调队列求窗口最大值
 **/

func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
	var n = len(chargeTimes)
	var sums = make([]int, n+1)
	var q = make([]int, 0)
	for i := 0; i < n; i++ {
		sums[i+1] = sums[i] + runningCosts[i]
	}
	var push = func(num int) int {
		for len(q) > 0 && q[len(q)-1] < num {
			q = q[0 : len(q)-1]
		}
		q = append(q, num)
		return q[0]
	}
	var pop = func(num int) {
		if num == q[0] {
			q = q[1:]
		}
	}
	res, mx := 0, 0
	l, r := 0, 0
	for r < n {
		cur := chargeTimes[r]
		push(cur)
		mx = q[0]
		var cost = int64(mx + (r-l+1)*(sums[r+1]-sums[l]))
		for cost > budget && l <= r {
			pop(chargeTimes[l])
			if len(q) == 0 {
				mx = 0
			} else {
				mx = q[0]
			}
			l++
			cost = int64(mx + (r-l+1)*(sums[r+1]-sums[l]))
		}
		res = max(res, r-l+1)
		r++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
