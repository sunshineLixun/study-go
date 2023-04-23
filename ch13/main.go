package main

import "fmt"

var c = make(chan int, 2)

func main() {
	c <- 1
	c <- 2
	close(c)
	// rever()
	rever2()
}

func rever() {

	for {
		msg, ok := <-c
		if !ok {
			fmt.Println("通道关闭", msg, ok)
			break
		}
		fmt.Println("接受成功:", msg, ok)
	}
}

func rever2() {
	for msg := range c {

		fmt.Println("接受成功:", msg)
	}
}
