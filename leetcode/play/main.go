package main

import (
	"fmt"
	"reflect"
)

var beanMap = make(map[string]interface{})

func getBean[T any](t T) T {
	fmt.Println(reflect.TypeOf(t).String())
	return beanMap[reflect.TypeOf(t).String()].(T)
}

type A struct {
	A int
}

func a() {
	beanMap[reflect.TypeOf(A{}).String()] = A{2}
	var T = getBean(A{})
	fmt.Println(T)
}

func main() {
	a()
}
