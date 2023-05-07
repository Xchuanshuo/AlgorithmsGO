package lcw337

/**
 * @Description https://leetcode.cn/problems/the-number-of-beautiful-subsets/
 * idea: 二进制枚举 对于不合法状态提前进行存储 对于所有组合能快速判断是否合法
 **/
func beautifulSubsets(nums []int, k int) int {
	var n = len(nums)
	var mp = make(map[int]bool)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]-nums[j] == k || nums[j]-nums[i] == k {
				mp[1<<i|1<<j] = true
			}
		}
	}
	var res = 0
	var r = 1 << n
	for s := 1; s < r; s++ {
		var valid = true
		for v := range mp {
			if v&s == v {
				valid = false
				break
			}
		}
		if valid {
			res++
		}
	}
	return res
}

func beautifulSubsets1(nums []int, k int) int {
	var res = -1
	var dfs func(i int)
	var cnt = make(map[int]int)
	dfs = func(i int) {
		if i == len(nums) {
			res++
			return
		}
		dfs(i + 1)
		if cnt[nums[i]-k] == 0 && cnt[nums[i]+k] == 0 {
			cnt[nums[i]]++
			dfs(i + 1)
			cnt[nums[i]]--
		}
	}
	dfs(0)
	return res
}
