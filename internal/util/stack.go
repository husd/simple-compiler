package util

import "container/list"

/**
 * 这里用切片，模拟一个简单的栈，栈的元素是字符串
 * 数组是定长的，所以不太适合做栈，这里是一个实验性质的代码，后续要删除 TODO delete later
 * 或者优化一下，底层用数组实现 性能可能会更高点
 * @author hushengdong
 */
type DefaultStack struct {
	data *list.List
	size int
}

//实现栈的接口
func (s *DefaultStack) Push(item interface{}) {

	s.data.PushBack(item)
	s.size++
}

func (s *DefaultStack) Pop() interface{} {

	if s.size <= 0 {
		return nil
	}
	e := s.data.Back()
	s.data.Remove(e)
	s.size--
	return e.Value
}

func (s *DefaultStack) IsEmpty() bool {

	return s.size <= 0
}

func (s *DefaultStack) Size() int {

	return s.size
}

func (s *DefaultStack) Peek() interface{} {

	if s.size <= 0 {
		return nil
	}
	e := s.data.Back()
	return e.Value
}

func (s *DefaultStack) Init() {
	s.data = list.New()
	s.size = 0
}
