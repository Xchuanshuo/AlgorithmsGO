package test1656

type OrderedStream struct {
	arr []string
	ptr int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{make([]string, n), 0}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	var pos = idKey - 1
	this.arr[pos] = value
	var res = make([]string, 0)
	for i := this.ptr; i < len(this.arr); i++ {
		if this.arr[i] == "" {
			break
		}
		res = append(res, this.arr[i])
	}
	this.ptr = this.ptr + len(res)
	return res
}
