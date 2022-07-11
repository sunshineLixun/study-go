package ch8

import "fmt"

type Move interface {
	move()
	say()
	setName(string) bool
}

type Run interface {
	run()
}

// 接口嵌套
type Any interface {
	Move
	Run
}

type Dog struct{}

type Cat struct{}

type People struct{}

func (dog Dog) move() {
	fmt.Println("dog move")
}

func (cat Cat) say() {
	fmt.Println("cat say")
}

func (p People) setName(name string) bool {
	fmt.Println(name)
	return true
}

func (p People) run() {
	fmt.Println("people run")
}

func init() {
	dog := Dog{}
	dog.move()

	cat := Cat{}

	cat.say()

	p := People{}

	ok := p.setName("go")
	if ok {
		fmt.Println("成功")
	}

	p.run()

}
