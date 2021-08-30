package main

import (
	"container/list"
	"fmt"
	jcComp "husd.com/v0/compiler"
	"husd.com/v0/jc"
	"husd.com/v0/parser"
	"husd.com/v0/util"
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
	context := util.NewContext()
	parseFiles(files, context)
	return SUCCESS
}

func parseFiles(files []string, context *util.Context) *list.List {

	res := list.New()
	len := len(files)
	for i := 0; i < len; i++ {
		f := files[i]
		if _, ok := fileMap[f]; !ok {
			fmt.Println("开始编译文件: ", f)
			fileMap[f] = 1
			parseFile(f, context)
		}
	}
	return res
}

func parseFile(path string, context *util.Context) jc.JCCompilationUnit {

	p := parser.GetParserFromFactory(path)
	res := p.ParseJCCompilationUnit()
	return res
}
