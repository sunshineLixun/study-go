package ch6

import "fmt"

// struct是值类型
// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

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

	p7 := newPerson("vscode", 12, true)

	p7.dreak()

	fmt.Println(p7)

	p8 := newPerson("123", 12, true)

	p8.setName("456")

	fmt.Println(p8) // &{456 12 true}

	p9 := newPerson("kk", 12, true)
	p9.setName2("nn")

	fmt.Println(p9) // &{kk 12 true}

}

// 构造函数。工厂函数
func newPerson(name string, age int, gender bool) *people {
	return &people{
		name:   name,
		age:    age,
		gender: gender,
	}
}

func (p people) dreak() {
	fmt.Printf("姓名是: %s", p.name)
}

// 指针类型 类似js 中 的 this。传递指针，修改的是指针指向的数据
func (p *people) setName(name string) {
	p.name = name
}

// 值类型 改操作只是针对副本，无法修改接收者变量本身
func (p people) setName2(name string) {
	p.name = name
}
