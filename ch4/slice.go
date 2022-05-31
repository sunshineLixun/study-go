package ch4

import "fmt"

func init() {
	// 声明
	var s1 []int
	s2 := []int{}
	s3 := make([]int, 2)

	arr := [5]int{1, 2, 3, 4, 5}
	// 前包后不包 [)
	s4 := arr[1:2]

	s5 := arr[1:4]

	//去掉切片的最后一个元素
	s6 := arr[0 : len(arr)-1]

	// 全部
	s7 := arr[:]

	s8 := arr[1:3]
	s8[0] = 100
	s8[1] = 200

	fmt.Println(s1, s2, s3, s4, s5, s6, s7, s8, arr)

}
