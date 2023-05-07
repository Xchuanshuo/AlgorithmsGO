package main

import "fmt"

func pyramidTransition(bottom string, allowed []string) bool {
	var mp = make(map[string][]byte)
	for _, v := range allowed {
		mp[v[0:2]] = append(mp[v[0:2]], v[2])
	}
	var mem = make(map[string]bool)
	var dfs func(last, cur []byte) bool
	dfs = func(last, cur []byte) bool {
		if len(last) == 1 {
			return true
		}
		var key = string(last) + "_" + string(cur)
		if v, exist := mem[key]; exist {
			return v
		}
		if len(cur)+1 == len(last) {
			return dfs(cur, []byte(""))
		}
		var i = len(cur)
		for _, c := range mp[string(last[i:i+2])] {
			if dfs(last, append(cur, c)) {
				mem[key] = true
				return true
			}
		}
		mem[key] = false
		return false
	}
	return dfs([]byte(bottom), []byte(""))
}

func main() {
	var bottom = "DAAAAD"
	var pat = []string{"DAD", "DAE", "DAB", "DAF", "DAC", "EAD", "EAE", "EAB", "EAF", "EAC", "BAD", "BAE", "BAB", "BAF", "BAC", "FAD", "FAE", "FAB", "FAF", "FAC", "CAD", "CAE", "CAB", "CAF", "CAC", "ADD", "ADE", "ADB", "ADF", "ADC", "AED", "AEE", "AEB", "AEF", "AEC", "ABD", "ABE", "ABB", "ABF", "ABC", "AFD", "AFE", "AFB", "AFF", "AFC", "ACD", "ACE", "ACB", "ACF", "ACC", "AAD", "AAE", "AAB", "AAF", "AAC", "AAA"}
	var res = pyramidTransition(bottom, pat)
	fmt.Println(res)
}
