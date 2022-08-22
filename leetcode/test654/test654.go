package test654

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return dfs(nums)
}

func dfs(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var maxIdx = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
	}
	var root = &TreeNode{Val: nums[maxIdx]}
	root.Left = dfs(nums[0:maxIdx])
	root.Right = dfs(nums[maxIdx+1:])
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
