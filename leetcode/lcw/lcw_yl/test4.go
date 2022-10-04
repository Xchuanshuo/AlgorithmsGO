package main

import (
	"container/heap"
	"fmt"
)

type VendingMachine struct {
	mp     map[string]*Heap
	disMap map[string]int
}

func Constructor() VendingMachine {
	return VendingMachine{map[string]*Heap{}, make(map[string]int)}
}

func (this *VendingMachine) AddItem(time int, number int, item string, price int, duration int) {
	if _, exist := this.mp[item]; !exist {
		this.mp[item] = &Heap{}
	}
	heap.Push(this.mp[item], Goods{number, price, time + duration})
}

func (this *VendingMachine) Sell(time int, customer string, item string, number int) int64 {
	if _, exist := this.mp[item]; !exist {
		return -1
	}
	var hp = this.mp[item]
	var total int64 = 0
	var buyCnt = number
	var temp = make([]Goods, 0)
	for hp.Len() > 0 {
		var cur = heap.Pop(hp).(Goods)
		if time > cur.duration {
			continue
		}
		if cur.number >= buyCnt {
			total += int64(buyCnt * cur.price)
			cur.number -= buyCnt
			buyCnt = 0
			if cur.number > 0 {
				heap.Push(hp, cur)
			}
			break
		}
		temp = append(temp, Goods{cur.number, cur.price, cur.duration})
		buyCnt -= cur.number
		total += int64(cur.number * cur.price)
	}
	if buyCnt > 0 {
		for _, v := range temp {
			fmt.Println(v)
			heap.Push(hp, v)
		}
		return -1
	}
	if _, exist := this.disMap[customer]; !exist {
		this.disMap[customer] = 100
	} else {
		this.disMap[customer] = max(this.disMap[customer]-1, 70)
	}
	return (total*int64(this.disMap[customer]) + 99) / 100
}

type Goods struct {
	number          int
	price, duration int
}

type GoodsSlice []Goods

type Heap struct {
	GoodsSlice
}

func (h Heap) Len() int {
	return len(h.GoodsSlice)
}

func (h Heap) Less(i, j int) bool {
	a, b := h.GoodsSlice[i], h.GoodsSlice[j]
	if a.price == b.price {
		return a.duration < b.duration
	}
	return a.price < b.price
}

func (h Heap) Swap(i, j int) {
	h.GoodsSlice[i], h.GoodsSlice[j] = h.GoodsSlice[j], h.GoodsSlice[i]
}

func (h *Heap) Push(v interface{}) {
	h.GoodsSlice = append(h.GoodsSlice, v.(Goods))
}

func (h *Heap) Pop() interface{} {
	a := h.GoodsSlice
	v := a[len(h.GoodsSlice)-1]
	h.GoodsSlice = a[0 : len(h.GoodsSlice)-1]
	return v
}

//func main() {
//	var m = Constructor()
//	m.mp["ss"] = &Heap{}
//	hp := &Heap{}
//	hp.Push(Goods{1, 1, 1})
//	fmt.Println(hp)
//	//m.mp["ss"] = hp
//	//m.mp["ss"].Push(&Goods{2, 2, 2})
//	//fmt.Println(m.mp)
//}
