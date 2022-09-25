package test315

var offset = 10001

func countSmaller(nums []int) []int {
	var n = len(nums)
	var tree = TreeArr{arr: make([]int, offset*2+10)}
	for _, num := range nums {
		tree.update(num+offset, 1)
	}
	var res = make([]int, n)
	for i, num := range nums {
		tree.update(offset+num, -1)
		res[i] = tree.query(offset + num - 1)
	}
	return res
}

type TreeArr struct {
	arr []int
}

func (t TreeArr) query(i int) int {
	var res = 0
	for i > 0 {
		res += t.arr[i]
		i = i - t.lowbit(i)
	}
	return res
}

func (t TreeArr) update(i int, val int) {
	var n = len(t.arr)
	for i <= n {
		t.arr[i] = t.arr[i] + val
		i = i + t.lowbit(i)
	}
}

// x最低位1的位置
func (t TreeArr) lowbit(x int) int {
	return x & -x
}
