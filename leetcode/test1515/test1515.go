package test1515

import "math"

/**
 * @Description https://leetcode.cn/problems/best-position-for-a-service-centre
 * idea: 随机算法，爬山法。凸函数求局部最优解即为全局最优。
		1.最优解从所有x、y坐标的中心点出发更容易找到
		2.对于当前fx,fy 尝试往4个不同方向走，若得到更小距离则继续走。得不到更小距离说明当前的步长过大，尝试缩短步长，
		  重复前面的过程，由于误差范围需要再1e-5内，所以ste下界p从更小开始，步长设定下边界为1e-7
 **/

var dirs = [][]float64{{1.0, 0.0}, {0.0, 1.0}, {-1.0, 0.0}, {0.0, -1.0}}

func getMinDistSum(positions [][]int) float64 {
	var n = len(positions)
	var calDis = func(x, y float64) float64 {
		var res = 0.0
		for _, p := range positions {
			res += math.Sqrt((x-float64(p[0]))*(x-float64(p[0])) + (y-float64(p[1]))*(y-float64(p[1])))
		}
		return res
	}
	fx, fy := 0.0, 0.0
	for _, p := range positions {
		fx += float64(p[0])
		fy += float64(p[1])
	}
	fx, fy = fx/float64(n), fy/float64(n)
	var ep = 1e-7
	var step = 1.0
	for step > ep {
		var found = false
		for _, d := range dirs {
			nx, ny := fx+step*d[0], fy+step*d[1]
			if calDis(nx, ny) < calDis(fx, fy) {
				fx, fy = nx, ny
				found = true
			}
		}
		if !found { // 没找到更优解，说明step过大
			step *= 0.5
		}
	}
	return calDis(fx, fy)
}
