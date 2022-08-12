package test1417

func reformat(s string) string {
	var a1 = make([]byte, 0)
	var a2 = make([]byte, 0)
	for _, v := range s {
		if v >= '0' && v <= '9' {
			a2 = append(a2, byte(v))
		} else {
			a1 = append(a1, byte(v))
		}
	}
	n1, n2 := len(a1), len(a2)
	if n2 > n1 {
		n1, n2 = n2, n1
		a1, a2 = a2, a1
	}
	if n1-n2 > 1 {
		return ""
	}
	var bytes = make([]byte, 0)
	i, j := 0, 0
	for i < n1 || j < n2 {
		if i < n1 {
			bytes = append(bytes, a1[i])
			i++
		}
		if j < n2 {
			bytes = append(bytes, a2[j])
			j++
		}
	}
	return string(bytes)
}
