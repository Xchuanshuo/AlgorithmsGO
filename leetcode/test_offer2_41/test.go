package test_offer2_41

type MovingAverage struct {
	size  int
	q     []float64
	total float64
	avg   float64
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size: size,
		q:    make([]float64, 0),
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.q) == this.size {
		this.total -= this.q[0]
		this.q = append(this.q[1:])
	}
	this.q = append(this.q, float64(val))
	this.total += float64(val)
	return this.total / float64(len(this.q))
}
