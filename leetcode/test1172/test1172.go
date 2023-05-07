package test1172

import "github.com/emirpasic/gods/trees/redblacktree"

/**
 * @Description https://leetcode.cn/problems/dinner-plate-stacks/
 * idea: 1.每个元素 -> 栈
         2.push, 未满 从左至右
         3.pop, 非空, 从右往左第一个
	     4.popAt，随机访问元素栈
		记录所有栈stacks和所有未满栈下标,
		push->取最小下标, 为空则新创建一个栈
		popAt->下标大于所有栈大小则返回，为最后一个则一次移除为空的位置，其它情况直接移除栈顶即可
		pop-> popAt(len(stacks)-1), 即最后一个不为空的位置的栈顶元素
 **/

type DinnerPlates struct {
	cpc     int
	notFull *redblacktree.Tree
	stacks  [][]int
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{capacity, redblacktree.NewWithIntComparator(), make([][]int, 0)}
}

func (this *DinnerPlates) Push(val int) {
	if this.notFull.Empty() {
		this.stacks = append(this.stacks, []int{val})
		if this.cpc > 1 {
			this.notFull.Put(len(this.stacks)-1, nil)
		}
	} else {
		idx, _ := this.notFull.Left().Key.(int)
		this.stacks[idx] = append(this.stacks[idx], val)
		if len(this.stacks[idx]) == this.cpc {
			this.notFull.Remove(idx)
		}
	}
}

func (this *DinnerPlates) Pop() int {
	return this.PopAtStack(len(this.stacks) - 1)
}

func (this *DinnerPlates) PopAtStack(index int) int {
	if index < 0 || index >= len(this.stacks) || len(this.stacks[index]) == 0 {
		return -1
	}
	val := this.stacks[index][len(this.stacks[index])-1]
	this.stacks[index] = this.stacks[index][:len(this.stacks[index])-1]
	if len(this.stacks[index]) > 0 || index != len(this.stacks)-1 {
		this.notFull.Put(index, nil)
	} else {
		for len(this.stacks) > 0 && len(this.stacks[len(this.stacks)-1]) == 0 {
			this.stacks = this.stacks[:len(this.stacks)-1]
			this.notFull.Remove(len(this.stacks))
		}
	}
	return val
}
