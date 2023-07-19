package lcw353

/**
 * @Description https://leetcode.cn/problems/apply-operations-to-make-all-array-elements-equal-to-zero/
 * idea: 差分数组. 从左到右对长度为k的区间进行操作，对区间所有元素减去当前剩余值，
		 当前剩余值为前面的差分和
 **/

func checkArray(nums []int, k int) bool {
	var n = len(nums)
	var dif = make([]int, n+1)
	for i := 1; i < n; i++ {
		dif[i] = nums[i] - nums[i-1]
	}
	var sum = nums[0]
	for i := 0; i < n; i++ {
		// 累计前缀和
		var cur = sum + dif[i]
		if cur == 0 {
			continue
		}
		// 当前已经小于0，或区间小于k无法操作
		if cur < 0 || i+k > n {
			return false
		}
		dif[i] -= cur
		dif[i+k] += cur
		sum += dif[i]
	}
	return true
}
