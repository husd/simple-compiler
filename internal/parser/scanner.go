package parser

import (
	"container/list"
	"fmt"
	"husd.com/v0/io"
)

//Scanner 这里先不管处理unicode的问题
type Scanner struct {
	token    *Token
	preToken *Token
	seq      *io.CharSequence
	length   int // 总长度
	// ahead token list
	tokenList *list.List
}

func NewScannerLexer(seq *io.CharSequence) Scanner {

	scanner := Scanner{}
	scanner.tokenList = list.New()
	scanner.seq = seq
	//设置了一个假的节点
	dummy := dummyToken()
	scanner.token = dummy
	scanner.preToken = dummy

	scanner.length = (*seq).Len()
	return scanner
}

func (scan Scanner) NextToken() {
	fmt.Println(scan.token.StartPos)
	panic("implement me")
}

func (scan Scanner) CurrentToken() *Token {

	fmt.Println("implement me")
	return dummyToken()
}

func (scan Scanner) Ahead(len int) *Token {
	fmt.Println("implement me")
	return dummyToken()
}

func (scan Scanner) PreToken() *Token {
	fmt.Println("implement me")
	return dummyToken()
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
