package main

/**
 * @Description https://leetcode.cn/problems/coin-bonus/
 * idea: 树状数组，使用两个差分数组 支持区间修改，与区间查询
 **/

func bonus1(n int, leadership [][]int, operations [][]int) []int {
	var g = make(map[int][]int)
	for _, ld := range leadership {
		g[ld[0]] = append(g[ld[0]], ld[1])
	}
	L, R := make(map[int]int), make(map[int]int)
	var pos = 0
	var dfs func(int)
	dfs = func(id int) {
		pos++
		L[id] = pos
		for _, s := range g[id] {
			dfs(s)
		}
		R[id] = pos
	}
	dfs(1)
	tree := NewTreeArr(n + 1)
	var res = make([]int, 0)
	for _, op := range operations {
		if op[0] == 1 {
			id, cnt := op[1], op[2]
			tree.UpdateRange(L[id], L[id], cnt)
		} else if op[0] == 2 {
			id, cnt := op[1], op[2]
			tree.UpdateRange(L[id], R[id], cnt)
		} else {
			res = append(res, tree.QueryRange(L[op[1]], R[op[1]])%M)
		}
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
