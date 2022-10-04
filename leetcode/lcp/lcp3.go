package main

import (
	"sort"
)

/**
 * @Description https://leetcode.cn/problems/programmable-robot/
 * idea: 1.写一个函数判断能否走到目标点tx,ty, 防止tx,ty过大，直接跳过重复步骤
 *       2.a.不考虑障碍情况 判断能否走到目标点，不能直接返回结果
 *         b.将每个障碍当做目标点，能走到则返回结果
 *		   不经过障碍 切能到终点 即表示最终能走到终点
 **/

func robot(command string, obstacles [][]int, tx int, ty int) bool {
	var isPathTo func(tx, ty int) bool
	isPathTo = func(tx, ty int) bool {
		x, y := 0, 0
		for true {
			for i := 0; i < len(command); i++ {
				if command[i] == 'U' {
					y += 1
				} else {
					x += 1
				}
				if x == tx && y == ty {
					return true
				}
				if x > tx || y > ty {
					return false
				}
			}
			if tx > x && ty > y {
				x, y = (tx-1)/x*x, (tx-1)/x*y
			}
			if x == tx && y == ty {
				return true
			}
		}
		return false
	}
	if !isPathTo(tx, ty) {
		return false
	}
	sort.Slice(obstacles, func(i, j int) bool {
		a, b := obstacles[i], obstacles[j]
		if a[0] == b[0] {
			return a[1] < b[1]
		}
		return a[0] < b[0]
	})
	for _, obs := range obstacles {
		if isPathTo(obs[0], obs[1]) && (obs[0] <= tx && obs[1] <= ty) {
			return false
		}
	}
	return true
}

//func main() {
//	// (0,1) (1,1) (2,1) (2,2) (3,2) (4,2) (5,2)
//	var command = "URRURRR"
//	var res = robot(command, [][]int{{7, 7}}, 4915, 1966)
//	fmt.Println(res)
//}
