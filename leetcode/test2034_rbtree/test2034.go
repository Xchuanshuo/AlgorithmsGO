package test2034_rbtree

import "github.com/emirpasic/gods/trees/redblacktree"

type StockPrice struct {
	prices   *redblacktree.Tree
	timesMap map[int]int
	maxT     int
}

func Constructor() StockPrice {
	return StockPrice{redblacktree.NewWithIntComparator(), make(map[int]int), 0}
}

func (this *StockPrice) Update(timestamp int, price int) {
	if timestamp > this.maxT {
		this.maxT = timestamp
	}
	if last, exist := this.timesMap[timestamp]; exist {
		res, _ := this.prices.Get(last)
		cnt := res.(int)
		if cnt <= 1 {
			this.prices.Remove(last)
		} else {
			this.prices.Put(last, cnt-1)
		}
	}
	res, found := this.prices.Get(price)
	if !found {
		this.prices.Put(price, 1)
	} else {
		this.prices.Put(price, res.(int)+1)
	}
	this.timesMap[timestamp] = price
}

func (this *StockPrice) Current() int {
	return this.timesMap[this.maxT]
}

func (this *StockPrice) Maximum() int {
	return this.prices.Right().Key.(int)
}

func (this *StockPrice) Minimum() int {
	return this.prices.Left().Key.(int)
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
