package main

import "fmt"

var c = make(chan int)

func main() {
	go getMsg()
	c <- 777

	fmt.Println("发送成功")
}

func getMsg() {
	msg := <-c
	fmt.Println("接受成功:", msg)
}
