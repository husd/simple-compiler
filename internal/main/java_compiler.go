package main

import (
	"container/list"
	"fmt"
	jcComp "husd.com/v0/compiler"
	jcIo "husd.com/v0/io"
	"husd.com/v0/parser"
	"husd.com/v0/tree"
)

type RES int

const (
	SUCCESS RES = 1
	FAIL    RES = 0
)

var state jcComp.CompileState      // 编译状态
var errorCount int                 // 错误的数量
var fileMap = make(map[string]int) // 已经解析过的文件的集合

func compiler(files []string) RES {

	//转换源代码文件
	parseFiles(files)
	return SUCCESS
}

func parseFiles(files []string) *list.List {

	res := list.New()
	len := len(files)
	for i := 0; i < len; i++ {
		f := files[i]
		if _, ok := fileMap[f]; !ok {
			fmt.Println("开始编译文件: ", f)
			fileMap[f] = 1
			res.PushBack(parseFile(f))
		}
	}
	return res
}

func parseFile(path string) tree.JCCompilationUnit {

	charSequence := jcIo.GetCharSequenceFromFactory(path)
	p := parser.GetParserFromFactory(charSequence)
	res := p.ParseJCCompilationUnit()
	return res
}