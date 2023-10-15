package test1993

type LockingTree struct {
	locks  map[int]int
	g      map[int][]int
	parent []int
}

func Constructor(parent []int) LockingTree {
	var g = make(map[int][]int)
	for i, p := range parent {
		g[p] = append(g[p], i)
	}
	return LockingTree{map[int]int{}, g, parent}
}

func (this *LockingTree) Lock(num int, user int) bool {
	if this.locks[num] > 0 {
		return false
	}
	this.locks[num] = user
	return true
}

func (this *LockingTree) Unlock(num int, user int) bool {
	if this.locks[num] != user {
		return false
	}
	this.locks[num] = 0
	return true
}

func (this *LockingTree) Upgrade(num int, user int) bool {
	// 1.指定节点没有任何上锁的祖先节点
	if this.hasLockedParent(num) {
		return false
	}
	// 2.指定节点至少有一个上锁状态的子孙节点 且全部解锁
	if !this.hasLockedChild(num) {
		return false
	}
	this.locks[num] = user
	return true
}

func (this *LockingTree) hasLockedChild(node int) bool {
	var valid = false
	var dfs func(cur int)
	dfs = func(cur int) {
		for _, nxt := range this.g[cur] {
			if this.locks[nxt] > 0 {
				this.locks[nxt] = 0
				valid = true
			}
			dfs(nxt)
		}
	}
	dfs(node)
	return valid
}

func (this *LockingTree) hasLockedParent(node int) bool {
	var parent = this.parent
	for node != -1 {
		if this.locks[node] > 0 {
			return true
		}
		node = parent[node]
	}
	return false
}
