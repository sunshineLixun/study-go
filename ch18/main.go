package main

import (
	"fmt"
	"time"
)

type Args struct {
	Name string
}

func debounce(fn func(interface{}), delay time.Duration) func(interface{}) {
	var timer *time.Timer = nil
	return func(args interface{}) {
		if timer != nil {
			timer.Stop()
			timer = nil
		}

		timer = time.AfterFunc(delay*time.Millisecond, func() {
			fn(args)
		})

	}
}

func main() {

	originFn := func(arg interface{}) {
		fmt.Println(arg)
	}

	deFn := debounce(originFn, 1000)

	last := &Args{Name: "args"}

	args := []interface{}{1, "Fn2", "Fn3", true, last}

	for _, arg := range args {
		deFn(arg)
	}

	// 模拟 阻塞下主线程，保存timer的延迟fn会执行
	time.Sleep(2 * time.Second)
}
