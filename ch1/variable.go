package ch1

import (
	"fmt"
	"time"
)

// 变量 常量

var Package string

const Package1 = "Package1"

var name string

var age int

var sex bool

var (
	_name  string
	_age   int
	gneder bool
)

var (
	m = 13
	o = int16(23)
	r = float32(3.14)
)

const (
	n  = 100
	n2 = 101
	n3 = "1233123"
)

var a, b, c = 12, 'A', "hello"

var pan = 'B'

func Get() {
	var a int16 = 5
	var b int = 8
	var c int16
	c = a + int16(b)
	fmt.Println(n, n2, n3, name, age, sex, c)

	// 局部变量 :=
	st := "12321"
	fmt.Println(st)

	f, x, y := 12, 'A', "hello"
	fmt.Println(f, x, y)
}

func init() {
	fmt.Println("var init")

	log()

	testMakePerson()
}

func log() {
	fmt.Println(n + n2)
}

// 声明聚类

var (
	isShow bool // false
	canUse bool // false
)

var (
	person       = ""
	son          = ""
	girl         = ""
	aLongTimeAgo = time.Unix(1, 0)
)

func makePerson() (name string, age int) {
	return "name", 20
}

// 短变量声明的就近原则
func testMakePerson() {
	if name, age := makePerson(); name != "" && age != 0 {
		fmt.Println(name, age)
	}
}
