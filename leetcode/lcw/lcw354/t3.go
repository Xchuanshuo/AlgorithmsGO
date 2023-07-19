package lcw354

func minimumIndex(nums []int) int {
	var n = len(nums)
	x, cnt := 0, 0
	for _, v := range nums {
		if v == x {
			cnt++
		} else {
			if cnt > 0 {
				cnt--
			} else {
				cnt = 1
				x = v
			}
		}
	}
	var cnts = make([]int, n+1)
	for i, v := range nums {
		cnts[i+1] = cnts[i]
		if v == x {
			cnts[i+1] += 1
		}
	}
	for i := 1; i <= n; i++ {
		if cnts[i]*2 > i && (cnts[n]-cnts[i])*2 > n-i {
			return i - 1
		}
	}
	return -1
}
