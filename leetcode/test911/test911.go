package test911


type TopVotedCandidate struct {
	res, times []int
}


func Constructor(persons []int, times []int) TopVotedCandidate {
	last, n := 0, len(persons)
	res := make([]int, n)
	cnts := make([]int, n)
	for idx, p := range persons {
		cnts[p]++
		if cnts[p] >= cnts[last] {
			last = p
		}
		res[idx] = last
	}
	return TopVotedCandidate{
		res: res, times: times,
	}
}


func (this *TopVotedCandidate) Q(t int) int {
	l, r := 0, len(this.res) - 1
	for l <= r {
		mid := l + (r-l)/2
		if t < this.times[mid] {
			r = mid - 1
		} else {
			if mid == r || this.times[mid+1] > t {
				return this.res[mid]
			}
			l = mid + 1
		}
	}
	return this.res[0]
}