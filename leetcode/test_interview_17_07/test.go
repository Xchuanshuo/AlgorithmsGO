package test_interview_17_07

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func trulyMostPopular(names []string, synonyms []string) []string {
	var cntMap = make(map[string]int)
	var uf = NewUF()
	for _, name := range names {
		var strs = strings.Split(name[0:len(name)-1], "(")
		v, _ := strconv.Atoi(strs[1])
		cntMap[strs[0]] = v
		uf.union(strs[0], strs[0])
	}
	for _, syn := range synonyms {
		var strs = strings.Split(syn[1:len(syn)-1], ",")
		uf.union(strs[0], strs[1])
	}
	var unionMap = make(map[string][]string)
	for k, _ := range uf.parent {
		var p = uf.find(k)
		unionMap[p] = append(unionMap[p], k)
	}
	var res = make([]string, 0)
	for _, names := range unionMap {
		sort.Strings(names)
		var v = 0
		for _, name := range names {
			v += cntMap[name]
		}
		res = append(res, fmt.Sprintf("%s(%d)", names[0], v))
	}
	return res
}

type UF struct {
	parent map[string]string
	sz     map[string]int
}

func NewUF() *UF {
	var parent = make(map[string]string)
	var sz = make(map[string]int)
	var uf = &UF{parent: parent, sz: sz}
	return uf
}

func (uf *UF) union(p, q string) {
	var pRoot = uf.find(p)
	var qRoot = uf.find(q)
	var sz = uf.sz
	if pRoot == qRoot {
		return
	}
	if sz[pRoot] > sz[qRoot] {
		sz[pRoot] += sz[qRoot]
		uf.parent[qRoot] = pRoot
	} else {
		sz[qRoot] += sz[pRoot]
		uf.parent[pRoot] = qRoot
	}
}

func (uf *UF) isConnected(p, q string) bool {
	return uf.find(p) == uf.find(q)
}

func (uf *UF) find(p string) string {
	if _, exist := uf.parent[p]; !exist {
		uf.parent[p] = p
	}
	var parent = uf.parent
	for parent[p] != p {
		parent[p] = parent[parent[p]]
		p = parent[p]
	}
	return p
}
