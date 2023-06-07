package test1012

/**
 * @Description https://leetcode.cn/problems/numbers-with-repeated-digits/
 * idea: 至少重复1位 => 全部数字 - 1位不重复的数字 => 求不重复的数位的数字个数
		若数字[1,N]数位为n，对于后n-1位, 为 C(9,1)*A(9, n-2)
					     对于n位数，为A(9-i, n-i-1)
 **/

func numDupDigitsAtMostN1(N int) int {
	var a = make([]int, 0)
	var t = N + 1
	for t != 0 {
		a = append(a, t%10)
		t /= 10
	}
	var n = len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
	var cnt = 0
	// 小于n位数时
	for i := 1; i < n; i++ {
		cnt += 9 * A(9, i-1)
	}
	var visited = make([]bool, 10)
	// 刚好为n位数时(首位不能取0)
	for i := 0; i < n; i++ {
		var cur = a[i]
		var j = 0
		if i == 0 {
			j = 1
		}
		for ; j < cur; j++ {
			if visited[j] {
				continue
			}
			cnt += A(9-i, n-i-1)
		}
		if visited[cur] {
			break
		}
		visited[cur] = true
	}
	return N - cnt
}

func A(m, n int) int {
	var res = 1
	for i := 0; i < n; i++ {
		res *= m - i
	}
	return res
}
