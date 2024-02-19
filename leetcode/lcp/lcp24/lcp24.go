package lcp24

import (
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

/**
 * @Description https://leetcode.cn/problems/5TxKeK
 * idea: 双堆，问题转换为【求数据流中的中位数】, 因为要前[0,i]的元素转换为 nums[a]+1 == nums[a+1]的形式
		 现在有一组数 v0、v1、v2、v3，假设将v0转换为x，则最终元素为x、x+1、x+2、x+3, 则变化量为
		 |v0-x| + |v1-(x+1)| + |v2-(x+2)| + |v3-(x+3)|，设ai = vi - i，则该式子可以写为
		 |a0-x| + |a1-x| + |a2-x| + |a3-x|，最终变成了求将ai转换为固定数x的最小操作数
		1.数x应该为多少? 实际将一组数变为该组数的【中位数】，变化量最小
		2.双堆分别保存一半的数即可
 **/

var M = int(1e9) + 7

func numsGame(nums []int) []int {
	var n = len(nums)
	var res = make([]int, n)
	var finder = NewMedianFinder()
	for i, v := range nums {
		var a = v - i
		finder.addNum(a)
		res[i] = finder.cal()
	}
	return res
}

type MedianFinder struct {
	small, large *priorityqueue.Queue
	sumL, sumR   int
}

func NewMedianFinder() MedianFinder {
	var small = priorityqueue.NewWith(func(a, b interface{}) int {
		av, bv := a.(int), b.(int)
		return bv - av
	}) // 大顶堆
	var large = priorityqueue.NewWith(utils.IntComparator) // 小顶堆
	return MedianFinder{small, large, 0, 0}
}

func (this *MedianFinder) addNum(num int) {
	if this.small.Size() < this.large.Size() {
		this.large.Enqueue(num)
		tv, _ := this.large.Dequeue()
		v := tv.(int)
		this.small.Enqueue(v)
		this.sumR += num - v
		this.sumL += v
	} else {
		this.small.Enqueue(num)
		tv, _ := this.small.Dequeue()
		v := tv.(int)
		this.large.Enqueue(v)
		this.sumL += num - v
		this.sumR += v
	}
}

func (this *MedianFinder) findMedian() int {
	if this.small.Size() < this.large.Size() {
		tv, _ := this.large.Peek()
		v := tv.(int)
		return v
	} else if this.small.Size() > this.large.Size() {
		tv, _ := this.small.Peek()
		v := tv.(int)
		return v
	}
	tv1, _ := this.small.Peek()
	tv2, _ := this.large.Peek()
	v1, v2 := tv1.(int), tv2.(int)
	return (v1 + v2) / 2
}

func (this *MedianFinder) cal() int {
	var middle = this.findMedian()
	cntL, cntR := this.small.Size(), this.large.Size()
	return (cntL*middle - this.sumL + this.sumR - cntR*middle) % M
}
