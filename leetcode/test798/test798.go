package test798

func bestRotation(nums []int) int {
	var n = len(nums)
	var dif = make([]int, n+1)
	for i := 0; i < n; i++ {
		l, r := (i+1)%n, (i-nums[i]+n)%n
		if l <= r {
			add(dif, l, r)
		} else {
			add(dif, 0, r)
			add(dif, l, n-1)
		}
	}
	res, cur, sum := 0, 0, 0
	for i := 0; i <= n; i++ {
		sum += dif[i]
		if sum > cur {
			cur = sum
			res = i
		}
	}
	return res
}

func add(dif []int, l, r int) {
	dif[l]++
	dif[r+1]--
}
