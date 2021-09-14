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
	c            *util.Context //
	ScannerLexer lexer         // 词法分析器
	source       code.JVersion // 当前JDK的版本
	token        Token

	endPosTable *SimpleEndPosTable

	TreeMaker   *jc.AstTreeMaker
	names       *util.Names
	symbolTable *SymbolTable

	mode     parseMode // 当前正在进行转换的模式
	lastMode parseMode // 上一个模式 lastMode = 2 then mode = 1  from 1 to 2
}

func NewJavacParser(path string, context *util.Context) *JavacParser {

	parser := JavacParser{}
	parser.ScannerLexer = GetScannerLexerFromFactory(path, context)
	parser.nextToken()

	parser.endPosTable = NewSimpleEndPosTable(&parser)
	parser.TreeMaker = jc.InstanceAstTreeMaker(context)
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
func (jp *JavacParser) ParseJCCompilationUnit() *jc.JCCompilationUnit {

	// for {
	// 	tok := jp.token
	// 	jp.symbolTable.PutToken(tok)
	// 	if compiler.DEBUG_TOKEN {
	// 		fmt.Println(tok.DebugToString())
	// 	}
	// 	if tok.GetTokenKind() == TOKEN_KIND_EOF {
	// 		break
	// 	}
	// 	jp.nextToken()
	// }
	jp.symbolTable.GetTokenByIndex(1000)
	seenPackage := false
	seenImport := false
	// firstToken := jp.Token
	var pid *jc.JCExpression
	var mods *jc.JCModifiers
	packageAnnotations := make([]jc.JCAnnotation, 0, 10)

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

	jcTreeList := make([]*jc.JCTree, 0, 100)
	checkForImports := true
	//firstTypeDecl := true

	for jp.token.GetTokenKind() != TOKEN_KIND_EOF {

		if jp.token.Pos() > 0 &&
			jp.token.Pos() <= jp.endPosTable.errorEndPos {
			jp.skip(checkForImports, false, false, false)
			if jp.token.GetTokenKind() == TOKEN_KIND_EOF {
				break
			}
		}
		if checkForImports &&
			mods == nil &&
			jp.token.GetTokenKind() == TOKEN_KIND_IMPORT {
			seenImport = true
			jcTreeList = append(jcTreeList, jp.importDeclaration())
		} else {
			def := jp.typeDeclaration(mods)
			jcTreeList = append(jcTreeList, def)
		}
	}

	// 防止报错才加的打印 todo delete later
	fmt.Println(pid)
	fmt.Println(seenPackage)
	fmt.Println(seenImport)
	fmt.Println(packageAnnotations)
	return &jc.JCCompilationUnit{}
}

type parseMode int

/** When terms are parsed, the mode determines which is expected:
 *     mode = EXPR        : an expression
 *     mode = TYPE        : a type
 *     mode = NOPARAMS    : no parameters allowed for type
 *     mode = TYPEARG     : type argument
 */
const mode_expr parseMode = 0x1
const mode_type parseMode = 0x2
const mode_noparams parseMode = 0x4
const mode_typearg parseMode = 0x8
const mode_diamond parseMode = 0x10

func (jp *JavacParser) ParseExpression() *jc.JCExpression {
	panic("implement me")
}

func (jp *JavacParser) termWithMode(newMode parseMode) *jc.JCExpression {

	preMode := jp.mode
	jp.mode = newMode
	t := jp.term()
	jp.lastMode = jp.mode
	jp.mode = preMode
	return t
}

/** Statement =
 *       Block
 *     | IF ParExpression Statement [ELSE Statement]
 *     | FOR "(" ForInitOpt ";" [Expression] ";" ForUpdateOpt ")" Statement
 *     | FOR "(" FormalParameter : Expression ")" Statement
 *     | WHILE ParExpression Statement
 *     | DO Statement WHILE ParExpression ";"
 *     | TRY Block ( Catches | [Catches] FinallyPart )
 *     | TRY "(" ResourceSpecification ";"opt ")" Block [Catches] [FinallyPart]
 *     | SWITCH ParExpression "{" SwitchBlockStatementGroups "}"
 *     | SYNCHRONIZED ParExpression Block
 *     | RETURN [Expression] ";"
 *     | THROW Expression ";"
 *     | BREAK [Ident] ";"
 *     | CONTINUE [Ident] ";"
 *     | ASSERT Expression [ ":" Expression ] ";"
 *     | ";"
 *     | ExpressionStatement
 *     | Ident ":" Statement
 */
func (jp *JavacParser) ParseStatement() *jc.JCStatement {
	panic("implement me")
}

func (jp *JavacParser) ParseType() *jc.JCExpression {
	panic("implement me")
}

/** ModifiersOpt = { Modifier }
 *  Modifier = PUBLIC | PROTECTED | PRIVATE | STATIC | ABSTRACT | FINAL
 *           | NATIVE | SYNCHRONIZED | TRANSIENT | VOLATILE | "@"
 *           | "@" Annotation
 */
func (jp *JavacParser) modifiersOpt() *jc.JCModifiers {

	//todo 暂时先不解析注解
	return &jc.JCModifiers{}
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
func (jp *JavacParser) qualident(allowAnnotations bool) *jc.JCExpression {

	tk := jp.token
	expression := jp.toExpression(jp.TreeMaker.At(tk.Pos()).Identify(jp.ident()).JCExpression.JCTree)
	// 解析这个逗号
	for jp.token.GetTokenKind() == TOKEN_KIND_DOT {
		pos := jp.token.Pos()
		jp.nextToken() // 查看点之后是什么
		//var annotations []ast_tree.JCAnnotation
		// if allowAnnotations {
		//	annotations = typeAnnotationsOpt()
		// }
		// todo
		expression = jp.toExpression(jp.TreeMaker.At(pos).Select(expression, jp.ident()).GetExpression().GetJCTree())
		// 我们这里没有注解 todo annotation
	}
	return expression
}

/* ---------- parsing -------------- */

// Ident = IDENTIFIER
func (jp *JavacParser) ident() *util.Name {

	tk := jp.token
	if jp.token.GetTokenKind() == TOKEN_KIND_IDENTIFIER {
		name := tk.GetName()
		jp.nextToken()
		return name
	} else if jp.token.GetTokenKind() == TOKEN_KIND_ASSERT {
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
	} else if jp.token.GetTokenKind() == TOKEN_KIND_ENUM {
		if allowEnums {
			jp.error(jp.token.Pos(), "enum as identifier")
			jp.nextToken()
			return jp.names.Error
		} else {
			jp.warn(jp.token.Pos(), "enum as identifier")
			_name := jp.token.GetName()
			jp.nextToken()
			return _name
		}
	} else if jp.token.GetTokenKind() == TOKEN_KIND_THIS {
		if allowThisIdent {
			_name := jp.token.GetName()
			jp.nextToken()
			return _name
		} else {
			jp.error(jp.token.Pos(), "this as identifier")
			jp.nextToken()
			return jp.names.Error
		}
	} else if jp.token.GetTokenKind() == TOKEN_KIND_UNDERSCORE {
		jp.warn(jp.token.Pos(), "underscore.as.identifier")
		_name := jp.token.GetName()
		jp.nextToken()
		return _name
	} else {
		jp.accept(TOKEN_KIND_IDENTIFIER)
		return jp.names.Error
	}
}

func (jp *JavacParser) toExpression(t *jc.JCTree) *jc.JCExpression {

	return jp.endPosTable.toP(t)
}

func (jp *JavacParser) term() *jc.JCExpression {

	return &jc.JCExpression{}
}

/** Skip forward until a suitable stop token is found.
 */
func (jp *JavacParser) skip(stopAtImport bool, stopAtMemberDecl bool,
	stopAtIdentifier bool, stopAtStatement bool) {
	for {
		switch jp.token.GetTokenKind() {
		case
			TOKEN_KIND_SEMI:
			jp.nextToken()
			return
		case TOKEN_KIND_PUBLIC, TOKEN_KIND_FINAL, TOKEN_KIND_ABSTRACT,
			TOKEN_KIND_MONKEYS_AT, TOKEN_KIND_EOF, TOKEN_KIND_CLASS,
			TOKEN_KIND_INTERFACE, TOKEN_KIND_ENUM:
			return
		case TOKEN_KIND_IMPORT:
			if stopAtImport {
				return
			}
			break
		case TOKEN_KIND_LBRACE, TOKEN_KIND_RBRACE, TOKEN_KIND_PRIVATE,
			TOKEN_KIND_PROTECTED, TOKEN_KIND_STATIC, TOKEN_KIND_TRANSIENT,
			TOKEN_KIND_NATIVE, TOKEN_KIND_VOLATILE, TOKEN_KIND_SYNCHRONIZED,
			TOKEN_KIND_STRICTFP, TOKEN_KIND_LT, TOKEN_KIND_BYTE, TOKEN_KIND_SHORT,
			TOKEN_KIND_CHAR, TOKEN_KIND_INT, TOKEN_KIND_LONG, TOKEN_KIND_FLOAT,
			TOKEN_KIND_DOUBLE, TOKEN_KIND_BOOLEAN, TOKEN_KIND_VOID:
			if stopAtMemberDecl {
				return
			}
			break
		case TOKEN_KIND_UNDERSCORE, TOKEN_KIND_IDENTIFIER:
			if stopAtIdentifier {
				return
			}
			break
		case TOKEN_KIND_CASE, TOKEN_KIND_DEF, TOKEN_KIND_IF, TOKEN_KIND_FOR, TOKEN_KIND_WHILE,
			TOKEN_KIND_DO, TOKEN_KIND_TRY, TOKEN_KIND_SWITCH, TOKEN_KIND_RETURN, TOKEN_KIND_THROW, TOKEN_KIND_BREAK,
			TOKEN_KIND_CONTINUE, TOKEN_KIND_ELSE, TOKEN_KIND_FINALLY, TOKEN_KIND_CATCH:
			if stopAtStatement {
				return
			}
			break
		}
		jp.nextToken()
	}
}

/** ImportDeclaration = IMPORT [ STATIC ] Ident { "." Ident } [ "." "*" ] ";"
 * 这个是import的语法，这个语法应该比较好解析，固定的格式。
 */
func (jp *JavacParser) importDeclaration() *jc.JCTree {

	t := &jc.JCTree{}
	pos := jp.token.Pos()
	jp.nextToken()
	importStatic := false
	// 这里先允许 import static com.husd; 这样的语法
	if jp.token.GetTokenKind() == TOKEN_KIND_STATIC {
		importStatic = true
		jp.nextToken()
	}
	var pid *jc.JCExpression = &jc.JCExpression{}
	for {
		pos1 := jp.token.Pos()
		jp.accept(TOKEN_KIND_DOT)
		if jp.token.GetTokenKind() == TOKEN_KIND_STAR {
			pid = &jc.JCExpression{}
			pos1++
			jp.nextToken()
			break
		} else {
			pid = &jc.JCExpression{}
			pos1++
		}
		if jp.token.GetTokenKind() != TOKEN_KIND_DOT {
			break
		}
	}
	jp.accept(TOKEN_KIND_SEMI)
	if compiler.DEBUG {
		fmt.Println("import static", importStatic, " pid:", pid, pos)
	}
	return t
}

func (jp *JavacParser) typeDeclaration(mods *jc.JCModifiers) *jc.JCTree {

	t := &jc.JCTree{}

	return t
}
