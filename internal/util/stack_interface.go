package util

type Stack interface {
	Push(item interface{})

	Pop() interface{}

	IsEmpty() bool

	Size() int

	Peek() interface{}
}
