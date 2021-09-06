package main

import (
	"container/list"
	"fmt"
	"husd.com/v0/ast_tree"
	jcComp "husd.com/v0/compiler"
	"husd.com/v0/parser"
	"husd.com/v0/util"
)

/**
 *
 * @author hushengdong
 */
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
			if !util.Exists(f) {
				fmt.Println("编译文件失败，文件不存在:", f)
				continue
			}
			fileMap[f] = 1
			context := util.NewContext()
			parseFile(f, context)
		}
	}
	return res
}

func parseFile(path string, context *util.Context) ast_tree.JCCompilationUnit {

	p := parser.GetParserFromFactory(path, context)
	res := p.ParseJCCompilationUnit()
	return res
}
