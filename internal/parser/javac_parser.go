package parser

import (
	"fmt"
	"husd.com/v0/code"
	"husd.com/v0/tree"
)

type JavacParser struct {
	lex    lexer         // 词法分析器
	source code.JVersion // 当前JDK的版本
	token  *Token
}

func NewJavacParser(path string) JavacParser {

	parser := JavacParser{}
	parser.lex = GetScannerLexerFromFactory(path)
	return parser
}

// ----------------- token 相关的方法
func (javaParser JavacParser) currentToken() *Token {

	return javaParser.token
}

// 设置下一个token
func (javaParser JavacParser) nextToken() {

	javaParser.lex.NextToken()
	javaParser.token = javaParser.lex.CurrentToken()
}

// ----------------- token 相关的方法

//core function
func (javaParser JavacParser) ParseJCCompilationUnit() tree.JCCompilationUnit {

	//seenImport := false
	//seenPackage := false
	lex := javaParser.lex
	currentToken := lex.CurrentToken()
	fmt.Println(currentToken)

	//TODO
	return tree.JCCompilationUnit{}
}

func (javaParser JavacParser) ParseExpression() tree.JCExpression {
	panic("implement me")
}

func (javaParser JavacParser) ParseStatement() tree.JCStatement {
	panic("implement me")
}

func (javaParser JavacParser) ParseType() tree.JCExpression {
	panic("implement me")
}
