package lcw366

var M = int(1e9) + 7

func maxSum(nums []int, k int) int {
	var cnt = make([]int, 32)
	for _, v := range nums {
		for i := 0; i < 32; i++ {
			if v&(1<<i) != 0 {
				cnt[i]++
			}
		}
	}
	var res = 0
	for i := 0; i < k; i++ {
		var x = 0
		for j := 31; j >= 0; j-- {
			if cnt[j] > 0 {
				x |= 1 << j
				cnt[j]--
			}
		}
		res = (res + x*x%M) % M
	}
	return res
}
