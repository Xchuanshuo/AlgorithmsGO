package test1475

func finalPrices(prices []int) []int {
	var s = make([]int, 0)
	var n = len(prices)
	var res = make([]int, n)
	for i := n - 1; i >= 0; i-- {
		res[i] = prices[i]
		for len(s) > 0 && prices[i] < s[len(s)-1] {
			s = s[0 : len(s)-1]
		}
		if len(s) > 0 {
			res[i] -= s[len(s)-1]
		}
		s = append(s, prices[i])
	}
	return res
}
