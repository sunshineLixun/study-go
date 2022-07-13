package ch11

import "fmt"

func init() {

	c := make(chan int)

	strC := make(chan string)

	go func() {
		c <- 777
	}()

	go func() {
		strC <- "111"
	}()

	num := <-c

	fmt.Println(num)

	str := <-strC

	fmt.Println(str)

}
