package parser

import (
	"fmt"
	"husd.com/v0/code"
	"husd.com/v0/compiler"
	"husd.com/v0/util"
)

/**
 *
 * @author hushengdong
 */

type JavacParser struct {
	c         *util.Context //
	S         lexer         // 词法分析器
	source    code.JVersion // 当前JDK的版本
	token     Token         // 当前读到的token
	tk        TokenKind     // 当前读到的token的TOKEN_KIND
	rowNum    int           // 多少行
	columnNum int           // 列

	endPosTable *SimpleEndPosTable

	names       *util.Names
	symbolTable *SymbolTable // 符号表

	allowGenerics bool // 是否允许泛型

	/**
	*如果解析器注意到额外的注释，那么它会立即
	*发出错误（如果此变量为false）或放置额外的
	*变量类型AnnotationSpushedBack中的注释（如果此变量
	*是真的）。
	 */
	permitTypeAnnotationsPushBack bool
	/** Type annotations that have already been read but have not yet been used. **/
	typeAnnotationsPushedBack *[]*TreeNode

	mode     parseMode // 当前正在进行转换的模式
	lastMode parseMode // 上一个模式 lastMode = 2 then mode = 1  from 1 to 2
}

func NewJavacParser(path string, context *util.Context) *JavacParser {

	parser := JavacParser{}
	parser.S = GetScannerLexerFromFactory(path, context)
	parser.nextToken()

	parser.endPosTable = NewSimpleEndPosTable(&parser)
	parser.names = util.InstanceNames(context)
	parser.symbolTable = InstanceSymbolTable(context)
	parser.source = code.JDK8
	parser.allowGenerics = code.AllowGenerics(parser.source)
	parser.permitTypeAnnotationsPushBack = false

	return &parser
}

func NewJavacParserWithString(str string, context *util.Context) *JavacParser {

	parser := JavacParser{}
	parser.S = NewScannerLexerWithString(str, context)
	parser.nextToken()

	parser.endPosTable = NewSimpleEndPosTable(&parser)
	parser.names = util.InstanceNames(context)
	parser.symbolTable = InstanceSymbolTable(context)
	parser.permitTypeAnnotationsPushBack = false

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
	// fmt.Println("----------------next token:", lex.Token())
	jp.token = lex.Token()
	jp.tk = jp.token.GetTokenKind()
	jp.rowNum = jp.token.GetRowNum()
	jp.columnNum = jp.token.GetColumnNum()
}

// ----------------- Token 相关的方法

//core function
func (jp *JavacParser) ParseJCCompilationUnit() *TreeNode {

	return GetEmptyTreeNode()
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
const term_mode_diamond parseMode = 0x10 // 这个就是 2 + 8 = 10

func (jp *JavacParser) ParseExpression() *TreeNode {

	// a > 10
	// a == 10 && b == 15
	// sum(1,5) == 20
	//
	// term有条件得意思 泛指Java里的表达式条件
	return jp.termWithMode(term_mode_expr)
	//return jp.term01()
}

func (jp *JavacParser) termWithMode(newMode parseMode) *TreeNode {

	preMode := jp.mode
	jp.mode = newMode
	var t *TreeNode // 这么做仅仅为了表达以下t的类型
	t = jp.term()
	jp.lastMode = jp.mode // 这个很容易理解 上一次的mode
	jp.mode = preMode     // 恢复之前的模式，临时离开preMode，在newMode下转换，在转换结束之后，再恢复到preMode
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
func (jp *JavacParser) ParseStatement() *TreeNode {

	res := GetEmptyTreeNode()
	// 先从一个简单的空语句开始转换，然后慢慢开始加case语句
	pos := jp.token.Pos()
	switch jp.tk {
	case SEMI: // ; 空语句
		jp.nextToken()
	case EOF: // EOF也是空语句
	case LBRACE: // {
		jp.nextToken()
		blockTree := NewBlockTreeNode(jp.token)
		jp.parseBlock(blockTree)
		res = blockTree
	case IF: // if 语句
		jp.nextToken()
		res = jp.parseIf()
	case IDENTIFIER: // 这里要解析 expression;
		res = jp.ParseExpression()
	default:
		// 其它情况都是错误的
		jp.reportSyntaxError(pos, "无效的token ", jp.tk)
	}
	// 空语句
	// jp.accept(common.SEMI)
	return res
}

func (jp *JavacParser) ParseType() *TreeNode {
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
func (jp *JavacParser) literal(pre *util.Name, pos int) *TreeNode {

	var res *TreeNode
	switch jp.tk {
	case INTLITERAL:
		num, err := util.String2int(jp.token.GetStringVal(), jp.token.GetRadix())
		if err != nil {
			jp.error(jp.token.Pos(), "int 类型数字太大溢出了")
		}
		res = NewLiteralTreeNode(jp.token, code.TYPE_TAG_INT, num, tt_int_literal)
	case LONGLITERAL:
		num, err := util.String2long(jp.token.GetStringVal(), jp.token.GetRadix())
		if err != nil {
			jp.error(jp.token.Pos(), "long 类型数字太大溢出了")
		}
		res = NewLiteralTreeNode(jp.token, code.TYPE_TAG_LONG, num, tt_long_literal)
	case TRUE:
		res = NewLiteralTreeNode(jp.token, code.TYPE_TAG_BOOLEAN, 1, tt_boolean_literal)
	case FALSE:
		res = NewLiteralTreeNode(jp.token, code.TYPE_TAG_BOOLEAN, 0, tt_boolean_literal)
	case NULL:
		res = NewLiteralTreeNode(jp.token, code.TYPE_TAG_NONE, nil, tt_null_literal)
	case STRINGLITERAL:
		res = NewLiteralTreeNode(jp.token, code.TYPE_TAG_CLASS, jp.token.GetStringVal(), tt_string_literal)
	case CHARLITERAL:
		res = NewLiteralTreeNode(jp.token, code.TYPE_TAG_CHAR, jp.token.GetStringVal()[0], tt_char_literal)
	default:
		res = NewErrorTreeNode(jp.token.Pos(), "不支持的常量类型")
	}
	jp.nextToken()
	return res
}

/** ModifiersOpt = { Modifier }
 *  Modifier = PUBLIC | PROTECTED | PRIVATE | STATIC | ABSTRACT | FINAL
 *           | NATIVE | SYNCHRONIZED | TRANSIENT | VOLATILE | "@"
 *           | "@" Annotation
 */
func (jp *JavacParser) modifiersOpt() *TreeNode {

	// todo 暂时先不解析注解
	return GetEmptyTreeNode()
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
func (jp *JavacParser) accept(tk TokenKind) {

	if jp.tk == tk {
		jp.nextToken()
	} else {
		jp.setErrorEndPos(jp.token.EndPos())
		jp.reportSyntaxError(jp.S.PreToken().EndPos(), "期望是", tk)
	}
}

/**
 * 适用于可选的后续节点 例如 if 语句，可以是：
 * if ( condition ) { statement }
 * 也可以是：
 * if ( condition ) statement
 * 这个时候 {} 就是可选的
 */
func (jp *JavacParser) acceptMaybe(tk TokenKind) bool {

	if jp.tk == tk {
		jp.nextToken()
		return true
	} else {
		return false
	}
}

// 是否是运算符号
func (jp *JavacParser) isOpCompare() bool {

	switch jp.tk {
	case EQEQ, BANGEQ,
		LTEQ, GTEQ,
		LT, GT:
		return true
	default:
		return false
	}
}

// 是否是运算符号
func (jp *JavacParser) isOp() bool {

	switch jp.tk {
	case EQEQ, BANGEQ,
		LTEQ, GTEQ,
		LT, GT:
		return true
	default:
		return false
	}
}

func (jp *JavacParser) reportSyntaxError(pos int, msg string, tk TokenKind) {

	// TODO 暂时先打印，应该有更好的方式来报告语法错误
	// 发送一个事件，通知所有监听这个事件的程序来处理语法错误。
	fmt.Println("---------------- reportSyntaxError，位置：", pos, " msg:", msg, " TokenKind:", tk)
	panic(fmt.Sprintf("------------ 语法错误 位置 %d msg : %s tokenKind:%v [%v]", pos, msg, tk, GetTokenString(tk)))
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
func (jp *JavacParser) qualident(allowAnnotations bool) *TreeNode {

	// tk := jp.token
	// expression := jp.toExpression(jp.F.At(tk.Pos()).Identify(jp.ident()))
	// // 解析这个逗号
	// for jp.tk == DOT {
	// 	pos := jp.token.Pos()
	// 	jp.nextToken() // 查看点之后是什么
	// 	//var annotations []ast_tree.JCAnnotation
	// 	// if allowAnnotations {
	// 	//	annotations = typeAnnotationsOpt()
	// 	// }
	// 	// todo
	// 	expression = jp.toExpression(jp.F.At(pos).Select(expression, jp.ident()).AbstractJCExpression)
	// 	// 我们这里没有注解 todo annotation
	// }
	return GetEmptyTreeNode()
}

/* ---------- parsing -------------- */

// Ident = IDENTIFIER
func (jp *JavacParser) ident() *util.Name {

	tk := jp.token
	if jp.tk == IDENTIFIER {
		name := tk.GetName()
		jp.nextToken()
		return name
	} else if jp.tk == ASSERT {
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
	} else if jp.tk == ENUM {
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
	} else if jp.tk == THIS {
		if allowThisIdent {
			_name := jp.token.GetName()
			jp.nextToken()
			return _name
		} else {
			jp.error(jp.token.Pos(), "this as identifier")
			jp.nextToken()
			return jp.names.Error
		}
	} else if jp.tk == UNDERSCORE {
		jp.warn(jp.token.Pos(), "underscore.as.identifier")
		_name := jp.token.GetName()
		jp.nextToken()
		return _name
	} else {
		jp.accept(IDENTIFIER)
		return jp.names.Error
	}
}

func (jp *JavacParser) toExpression(t *TreeNode) *TreeNode {

	// jp.endPosTable.toP(t.AbstractJCTree)
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
func (jp *JavacParser) term() *TreeNode {

	e := jp.term1()
	if (jp.mode&term_mode_expr) != 0 &&
		jp.tk == EQ ||
		jp.tk >= PLUSEQ &&
			jp.tk <= GTGTGTEQ {
		return jp.termRest(e)
	} else {
		return e
	}
}

/** Expression1   = Expression2 [Expression1Rest]
 *  Type1         = Type2
 *  TypeNoParams1 = TypeNoParams2
 */
func (jp *JavacParser) term1() *TreeNode {

	e := jp.term2()
	if (jp.mode&term_mode_expr) != 0 &&
		jp.tk == QUES { // ？号表达式
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
func (jp *JavacParser) term2() *TreeNode {

	e := jp.term3()
	if (jp.mode&term_mode_expr) != 0 &&
		prec(jp.tk) >= orPrec {
		jp.mode = term_mode_expr
		return jp.term2Rest(e, orPrec)
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
func (jp *JavacParser) term3() *TreeNode {

	pos := jp.token.Pos()
	var t *TreeNode
	// 处理这样的范型表达式的，这里先忽略范型
	// TypeArguments  = "<" TypeArgument {"," TypeArgument} ">"
	// List<JCExpression> typeArgs = typeArgumentsOpt(EXPR);
	typeArgs := jp.typeArgumentsOpt(term_mode_expr)
	switch jp.tk {
	case QUES: // ?
		// todo ? 表达式
	case PLUSPLUS, SUBSUB, BANG, TILDE, PLUS, SUB: // ++ -- ! ~ + -
		if typeArgs == nil && (jp.mode&term_mode_expr) != 0 { // 说明当前是 表达式模式
			preTk := jp.tk
			jp.nextToken()
			jp.mode = term_mode_expr
			if preTk == SUB &&
				(jp.tk == INTLITERAL || jp.tk == LONGLITERAL) &&
				jp.token.GetRadix() == 10 { // int 或者 long 类型的 例如： -10 -100L 这种是一个负数的常量，词法分析
				// 没有识别出来，在这里做的识别
				jp.mode = term_mode_expr
				t = jp.literal(jp.names.Hyphen, jp.token.Pos())
			} else {
				t = jp.term3()                          // 这里还是递归
				res := NewUnaryTreeNode(unOpTag(jp.tk)) // 一元操作符号 针对一个元素进行操作
				res.Append(t)
				return res
			}
		} else {
			return jp.illegal("无效的 ++ -- ! ~ + - 等表达式")
		}
	case LPAREN: // (
		if typeArgs == nil && (jp.mode&term_mode_expr) != 0 { // 说明当前是 表达式模式
			var pres parenthesesResult
			pres = jp.analyzeParens() // 分析下括号的作用
			switch pres {
			case CAST:
			case IMPLICIT_LAMBDA:
			case EXPLICIT_LAMBDA:
			default:
			}
		}
	case THIS: // this
	case SUPER: // super
	case INTLITERAL, LONGLITERAL, FLOATLITERAL,
		DOUBLELITERAL, CHARLITERAL, STRINGLITERAL,
		TRUE, FALSE, NULL: // 最简单的 boolean a = false;

		if (jp.mode & term_mode_expr) != 0 {
			jp.mode = term_mode_expr
			t = jp.literal(jp.names.Empty, jp.token.Pos())
		} else {
			return jp.illegal("无效的表达式")
		}
	case NEW:
	case MONKEYS_AT:
	case UNDERSCORE, IDENTIFIER, ASSERT, ENUM: //
		//  ->  lambda表达式 如果前面一个token是 -> 表示接下来要解析的是lambda表达式
		if jp.termExpr() && jp.peekToken(ARROW) {
			return jp.lambdaExpressionOrStatement(false, false, pos)
		} else {
			t = NewIdentifyTreeNode(jp.token, jp.ident(), nil)
		loop:
			for {
				annos := jp.typeAnnotationsOpt() // 注解无处不在，这里先不处理注解 这里是类型注解
				// need to report an error later if LBRACKET is for array
				// index access rather than array creation level 可以是 @Some [] ，如果是 @Some [1] 就是错误的
				if len(*annos) > 0 && jp.tk != LBRACKET && jp.tk != ELLIPSIS {
					return jp.illegal("无效的对数组的注解")
				}
				switch jp.tk {
				case LBRACKET: // [
					jp.nextToken()
					if jp.tk == RBRACKET {
						// 读到了 []
						fmt.Println("读到了 [] ignore ..................")
					} else {
						if jp.termExpr() {
							jp.mode = term_mode_expr
							// t1 := jp.term()
							if len(*annos) > 0 {
								t = jp.illegal("无效的.............")
							}
							// arrayAccess := jp.F.At(pos).Indexed(t, t1)
							// jp.to(arrayAccess)
							// t = arrayAccess.AbstractJCExpression
							t = GetEmptyTreeNode()
						}
					}
				case LPAREN: // (
				case DOT: // .
				case ELLIPSIS: // ... 多个参数
				case LT: // <
				default:
					break loop
				}
				goto loop
			}
		}
	case BYTE, SHORT, CHAR, INT, LONG, FLOAT, DOUBLE, BOOLEAN:
		if typeArgs != nil {
			jp.illegal("")
		}
		// 分开写是为了表明代码的意图
		t = jp.basicType()    // 这是一个常量类型，最常见的
		t = jp.bracketsOpt(t) // 可能是数组 所以需要处理方括号，例如： int[] a = {} 这里我们先不处理数组
		// 这里先不处理 ，忽略 TODO
		//t = jp.bracketsSuffix(t)
	case VOID:
	default:
		jp.illegal("无效的token... term3()")
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
func (jp *JavacParser) termRest(t *TreeNode) *TreeNode {

	switch jp.tk {
	case EQ: // = 表示是赋值语句
		// pos := jp.token.Pos()
		jp.nextToken()
		jp.mode = term_mode_expr
		t1 := jp.term() // 这里就是递归了 结果值也可以是一个表达式，例如 int a = sum(10);
		// return jp.toP(jp.F.At(pos).Assign(t, t1).AbstractJCExpression)
		return t1
	case PLUSEQ, // ++ -- 这样的，例如 a++
		SUBEQ,
		STAREQ,
		SLASHEQ,
		PERCENTEQ,
		AMPEQ,
		BAREQ,
		CARETEQ,
		LTLTEQ,
		GTGTEQ,
		GTGTGTEQ:
		// pos := jp.token.Pos()
		// tk := jp.tk
		jp.nextToken()
		jp.mode = term_mode_expr
		t1 := jp.term()
		// return jp.F.At(pos).Assignop(opTag(tk), t, t1).AbstractJCExpression
		return t1
	default:
		return t
	}
}

/** Expression1Rest = ["?" Expression ":" Expression1]
 */
func (jp *JavacParser) term1Rest(t *TreeNode) *TreeNode {

	if jp.tk == QUES {
		// pos := jp.token.Pos()
		jp.nextToken()
		// t1 := jp.term()
		jp.accept(COLON) // 期望下一个字符是冒号 (:)
		t2 := jp.term1()
		// return jp.F.At(pos).Conditional(t, t1, t2).AbstractJCExpression
		return t2
	} else {
		return t
	}
}

// precedences 的意思是优先级
func (jp *JavacParser) term2Rest(t *TreeNode, precedences PrecOp) *TreeNode {

	// TODO
	return t
}

func prec(tk TokenKind) PrecOp {

	// treeTag := opTag(tk)
	// if treeTag != jc.TREE_TAG_NO_TAG {
	// 	return jc.OpPrec(treeTag)
	// } else {
	// 	return -1
	// }

	// todo
	return -1
}

/** Skip forward until a suitable stop token is found.
 */
func (jp *JavacParser) skip(stopAtImport bool, stopAtMemberDecl bool,
	stopAtIdentifier bool, stopAtStatement bool) {
	for {
		switch jp.tk {
		case
			SEMI:
			jp.nextToken()
			return
		case PUBLIC, FINAL, ABSTRACT,
			MONKEYS_AT, EOF, CLASS,
			INTERFACE, ENUM:
			return
		case IMPORT:
			if stopAtImport {
				return
			}
			break
		case LBRACE, RBRACE, PRIVATE,
			PROTECTED, STATIC, TRANSIENT,
			NATIVE, VOLATILE, SYNCHRONIZED,
			STRICTFP, LT, BYTE, SHORT,
			CHAR, INT, LONG, FLOAT,
			DOUBLE, BOOLEAN, VOID:
			if stopAtMemberDecl {
				return
			}
			break
		case UNDERSCORE, IDENTIFIER:
			if stopAtIdentifier {
				return
			}
			break
		case CASE, DEF, IF, FOR, WHILE,
			DO, TRY, SWITCH, RETURN, THROW, BREAK,
			CONTINUE, ELSE, FINALLY, CATCH:
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
func (jp *JavacParser) importDeclaration() *TreeNode {

	t := GetEmptyTreeNode()
	pos := jp.token.Pos()
	jp.nextToken()
	importStatic := false
	// 这里先允许 import static com.husd; 这样的语法
	if jp.tk == STATIC {
		importStatic = true
		jp.nextToken()
	}
	var pid *TreeNode = GetEmptyTreeNode()
	for {
		pos1 := jp.token.Pos()
		jp.accept(DOT)
		if jp.tk == STAR {
			pid = GetEmptyTreeNode()
			pos1++
			jp.nextToken()
			break
		} else {
			pid = GetEmptyTreeNode()
			pos1++
		}
		if jp.tk != DOT {
			break
		}
	}
	jp.accept(SEMI)
	if compiler.DEBUG {
		fmt.Println("import static", importStatic, " pid:", pid, pos)
	}
	return t
}

func (jp *JavacParser) typeDeclaration(mods *TreeNode) *TreeNode {

	return mods
}

func (jp *JavacParser) illegal(msg string) *TreeNode {

	jp.reportSyntaxError(jp.token.Pos(), msg, jp.tk)
	return jp.syntaxError(jp.token.Pos(), msg)
}

func (jp *JavacParser) syntaxError(pos int, msg string) *TreeNode {

	return NewErrorTreeNode(pos, msg)
}

/** BasicType = byte | short | char | int | long | float | double | boolean
 * 注意，包装类，不是语法解析器的事情，自动拆箱和装箱也不是
 */
func (jp *JavacParser) basicType() *TreeNode {

	var tt *code.TypeTag
	tt = typeTag(jp.tk)
	inx := jp.token.GetSymbolTableIndex()
	res := NewPrimitiveTypeTree(inx, tt)
	jp.nextToken()
	return res
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
// func (jp *JavacParser) bracketsOpt(expression *jc.AbstractJCExpression, annotations *[]jc.JCAnnotation) *jc.AbstractJCExpression {
//
// 	// nextLevelAnnotations := jp.typeAnnotationsOpt()
// 	// 这里我们不处理注解，所以先返回空
// 	return expression
// }

/**
 * 类型注解，这里先忽略。
 *
 * 要解析出来注解 ，这里我们暂时不支持注解，先忽略，返回空 TODO annotation
 * 这里要区分下普通注解和类型注解
 * 在 Java 8 之前的版本中，只能允许在声明式前使用 Annotation。
 * 而在 Java 8 版本中，Annotation 可以被用在任何使用 Type 的地方，例如：初始化对象时 (new)，
 * 对象类型转化时，使用 implements 表达式时，或者使用 throws 表达式时。
 * 类型注解，是为了增加编译器对于一些运行期的代码检查，例如：
 * /** @NonNull * / Object my = new Object(); 注意这里的 @NonNull是注释的，可以兼容JDK8之前的代码
 * 这里我们就先忽略这种注解了，词法分析阶段，也没有解析这种类型的注解。
 */
func (jp *JavacParser) typeAnnotationsOpt() *[]*TreeNode {

	return GetEmptyTreeNodeArray()
}

/** BracketsSuffixExpr = "." CLASS
 *  BracketsSuffixType =
 *
 * TODO 先不处理
 */
func (jp *JavacParser) bracketsSuffix(opt *TreeNode) *TreeNode {

	if (jp.mode&term_mode_expr) != 0 &&
		jp.tk == DOT {
		jp.mode = term_mode_expr
		// newPos := jp.token.Pos()
		jp.nextToken()
		jp.accept(CLASS)
		// TODO
	} else if (jp.mode & term_mode_type) != 0 {
		if jp.tk != COLCOL {
			jp.mode = term_mode_type
		}
	} else if jp.tk != COLCOL {
		jp.syntaxError(jp.token.Pos(), "期望.class")
	}
	return opt
}

/**
 * 向前看0个token，是不是指定的token，是就返回true
 */
func (jp *JavacParser) peekToken(tk TokenKind) bool {

	lookahead := 0
	return jp.peekTokenLookahead(lookahead, tk)
}

/**
 * 向前看指定数量个token，是不是指定的token，是就返回true
 */
func (jp *JavacParser) peekTokenLookahead(lookahead int, tk TokenKind) bool {

	return AcceptTokenKind(tk, jp.S.LookAheadByIndex(lookahead+1).GetTokenKind())
}

func (jp *JavacParser) peekTokenLookahead2(lookahead int, tk1 TokenKind, tk2 TokenKind) bool {

	return AcceptTokenKind(tk1, jp.S.LookAheadByIndex(lookahead+1).GetTokenKind()) &&
		AcceptTokenKind(tk2, jp.S.LookAheadByIndex(lookahead+2).GetTokenKind())
}

/**
 * 传一个返回bool类型的filter函数，来接纳多个token
 */
func (jp *JavacParser) peekTokenLookaheadByFilter(lookahead int, f func(TokenKind) bool) bool {

	return f(jp.S.LookAheadByIndex(lookahead + 1).GetTokenKind())
}

// 宽泛的标识符号
func acceptLaxIdentifier(t TokenKind) bool {
	return t == IDENTIFIER || t == UNDERSCORE ||
		t == ASSERT || t == ENUM
}

// 暂时不实现lambda表达式 todo lambda
func (jp *JavacParser) lambdaExpressionOrStatement(hasParens bool, explicitParams bool, pos int) *TreeNode {

	panic("implement me lambda")
}

// 记录 pos 到 endPosTable
func (jp *JavacParser) toP(expr *TreeNode) *TreeNode {

	// jp.endPosTable.toP(expr.AbstractJCTree)
	return expr
}

func (jp *JavacParser) to(expr *TreeNode) *TreeNode {

	// jp.endPosTable.toP(expr.AbstractJCTree)
	return expr
}

/**
 * 泛型 先不支持 TODO 泛型
 */
func (jp *JavacParser) typeArgumentsOpt(pm parseMode) *[]*TreeNode {

	if true {
		return nil
	}
	if jp.tk == LT {
		jp.checkGenerics()
		if (jp.mode&pm) == 0 ||
			(jp.mode&term_mode_noparams) != 0 {
			jp.illegal("")
		}
		jp.mode = pm
		return typeArguments(false)
	}
	return GetEmptyTreeNodeArray()
}

/**
 * TypeArguments  = "<" TypeArgument {"," TypeArgument} ">"
 * 目前先不支持
 */
func typeArguments(b bool) *[]*TreeNode {

	return GetEmptyTreeNodeArray()
}

// 检查是否支持泛型
func (jp *JavacParser) checkGenerics() {

	if jp.allowGenerics == false {
		jp.error(jp.token.Pos(), "不支持泛型")
	}
}

func (jp *JavacParser) term3Rest(t *TreeNode, args *[]*TreeNode) *TreeNode {

	// TODO
	return t
}

// { statement }
// 这里还是递归的思路，考虑下嵌套的 { { { } } } 这样的语句
func (jp *JavacParser) parseBlock(father *TreeNode) {

	res := GetEmptyTreeNode()
	for jp.tk != RBRACE {
		res = jp.ParseStatement()
		father.Append(res)
	}
	jp.accept(RBRACE) // 读取到 }
}

/**
 * 转换if语句 if ( condition ) { truePart } else { falsePart }
 * IF ParExpression Statement [ELSE Statement]
 */
func (jp *JavacParser) parseIf() *TreeNode {

	res := NewIfTreeNode(jp.token)
	condition := jp.parExpression()
	res.Append(condition)
	truePart := jp.parseIfElseStatement()
	res.Append(truePart)

	// 看看是否有else   else 是可选的
	// if (true) {} else {}
	// if (true) {}
	// if (true) {} else if (a == 1) {} else {}
	hasElse := jp.acceptMaybe(ELSE)
	if hasElse {
		falsePart := NewBlockTreeNode(jp.token)
		hasIf := jp.acceptMaybe(IF)
		if hasIf {
			falsePart.Append(jp.parseIf())
		} else {
			falsePart.Append(jp.parseIfElseStatement())
		}
	}
	// else 是可选的
	return res
}

/**
 * 看看是哪种 这个时候就需要决定是向前看，还是向后看了，向后读了之后，还能不能
 * 恢复到原来的状态
 */
func (jp *JavacParser) parseCompareExpression() TreeNodeTag {

	res := erroneous
	switch jp.tk {
	case EQEQ:
		res = eq
	case BANGEQ:
		res = ne
	case LTEQ:
		res = le
	case GTEQ:
		res = ge
	case LT:
		res = lt
	case GT:
		res = gt
	default:
		// error
		jp.reportSyntaxError(jp.token.Pos(), "错误的比较符号", jp.tk)
	}
	jp.nextToken()
	return res
}

/**
 * 这个要返回一个最小的表达式单元：
 * BNF：
 * COM = EXP1 OP EXP2
 * OP = "==" | "!=" | ">" | "<" | ">=" | "<="
 * int_literal|long_literal
 *
 */
func (jp *JavacParser) parseExpression1() *TreeNode {

	// a == 10 | 10 == a
	res := GetEmptyTreeNode()
	switch jp.tk {
	case IDENTIFIER:
		left := NewIdentifyTreeNode(jp.token)
		jp.nextToken()
		opC := jp.isOpCompare()
		if !opC {
			jp.reportSyntaxError(jp.token.Pos(), "期望是 == != 等", jp.tk)
		}
		res = NewCompareConditionTreeNode(jp.token, toOpTag(jp.tk))
		jp.nextToken()
		right := jp.parseExpression1()
		res.Append(left)
		res.Append(right)
	case INTLITERAL, LONGLITERAL, FLOATLITERAL,
		DOUBLELITERAL, CHARLITERAL, STRINGLITERAL,
		TRUE, FALSE, NULL:
		// 这里都是字面量类型 需要注意，包含了 true false null 这3个，不要漏了 暂时先这么写 为了不报错 TODO
		res = NewLiteralTreeNode(jp.token, code.TYPE_TAG_LONG, 100, tt_long_literal)
		jp.nextToken()
	default:
		// error
		jp.reportSyntaxError(jp.token.Pos(), "错误的表达式", jp.tk)
	}
	return res
}

/**
 * 这个方法和 parseBlock的区别在于，这个方法可以没有 {
 */
func (jp *JavacParser) parseIfElseStatement() *TreeNode {

	hasLeftBrace := jp.acceptMaybe(LBRACE)
	// 如果有左括号 { ，就必须有右括号 }
	if hasLeftBrace {
		stat := NewBlockTreeNode(jp.token)
		jp.parseBlock(stat)
		return stat
	} else {
		// 没有左括号，就只读一行代码
		return jp.ParseExpression()
	}
}

/** ParExpression = "(" Expression ")"
 */
func (jp *JavacParser) parExpression() *TreeNode {

	jp.accept(LPAREN)
	res := jp.ParseExpression()
	jp.accept(RPAREN)
	return res
}

/**
 * If we see an identifier followed by a '&lt;' it could be an unbound
 * method reference or a binary expression. To disambiguate, look for a
 * matching '&gt;' and see if the subsequent terminal is either '.' or '::'.
 *
 * 主要讨论 () 的作用，目前括号主要有4个用途，正好是 parenthesesResult
 * 对应的值
 * 1、类型转换 int a = (short)1;
 * 2、lamdba表达式 MathOperation division = (int a, int b) -> a / b;
 * 3、括号表明优先级的，例如： a = (b + 10) * 100
 */
func (jp *JavacParser) analyzeParens() parenthesesResult {

	depth := 0
	typeCast := false
	for lookahead := 0; ; lookahead++ { // 预测分析法，看看之后的token是什么类型，来决定
		currentTk := jp.S.LookAheadByIndex(lookahead).GetTokenKind()
		switch currentTk {
		case COMMA: // , 括号后面跟着逗号是什么语法？ MathOperation subtraction = (a, b) -> a - b;
			typeCast = true
		case EXTENDS, DOT,
			SUPER, AMP: // extends . super & 忽略
			break
		case QUES:
			if jp.peekTokenLookahead(lookahead, EXTENDS) ||
				jp.peekTokenLookahead(lookahead, SUPER) {
				typeCast = true
			}
		case BYTE, SHORT, INT,
			LONG, FLOAT,
			DOUBLE, BOOLEAN, CHAR: // 这几个基本类型的关键字
			// ( int a = 10 ) 这种类型
			if jp.peekTokenLookahead(lookahead, RPAREN) {
				// Type, ')' -> cast 类型转换
				return CAST
			} else if jp.peekTokenLookaheadByFilter(lookahead, acceptLaxIdentifier) { // 这里把函数的指针传过去
				// Type, Identifier/'_'/'assert'/'enum' -> explicit lambda
				return EXPLICIT_LAMBDA
			}
		case LPAREN: // ((
			if lookahead != 0 {
				// '(' in a non-starting position -> parens
				return PARENS
			} else if jp.peekTokenLookahead(lookahead, RPAREN) {
				// () -> explicit lambda 明确的lambda表达式
				return EXPLICIT_LAMBDA
			}
		case RPAREN: // )
			// if we have seen something that looks like a type,
			// then it's a cast expression  这样的：(int)100L 两个括号的中间，有一个类型，那么就是强制类型转换
			if typeCast {
				return CAST
			}
			// otherwise, disambiguate cast vs. parenthesized expression
			// based on subsequent token.
			switch jp.S.LookAheadByIndex(lookahead + 1).GetTokenKind() {
			case BANG, TILDE,
				LPAREN, THIS, SUPER,
				INTLITERAL, LONGLITERAL, FLOATLITERAL,
				DOUBLELITERAL, CHARLITERAL, STRINGLITERAL,
				TRUE, FALSE, NULL,
				NEW, IDENTIFIER, ASSERT, ENUM, UNDERSCORE,
				BYTE, SHORT, CHAR, INT,
				LONG, FLOAT, DOUBLE, BOOLEAN, VOID:
				return CAST
			default:
				return PARENS
			}
		case UNDERSCORE, ASSERT,
			ENUM, IDENTIFIER:
			if jp.peekTokenLookaheadByFilter(lookahead, acceptLaxIdentifier) {
				// Identifier, Identifier/'_'/'assert'/'enum' -> explicit lambda
				return EXPLICIT_LAMBDA
			} else if jp.peekTokenLookahead2(lookahead, RPAREN, ARROW) {
				// Identifier, ')' '->' -> implicit lambda
				return IMPLICIT_LAMBDA
			}
			// 确认不是类型转换了
			typeCast = false
		case FINAL, ELLIPSIS:
			//those can only appear in explicit lambdas
			return EXPLICIT_LAMBDA
		case MONKEYS_AT:
			// 先不管注解
			typeCast = true
			lookahead = lookahead + 1 // 跳过@符号 进入注解解析
			for jp.peekTokenLookahead(lookahead, DOT) {
				lookahead = lookahead + 2
			}
			if jp.peekTokenLookahead(lookahead, LPAREN) {
				lookahead++
				//skip annotation values
				var nesting int = 0
				for ; ; lookahead++ {
					tk2 := jp.S.LookAheadByIndex(lookahead).GetTokenKind()
					switch tk2 {
					case EOF:
						return PARENS
					case LPAREN:
						nesting++
					case RPAREN:

					}
				}
			}
		case LBRACKET:
		case LT:
			depth++
		case GTGTGT:
			depth--
		case GTGT:
			depth--
		case GT:
		default:
			return PARENS
		}
	}
	// goto outer
	// todo delete later
	return CAST
}

/**
 * 处理数组类型  TODO
 */
func (jp *JavacParser) bracketsOpt(tree *TreeNode) *TreeNode {

	empty := &([]*TreeNode{})
	return jp.bracketsOptAnnotation(tree, empty)
}

func (jp *JavacParser) bracketsOptAnnotation(tree *TreeNode, annotationArray *[]*TreeNode) *TreeNode {

	var t *TreeNode = tree
	var nextLevelAnnotations *[]*TreeNode
	nextLevelAnnotations = jp.typeAnnotationsOpt()
	if jp.tk == LBRACKET { // [
		pos := jp.token.Pos()
		jp.nextToken()
		t = jp.bracketsOptCont(t, pos, nextLevelAnnotations)
	} else if len(*nextLevelAnnotations) > 0 {
		if jp.permitTypeAnnotationsPushBack {
			jp.typeAnnotationsPushedBack = nextLevelAnnotations
		} else {
			jp.illegal("无效的 nextLevelAnnotations")
		}
	}
	if len(*annotationArray) > 0 {
		//t = toP(F.at(token.pos).AnnotatedType(annotations, t));
	}
	return t
}

func (jp *JavacParser) bracketsOptCont(t *TreeNode, p int, annotations *[]*TreeNode) *TreeNode {

	jp.accept(RBRACKET) // 先解析 ] 此时就是 []这样的状态
	t = jp.bracketsOpt(t)
	t = NewTypeArrayTreeNode(t)
	if len(*annotations) > 0 {
		// 这里类型注解永远是空，先不支持，所以不用考虑这段代码 TODO
	}
	return t
}

// 返回none就是没有类型
func typeTag(tk TokenKind) *code.TypeTag {

	switch tk {
	case BYTE:
		return code.TYPE_TAG_BYTE
	case CHAR:
		return code.TYPE_TAG_CHAR
	case SHORT:
		return code.TYPE_TAG_SHORT
	case INT:
		return code.TYPE_TAG_INT
	case LONG:
		return code.TYPE_TAG_LONG
	case FLOAT:
		return code.TYPE_TAG_FLOAT
	case DOUBLE:
		return code.TYPE_TAG_DOUBLE
	case BOOLEAN:
		return code.TYPE_TAG_BOOLEAN
	default:
		return code.TYPE_TAG_NONE
	}
}
