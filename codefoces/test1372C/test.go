package main

/**
 * @Description https://codeforces.com/problemset/problem/1399/D
 * idea: 构造 + 贪心。
		1.若每个元素在原来的位置则无需交换，答案为0
		2.若元素的前缀和后缀都在原来的位置，中间【完全乱序】的元素只需要一次交换
		3.否则全部为2次。去除情况2后，对于中间有元素在原来的位置，至少需要1次交换使得在原来位置的元素交换到其它位置
		  再用1次使得整个数组有序。因为只要存在不在原来位置的元素，一定能通过和在原来位置的元素交换、移动，满足每
		 个元素位置均进行了改变。
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
			a[i] = v
		}
		fmt.Fscanln(in)
		fmt.Fprintln(out, cal(a))
		out.Flush()
	}
}

func cal(a []int) int {
	var n = len(a)
	i, j := 0, n-1
	for i < n {
		if a[i] != i+1 {
			break
		}
		i++
	}
	for j >= 0 && j > i {
		if a[j] != j+1 {
			break
		}
		j--
	}
	if i >= j {
		return 0
	}
	for l := i; l <= j; l++ {
		if a[l] == l+1 {
			return 2
		}
	}
	return 1
}
