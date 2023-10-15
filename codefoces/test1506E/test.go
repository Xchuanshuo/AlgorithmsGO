package main

/**
 * @Description https://codeforces.com/problemset/problem/1506/E
 * idea: 构造 + 贪心。
				1.最大值第一次出现的位置为固定值，其它地方按照贪心策略插空
				2.最小数组: 固定值后剩余的空位置从小到大插入id即可
				  最大数组: 从固定值位置由大到小插入，但可能出现当前将插入的id已经使用过，
						所以可以用一个hash表保存每个固定位置最后没有访问到的id，当前id
						使用过时直接从使用过的位置查找即可
 **/

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	run(os.Stdin, os.Stdout)
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	cases := 1
	fmt.Fscanln(in, &cases)
	for curCase := 1; curCase <= cases; curCase++ {
		var n int
		fmt.Fscanln(in, &n)
		var a = make([]int, n)
		for i := 0; i < n; i++ {
			var v int
			fmt.Fscan(in, &v)
			fmt.Fscan(in, " ")
			a[i] = v
		}
		fmt.Fscanln(in)
		for _, v := range calMin(a) {
			fmt.Fprint(out, v)
			fmt.Fprint(out, " ")
		}
		fmt.Fprintln(out)
		for _, v := range calMax(a) {
			fmt.Fprint(out, v)
			fmt.Fprint(out, " ")
		}
		fmt.Fprintln(out)
		out.Flush()
	}
}

func calMin(a []int) []int {
	var n = len(a)
	var res = make([]int, n)
	var mp = make(map[int]bool)
	for i, v := range a {
		if i == 0 || v != a[i-1] {
			res[i] = v
			mp[v] = true
		}
	}
	var id = 1
	for i := 0; i < n; i++ {
		if res[i] != 0 {
			continue
		}
		for mp[id] {
			id += 1
		}
		res[i] = id
		id += 1
	}
	return res
}

func calMax(a []int) []int {
	var n = len(a)
	var res = make([]int, n)
	var mp = make(map[int]bool)
	var mnMap = make(map[int]int)
	for i, v := range a {
		if i == 0 || v != a[i-1] {
			res[i] = v
			mp[v] = true
			mnMap[v] = v - 1
		}
	}
	for i := 0; i < n; {
		if res[i] != 0 {
			i++
			continue
		}
		var pre = res[i-1]
		var id = res[i-1] - 1
		for i < n && res[i] == 0 {
			for mp[id] {
				id = mnMap[id]
			}
			res[i] = id
			mp[id] = true
			id--
			i++
		}
		mnMap[pre] = id
	}
	return res
}
