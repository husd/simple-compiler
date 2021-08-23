package io

import (
	"fmt"
)

type StringCharSequence struct {
	length  int    //长度
	content string //内容
}

func NewStringCharSequence(str string) (f CharSequence) {
	fcs := StringCharSequence{}
	fcs.content = str
	fcs.length = len(str)
	return fcs
}

func (f StringCharSequence) Len() int {
	return f.length
}

func (f StringCharSequence) CharAt(index int) rune {

	if index < 0 || index >= f.length {
		panic(fmt.Sprintf("out of index %d", index))
	}
	return (rune)(f.content[index])
}

func (f StringCharSequence) ByteAt(index int) uint8 {

	if index < 0 || index >= f.length {
		panic(fmt.Sprintf("out of index %d", index))
	}
	return f.content[index]
}

// 左闭右开
func (f StringCharSequence) SubCharSequence(start int, end int) string {

	checkScope(start, end, f.length)
	return f.content[start:end]
}
