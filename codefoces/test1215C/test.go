package main

/**
 * @Description https://codeforces.com/problemset/problem/1215/C
 * idea: 构造 + 贪心。
		1.只交换不相等的位置 2.优先使用一次交换的方法
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
	var s, t string
	fmt.Fscanln(in, &s)
	fmt.Fscanln(in, &t)
	l, res := cal(s, t)
	fmt.Fprintln(out, l)
	for _, v := range res {
		fmt.Fprintf(out, "%d %d\n", v[0], v[1])
	}
	out.Flush()
}

// a a 交叉交换   a b 使用一次交换变为 a a 第二次交叉交换
// b b			 b a 			   b b
func cal(s, t string) (int, [][]int) {
	var ab = make([]int, 0)
	var ba = make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == t[i] {
			continue
		}
		if s[i] == 'a' {
			ab = append(ab, i+1)
		} else {
			ba = append(ba, i+1)
		}
	}
	// 必须包含偶数个位置不相等
	if (len(ab)+len(ba))%2 != 0 {
		return -1, [][]int{}
	}
	var res = make([][]int, 0)
	for i := 0; i < len(ab)-1; i += 2 {
		res = append(res, []int{ab[i], ab[i+1]})
	}
	for i := 0; i < len(ba)-1; i += 2 {
		res = append(res, []int{ba[i], ba[i+1]})
	}

	if len(ab)%2 == 1 {
		res = append(res, []int{ab[len(ab)-1], ab[len(ab)-1]})
		res = append(res, []int{ab[len(ab)-1], ba[len(ba)-1]})
	}
	return len(res), res
}
