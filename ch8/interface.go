package ch8

import (
	"fmt"
	"reflect"
)

// interface{}  万能类型

func showSomeThing(sm interface{}) {
	// 断言
	// va := sm.(string)
	// fmt.Println(va)

	i, ok := sm.(int)
	if !ok {
		fmt.Println("not int")
	} else {
		fmt.Println(i)
	}
}

type Type struct {
	Name string
}

func init() {
	// showSomeThing("111")

	// showSomeThing(123)

	t := &Type{
		"aaa",
	}

	reflectFunc(t)
}

func reflectFunc(value interface{}) {
	vType := reflect.TypeOf(value)
	fmt.Println(vType)
}
