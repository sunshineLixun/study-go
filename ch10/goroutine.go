package ch10

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Println("i =", i)
		time.Sleep(1 * time.Second)
	}
}

func init() {
	go newTask()

	i := 0
	for {
		i++
		fmt.Printf("main goroutine i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
	// fmt.Println("main exit")
}
