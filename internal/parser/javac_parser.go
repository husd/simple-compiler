package parser

import (
	"fmt"
	"husd.com/v0/code"
	"husd.com/v0/compiler"
	"husd.com/v0/jc"
	"husd.com/v0/util"
)

/**
 *
 * @author hushengdong
 */

type JavacParser struct {
	lex    lexer         // 词法分析器
	source code.JVersion // 当前JDK的版本
	token  token
}

func NewJavacParser(path string, context *util.Context) *JavacParser {

	parser := JavacParser{}
	parser.lex = GetScannerLexerFromFactory(path, context)
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
	for {
		tok := jp.token
		if compiler.DEBUG_TOKEN {
			fmt.Println(tok.DebugToString())
		}
		if tok.GetTokenKind() == TOKEN_KIND_EOF {
			break
		}
		jp.nextToken()
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
