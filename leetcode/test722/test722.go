package test722

const (
	INIT = iota
	ML
)

func removeComments(source []string) []string {
	var res = make([]string, 0)
	var state = INIT
	var bs = make([]byte, 0)
	for _, s := range source {
		for i := 0; i < len(s); i++ {
			if state == INIT {
				if i != len(s)-1 && s[i:i+2] == "//" {
					break
				}
				if i != len(s)-1 && s[i:i+2] == "/*" {
					i++
					state = ML
					continue
				}
				bs = append(bs, s[i])
			}
			if state == ML {
				if i != len(s)-1 && s[i:i+2] == "*/" {
					i++
					state = INIT
				}
			}
		}
		if len(bs) != 0 && state != ML {
			res = append(res, string(bs))
			bs = make([]byte, 0)
		}
	}
	return res
}
