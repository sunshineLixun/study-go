package ch7

import "fmt"

// defer 压栈 先进后出 是在函数体结束之前调用。如遇到return语句，则会在return语句之后执行，
// 使用场景：释放资源

func deferFunc() {
	fmt.Println("defer call")
}

func returnFunc() int {
	defer deferFunc()
	fmt.Println("return call")
	return 1
}

func init() {
	// var wathever [5]struct{}
	// for index := range wathever {
	// 	defer fmt.Println(index) // 4 3 2 1 0
	// }

	// for i := range wathever {
	// 	defer func() { fmt.Println(i) }() // 4 4 4 4
	// }

	returnFunc()
}
