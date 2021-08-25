package parser

import (
	"container/list"
	"fmt"
)

//Scanner 在CharSequence里处理好unicode的事情
type Scanner struct {
	token    *Token
	preToken *Token
	// ahead token list
	tokenList     *list.List
	javaTokenizer *JavaTokenizer
}

func NewScannerLexer(path string) Scanner {

	scanner := Scanner{}
	scanner.tokenList = list.New()
	//设置了一个假的节点
	dummy := dummyToken()
	scanner.token = dummy
	scanner.preToken = dummy
	scanner.javaTokenizer = NewJavaTokenizer(path)

	return scanner
}

func (scan Scanner) NextToken() {

	// TODO husd
}

func (scan Scanner) CurrentToken() *Token {

	return scan.token
}

func (scan Scanner) Ahead(len int) *Token {
	fmt.Println("implement me")
	return dummyToken()
}

func (scan Scanner) PreToken() *Token {

	return scan.preToken
}

func (scan Scanner) ErrPos() int {
	fmt.Println("implement me")
	return 10
}

func (scan Scanner) SetErrPos(pos int) {
	fmt.Println("implement me")
}

func (scan Scanner) GetLineMap() *lineMap {
	fmt.Println("implement me")
	return NewLineMap()
}

func dummyToken() *Token {

	token := Token{}
	token.TokenKind = TOKEN_KIND_ERROR
	token.StartPos = 0
	token.EndPos = 0
	return &token
}
