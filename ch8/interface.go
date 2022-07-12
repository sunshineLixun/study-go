package ch8

import (
	"fmt"
	"reflect"
)

// interface{} 空接口 万能类型

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

	// 空接口实现可以保存任意值的字典
	var imap = make(map[string]interface{})

	imap["name"] = "go"

	imap["id"] = [6]int{
		1,
	}

	imap["infos"] = make([]Type, 5)

	imap["stu"] = [...]Type{
		{
			Name: "12312",
		},
	}

	fmt.Println(imap)
}

func reflectFunc(value interface{}) {
	vType := reflect.TypeOf(value)
	fmt.Println(vType)
}
