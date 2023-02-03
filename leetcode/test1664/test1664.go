package test1664

// s1, s2分别为全部奇数和、全部偶数和; l1, l2 分别左边奇数和、左边偶数和
// 1.枚举每一个数v，删除后，其后坐标的奇偶性发生变化
// 2.若当前下标为奇数 odd=l1+s2-l2, even=l2+s1-l1-v
//   若当前下标为偶数,odd=l1+s2-l2-v,  even=l2+s1-l1
func waysToMakeFair(nums []int) int {
	var n = len(nums)
	s1, s2 := 0, 0
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			s1 += nums[i]
		} else {
			s2 += nums[i]
		}
	}
	res, l1, l2 := 0, 0, 0
	for i, v := range nums {
		odd, even := 0, 0
		if i%2 == 1 {
			odd = l1 + s2 - l2
			even = l2 + s1 - l1 - v
			l1 += v
		} else {
			odd = l1 + s2 - l2 - v
			even = l2 + s1 - l1
			l2 += v
		}
		if odd == even {
			res++
		}
	}
	return res
}
