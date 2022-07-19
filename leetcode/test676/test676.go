package test676

type MagicDictionary struct {
	d []string
}

func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	this.d = dictionary
}

func (this *MagicDictionary) Search(searchWord string) bool {
	for _, d := range this.d {
		if len(d) != len(searchWord) {
			continue
		}
		var dif = 0
		for idx, _ := range searchWord {
			if d[idx] != searchWord[idx] {
				dif++
			}
			if dif > 1 {
				break
			}
		}
		if dif == 1 {
			return true
		}
	}
	return false
}
