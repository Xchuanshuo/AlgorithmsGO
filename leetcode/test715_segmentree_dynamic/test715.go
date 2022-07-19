package test715_segmentree_dynamic

type RangeModule struct {
	root *Node
	n    int
}

func Constructor() RangeModule {
	return RangeModule{root: &Node{}, n: 1e9}
}

func (this *RangeModule) AddRange(left int, right int) {
	this.update(this.root, 1, this.n, left, right-1, 1)
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	return this.query(this.root, 1, this.n, left, right-1)
}

func (this *RangeModule) RemoveRange(left int, right int) {
	this.update(this.root, 1, this.n, left, right-1, -1)
}

type Node struct {
	left, right *Node
	cover       bool
	val         int
}

func (this *RangeModule) update(node *Node, start, end, l, r, val int) {
	if start >= l && end <= r {
		node.cover = val == 1
		node.val = val
		return
	}
	mid := (start + end) >> 1
	// 动态开点,下推标记
	this.pushDown(node, mid-start+1, end-mid)
	if l <= mid {
		this.update(node.left, start, mid, l, r, val)
	}
	if r > mid {
		this.update(node.right, mid+1, end, l, r, val)
	}
	this.pushUp(node)
}

func (this *RangeModule) query(node *Node, start, end, l, r int) bool {
	if start >= l && end <= r {
		return node.cover
	}
	mid := (start + end) >> 1
	// 动态开点, 下推标记
	this.pushDown(node, mid-start+1, end-mid)
	var res = true
	if l <= mid {
		res = res && this.query(node.left, start, mid, l, r)
	}
	if r > mid {
		res = res && this.query(node.right, mid+1, end, l, r)
	}
	return res
}

func (this *RangeModule) pushDown(node *Node, leftNum, rightNum int) {
	if node.left == nil {
		node.left = &Node{}
	}
	if node.right == nil {
		node.right = &Node{}
	}
	// 没有标记 无需再覆盖
	if node.val == 0 {
		return
	}
	// 1为覆盖, -1为取消覆盖
	node.left.cover = node.val == 1
	node.right.cover = node.val == 1
	node.left.val = node.val
	node.right.val = node.val
	// 取消标记 防止重复下推
	node.val = 0
}

func (this *RangeModule) pushUp(node *Node) {
	node.cover = node.left.cover && node.right.cover
}
