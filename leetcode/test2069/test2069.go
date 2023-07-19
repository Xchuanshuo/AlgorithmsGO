package test2069

var s = []string{"North", "East", "South", "West"}

type Robot struct {
	dir  [][]int
	cur  int
	move bool
}

func Constructor(width int, height int) Robot {
	var n = 2*width + 2*(height-2)
	var dir = make([][]int, n)
	dir[0] = []int{0, 0, 2}
	for i := 1; i < width; i++ {
		dir[i] = []int{i, 0, 1}
	}
	for i := width; i < width+height-1; i++ {
		dir[i] = []int{width - 1, i - width + 1, 0}
	}
	var t = width + height - 1
	for i := t; i < t+width-1; i++ {
		dir[i] = []int{width - (i - t + 2), height - 1, 3}
	}
	t = t + width - 1
	for i := t; i < n; i++ {
		dir[i] = []int{0, height - (i - t + 2), 2}
	}
	return Robot{dir, 0, false}
}

func (this *Robot) Step(num int) {
	this.move = true
	this.cur = (this.cur + num) % len(this.dir)
}

func (this *Robot) GetPos() []int {
	var d = this.dir[this.cur]
	return []int{d[0], d[1]}
}

func (this *Robot) GetDir() string {
	if !this.move {
		return s[1]
	}
	var d = this.dir[this.cur]
	return s[d[2]]
}
