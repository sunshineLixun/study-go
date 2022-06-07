package ch6

import "fmt"

// 嵌套
type Person1 struct {
	name string
	age  int
}

type Son struct {
	name string
	age  int
}

// 具名
// type User1 struct {
// 	person Person1
// 	son    Son
// }

// 匿名
type User struct {
	Person1
	Son
}

func init() {
	user := User{
		Person1{
			name: "Persson",
			age:  30,
		},
		Son{
			name: "son",
			age:  2,
		},
	}

	// 匿名
	user2 := User{
		Person1{
			"Persson",
			30,
		},
		Son{
			"son",
			2,
		},
	}

	fmt.Println(user, user2)

	fmt.Println(user.Son.age, user.Person1.name)

	user3 := User{} // {{ 0} { 0}}
	fmt.Println(user3)

	d := &Dog{
		age: 2,
		Animal: &Animal{ // 指针类型不能使用匿名嵌套
			name: "wangwang",
		},
	}

	fmt.Println(d.name, d.age)

	d.dreaking()
	d.runing()

	d.setName("haha")

	d.dreaking() // haha正在喝水
	d.runing()   // hhaha正在跑步

}

// 继承
type Animal struct {
	name string
}

func (a *Animal) runing() {
	fmt.Printf("%s正在跑步\n", a.name)
}

type Dog struct {
	age     int
	*Animal // 指针类型 -- 通过嵌套匿名结构体实现继承
}

func (d *Dog) dreaking() {
	fmt.Printf("%s正在喝水\n", d.name)
}

func (d *Dog) setName(name string) {
	d.name = name
}
