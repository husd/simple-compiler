package parser

import (
	"container/list"
	"fmt"
	"husd.com/v0/util"
)

//Scanner 在CharSequence里处理好unicode的事情
type Scanner struct {
	token    Token
	preToken Token
	// ahead token list
	tokenList     *list.List
	javaTokenizer *JavaTokenizer
}

func InstanceScannerLexer(path string, c *util.Context) *Scanner {

	ok, obj := c.Get(util.C_LEXER)
	if ok {
		return obj.(*Scanner)
	}
	return NewScannerLexer(path, c)
}

func NewScannerLexer(path string, c *util.Context) *Scanner {

	scanner := Scanner{}
	scanner.tokenList = list.New()
	//设置了一个假的节点
	dummy := dummyToken()
	scanner.token = dummy
	scanner.preToken = dummy
	scanner.javaTokenizer = NewJavaTokenizer(path, c)

	c.Put(util.C_LEXER, &scanner)
	return &scanner
}

func NewScannerLexerWithString(str string, c *util.Context) *Scanner {

	scanner := Scanner{}
	scanner.tokenList = list.New()
	//设置了一个假的节点
	dummy := dummyToken()
	scanner.token = dummy
	scanner.preToken = dummy
	scanner.javaTokenizer = NewJavaTokenizerWithString(str, c)

	c.Put(util.C_LEXER, &scanner)
	return &scanner
}

func (scan *Scanner) NextToken() {

	scan.preToken = scan.token
	list := scan.tokenList
	if list.Len() > 0 {
		first := list.Front()
		list.Remove(first)
		scan.token = first.Value.(Token)
	} else {
		scan.token = scan.javaTokenizer.readToken()
	}
}

func (scan *Scanner) Token() Token {

	return scan.token
}

// 提前读取token
func (scan *Scanner) LookAhead() Token {

	return scan.LookAheadByIndex(0)
}

// 提前读取token
func (scan *Scanner) LookAheadByIndex(inx int) Token {

	if inx == 0 {
		return scan.token
	} else {
		scan.ensureLookahead(inx)
		return scan.tokenList.Back().Value.(Token)
	}
}

func (scan *Scanner) PreToken() Token {

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

func dummyToken() Token {

	token := newDefaultToken(TOKEN_KIND_ERROR, 0, 0, 0, 0)
	return token
}
