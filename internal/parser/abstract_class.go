package parser

import "fmt"

/**
 * 抽象类得关键是这个函数指针
 * @author hushengdong
 */
type AbstractPerson struct {
	Sleep func() // 函数指针
}

func (this *AbstractPerson) eat() {

	fmt.Println("AbstractPerson eat")
	this.Sleep()
}

func (this *AbstractPerson) run() {

	fmt.Println("AbstractPerson run")
}

func eatAndRun(p *AbstractPerson) {

	fmt.Println("eat and run -----------------")
	p.eat()
	p.run()
	fmt.Println("eat and run end -----------------")
}

// Man 继承了抽象类
type Man struct {
	*AbstractPerson
}

func NewMan() *Man {

	p := &Man{}
	p.AbstractPerson = &AbstractPerson{}
	p.Sleep = p.sleep
	return p
}

func (this *Man) sleep() {

	fmt.Println("Man sleep")
}

// Women 继承了抽象类
type Women struct {
	*AbstractPerson
}

func NewWomen() *Women {

	p := &Women{}
	p.AbstractPerson = &AbstractPerson{}
	p.Sleep = p.sleep
	return p
}

func (this *Women) sleep() {

	fmt.Println("Women sleep")
}

// ---
