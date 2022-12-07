package main

func twoEditWords(queries []string, dictionary []string) []string {
	var res = make([]string, 0)
	for _, q := range queries {
		var can = false
		for _, d := range dictionary {
			if len(q) != len(d) {
				continue
			}
			var cnt = 0
			var isCan = true
			for i := 0; i < len(q); i++ {
				if q[i] != d[i] {
					cnt++
				}
				if cnt > 2 {
					isCan = false
					break
				}
			}
			if isCan {
				can = true
				break
			}
		}
		if can {
			res = append(res, q)
		}
	}
	return res
}
