package test732

import "github.com/emirpasic/gods/trees/redblacktree"

type MyCalendarThree struct {
	mp *redblacktree.Tree
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		mp: redblacktree.NewWithIntComparator(),
	}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	if v, ok := this.mp.Get(start); ok {
		this.mp.Put(start, v.(int)+1)
	} else {
		this.mp.Put(start, 1)
	}
	if v, ok := this.mp.Get(end); ok {
		this.mp.Put(end, v.(int)-1)
	} else {
		this.mp.Put(end, -1)
	}
	cur, max := 0, 0
	for _, v := range this.mp.Values() {
		cur += v.(int)
		if cur > max {
			max = cur
		}
	}
	return max
}
