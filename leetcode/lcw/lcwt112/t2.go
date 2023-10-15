package lcwt112

import "sort"

func checkStrings(s1 string, s2 string) bool {
	var bse1 = make([]byte, 0)
	var bso1 = make([]byte, 0)
	var bse2 = make([]byte, 0)
	var bso2 = make([]byte, 0)
	for i, v := range s1 {
		if i%2 == 0 {
			bse1 = append(bse1, byte(v))
			bse2 = append(bse2, s2[i])
		} else {
			bso1 = append(bso1, byte(v))
			bso2 = append(bso2, s2[i])
		}
	}
	sort.Slice(bse1, func(i, j int) bool {
		return bse1[i] < bse1[j]
	})
	sort.Slice(bso1, func(i, j int) bool {
		return bso1[i] < bso1[j]
	})
	sort.Slice(bse2, func(i, j int) bool {
		return bse2[i] < bse2[j]
	})
	sort.Slice(bso2, func(i, j int) bool {
		return bso2[i] < bso2[j]
	})
	return string(bse1) == string(bse2) && string(bso1) == string(bso2)
}
