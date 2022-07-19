package test1994

var p = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

func numberOfGoodSubsets(nums []int) int {
	var cnts = make(map[int]int)
	mask, M := 1<<10, int(1e9)+7
	for _, v := range nums {
		cnts[v]++
	}
	var dp = make([]int, mask)
	dp[0] = 1
	for i := 2; i <= 30; i++ {
		cur, v := 0, i
		if cnts[v] == 0 {
			continue
		}
		var isValid = true
		for j := 0; j < 10; j++ {
			t, c := p[j], 0
			for v%t == 0 {
				cur |= 1 << j
				v /= t
				c++
			}
			if c > 1 {
				isValid = false
				break
			}
		}
		if !isValid {
			continue
		}
		for pre := mask - 1; pre >= 0; pre-- {
			if cur&pre != 0 {
				continue
			}
			dp[pre|cur] = (dp[pre|cur] + cnts[i]*dp[pre]) % M
		}
	}
	var res = 0
	for i := 1; i < mask; i++ {
		res = (res + dp[i]) % M
	}
	for i := 0; i < cnts[1]; i++ {
		res = (res * 2) % M
	}
	return res
}
