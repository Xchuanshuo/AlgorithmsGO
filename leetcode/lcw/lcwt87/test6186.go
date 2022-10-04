package main

import "fmt"

/**
 * @Description
 * idea: 使用数字 存储对指定二进制位有贡献的最后位置
 **/

func smallestSubarrays(nums []int) []int {
	var n = len(nums)
	var f = make([]int, 32)
	var res = make([]int, n)
	for i := n - 1; i >= 0; i-- {
		var num = nums[i]
		var mx = i
		for j := 0; j < 32; j++ {
			if (num & (1 << j)) > 0 {
				f[j] = i
			}
			mx = max(mx, f[j])
		}
		res[i] = mx - i + 1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//输入：
//[9,5,0,10,7,2,9,2,4]
//输出：
//[4,4,4,4,4,4,3,2,1]

//[4,3,3,2,3,4,3,2,1]

func main() {
	var arr = []int{0}
	var res = smallestSubarrays(arr)
	fmt.Println(res)
}
