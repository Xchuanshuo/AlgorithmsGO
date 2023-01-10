package lcwt93

/**
 * @Description https://leetcode.cn/problems/minimum-total-cost-to-make-arrays-unequal/
 * idea: 分情况谈论, 统计所有值相等的数对个数swapCnt，
		1.若众数 <= swapCnt/2, 且数对为偶数时 所有数对可以两两交换，总代价为所有【相等】位置的索引和
					数对为奇数时，至少有【三个】不同的数 将其中一个数与位置0交换即可，总代价同偶数对
		2.若众数 > swapCnt/2, 则数对之间没法直接完成交换，需要与原先【不相等】的位置交换, 此时不相等
		  的数对相当于实际被计算在了相等的里面，保证结果最小，需要从左至右计算，看最后是否
		  满足众数小于总数对的一半即可
 **/

func minimumTotalCost(nums1 []int, nums2 []int) int64 {
	mode, modeCnt, swapCnt, res := 0, 0, 0, int64(0)
	var cnts = make(map[int]int)
	var n = len(nums1)
	for i := 0; i < n; i++ {
		x, y := nums1[i], nums2[i]
		if x != y {
			continue
		}
		swapCnt++
		res += int64(i)
		cnts[x]++
		if cnts[x] > modeCnt {
			mode, modeCnt = x, cnts[x]
		}
	}
	for i := 0; i < n; i++ {
		if 2*modeCnt <= swapCnt {
			break
		}
		x, y := nums1[i], nums2[i]
		if x != y && x != mode && y != mode {
			res += int64(i)
			swapCnt++
		}
	}
	if 2*modeCnt > swapCnt {
		return -1
	}
	return res
}
