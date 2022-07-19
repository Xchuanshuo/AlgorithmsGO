package test2013

type DetectSquares struct {
	mp map[int]map[int]int
}

func Constructor() DetectSquares {
	return DetectSquares{mp: make(map[int]map[int]int)}
}


func (this *DetectSquares) Add(point []int)  {
	x, y := point[0], point[1]
	if _, exist := this.mp[x];!exist {
		this.mp[x] = make(map[int]int)
	}
	this.mp[x][y]++
}

func (this *DetectSquares) Count(point []int) int {
	res, x, y := 0, point[0], point[1]
	if _, exist := this.mp[x];!exist {
		return res
	}
	for nx, _ := range this.mp {
		if nx != x {
			d := x - nx
			res += this.mp[x-d][y] * this.mp[x][y+d] * this.mp[x-d][y+d]
			res += this.mp[x-d][y] * this.mp[x][y-d] * this.mp[x-d][y-d]
		}
	}
	return res
}