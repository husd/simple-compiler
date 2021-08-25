package io

import (
	"fmt"
	"unicode/utf8"
)

type StringCharSequence struct {
	length  int    //长度 这个是字符串的长度
	content string //内容
	pos     int    //当前位置
}

func NewStringCharSequence(str string) (f CharSequence) {
	fcs := StringCharSequence{}
	fcs.content = str
	fcs.length = utf8.RuneCountInString(str)
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

func (f StringCharSequence) ReadRune() rune {

	r := f.content[f.pos]
	f.pos = f.pos + 1
	return rune(r)
}

func (f StringCharSequence) CurrentPos() int {

	return f.pos
}
