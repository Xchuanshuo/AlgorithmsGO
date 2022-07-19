package test729

import (
	"github.com/emirpasic/gods/trees/redblacktree"
	"math"
)

type MyCalendar struct {
	*redblacktree.Tree
}

func Constructor() MyCalendar {
	t := redblacktree.NewWithIntComparator()
	t.Put(math.MaxInt32, nil) // 哨兵，简化代码
	return MyCalendar{t}
}

func (c MyCalendar) Book(start, end int) bool {
	node, _ := c.Ceiling(end)
	it := c.IteratorAt(node)
	if !it.Prev() || it.Value().(int) <= start {
		c.Put(start, end)
		return true
	}
	return false
}
