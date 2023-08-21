package test1348

/**
 * @Description https://leetcode.cn/problems/tweet-counts-per-frequency/
 * idea: 动态开点线段树
 **/

type TweetCounts struct {
	mp map[string]*SegmentTree
}

func Constructor() TweetCounts {
	return TweetCounts{map[string]*SegmentTree{}}
}

func (this *TweetCounts) RecordTweet(tweetName string, time int) {
	if _, exist := this.mp[tweetName]; !exist {
		this.mp[tweetName] = NewSegmentTree(int(1e9)+1, func(lv, rv int) int {
			return lv + rv
		})
	}
	this.mp[tweetName].update(time, time, 1)
}

func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	var sz = getL(freq)
	var res = make([]int, 0)
	for i := startTime; i <= endTime; i += sz {
		var cnt = 0
		l, r := i, min(endTime, i+sz-1)
		if tree, exist := this.mp[tweetName]; exist {
			cnt = tree.query(l, r)
		}
		res = append(res, cnt)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getL(f string) int {
	if f == "minute" {
		return 60
	}
	if f == "hour" {
		return 3600
	}
	return 86400
}

func NewSegmentTree(n int, f func(lv, rv int) int) *SegmentTree {
	return &SegmentTree{&Node{}, n, f}
}

type SegmentTree struct {
	root *Node
	n    int
	f    func(lv, rv int) int
}

// Node 动态开点线段树节点信息
type Node struct {
	left, right *Node
	val         int
	lazy        int
}

func (s *SegmentTree) query(curL, curR int) int {
	return s._query(s.root, 0, s.n, curL, curR)
}

func (s *SegmentTree) update(curL, curR int, val int) {
	s._update(s.root, 0, s.n, curL, curR, val)
}

func (s *SegmentTree) _query(node *Node, l, r, curL, curR int) int {
	if r < curL || l > curR || node == nil {
		return 0
	}
	if l >= curL && r <= curR {
		return node.val
	}
	s.pushDown(node)
	var mid = l + (r-l)/2
	var left = s._query(node.left, l, mid, curL, curR)
	var right = s._query(node.right, mid+1, r, curL, curR)
	return s.f(left, right)
}

// 在区间[l,r]内更新[curL,curR]
func (s *SegmentTree) _update(node *Node, l, r, curL, curR int, val int) {
	if r < curL || l > curR {
		return
	}
	if l >= curL && r <= curR {
		s.do(node, val)
		return
	}
	s.pushDown(node)
	var mid = l + (r-l)/2
	s._update(node.left, l, mid, curL, curR, val)
	s._update(node.right, mid+1, r, curL, curR, val)
	node.val = s.f(node.left.val, node.right.val)
}

func (s *SegmentTree) pushDown(node *Node) {
	if node.left == nil {
		node.left = &Node{}
	}
	if node.right == nil {
		node.right = &Node{}
	}
	if node.lazy == 0 {
		return
	}
	s.do(node.left, node.lazy)
	s.do(node.right, node.lazy)
	node.lazy = 0
}

// 处理lazy逻辑
func (s *SegmentTree) do(node *Node, val int) {
	node.lazy += val
	node.val += val
}
