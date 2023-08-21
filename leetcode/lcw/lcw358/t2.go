package lcw358

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
	var bs = []byte{}
	var h = head
	for h != nil {
		bs = append(bs, byte(h.Val+'0'))
		h = h.Next
	}
	var res = calc(string(bs), string(bs))
	var dummy = &ListNode{}
	var cur = dummy
	for _, v := range res {
		cur.Next = &ListNode{int(v - '0'), nil}
		cur = cur.Next
	}
	return dummy.Next
}

func calc(num1 string, num2 string) string {
	n, m := len(num1), len(num2)
	var bs = make([]byte, 0)
	i, j := n-1, m-1
	var c = 0
	for i >= 0 || j >= 0 || c > 0 {
		va, vb := 0, 0
		if i >= 0 {
			va = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			vb = int(num2[j] - '0')
			j--
		}
		var t = va + vb + c
		if t >= 10 {
			c = 1
		} else {
			c = 0
		}
		bs = append(bs, byte('0'+t%10))
	}
	for i := 0; i < len(bs)/2; i++ {
		bs[i], bs[len(bs)-i-1] = bs[len(bs)-i-1], bs[i]
	}
	return string(bs)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
