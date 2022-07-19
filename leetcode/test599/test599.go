package test599

func findRestaurant(list1 []string, list2 []string) []string {
	var mp = make(map[string]int)
	for idx, s1 := range list1 {
		mp[s1] = idx
	}
	var min = len(list1) + len(list2)
	for idx, s2 := range list2 {
		if v, exist := mp[s2]; exist {
			val := v + idx
			if val < min {
				min = val
			}
		}
	}
	var res = make([]string, 0)
	for idx, s2 := range list2 {
		if v, exist := mp[s2]; exist && v+idx == min {
			res = append(res, s2)
		}
	}
	return res
}
