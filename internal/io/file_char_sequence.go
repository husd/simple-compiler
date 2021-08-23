package io

import (
	"fmt"
	"io/ioutil"
)

type FileCharSequence struct {
	// 文件的路径
	Path    string
	length  int    //长度
	content string //内容
}

func NewFileCharSequence(path string) (f CharSequence) {
	fcs := FileCharSequence{}
	//init
	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		panic("读取文件错误：" + path)
	}
	fcs.content = string(buffer)
	fcs.Path = path
	fcs.length = len(fcs.content)
	return fcs
}

func (f FileCharSequence) Len() int {
	return f.length
}

func (f FileCharSequence) CharAt(index int) rune {

	if index < 0 || index >= f.length {
		panic(fmt.Sprintf("out of index %d", index))
	}
	return (rune)(f.content[index])
}

// 左闭右开
func (f FileCharSequence) SubCharSequence(start int, end int) string {

	checkScope(start, end, f.length)
	return f.content[start:end]
}
