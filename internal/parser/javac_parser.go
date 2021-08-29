package parser

import (
	"fmt"
	"husd.com/v0/code"
	"husd.com/v0/jc"
)

type JavacParser struct {
	lex    lexer         // 词法分析器
	source code.JVersion // 当前JDK的版本
	token  token
}

func NewJavacParser(path string) *JavacParser {

	parser := JavacParser{}
	parser.lex = GetScannerLexerFromFactory(path)
	parser.nextToken()
	return &parser
}

// ----------------- token 相关的方法
func (jp *JavacParser) currentToken() token {

	return jp.token
}

// 设置下一个token
func (jp *JavacParser) nextToken() {
	lex := jp.lex
	lex.NextToken()
	jp.token = lex.Token()
}

// ----------------- token 相关的方法

//core function
func (jp *JavacParser) ParseJCCompilationUnit() jc.JCCompilationUnit {

	//seenImport := false
	//seenPackage := false
	//consumedToplevelDoc := false
	lex := jp.lex
	tok := lex.Token()
	fmt.Println("current token is : ", tok.GetTokenKind())

	for jp.token != nil {
		jp.nextToken()
		fmt.Println("current token is : ", jp.token)
	}

	if tok.GetTokenKind() == TOKEN_KIND_PACKAGE {

	}

	return jc.JCCompilationUnit{}
}

func (jp *JavacParser) ParseExpression() jc.JCExpression {
	panic("implement me")
}

func (jp *JavacParser) ParseStatement() jc.JCStatement {
	panic("implement me")
}

func (jp *JavacParser) ParseType() jc.JCExpression {
	panic("implement me")
}
