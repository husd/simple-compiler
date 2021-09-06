package parser

import (
	"fmt"
	"husd.com/v0/ast_tree"
	"husd.com/v0/code"
	"husd.com/v0/compiler"
	"husd.com/v0/util"
)

/**
 *
 * @author hushengdong
 */

type JavacParser struct {
	c            *util.Context //
	ScannerLexer lexer         // 词法分析器
	source       code.JVersion // 当前JDK的版本
	token        Token

	endPosTable *SimpleEndPosTable

	TreeMaker   *ast_tree.AstTreeMaker
	names       *util.Names
	symbolTable *SymbolTable
}

func NewJavacParser(path string, context *util.Context) *JavacParser {

	parser := JavacParser{}
	parser.ScannerLexer = GetScannerLexerFromFactory(path, context)
	parser.nextToken()

	parser.endPosTable = NewSimpleEndPosTable(&parser)
	parser.TreeMaker = ast_tree.InstanceAstTreeMaker(context)
	parser.names = util.InstanceNames(context)
	parser.symbolTable = InstanceSymbolTable(context)

	return &parser
}

// ----------------- Token 相关的方法
func (jp *JavacParser) currentToken() Token {

	return jp.token
}

// 设置下一个token
func (jp *JavacParser) nextToken() {
	lex := jp.ScannerLexer
	lex.NextToken()
	jp.token = lex.Token()
}

// ----------------- Token 相关的方法

//core function
func (jp *JavacParser) ParseJCCompilationUnit() ast_tree.JCCompilationUnit {

	//seenImport := false
	//consumedToplevelDoc := false
	for {
		tok := jp.token
		jp.symbolTable.PutToken(tok)
		if compiler.DEBUG_TOKEN {
			fmt.Println(tok.DebugToString())
		}
		if tok.GetTokenKind() == TOKEN_KIND_EOF {
			break
		}
		jp.nextToken()
	}
	jp.symbolTable.GetTokenByIndex(1000)
	seenPackage := false
	//firstToken := jp.Token
	var pid *ast_tree.JCExpression
	var mods *ast_tree.JCModifiers
	packageAnnotations := make([]ast_tree.JCAnnotation, 0, 10)

	/**
	 * 读到了 @ 这里的注解，指的是包的注释内容 package-info 例如：
	 * @java.lang.Deprecated
	 * package com.husd;
	 * 表示对包 com.husd;的注释内容 必须放在package的上方
	 */
	if jp.token.GetTokenKind() == TOKEN_KIND_MONKEYS_AT {
		mods = jp.modifiersOpt()
	}
	// 读到了 package 关键字
	if jp.token.GetTokenKind() == TOKEN_KIND_PACKAGE {
		seenPackage = true
		if mods != nil {
			jp.checkNoMods(mods.Flags)
			packageAnnotations = mods.Annotations
			mods = nil
		}
		jp.nextToken()
		pid = jp.qualident(false)
		jp.accept(TOKEN_KIND_SEMI) // 处理分号
	}

	//防止报错才加的打印 todo delete later
	fmt.Println(pid)
	fmt.Println(seenPackage)
	fmt.Println(packageAnnotations)
	return ast_tree.JCCompilationUnit{}
}

func (jp *JavacParser) ParseExpression() ast_tree.JCExpression {
	panic("implement me")
}

func (jp *JavacParser) ParseStatement() ast_tree.JCStatement {
	panic("implement me")
}

func (jp *JavacParser) ParseType() ast_tree.JCExpression {
	panic("implement me")
}

/** ModifiersOpt = { Modifier }
 *  Modifier = PUBLIC | PROTECTED | PRIVATE | STATIC | ABSTRACT | FINAL
 *           | NATIVE | SYNCHRONIZED | TRANSIENT | VOLATILE | "@"
 *           | "@" Annotation
 */
func (jp *JavacParser) modifiersOpt() *ast_tree.JCModifiers {

	//todo 暂时先不解析注解
	return &ast_tree.JCModifiers{}
}

func (jp *JavacParser) checkNoMods(flags int64) {

	if flags != 0 {
		lowestMod := flags & -flags
		jp.error(jp.token.Pos(), "mod.not.allowed.here", lowestMod)
	}
}

/* ---------- auxiliary methods -------------- *
 */

func (jp *JavacParser) error(bp int, msg ...interface{}) {

	fmt.Println("---------------- parse error，位置：", bp, " msg:", msg)
}

func (jp *JavacParser) warn(bp int, msg ...interface{}) {

	fmt.Println("---------------- parse warn，位置：", bp, " msg:", msg)
}

/**
 * 如果下一个token的tokenKind相同，那么就跳过去，否则报错，这个方法就是为了检查下一个token是否符合
 * 要求的，例如表达式要用分号结尾  int a = 10; 没有分号就报错
 */
func (jp *JavacParser) accept(tk *tokenKind) {

	//因为不能保证都是同一个对象，所以用tokenKind的索引比较是否是同一个tokenKind
	if jp.token.GetTokenKind() == tk ||
		jp.token.GetTokenKind().Index == tk.Index {
		jp.nextToken()
	} else {
		jp.setErrorEndPos(jp.token.EndPos())
		jp.reportSyntaxError(jp.ScannerLexer.PreToken().EndPos(), "期望是", tk)
	}
}

func (jp *JavacParser) reportSyntaxError(pos int, msg string, tk *tokenKind) {

	//TODO 暂时先打印，应该有更好的方式来报告语法错误
	// 发送一个事件，通知所有监听这个事件的程序来处理语法错误。
	fmt.Println("---------------- reportSyntaxError，位置：", pos, " msg:", msg, " tokenkind:", tk)
}

func (jp *JavacParser) setErrorEndPos(pos int) {

	jp.endPosTable.SetErrorPos(pos)
}

/**
 * Qualident = Ident { DOT [Annotations] Ident }
 * 这里先忽略注解的因素，就是要解析出来包 ，把解析出来的包，转换为 JCExpression
 * 例如： package com.husd;
 * 要读取 com.husd 分号不处理
 * 这里先不处理注解 todo Annotations
 */
func (jp *JavacParser) qualident(allowAnnotations bool) *ast_tree.JCExpression {

	tk := jp.token
	expression := jp.toExpression(jp.TreeMaker.At(tk.Pos()).Identify(jp.ident()).JcTree)
	// 解析这个逗号
	for jp.token.GetTokenKind() == TOKEN_KIND_DOT {
		pos := jp.token.Pos()
		jp.nextToken() // 查看点之后是什么
		//var annotations []ast_tree.JCAnnotation
		//if allowAnnotations {
		//	annotations = typeAnnotationsOpt()
		//}

		// todo
		expression = jp.toExpression(jp.TreeMaker.At(pos).Select(expression, jp.ident()).GetExpression().GetJCTree())
		//我们这里没有注解 todo annotation
	}
	return expression
}

//
func (jp *JavacParser) ident() *util.Name {

	tk := jp.token
	if tk.GetTokenKind() == TOKEN_KIND_IDENTIFIER {
		name := tk.GetName()
		jp.nextToken()
		return name
	} else if tk.GetTokenKind() == TOKEN_KIND_ASSERT {
		if allowAssert {
			jp.error(tk.Pos(), "错误的assert位置")
			jp.nextToken()
			return jp.names.Error
		} else {
			jp.warn(tk.Pos(), "assert.as.identifier")
			name := tk.GetName()
			jp.nextToken()
			return name
		}
	} else if tk.GetTokenKind() == TOKEN_KIND_ENUM {

	}
	//todo next
	return nil
}

func (jp *JavacParser) toExpression(t *ast_tree.JCTree) *ast_tree.JCExpression {

	return jp.endPosTable.toP(t)
}
