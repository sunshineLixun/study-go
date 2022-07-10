package ch7

import "fmt"

func init() {
	var wathever [5]struct{}
	for index := range wathever {
		defer fmt.Println(index) // 4 3 2 1 0
	}

	for i := range wathever {
		defer func() { fmt.Println(i) }() // 4 4 4 4
	}
}
