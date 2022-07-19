package test1305

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	arr1, arr2 := dfs(root1), dfs(root2)
	n1, n2 := len(arr1), len(arr2)
	var res = make([]int, n1+n2)
	i, j := 0, 0
	for i < n1 && j < n2 {
		if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}
	if i < n1 {
		res = append(res, arr1[i:n1]...)
	}
	if j < n2 {
		res = append(res, arr2[j:n2]...)
	}
	return res
}

func dfs(root *TreeNode) (arr []int) {
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		arr = append(arr, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
