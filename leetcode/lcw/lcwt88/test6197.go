package main

type LUPrefix struct {
	cur int
	q   []int
	mp  map[int]bool
}

func Constructor(n int) LUPrefix {
	var q = make([]int, 0)
	for i := 1; i <= n; i++ {
		q = append(q, i)
	}
	return LUPrefix{0, q, make(map[int]bool)}
}

func (this *LUPrefix) Upload(video int) {
	this.mp[video] = true
}

func (this *LUPrefix) Longest() int {
	for len(this.q) > 0 {
		var cur = this.q[0]
		if !this.mp[cur] {
			break
		}
		this.cur = cur
		this.q = this.q[1:]
	}
	return this.cur
}

/**
 * Your LUPrefix object will be instantiated and called as such:
 * obj := Constructor(n);
 * obj.Upload(video);
 * param_2 := obj.Longest();
 */
