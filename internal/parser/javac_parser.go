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
	c      *util.Context //
	S      lexer         // 词法分析器
	source code.JVersion // 当前JDK的版本
	token  Token         //当前读到的token

	endPosTable *SimpleEndPosTable

	F           *jc.AstTreeMaker
	names       *util.Names
	symbolTable *SymbolTable // 符号表

	allowGenerics bool // 是否允许泛型

	mode     parseMode // 当前正在进行转换的模式
	lastMode parseMode // 上一个模式 lastMode = 2 then mode = 1  from 1 to 2
}

func NewJavacParser(path string, context *util.Context) *JavacParser {

	parser := JavacParser{}
	parser.S = GetScannerLexerFromFactory(path, context)
	parser.nextToken()

	parser.endPosTable = NewSimpleEndPosTable(&parser)
	parser.F = jc.InstanceAstTreeMaker(context)
	parser.names = util.InstanceNames(context)
	parser.symbolTable = InstanceSymbolTable(context)
	parser.source = code.JDK8
	parser.allowGenerics = code.AllowGenerics(parser.source)

	return &parser
}

func NewJavacParserWithString(str string, context *util.Context) *JavacParser {

	parser := JavacParser{}
	parser.S = NewScannerLexerWithString(str, context)
	parser.nextToken()

	parser.endPosTable = NewSimpleEndPosTable(&parser)
	parser.F = jc.InstanceAstTreeMaker(context)
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
	lex := jp.S
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
	var pid *jc.AbstractJCExpression
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

	jcTreeList := make([]*jc.AbstractJCTree, 0, 100)
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

/**  条件有4种情况，分别是以下4种情况：
 *     mode = EXPR        : an expression
 *     mode = TYPE        : a type
 *     mode = NOPARAMS    : no parameters allowed for type
 *     mode = TYPEARG     : type argument
 */
const term_mode_expr parseMode = 0x1
const term_mode_type parseMode = 0x2
const term_mode_noparams parseMode = 0x4
const term_mode_typearg parseMode = 0x8
const term_mode_diamond parseMode = 0x10

func (jp *JavacParser) ParseExpression() *jc.AbstractJCExpression {

	// a > 10
	// a == 10 && b == 15
	// sum(1,5) == 20
	//
	// term有条件得意思 泛指Java里的表达式条件
	return jp.termWithMode(term_mode_expr)
}

func (jp *JavacParser) termWithMode(newMode parseMode) *jc.AbstractJCExpression {

	preMode := jp.mode
	jp.mode = newMode
	var t *jc.AbstractJCExpression //这么做仅仅为了表达以下t的类型
	t = jp.term()
	jp.lastMode = jp.mode //jp.lastMode =  newMode 是最终生效的mode
	jp.mode = preMode     //  jp.mode =
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
func (jp *JavacParser) ParseStatement() *jc.AbstractJCStatement {
	panic("implement me")
}

func (jp *JavacParser) ParseType() *jc.AbstractJCExpression {
	panic("implement me")
}

/**
 * Literal =
 *     INTLITERAL
 *   | LONGLITERAL
 *   | FLOATLITERAL
 *   | DOUBLELITERAL
 *   | CHARLITERAL
 *   | STRINGLITERAL
 *   | TRUE
 *   | FALSE
 *   | NULL
 */
func (jp *JavacParser) literal(pre *util.Name, pos int) *jc.AbstractJCExpression {

	var t *jc.AbstractJCExpression
	t = jc.NewJCError(pos, "默认错误")
	switch jp.token.GetTokenKind() {
	case TOKEN_KIND_INT_LITERAL:
		num, err := util.String2int(jp.token.GetStringVal(), jp.token.GetRadix(), 32)
		if err != nil {
			jp.error(jp.token.Pos(), err.Error())
		}
		literal := jp.F.At(pos).Literal(code.TYPE_TAG_INT, int(num))
		return literal.AbstractJCExpression
	}
	return t
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
func (jp *JavacParser) accept(tk tokenKind) {

	// 因为不能保证都是同一个对象，所以用tokenKind的索引比较是否是同一个tokenKind
	if jp.token.GetTokenKind() == tk {
		jp.nextToken()
	} else {
		jp.setErrorEndPos(jp.token.EndPos())
		jp.reportSyntaxError(jp.S.PreToken().EndPos(), "期望是", tk)
	}
}

func (jp *JavacParser) reportSyntaxError(pos int, msg string, tk tokenKind) {

	// TODO 暂时先打印，应该有更好的方式来报告语法错误
	// 发送一个事件，通知所有监听这个事件的程序来处理语法错误。
	fmt.Println("---------------- reportSyntaxError，位置：", pos, " msg:", msg, " tokenkind:", tk)
}

func (jp *JavacParser) setErrorEndPos(pos int) {

	jp.endPosTable.SetErrorPos(pos)
}

/**
 * Qualident = Ident { DOT [Annotations] Ident }
 * 这里先忽略注解的因素，就是要解析出来包 ，把解析出来的包，转换为 AbstractJCExpression
 * 例如： package com.husd;
 * 要读取 com.husd 分号不处理
 * 这里先不处理注解 todo Annotations
 */
func (jp *JavacParser) qualident(allowAnnotations bool) *jc.AbstractJCExpression {

	tk := jp.token
	expression := jp.toExpression(jp.F.At(tk.Pos()).Identify(jp.ident()))
	// 解析这个逗号
	for jp.token.GetTokenKind() == TOKEN_KIND_DOT {
		pos := jp.token.Pos()
		jp.nextToken() // 查看点之后是什么
		//var annotations []ast_tree.JCAnnotation
		// if allowAnnotations {
		//	annotations = typeAnnotationsOpt()
		// }
		// todo
		expression = jp.toExpression(jp.F.At(pos).Select(expression, jp.ident()).AbstractJCExpression)
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

func (jp *JavacParser) toExpression(t *jc.AbstractJCExpression) *jc.AbstractJCExpression {

	jp.endPosTable.toP(t.AbstractJCTree)
	return t
}

/**
 * 这里是整个表达式的BNF
 *  {@literal
 *  Expression = Expression1 [ExpressionRest]
 *  ExpressionRest = [AssignmentOperator Expression1]
 *  AssignmentOperator = "=" | "+=" | "-=" | "*=" | "/=" |
 *                       "&=" | "|=" | "^=" |
 *                       "%=" | "<<=" | ">>=" | ">>>="
 *  Type = Type1
 *  TypeNoParams = TypeNoParams1
 *  StatementExpression = Expression
 *  ConstantExpression = Expression
 *  }
 */
func (jp *JavacParser) term() *jc.AbstractJCExpression {

	e := jp.term1()
	if (jp.mode&term_mode_expr) != 0 &&
		jp.token.GetTokenKind() == TOKEN_KIND_EQ ||
		jp.token.GetTokenKind() >= TOKEN_KIND_PLUSEQ &&
			jp.token.GetTokenKind() <= TOKEN_KIND_GTGTGTEQ {
		return jp.termRest(e)
	} else {
		return e
	}
}

/** Expression1   = Expression2 [Expression1Rest]
 *  Type1         = Type2
 *  TypeNoParams1 = TypeNoParams2
 */
func (jp *JavacParser) term1() *jc.AbstractJCExpression {

	e := jp.term2()
	if (jp.mode&term_mode_expr) != 0 &&
		jp.token.GetTokenKind() == TOKEN_KIND_QUES {
		jp.mode = term_mode_expr
		return jp.term1Rest(e)
	} else {
		return e
	}
}

/** Expression2   = Expression3 [Expression2Rest]
 *  Type2         = Type3
 *  TypeNoParams2 = TypeNoParams3
 */
func (jp *JavacParser) term2() *jc.AbstractJCExpression {

	const OR_PREC int = 4
	e := jp.term3()
	if (jp.mode&term_mode_expr) != 0 &&
		prec(jp.token.GetTokenKind()) >= OR_PREC {
		jp.mode = term_mode_expr
		return jp.term2Rest(e, OR_PREC)
	} else {
		return e
	}
}

/**
 *  Expression3    = PrefixOp Expression3
 *                 | "(" Expr | TypeNoParams ")" Expression3
 *                 | Primary {Selector} {PostfixOp}
 *
 *  {@literal
 *  Primary        = "(" Expression ")"
 *                 | Literal
 *                 | [TypeArguments] THIS [Arguments]
 *                 | [TypeArguments] SUPER SuperSuffix
 *                 | NEW [TypeArguments] Creator
 *                 | "(" Arguments ")" "->" ( Expression | Block )
 *                 | Ident "->" ( Expression | Block )
 *                 | [Annotations] Ident { "." [Annotations] Ident }
 *                 | Expression3 MemberReferenceSuffix
 *                   [ [Annotations] "[" ( "]" BracketsOpt "." CLASS | Expression "]" )
 *                   | Arguments
 *                   | "." ( CLASS | THIS | [TypeArguments] SUPER Arguments | NEW [TypeArguments] InnerCreator )
 *                   ]
 *                 | BasicType BracketsOpt "." CLASS
 *  }
 *
 *  PrefixOp       = "++" | "--" | "!" | "~" | "+" | "-"
 *  PostfixOp      = "++" | "--"
 *  Type3          = Ident { "." Ident } [TypeArguments] {TypeSelector} BracketsOpt
 *                 | BasicType
 *  TypeNoParams3  = Ident { "." Ident } BracketsOpt
 *  Selector       = "." [TypeArguments] Ident [Arguments]
 *                 | "." THIS
 *                 | "." [TypeArguments] SUPER SuperSuffix
 *                 | "." NEW [TypeArguments] InnerCreator
 *                 | "[" Expression "]"
 *  TypeSelector   = "." Ident [TypeArguments]
 *  SuperSuffix    = Arguments | "." Ident [Arguments]
 */
func (jp *JavacParser) term3() *jc.AbstractJCExpression {

	pos := jp.token.Pos()
	var t *jc.AbstractJCExpression
	typeArgs := jp.typeArgumentsOpt(term_mode_expr)

	// 处理这样的范型表达式的，这里先忽略范型
	// TypeArguments  = "<" TypeArgument {"," TypeArgument} ">"
	// List<JCExpression> typeArgs = typeArgumentsOpt(EXPR);
	switch jp.token.GetTokenKind() {

	case TOKEN_KIND_INT_LITERAL, TOKEN_KIND_LONG_LITERAL, TOKEN_KIND_FLOAT_LITERAL,
		TOKEN_KIND_DOUBLE_LITERAL, TOKEN_KIND_CHAR_LITERAL, TOKEN_KIND_STRING_LITERAL,
		TOKEN_KIND_TRUE, TOKEN_KIND_FALSE, TOKEN_KIND_NULL: // 最简单的 boolean a = false;

		if (jp.mode & term_mode_expr) != 0 {
			jp.mode = term_mode_expr
			t = jp.literal(jp.names.Empty, jp.token.Pos())
		} else {
			return jp.illegal("无效的表达式")
		}

	case TOKEN_KIND_BYTE, TOKEN_KIND_SHORT, TOKEN_KIND_CHAR, TOKEN_KIND_INT,
		TOKEN_KIND_LONG, TOKEN_KIND_FLOAT, TOKEN_KIND_DOUBLE, TOKEN_KIND_BOOLEAN: // 最简单的 boolean a = false;

		emptyAnnotations := &[]jc.JCAnnotation{}
		primitiveTypeTree := jp.basicType()
		t = jp.bracketsSuffix(jp.bracketsOpt(primitiveTypeTree.AbstractJCExpression, emptyAnnotations))

	case TOKEN_KIND_UNDERSCORE, TOKEN_KIND_IDENTIFIER,
		TOKEN_KIND_ASSERT, TOKEN_KIND_ENUM: //

		//  ->  lambda表达式 如果前面一个token是 -> 表示接下来要解析的是lambda表达式
		if jp.termExpr() && jp.peekToken(TOKEN_KIND_ARROW) {
			return jp.lambdaExpressionOrStatement(false, false, pos)
		} else {
			t = jp.toP(jp.F.At(pos).Identify(jp.ident()))
		loop:
			for {
				pos := jp.token.Pos()
				annos := jp.typeAnnotationsOpt() // 注解无处不在，这里先不处理注解
				// need to report an error later if LBRACKET is for array
				// index access rather than array creation level 可以是 @Some [] ，如果是 @Some [1] 就是错误的
				if len(*annos) > 0 &&
					jp.token.GetTokenKind() != TOKEN_KIND_LBRACKET &&
					jp.token.GetTokenKind() != TOKEN_KIND_ELLIPSIS {
					return jp.illegal("无效的对数组的注解")
				}
				switch jp.token.GetTokenKind() {
				case TOKEN_KIND_LBRACKET: // [
					jp.nextToken()
					if jp.token.GetTokenKind() == TOKEN_KIND_RBRACKET {
						// 读到了 []
						fmt.Println("读到了 [] ignore ..................")
					} else {
						if jp.termExpr() {
							jp.mode = term_mode_expr
							t1 := jp.term()
							if len(*annos) > 0 {
								t = jp.illegal("无效的.............")
							}
							arrayAccess := jp.F.At(pos).Indexed(t, t1)
							jp.to(arrayAccess)
							t = arrayAccess.AbstractJCExpression
						}
					}
				case TOKEN_KIND_LPAREN: // (
				case TOKEN_KIND_DOT: // .
				case TOKEN_KIND_ELLIPSIS: // ... 多个参数
				case TOKEN_KIND_LT: // <
				default:
					break loop
				}
				goto loop
			}
		}
	default:
		jp.illegal("")
	}
	fmt.Println("delete after .... pos is:", pos)
	return jp.term3Rest(t, typeArgs)
}

func (jp *JavacParser) termExpr() bool {

	return (jp.mode & term_mode_expr) != 0
}

/**
 * 解析表达式的结果部分，例如： int a = 10; 就是解析10
 */
func (jp *JavacParser) termRest(t *jc.AbstractJCExpression) *jc.AbstractJCExpression {

	switch jp.token.GetTokenKind() {
	case TOKEN_KIND_EQ: // = 表示是赋值语句
		pos := jp.token.Pos()
		jp.nextToken()
		jp.mode = term_mode_expr
		t1 := jp.term() // 这里就是递归了 结果值也可以是一个表达式，例如 int a = sum(10);
		return jp.toP(jp.F.At(pos).Assign(t, t1).AbstractJCExpression)
	case TOKEN_KIND_PLUSEQ, // ++ -- 这样的，例如 a++
		TOKEN_KIND_SUBEQ,
		TOKEN_KIND_STAREQ,
		TOKEN_KIND_SLASHEQ,
		TOKEN_KIND_PERCENTEQ,
		TOKEN_KIND_AMPEQ,
		TOKEN_KIND_BAREQ,
		TOKEN_KIND_CARETEQ,
		TOKEN_KIND_LTLTEQ,
		TOKEN_KIND_GTGTEQ,
		TOKEN_KIND_GTGTGTEQ:
		pos := jp.token.Pos()
		tk := jp.token.GetTokenKind()
		jp.nextToken()
		jp.mode = term_mode_expr
		t1 := jp.term()
		return jp.F.At(pos).Assignop(opTag(tk), t, t1).AbstractJCExpression
	default:
		return t
	}
}

/** Expression1Rest = ["?" Expression ":" Expression1]
 */
func (jp *JavacParser) term1Rest(t *jc.AbstractJCExpression) *jc.AbstractJCExpression {

	if jp.token.GetTokenKind() == TOKEN_KIND_QUES {
		pos := jp.token.Pos()
		jp.nextToken()
		t1 := jp.term()
		jp.accept(TOKEN_KIND_COLON) // 期望下一个字符是冒号 (:)
		t2 := jp.term1()
		return jp.F.At(pos).Conditional(t, t1, t2).AbstractJCExpression
	} else {
		return t
	}
}

func (jp *JavacParser) term2Rest(t *jc.AbstractJCExpression, perc int) *jc.AbstractJCExpression {

	//TODO
	return t
}

func prec(tk tokenKind) int {

	treeTag := opTag(tk)
	if treeTag != jc.TREE_TAG_NO_TAG {
		return jc.OpPrec(treeTag)
	} else {
		return -1
	}
}

func opTag(tk tokenKind) jc.JCTreeTag {

	switch tk {
	case TOKEN_KIND_BARBAR:
		return jc.TREE_TAG_OR
	case TOKEN_KIND_AMPAMP:
		return jc.TREE_TAG_AND
	case TOKEN_KIND_BAR:
		return jc.TREE_TAG_BITOR
	case TOKEN_KIND_BAREQ:
		return jc.TREE_TAG_BITOR_ASG
	case TOKEN_KIND_CARET:
		return jc.TREE_TAG_BITXOR
	case TOKEN_KIND_CARETEQ:
		return jc.TREE_TAG_BITXOR_ASG
	case TOKEN_KIND_AMP:
		return jc.TREE_TAG_BITAND
	case TOKEN_KIND_AMPEQ:
		return jc.TREE_TAG_BITAND_ASG
	case TOKEN_KIND_EQEQ:
		return jc.TREE_TAG_EQ
	case TOKEN_KIND_BANGEQ:
		return jc.TREE_TAG_NE
	case TOKEN_KIND_LT:
		return jc.TREE_TAG_LT
	case TOKEN_KIND_GT:
		return jc.TREE_TAG_GT
	case TOKEN_KIND_LTEQ:
		return jc.TREE_TAG_LE
	case TOKEN_KIND_GTEQ:
		return jc.TREE_TAG_GE
	case TOKEN_KIND_LTLT:
		return jc.TREE_TAG_SL
	case TOKEN_KIND_LTLTEQ:
		return jc.TREE_TAG_SL_ASG
	case TOKEN_KIND_GTGT:
		return jc.TREE_TAG_SR
	case TOKEN_KIND_GTGTEQ:
		return jc.TREE_TAG_SR_ASG
	case TOKEN_KIND_GTGTGT:
		return jc.TREE_TAG_USR
	case TOKEN_KIND_GTGTGTEQ:
		return jc.TREE_TAG_USR_ASG
	case TOKEN_KIND_PLUS:
		return jc.TREE_TAG_PLUS
	case TOKEN_KIND_PLUSEQ:
		return jc.TREE_TAG_PLUS_ASG
	case TOKEN_KIND_SUB:
		return jc.TREE_TAG_MINUS
	case TOKEN_KIND_SUBEQ:
		return jc.TREE_TAG_MINUS_ASG
	case TOKEN_KIND_STAR:
		return jc.TREE_TAG_MUL
	case TOKEN_KIND_STAREQ:
		return jc.TREE_TAG_MUL_ASG
	case TOKEN_KIND_SLASH:
		return jc.TREE_TAG_DIV
	case TOKEN_KIND_SLASHEQ:
		return jc.TREE_TAG_DIV_ASG
	case TOKEN_KIND_PERCENT:
		return jc.TREE_TAG_MOD
	case TOKEN_KIND_PERCENTEQ:
		return jc.TREE_TAG_MOD_ASG
	case TOKEN_KIND_INSTANCEOF:
		return jc.TREE_TAG_TYPETEST
	default:
		return jc.TREE_TAG_NO_TAG
	}
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
func (jp *JavacParser) importDeclaration() *jc.AbstractJCTree {

	t := &jc.AbstractJCTree{}
	pos := jp.token.Pos()
	jp.nextToken()
	importStatic := false
	// 这里先允许 import static com.husd; 这样的语法
	if jp.token.GetTokenKind() == TOKEN_KIND_STATIC {
		importStatic = true
		jp.nextToken()
	}
	var pid *jc.AbstractJCExpression = &jc.AbstractJCExpression{}
	for {
		pos1 := jp.token.Pos()
		jp.accept(TOKEN_KIND_DOT)
		if jp.token.GetTokenKind() == TOKEN_KIND_STAR {
			pid = &jc.AbstractJCExpression{}
			pos1++
			jp.nextToken()
			break
		} else {
			pid = &jc.AbstractJCExpression{}
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

func (jp *JavacParser) typeDeclaration(mods *jc.JCModifiers) *jc.AbstractJCTree {

	t := &jc.AbstractJCTree{}
	return t
}

func (jp *JavacParser) illegal(msg string) *jc.AbstractJCExpression {

	jp.reportSyntaxError(jp.token.Pos(), msg, jp.token.GetTokenKind())
	return jp.syntaxError(jp.token.Pos(), msg)
}

func (jp *JavacParser) syntaxError(pos int, msg string) *jc.AbstractJCExpression {

	jp.F.At(jp.token.Pos())
	err := jc.NewJCError(pos, msg)
	return err
}

func (jp *JavacParser) basicType() *jc.JCPrimitiveTypeTree {

	jp.F.At(jp.token.Pos())
	tree := jp.F.TypeIdent(typeTag(jp.token.GetTokenKind()))
	jp.nextToken()
	return tree
}

/**
 * 解析方括号里的内容
 * BracketsOpt = [ "[" "]" { [Annotations] "[" "]"} ]
 *
 * 考虑这样的代码
 * void m(String [] m) { }
 * void m(String ... m) { }
 * void m(String @A [] m) { }
 * void m(String @A ... m) { }
 */
func (jp *JavacParser) bracketsOpt(expression *jc.AbstractJCExpression, annotations *[]jc.JCAnnotation) *jc.AbstractJCExpression {

	// nextLevelAnnotations := jp.typeAnnotationsOpt()
	// 这里我们不处理注解，所以先返回空
	return expression
}

/**
 * 要解析出来注解 ，这里我们暂时不支持注解，先忽略，返回空 TODO annotation
 */
func (jp *JavacParser) typeAnnotationsOpt() *[]jc.JCAnnotation {

	return &[]jc.JCAnnotation{}
}

/** BracketsSuffixExpr = "." CLASS
 *  BracketsSuffixType =
 *
 *
 *
 * TODO 先不处理
 */
func (jp *JavacParser) bracketsSuffix(opt *jc.AbstractJCExpression) *jc.AbstractJCExpression {

	if (jp.mode&term_mode_expr) != 0 &&
		jp.token.GetTokenKind() == TOKEN_KIND_DOT {
		jp.mode = term_mode_expr
		// newPos := jp.token.Pos()
		jp.nextToken()
		jp.accept(TOKEN_KIND_CLASS)
		// TODO
	} else if (jp.mode & term_mode_type) != 0 {

	} else if jp.token.GetTokenKind() != TOKEN_KIND_COLCOL {
		jp.syntaxError(jp.token.Pos(), "期望.class")
	}
	return opt
}

/**
 * 向前看0个token，是不是指定的token，是就返回true
 */
func (jp *JavacParser) peekToken(tk tokenKind) bool {

	lookahead := 0
	return jp.peekTokenLookahead(lookahead, tk)
}

/**
 * 向前看指定数量个token，是不是指定的token，是就返回true
 */
func (jp *JavacParser) peekTokenLookahead(lookahead int, tk tokenKind) bool {

	return AcceptTokenKind(tk, jp.S.LookAheadByIndex(lookahead+1).GetTokenKind())
}

// 暂时不实现lambda表达式 todo lambda
func (jp *JavacParser) lambdaExpressionOrStatement(hasParens bool, explicitParams bool, pos int) *jc.AbstractJCExpression {

	panic("implement me lambda")
}

//记录 pos 到 endPosTable
func (jp *JavacParser) toP(expr *jc.AbstractJCExpression) *jc.AbstractJCExpression {

	jp.endPosTable.toP(expr.AbstractJCTree)
	return expr
}

func (jp *JavacParser) to(expr *jc.JCArrayAccess) interface{} {

	jp.endPosTable.toP(expr.AbstractJCTree)
	return expr
}

/**
 * 泛型 先不支持 TODO 泛型
 */
func (jp *JavacParser) typeArgumentsOpt(pm parseMode) *[]jc.AbstractJCExpression {

	if jp.token.GetTokenKind() == TOKEN_KIND_LT {
		jp.checkGenerics()
		if (jp.mode&pm) == 0 ||
			(jp.mode&term_mode_noparams) != 0 {
			jp.illegal("")
		}
		jp.mode = pm
		return typeArguments(false)
	}
	return &[]jc.AbstractJCExpression{}
}

/**
 * TypeArguments  = "<" TypeArgument {"," TypeArgument} ">"
 * 目前先不支持
 */
func typeArguments(b bool) *[]jc.AbstractJCExpression {

	return &[]jc.AbstractJCExpression{}
}

// 检查是否支持泛型
func (jp *JavacParser) checkGenerics() {

	if jp.allowGenerics == false {
		jp.error(jp.token.Pos(), "不支持泛型")
	}
}

func (jp *JavacParser) term3Rest(t *jc.AbstractJCExpression, args *[]jc.AbstractJCExpression) *jc.AbstractJCExpression {

	//TODO
	return t
}

// 返回none就是没有类型
func typeTag(tk tokenKind) *code.TypeTag {

	switch tk {
	case TOKEN_KIND_BYTE:
		return code.TYPE_TAG_BYTE
	case TOKEN_KIND_CHAR:
		return code.TYPE_TAG_CHAR
	case TOKEN_KIND_SHORT:
		return code.TYPE_TAG_SHORT
	case TOKEN_KIND_INT:
		return code.TYPE_TAG_INT
	case TOKEN_KIND_LONG:
		return code.TYPE_TAG_LONG
	case TOKEN_KIND_FLOAT:
		return code.TYPE_TAG_FLOAT
	case TOKEN_KIND_DOUBLE:
		return code.TYPE_TAG_DOUBLE
	case TOKEN_KIND_BOOLEAN:
		return code.TYPE_TAG_BOOLEAN
	default:
		return code.TYPE_TAG_NONE
	}
}
