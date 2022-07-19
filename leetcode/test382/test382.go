package test382

import "math/rand"

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head: head}
}


func (this *Solution) GetRandom() int {
	p := this.head
	idx, val := 0, p.Val
	for p != nil {
		r := rand.Intn(idx+1)
		if r == 0 {
			val = p.Val
		}
		idx++
		p = p.Next
	}
	return val
}




type ListNode struct {
	Val int
	Next *ListNode
}
