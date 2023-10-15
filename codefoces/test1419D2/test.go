package main

/**
 * @Description https://codeforces.com/problemset/problem/1419/D2
 * idea: 构造 + 贪心 + 二分
		1.贪心，优先使用较小的元素往较大的元素中间插空
		2.可插入元素数量切分点满足单调性所以可以使用二分查找
		边界: 总元素数量 <= 2, 或得重复元素过多直接返回原数组即可
 **/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	run(os.Stdin, os.Stdout)
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	cases := 1
	fmt.Fscanln(in, &cases)
	var a = make([]int, 0)
	for curCase := 1; curCase <= cases; curCase++ {
		var v int
		fmt.Fscanf(in, "%d ", &v)
		a = append(a, v)
	}
	cnt, res := cal(a)
	fmt.Fprintln(out, cnt)
	for _, v := range res {
		fmt.Fprint(out, v)
		fmt.Fprint(out, " ")
	}
	fmt.Fprintln(out)
	out.Flush()
}

func cal(a []int) (int, []int) {
	var n = len(a)
	if n <= 2 {
		return 0, a
	}
	sort.Ints(a)
	var check = func(t int) bool {
		la, ra := a[0:t], a[t:]
		var j = 0
		for i := 1; i < len(ra); i++ {
			if j < t && la[j] < ra[i-1] && la[j] < ra[i] {
				j++
			}
		}
		return j >= t
	}
	var calAns = func(t int) []int {
		la, ra := a[0:t], a[t:]
		var res = []int{ra[0]}
		var j = 0
		for i := 1; i < len(ra); i++ {
			if j < t && la[j] < ra[i-1] && la[j] < ra[i] {
				res = append(res, la[j])
				j++
			}
			res = append(res, ra[i])
		}
		return res
	}
	l, r := 1, n/2
	for l <= r {
		var mid = l + (r-l)/2
		if !check(mid) {
			r = mid - 1
		} else {
			if mid == n/2 || !check(mid+1) { // 找到答案
				return mid, calAns(mid)
			}
			l = mid + 1
		}
	}
	return 0, a
}
