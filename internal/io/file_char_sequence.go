package io

import (
	"fmt"
	"husd.com/v0/parser"
	"io/ioutil"
)

type FileCharSequence struct {
	// 文件的路径
	path   string
	length int                   //长度 这个是byte的长度
	reader *parser.UnicodeReader //reader 数据在这里
}

func NewFileCharSequence(path string) (f CharSequence) {
	fcs := FileCharSequence{}
	//init
	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		panic("读取文件错误：" + path)
	}
	fcs.path = path
	fcs.length = len(buffer)
	fcs.reader = parser.NewUnicodeReader(buffer)
	return fcs
}

//这个返回的只真实的字节的长度
func (f FileCharSequence) Len() int {
	return f.length
}

func (f FileCharSequence) CharAt(index int) rune {

	panic("这个方法没有实现，请检查")
}

func (f FileCharSequence) ByteAt(index int) uint8 {

	if index < 0 || index >= f.length {
		panic(fmt.Sprintf("out of index %d", index))
	}
	return f.reader.byteAt(index)
}

// 左闭右开
func (f FileCharSequence) SubCharSequence(start int, end int) string {

	checkScope(start, end, f.length)
	return string(f.reader.subByteArray(start, end))
}

func (f FileCharSequence) ReadRune() rune {

	r, _ := f.reader.readRune()
	return r
}

func (f FileCharSequence) CurrentPos() int {

	return f.reader.currentPos()
}
