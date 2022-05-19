package main

import (
	"fmt"
	ch1 "study-go/ch1"
)

func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Println("hello world!")
	ch1.Get()
	ch1.Log()

	fmt.Println(ch1.Package)
}
