package main

func sumOfEncryptedInt(nums []int) int {
	var res = 0
	for _, v := range nums {
		var c = 0
		var t = v
		var cnt = 0
		for t != 0 {
			c = max(c, t%10)
			t /= 10
			cnt++
		}
		var val = 0
		for i := 0; i < cnt; i++ {
			val = val*10 + c
		}
		res += val
	}
	return res
}
