package lexical

import (
	"container/list"
	"fmt"
)

type DDLLexical struct{}

func (a *DDLLexical) lexicalAnalysis(str string) *list.List {

	res := list.List{}

	//遍历这个字符串
	for _, ch := range str {

		fmt.Println(ch)
	}

	return &res
}

func (a *DDLLexical) tokenTag(str string) int {

	return 0
}
