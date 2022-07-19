package test1108

func defangIPaddr(address string) string {
	var bytes = make([]byte, 0)
	var p = []byte{'[', '.', ']'}
	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			bytes = append(bytes, p...)
		} else {
			bytes = append(bytes, address[i])
		}
	}
	return string(bytes)
}
