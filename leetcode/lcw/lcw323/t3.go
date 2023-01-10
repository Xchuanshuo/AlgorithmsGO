package lcw323

type Allocator struct {
	n       int
	tree    *TreeArr
	updates map[int][][]int
}

func Constructor(n int) Allocator {
	return Allocator{n: n, tree: NewTreeArr(n + 20), updates: map[int][][]int{}}
}

func (this *Allocator) Allocate(size int, mID int) int {
	var tIdx = -1
	for i := 1; i <= this.n-size+1; i++ {
		if this.tree.QueryRange(i, i+size-1) == 0 {
			tIdx = i
			break
		}
	}
	if tIdx == -1 {
		return -1
	}
	this.tree.UpdateRange(tIdx, tIdx+size-1, mID)
	this.updates[mID] = append(this.updates[mID], []int{tIdx, tIdx + size - 1})
	return tIdx - 1
}

func (this *Allocator) Free(mID int) int {
	l, res := len(this.updates[mID]), 0
	for i := 0; i < l; i++ {
		var u = this.updates[mID][0]
		this.tree.UpdateRange(u[0], u[1], -mID)
		res += u[1] - u[0] + 1
		this.updates[mID] = this.updates[mID][1:]
	}
	return res
}

type TreeArr struct {
	dif1, dif2 []int
}

func NewTreeArr(n int) *TreeArr {
	return &TreeArr{dif1: make([]int, n), dif2: make([]int, n)}
}
func (t *TreeArr) QueryRange(l, r int) int {
	return t.query(r) - t.query(l-1)
}
func (t *TreeArr) UpdateRange(l, r int, val int) {
	t.update(l, val)
	t.update(r+1, -val)
}
func (t *TreeArr) query(i int) int {
	var x = i
	var res = 0
	for i > 0 {
		res += (x+1)*t.dif1[i] - t.dif2[i]
		i -= t.lowbit(i)
	}
	return res
}
func (t *TreeArr) update(i int, val int) {
	x, n := i, len(t.dif1)
	for i < n {
		t.dif1[i] += val
		t.dif2[i] += x * val
		i += t.lowbit(i)
	}
}
func (t *TreeArr) lowbit(x int) int {
	return x & -x
}
