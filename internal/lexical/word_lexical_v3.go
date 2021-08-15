package lexical

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
)

// WordLexicalV3 非常简单的程序，把字符串分割成单词，这里已经要开始处理文件和目录了。
type WordLexicalV3 struct{}

// LexicalAnalysis 解析单个文件 按行读取单个文件，把解析到的数据，存入到链表里。
//
func (a *WordLexicalV3) LexicalAnalysis(path string) *list.List {

	res := list.New()

	lexer := WordLexical{}
	fd, err := os.Open(path)
	defer fd.Close()
	if err != nil {
		panic("读取文件失败 " + path)
	}
	buf := bufio.NewReader(fd)
	lineNum := 0

	// 多行注释的标记，如果正处在多行注释中，那么就需要忽略这行的内容
	commentFlag := false
	for {
		data, _, eof := buf.ReadLine()
		if eof == io.EOF {
			break
		}
		lineNum++
		line := string(data)
		if commentFlag {
			if commentEnd(line) {
				commentFlag = false
			} else {
				c := 100 //注释中的话，只要不是结束，就忽略当前行
				b := 12  /**
				 *
				 */a := 12
				fmt.Println("a is :")
				b++
				a++
				c++
				continue
			}
		} else {
			//判断是不是多行注释的
			if commentStart(line) {
				commentFlag = true
			}
		}

		lexer.LexicalAnalysis(line)
	}

	return res
}

func (a *WordLexicalV3) TokenTag(str string) int {

	// 由于只是简单的切割单词，所以不需要实现这个方法
	return 0
}

// 多上注释开头 /*
func commentStart(str string) bool {

	length := len(str)
	start := 0
	//去掉了开头的空格
	for ; start < length && endChar(str[start]); start++ {
	}
	//判断是不是单行注释 //
	if start+1 <= length-1 && str[start] == '/' && str[start+1] == '*' {
		//如果是单行的注释，那么忽略
		return true
	}
	return false
}

// 多上注释结尾 /*
func commentEnd(str string) bool {

	length := len(str)
	start := 0
	//去掉了开头的空格
	for ; start < length && endChar(str[start]); start++ {
	}
	//判断是不是单行注释 //
	if start+1 <= length-1 && str[start] == '*' && str[start+1] == '/' {
		//如果是单行的注释，那么忽略
		return true
	}
	return false
}
