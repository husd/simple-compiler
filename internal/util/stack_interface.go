package util

//Stack
/**
 * 之前设计这个栈的接口，是因为要用栈实现遍历目录下的所有文件的功能，就目前而言，还不需要这么做，
 * 所以暂时先不管这个接口。
 */
type Stack interface {
	Push(item interface{})

	Pop() interface{}

	IsEmpty() bool

	Size() int

	Peek() interface{}
}
