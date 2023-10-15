package test2102

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"strings"
)

/**
 * @Description https://leetcode.cn/problems/sequentially-ordinal-rank-tracker/
 * idea: 双堆。小顶堆存储topk，即给查询使用的元素；大顶堆存储其它元素;
			小顶堆的元素小于查询次数时，从大顶堆入队
 **/

type SORTracker struct {
	minPQ *priorityqueue.Queue
	maxPQ *priorityqueue.Queue
	cntQ  int
}

func Constructor() SORTracker {
	var minPQ = priorityqueue.NewWith(func(a, b interface{}) int {
		ta, tb := a.(*T), b.(*T)
		if ta.score != tb.score {
			return ta.score - tb.score
		}
		return strings.Compare(tb.name, ta.name)
	})
	var maxPQ = priorityqueue.NewWith(func(a, b interface{}) int {
		ta, tb := a.(*T), b.(*T)
		if ta.score != tb.score {
			return tb.score - ta.score
		}
		return strings.Compare(ta.name, tb.name)
	})
	return SORTracker{minPQ: minPQ, maxPQ: maxPQ, cntQ: 1}
}

func (this *SORTracker) Add(name string, score int) {
	var cur = &T{name, score}
	if this.minPQ.Empty() {
		this.minPQ.Enqueue(cur)
		return
	}
	if this.minTop().less(cur) { // 当前大于最小堆堆顶元素 替换
		t, _ := this.minPQ.Dequeue()
		this.maxPQ.Enqueue(t)
		this.minPQ.Enqueue(cur)
		return
	}
	this.maxPQ.Enqueue(cur)
}

func (this *SORTracker) Get() string {
	if this.cntQ > this.minPQ.Size() {
		t, _ := this.maxPQ.Dequeue()
		this.minPQ.Enqueue(t)
	}
	this.cntQ += 1
	return this.minTop().name
}

func (this *SORTracker) minTop() *T {
	if this.minPQ.Empty() {
		return nil
	}
	t, _ := this.minPQ.Peek()
	return t.(*T)
}

func (this *SORTracker) maxTop() *T {
	if this.maxPQ.Empty() {
		return nil
	}
	t, _ := this.maxPQ.Peek()
	return t.(*T)
}

type T struct {
	name  string
	score int
}

func (this *T) less(t *T) bool {
	ta, tb := this, t
	if ta.score != tb.score {
		return ta.score < tb.score
	}
	return ta.name > tb.name
}
