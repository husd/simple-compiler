package parser

import "fmt"

/**
 * 抽象类得关键是这个函数指针
 * @author hushengdong
 */
type abstractPerson struct {
	Sleep func() //函数指针
}

func (this *abstractPerson) eat() {

	fmt.Println("abstractPerson eat")
	this.Sleep()
}

func (this *abstractPerson) run() {

	fmt.Println("abstractPerson run")
}

// Man 继承了抽象类
type Man struct {
	abstractPerson
}

func NewMan() *Man {

	p := &Man{}
	p.Sleep = p.sleep
	return p
}

func (this *Man) sleep() {

	fmt.Println("Man sleep")
}

// Women 继承了抽象类
type Women struct {
	abstractPerson
}

func NewWomen() *Women {

	p := &Women{}
	p.Sleep = p.sleep
	return p
}

func (this *Women) sleep() {

	fmt.Println("Women sleep")
}
