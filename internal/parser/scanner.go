package parser

import (
	"container/list"
	"fmt"
)

//Scanner 在CharSequence里处理好unicode的事情
type Scanner struct {
	token    token
	preToken token
	// ahead token list
	tokenList     *list.List
	javaTokenizer *JavaTokenizer
}

func NewScannerLexer(path string) *Scanner {

	scanner := Scanner{}
	scanner.tokenList = list.New()
	//设置了一个假的节点
	dummy := dummyToken()
	scanner.token = dummy
	scanner.preToken = dummy
	scanner.javaTokenizer = NewJavaTokenizer(path)

	return &scanner
}

func (scan *Scanner) NextToken() {

	scan.preToken = scan.token
	list := scan.tokenList
	if list.Len() > 0 {
		first := list.Front()
		list.Remove(first)
		scan.token = first.Value.(token)
	} else {
		scan.token = scan.javaTokenizer.readToken()
	}
}

func (scan *Scanner) Token() token {

	return scan.token
}

// 提前读取token
func (scan *Scanner) LookAhead() token {

	return scan.LookAheadByIndex(0)
}

// 提前读取token
func (scan *Scanner) LookAheadByIndex(inx int) token {

	if inx == 0 {
		return scan.token
	} else {
		scan.ensureLookahead(inx)
		return scan.tokenList.Back().Value.(token)
	}
}

func (scan *Scanner) Ahead(len int) token {
	fmt.Println("implement me")
	dummy := dummyToken()
	return dummy
}

func (scan *Scanner) PreToken() token {

	return scan.preToken
}

func (scan *Scanner) ErrPos() int {
	fmt.Println("implement me")
	return 10
}

func (scan *Scanner) SetErrPos(pos int) {
	fmt.Println("implement me")
}

func (scan *Scanner) GetLineMap() *lineMap {
	fmt.Println("implement me")
	return NewLineMap()
}

func (scan *Scanner) ensureLookahead(lookahead int) {

	savedTokens := scan.tokenList
	for i := savedTokens.Len(); i < lookahead; i++ {
		savedTokens.PushBack(scan.javaTokenizer.readToken())
	}
}

func dummyToken() token {

	token := newDefaultToken(TOKEN_KIND_ERROR, 0, 0)
	return token
}
