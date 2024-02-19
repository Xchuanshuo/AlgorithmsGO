package lcwt124

func maxOperations1(nums []int) int {
	var score = nums[0] + nums[1]
	var cnt = 1
	for i := 2; i < len(nums); i += 2 {
		if i+1 >= len(nums) {
			break
		}
		var cur = nums[i] + nums[i+1]
		if cur != score {
			break
		}
		cnt++
	}
	return cnt
}
