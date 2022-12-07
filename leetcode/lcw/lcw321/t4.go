package lcw321

/**
 * @Description https://leetcode.cn/contest/weekly-contest-321/problems/count-subarrays-with-median-k/
 * idea: 记录大与k的元素与小于k的元素个数的差值，以k为界限
 		 若[idx(k),n]的元素还出现了同样的差值，说明
		 前后两个位置的中间【大于和小于k】的个数刚好【抵消】，记录答案
 **/

func countSubarrays(nums []int, k int) int {
	var mp = make(map[int]int)
	mp[0] = 1
	res, cnt := 0, 0
	var flag = false
	for _, v := range nums {
		if v > k {
			cnt++
		} else if v < k {
			cnt--
		} else if v == k {
			flag = true
		}
		if flag {
			res += mp[cnt]
			res += mp[cnt-1]
		} else {
			mp[cnt]++
		}
	}
	return res
}

//输入：
//[3,2,1,4,5]
//4
//输出：
//4
//预期：
//3
