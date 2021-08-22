package parser

import (
	"fmt"
	jcIo "husd.com/v0/io"
	"husd.com/v0/tree"
)

type JavacParser struct {
	sequence *jcIo.CharSequence //注意这里定义的是一个接口，而不是一个string
	lex      lexer              // 词法分析器
}

func NewJavacParser(sequence *jcIo.CharSequence) JavacParser {

	parser := JavacParser{}
	parser.sequence = sequence
	parser.lex = GetScannerLexerFromFactory(sequence)
	return parser
}

func (javaParser JavacParser) ParseJCCompilationUnit() tree.JCCompilationUnit {

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
