package test1835

/**
 * @Description https://leetcode.cn/problems/find-xor-sum-of-all-pairs-bitwise-and
 * idea: 按位与或运算满足分配率 即 (a&b)^(a&c) = a&(b^c), 化简后答案为
		(a1 & xor(arr2)) ^ (a2 & xor(arr2)) ^ ... ^ (an & xor(arr2))
 **/

func getXORSum(arr1 []int, arr2 []int) int {
	var xv = 0
	for _, v := range arr2 {
		xv ^= v
	}
	var res = 0
	for _, v := range arr1 {
		res ^= v & xv
	}
	return res
}
