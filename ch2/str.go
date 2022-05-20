package ch2

import (
	"fmt"
	"strings" // 所有的字符串操作都在strings的包中
)

const name string = "go"

func init() {
	const length = len(name)
	fmt.Println("name字符串长度：", length)

	name1 := name + "js"
	fmt.Println(name1)

	cloneName := strings.Clone(name1)

	fmt.Println("cloneName:", cloneName)

	splitName := strings.Split(name1, ",")

	fmt.Println("splitName:", splitName)

	fmt.Println(strings.Contains("go", "b"))

	fmt.Println(strings.HasPrefix(name1, "g"))

	fmt.Println(strings.HasSuffix(name1, "gs"))

	index := strings.LastIndex(name1, "0") // -1
	fmt.Println("index:", index)
}
