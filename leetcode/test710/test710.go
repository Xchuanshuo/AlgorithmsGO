package test710

import "math/rand"

type Solution struct {
	size int
	mp   map[int]int
}

func Constructor(n int, blacklist []int) Solution {
	s := Solution{size: n - len(blacklist)}
	var mp = make(map[int]int)
	// 将在黑名单中的数与不在黑名单的数做一个映射
	for _, b := range blacklist {
		mp[b] = 0
	}
	var last = n - 1
	for _, b := range blacklist {
		if b >= s.size {
			continue
		}
		for {
			if _, exist := mp[last]; !exist {
				break
			}
			last--
		}
		mp[b] = last
		last--
	}
	s.mp = mp
	return s
}

func (this *Solution) Pick() int {
	var pos = rand.Intn(this.size)
	// 在黑名单中则返回映射的数, 不在则直接返回原位置的数
	if v, exist := this.mp[pos]; exist {
		return v
	}
	return pos
}
