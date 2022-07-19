package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanln(&n)
	var tArr = make([]float64, 0)
	var dArr = make([]float64, 0)
	for i := 0;i < n;i++ {
		var t byte
		var v float64
		fmt.Scanf("%c %f", &t, &v)
		if t == 'T' {
			tArr = append(tArr, v)
		} else {
			dArr = append(dArr, v)
		}
	}
	sort.Float64s(tArr)
	sort.Float64s(dArr)
	i, j, m, n := 0, 0, len(tArr), len(dArr)
	v, d, t := 1.0, 0.0, 0.0
	for i<m && j<n {
		t1 := tArr[i] - t
		t2 := (dArr[j] - d) * v
		if t1 == t2 {
			t = tArr[i]
			d = dArr[j]
			v += 2
			i++
			j++
		} else if t1 < t2 {
			d += (tArr[i]-t)/v
			t = tArr[i]
			v++
			i++
		} else {
			t += (dArr[j]-d) * v
			d = dArr[j]
			v++
			j++
		}
	}
	for i < m {
		d += (tArr[i]-t)/v
		t = tArr[i]
		v++
		i++
	}
	for j < n {
		t += (dArr[j]-d) * v
		d = dArr[j]
		v++
		j++
	}
	fmt.Println(int(t + (1000-d)*v + 0.5))
}