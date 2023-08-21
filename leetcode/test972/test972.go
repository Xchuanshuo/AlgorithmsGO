package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/**
 * @Description https://leetcode.cn/problems/equal-rational-numbers/
 * idea: 小数化成最简分数后比较即可 转换方式 x.d(s), 整数部分(x/1) + 非循环节部分(d/len(d)) + 循环节部分(s/(len(d)*(len(s)-1)))
 **/

func isRationalEqual(s string, t string) bool {
	var convert = func(s string) Fraction {
		var strs = strings.Split(s, ".")
		v0, _ := strconv.Atoi(strs[0])
		var f = NewFrac(v0, 1)
		if len(strs) == 1 {
			return f
		}
		var strs1 = strings.Split(strs[1], "(")
		var sz1 = 0
		if strs1[0] != "" {
			v1, _ := strconv.Atoi(strs1[0]) // 小数部分
			sz1 = len(strs1[0])
			f.add(NewFrac(v1, int(math.Pow10(sz1))))
		}
		if len(strs1) > 1 && strs1[1] != "" {
			sz2 := len(strs1[1]) - 1
			v2, _ := strconv.Atoi(strs1[1][0:sz2])
			f.add(NewFrac(v2, int(math.Pow10(sz1)*(math.Pow10(sz2)-1))))
		}
		return f
	}
	var f1 = convert(s)
	var f2 = convert(t)
	return f1.m == f2.m && f1.d == f2.d
}

type Fraction struct {
	m, d int // m-分子 d-分母
}

func NewFrac(m, d int) Fraction {
	var c = gcd(m, d)
	return Fraction{m / c, d / c}
}

func (this *Fraction) add(other Fraction) {
	m1, d1 := this.m, this.d
	m2, d2 := other.m, other.d
	var lcm = d1 * d2 / gcd(d1, d2)
	nm, nd := lcm/d1*m1+lcm/d2*m2, lcm
	var t = NewFrac(nm, nd)
	this.m, this.d = t.m, t.d
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var s = "0.(52)"
	var t = "0.5(25)"
	fmt.Println(isRationalEqual(s, t))
}
