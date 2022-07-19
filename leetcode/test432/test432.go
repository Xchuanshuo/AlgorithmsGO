package test432

const INF = int(1e9)

type AllOne struct {
	mp           map[string]int
	fMax, sMax   int
	fMaxS, sMaxS string
	fMin, sMin   int
	fMinS, sMinS string
}

func Constructor() AllOne {
	return AllOne{
		mp:   make(map[string]int),
		fMin: INF, sMin: INF,
		fMax: 0, sMax: 0,
	}
}

func (this *AllOne) Inc(key string) {
	this.mp[key]++
	var cnt = this.mp[key]
	this.update(key, cnt, true)
}

func (this *AllOne) Dec(key string) {
	this.mp[key]--
	var cnt = this.mp[key]
	if cnt == 0 {
		delete(this.mp, key)
	}
	this.update(key, cnt, false)
}

func (this *AllOne) update(key string, cnt int, isAdd bool) {
	if isAdd {
		if cnt >= this.fMax {
			this.sMax = this.fMax
			this.sMaxS = this.fMaxS
			this.fMax = cnt
			this.fMaxS = key
		} else if cnt > this.sMax {
			this.sMax = cnt
		} else if key == this.fMinS {
			this.fMinS = this.sMinS
			this.fMin = this.sMin
			this.sMinS = key
			this.sMin = cnt
		}
	} else {
		if cnt <= this.fMin {
			this.sMin = this.fMin
			this.sMinS = this.fMinS
			this.fMin = cnt
			this.fMinS = key
		} else if cnt < this.sMin {
			this.sMin = cnt
		} else if key == this.fMaxS {
			this.fMaxS = this.sMaxS
			this.fMax = this.sMax
			this.sMaxS = key
			this.sMin = cnt
		}
	}
}

func (this *AllOne) GetMaxKey() string {
	return this.fMaxS
}

func (this *AllOne) GetMinKey() string {
	return this.fMinS
}
