package s1

type Encrypter struct {
	enMap  map[byte]string
	cntMap map[string]int
}

func Constructor(keys []byte, values []string, dictionary []string) Encrypter {
	var enMap = make(map[byte]string)
	var cntMap = make(map[string]int)
	for i, v := range keys {
		enMap[v] = values[i]
	}
	var encrypter = Encrypter{cntMap: cntMap, enMap: enMap}
	for _, d := range dictionary {
		encrypter.cntMap[encrypter.Encrypt(d)]++
	}
	return encrypter
}

func (this *Encrypter) Encrypt(word1 string) string {
	var bs = make([]byte, 0)
	for _, c := range word1 {
		if this.enMap[byte(c)] == "" {
			return "1"
		}
		bs = append(bs, this.enMap[byte(c)]...)
	}
	return string(bs)
}

func (this *Encrypter) Decrypt(word2 string) int {
	return this.cntMap[word2]
}
