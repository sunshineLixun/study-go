package ch6

import "fmt"

type people struct {
	name   string
	age    int
	gender bool
}

func init() {
	// 键值对初始化
	peo := people{
		name:   "go",
		age:    10,
		gender: true,
	}
	fmt.Println(peo)

	// 基础实例化
	var p1 people
	p1.name = "go"
	p1.age = 10
	p1.gender = true

	fmt.Println(p1)

	// new 关键字 结构体的地址
	var p2 = new(people)
	p2.name = "js"
	p2.age = 20
	p2.gender = false

	fmt.Println(p2)            // &{js 20 false}
	fmt.Printf("%T\n", p2)     // *ch6.people
	fmt.Printf("p2=%#v\n", p2) // p2=&ch6.people{name:"js", age:20, gender:false}

	// 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作
	var p3 = &people{}
	p3.name = "swift"

	fmt.Println(p3) // &{swift 0 false}

	var p4 = &people{}

	p0 := &people{
		"name",
		12,
		true,
	}

	fmt.Println(p0)

	fmt.Println(p4) // &{oc 0 false}

	fmt.Printf("%p\n", p4) // 0x1400007a0e0

	var p5 = &people{}

	fmt.Println(p5) // &{ 0 false}

	//简写 严格按照字段顺序
	p6 := &people{
		"go",
		12,
		true,
	}

	fmt.Println(p6)

}
