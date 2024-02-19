package main

/**
 * @Description https://leetcode.cn/contest/weekly-contest-380/problems/maximum-number-that-sum-of-the-prices-is-less-than-or-equal-to-k/
 * idea: 二分 + 数位dp(统计满足限制条件下[0,二进制序列]中1的个数)
		易忽略点: 1.二进制位从1开始 2.从左到右的顺序计数时，需要转换成限制条件(从右到左)去过滤
 **/

var bits = make(map[int]bool)

func findMaximumNumber(k int64, x int) int64 {
	bits = make(map[int]bool)
	for i := 1; i <= 64; i++ {
		if i%x == 0 {
			bits[i] = true
		}
	}
	var INF = int64(1e15)
	l, r := int64(1), INF
	for l <= r {
		var mid = l + (r-l)/2
		if cal(mid) > k {
			r = mid - 1
		} else {
			if mid == INF || cal(mid+1) > k {
				return mid
			}
			l = mid + 1
		}
	}
	return 0
}

var cal = func(t int64) int64 {
	var s = toBinStr(t)
	var l = len(s)
	var mem = make([][]int64, l)
	for i := 0; i < l; i++ {
		mem[i] = make([]int64, l)
		for j := 0; j < l; j++ {
			mem[i][j] = -1
		}
	}
	var dfs func(i, cnt1 int64, limit, isNum bool) int64
	dfs = func(i, cnt1 int64, limit, isNum bool) int64 {
		if i == int64(l) {
			if isNum {
				return cnt1
			}
			return 0
		}
		if !limit && isNum && mem[i][cnt1] > 0 {
			return mem[i][cnt1]
		}
		var up = 1
		if limit {
			up = int(s[i] - '0')
		}
		var res = int64(0)
		for x := 0; x <= up; x++ {
			var add = int64(0)
			if x == 1 && bits[l-int(i)] {
				add = 1
			}
			res += dfs(i+1, cnt1+add, limit && x == up, isNum || x > 0)
		}
		if !limit && isNum {
			mem[i][cnt1] = res
		}
		return res
	}
	var res = dfs(0, 0, true, false)
	return res
}

func toBinStr(x int64) string {
	var bs = make([]byte, 0)
	for x != 0 {
		bs = append(bs, byte('0'+x%2))
		x /= 2
	}
	var n = len(bs)
	for i := 0; i < n/2; i++ {
		bs[i], bs[n-i-1] = bs[n-i-1], bs[i]
	}
	return string(bs)
}

//func main() {
//	k, x := 9, 1
//	var res = findMaximumNumber(int64(k), x)
//	fmt.Println(res)
//	//fmt.Println(cal(6))
//}
