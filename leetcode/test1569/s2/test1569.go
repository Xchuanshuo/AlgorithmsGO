package s1

/**
 * @Description https://leetcode.cn/problems/number-of-ways-to-reorder-array-to-get-same-bst/
 * idea: 乘法逆元 + 组合数 + 分治.
		1.二分查找树的形状与插入元素的相对位置有关。对于一个元素数组，第一个数为根节点，
		比根节点小的元素【只】会在左子树，比根大的【只】会在右子树。

		2.设比根nums[0]小的元素有n个，大的有m个。为保持相对位置不变，在n+m个元素任选n个空位插入比根小的元素，
		其它插入m个元素，则不考虑左右两边内部的排列情况，以nums[0]为根的方案数总共有C(n+m, n)个；

		3.对于左半边和右半边的元素实际是一个子问题，用上述同样的方式算即可，以nums[0]为根总的方案数则为
         C(n+m,n) * numOfWays(smaller) * numOfWays(larger),

	 注意: 本题数据范围为1000，可以使用杨辉三角求组合；也可以根据组合数公式，使用乘法逆元计算，能支持1e5的数据范围
 **/

var M = int(1e9) + 7

func numOfWays(nums []int) int {
	var n = len(nums)
	var cb = NewCombine(n, M)
	cb.init()
	var dfs func(nums []int) int
	dfs = func(nums []int) int {
		if len(nums) <= 1 {
			return 1
		}
		var smaller = make([]int, 0)
		var larger = make([]int, 0)
		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[0] {
				smaller = append(smaller, nums[i])
			} else {
				larger = append(larger, nums[i])
			}
		}
		n, m := len(smaller), len(larger)
		return cb.C(n+m, n) * dfs(smaller) % M * dfs(larger) % M
	}
	return (dfs(nums) - 1 + M) % M
}

type Combine struct {
	n int // 求C(n,1..n)组合数
	M int
	c [][]int
}

func NewCombine(n, mod int) Combine {
	var c = make([][]int, n)
	for i := 0; i < n; i++ {
		c[i] = make([]int, n)
	}
	return Combine{n, mod, c}
}

func (ins Combine) C(n, m int) int {
	return ins.c[n][m]
}

func (ins Combine) init() {
	var c = ins.c
	n, M := ins.n, ins.M
	c[0][0] = 1
	for i := 1; i < n; i++ {
		c[i][0] = 1
		for j := 1; j < n; j++ {
			c[i][j] = (c[i-1][j-1] + c[i-1][j]) % M
		}
	}
}
