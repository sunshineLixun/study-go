package ch3

import "fmt"

//是同一种数据类型的固定长度的序列

// 申明
var a [5]int = [5]int{1, 2, 3, 4, 5}

var b [6]string = [6]string{"1", "2", "3", "4", "5", "6"}

var c = [...]string{"1", "2", "3"}

func loop() {
	for i := 0; i < len(a); i++ {
		fmt.Println(i)
	}

	for index, v := range a {
		fmt.Println(index, v)
	}
}

func setUp() {
	a := [...]string{"1", "2"}
	b := [5]int{1, 2}         // [1, 2, 0, 0, 0]
	c := [5]int{2: 10, 3: 11} // 使用索引号初始化元素。[0 0 10 11 0]

	d := [...]struct {
		name string
		age  int
	}{
		{
			name: "js",
			age:  20,
		},
		{
			name: "go",
			age:  20,
		},
	} // [{js 20} {go 20}]

	f := [...]struct {
		name string
		age  int
	}{
		{"js", 20},
		{"go", 20},
	} // 可省略元素类型
	fmt.Println(a, b, c, d, f)
}

// 多维数组
func two() {
	var a1 [3][3]int                           // [[0 0 0] [0 0 0] [0 0 0]]
	a := [5][3]int{{1, 2, 3}, {4, 5, 6}}       // [[1 2 3] [4 5 6] [0 0 0] [0 0 0] [0 0 0]]
	b := [5][3]int{2: {1, 2, 3}, 3: {4, 5, 6}} // [[0 0 0] [0 0 0] [1 2 3] [4 5 6] [0 0 0]]
	fmt.Println(a1, a, b)

	// 多维数组遍历
	for _, v := range a {
		for i := 0; i < len(v); i++ {
			fmt.Println(i)
		}
	}

	for _, value := range b {
		for _, value2 := range value {
			fmt.Println(value2)
		}
	}
}

// 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
func copy(value [2]int) {
	fmt.Printf("value: %p\n", &value)
	value[0] = 200
	fmt.Println(value)
}

func pointerCopy(value *[2]int) {
	value[0] = 200
	fmt.Println(value)
}

func init() {
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	// loop()

	// setUp()

	// two()

	a := [2]int{1, 2}
	fmt.Printf("a: %p\n", &a)
	// copy(a)
	pointerCopy(&a)
	fmt.Println("a:", a)
}
