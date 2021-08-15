package lexical

import (
	"bufio"
	"container/list"
	"fmt"
	_util "husd.com/v0/util"
	"io"
	"os"
)

// WordLexicalV2 非常简单的程序，把字符串分割成单词，这里已经要开始处理文件和目录了。
type WordLexicalV2 struct{}

// LexicalAnalysis 这个版本的词法分析器，想支持一次编译一个目录下的所有文件，但是设计起来不太好放到这里
//所以先暂停这个方法的开发，先开发解析单个文件的代码。
func (a *WordLexicalV2) LexicalAnalysis(name string) *list.List {

	lexer := WordLexical{}
	files := _util.Walk(name)
	for i := files.Front(); i != nil; i = i.Next() {
		path := fmt.Sprintf("%v", i.Value)
		fd, err := os.Open(path)
		defer fd.Close()
		if err != nil {
			panic("读取文件失败 " + path)
		}
		buf := bufio.NewReader(fd)
		for {
			data, _, eof := buf.ReadLine()
			if eof == io.EOF {
				break
			}
			tokens := lexer.LexicalAnalysis(string(data))
			for i := tokens.Front(); i != nil; i = i.Next() {
				fmt.Println(i.Value)
			}
		}
	}
	return list.New()
}

func (a *WordLexicalV2) TokenTag(str string) int {

	// 由于只是简单的切割单词，所以不需要实现这个方法
	return 0
}
