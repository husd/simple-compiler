package parser

import (
	"fmt"
	"husd.com/v0/code"
	"husd.com/v0/common"
	"husd.com/v0/compiler"
	"husd.com/v0/util"
)

/**
 *
 * @author hushengdong
 */

type JavacParser struct {
	c         *util.Context    //
	S         lexer            // 词法分析器
	source    code.JVersion    // 当前JDK的版本
	token     Token            // 当前读到的token
	tk        common.TokenKind // 当前读到的token的TOKEN_KIND
	rowNum    int              // 多少行
	columnNum int              // 列

	endPosTable *SimpleEndPosTable

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
const term_mode_diamond parseMode = 0x10

func (jp *JavacParser) ParseExpression() *TreeNode {

	// a > 10
	// a == 10 && b == 15
	// sum(1,5) == 20
	//
	// term有条件得意思 泛指Java里的表达式条件
	return jp.termWithMode(term_mode_expr)
}

func (jp *JavacParser) termWithMode(newMode parseMode) *TreeNode {

	preMode := jp.mode
	jp.mode = newMode
	var t *TreeNode // 这么做仅仅为了表达以下t的类型
	t = jp.term()
	jp.lastMode = jp.mode // jp.lastMode =  newMode 是最终生效的mode
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
func (jp *JavacParser) ParseStatement() *TreeNode {

	res := GetEmptyTreeNode()
	// 先从一个简单的空语句开始转换，然后慢慢开始加case语句
	pos := jp.token.Pos()
	switch jp.tk {
	case common.TOKEN_KIND_SEMI: // ; 空语句
		jp.nextToken()
	case common.TOKEN_KIND_EOF: // EOF也是空语句
	case common.TOKEN_KIND_LBRACE: // {
		jp.nextToken()
		blockTree := NewBlockTreeNode(jp.token)
		jp.parseBlock(blockTree)
		res = blockTree
	case common.TOKEN_KIND_IF: // if 语句
		jp.nextToken()
		res = jp.parseIf()
	default:
		// 其它情况都是错误的
		jp.reportSyntaxError(pos, "无效的token ", jp.tk)
	}
	// 空语句
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

	// var t *jc.AbstractJCExpression
	// t = jc.NewJCError(pos, "默认错误")
	// switch jp.tk {
	// case TOKEN_KIND_INT_LITERAL:
	// 	num, err := util.String2int(jp.token.GetStringVal(), jp.token.GetRadix(), 32)
	// 	if err != nil {
	// 		jp.error(jp.token.Pos(), err.Error())
	// 	}
	// 	literal := jp.F.At(pos).Literal(code.TYPE_TAG_INT, int(num))
	// 	return literal.AbstractJCExpression
	// }
	return GetEmptyTreeNode()
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
func (jp *JavacParser) accept(tk common.TokenKind) {

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
func (jp *JavacParser) acceptMaybe(tk common.TokenKind) bool {

	if jp.tk == tk {
		jp.nextToken()
		return true
	} else {
		return false
	}
}

func (jp *JavacParser) reportSyntaxError(pos int, msg string, tk common.TokenKind) {

	// TODO 暂时先打印，应该有更好的方式来报告语法错误
	// 发送一个事件，通知所有监听这个事件的程序来处理语法错误。
	fmt.Println("---------------- reportSyntaxError，位置：", pos, " msg:", msg, " TokenKind:", tk)
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
	// for jp.tk == TOKEN_KIND_DOT {
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
	if jp.tk == common.TOKEN_KIND_IDENTIFIER {
		name := tk.GetName()
		jp.nextToken()
		return name
	} else if jp.tk == common.TOKEN_KIND_ASSERT {
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
	} else if jp.tk == common.TOKEN_KIND_ENUM {
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
	} else if jp.tk == common.TOKEN_KIND_THIS {
		if allowThisIdent {
			_name := jp.token.GetName()
			jp.nextToken()
			return _name
		} else {
			jp.error(jp.token.Pos(), "this as identifier")
			jp.nextToken()
			return jp.names.Error
		}
	} else if jp.tk == common.TOKEN_KIND_UNDERSCORE {
		jp.warn(jp.token.Pos(), "underscore.as.identifier")
		_name := jp.token.GetName()
		jp.nextToken()
		return _name
	} else {
		jp.accept(common.TOKEN_KIND_IDENTIFIER)
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
		jp.tk == common.TOKEN_KIND_EQ ||
		jp.tk >= common.TOKEN_KIND_PLUSEQ &&
			jp.tk <= common.TOKEN_KIND_GTGTGTEQ {
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
		jp.tk == common.TOKEN_KIND_QUES {
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

	const OR_PREC int = 4
	e := jp.term3()
	if (jp.mode&term_mode_expr) != 0 &&
		prec(jp.tk) >= OR_PREC {
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
func (jp *JavacParser) term3() *TreeNode {

	pos := jp.token.Pos()
	var t *TreeNode
	typeArgs := jp.typeArgumentsOpt(term_mode_expr)

	// 处理这样的范型表达式的，这里先忽略范型
	// TypeArguments  = "<" TypeArgument {"," TypeArgument} ">"
	// List<JCExpression> typeArgs = typeArgumentsOpt(EXPR);
	switch jp.tk {

	case common.TOKEN_KIND_INT_LITERAL, common.TOKEN_KIND_LONG_LITERAL, common.TOKEN_KIND_FLOAT_LITERAL,
		common.TOKEN_KIND_DOUBLE_LITERAL, common.TOKEN_KIND_CHAR_LITERAL, common.TOKEN_KIND_STRING_LITERAL,
		common.TOKEN_KIND_TRUE, common.TOKEN_KIND_FALSE, common.TOKEN_KIND_NULL: // 最简单的 boolean a = false;

		if (jp.mode & term_mode_expr) != 0 {
			jp.mode = term_mode_expr
			t = jp.literal(jp.names.Empty, jp.token.Pos())
		} else {
			return jp.illegal("无效的表达式")
		}

	case common.TOKEN_KIND_BYTE, common.TOKEN_KIND_SHORT, common.TOKEN_KIND_CHAR, common.TOKEN_KIND_INT,
		common.TOKEN_KIND_LONG, common.TOKEN_KIND_FLOAT, common.TOKEN_KIND_DOUBLE, common.TOKEN_KIND_BOOLEAN: // 最简单的 boolean a = false;

		// emptyAnnotations := ast.GetEmptyTreeNode()
		// primitiveTypeTree := jp.basicType()
		// t = jp.bracketsSuffix(jp.bracketsOpt(primitiveTypeTree.AbstractJCExpression, emptyAnnotations))
		t = GetEmptyTreeNode()
	case common.TOKEN_KIND_UNDERSCORE, common.TOKEN_KIND_IDENTIFIER,
		common.TOKEN_KIND_ASSERT, common.TOKEN_KIND_ENUM: //

		//  ->  lambda表达式 如果前面一个token是 -> 表示接下来要解析的是lambda表达式
		if jp.termExpr() && jp.peekToken(common.TOKEN_KIND_ARROW) {
			return jp.lambdaExpressionOrStatement(false, false, pos)
		} else {
			// t = jp.toP(jp.F.At(pos).Identify(jp.ident()))
			t = GetEmptyTreeNode()
		loop:
			for {
				// pos := jp.token.Pos()
				annos := jp.typeAnnotationsOpt() // 注解无处不在，这里先不处理注解
				// need to report an error later if LBRACKET is for array
				// index access rather than array creation level 可以是 @Some [] ，如果是 @Some [1] 就是错误的
				if len(*annos) > 0 &&
					jp.tk != common.TOKEN_KIND_LBRACKET &&
					jp.tk != common.TOKEN_KIND_ELLIPSIS {
					return jp.illegal("无效的对数组的注解")
				}
				switch jp.tk {
				case common.TOKEN_KIND_LBRACKET: // [
					jp.nextToken()
					if jp.tk == common.TOKEN_KIND_RBRACKET {
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
				case common.TOKEN_KIND_LPAREN: // (
				case common.TOKEN_KIND_DOT: // .
				case common.TOKEN_KIND_ELLIPSIS: // ... 多个参数
				case common.TOKEN_KIND_LT: // <
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
func (jp *JavacParser) termRest(t *TreeNode) *TreeNode {

	switch jp.tk {
	case common.TOKEN_KIND_EQ: // = 表示是赋值语句
		// pos := jp.token.Pos()
		jp.nextToken()
		jp.mode = term_mode_expr
		t1 := jp.term() // 这里就是递归了 结果值也可以是一个表达式，例如 int a = sum(10);
		// return jp.toP(jp.F.At(pos).Assign(t, t1).AbstractJCExpression)
		return t1
	case common.TOKEN_KIND_PLUSEQ, // ++ -- 这样的，例如 a++
		common.TOKEN_KIND_SUBEQ,
		common.TOKEN_KIND_STAREQ,
		common.TOKEN_KIND_SLASHEQ,
		common.TOKEN_KIND_PERCENTEQ,
		common.TOKEN_KIND_AMPEQ,
		common.TOKEN_KIND_BAREQ,
		common.TOKEN_KIND_CARETEQ,
		common.TOKEN_KIND_LTLTEQ,
		common.TOKEN_KIND_GTGTEQ,
		common.TOKEN_KIND_GTGTGTEQ:
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

	if jp.tk == common.TOKEN_KIND_QUES {
		// pos := jp.token.Pos()
		jp.nextToken()
		// t1 := jp.term()
		jp.accept(common.TOKEN_KIND_COLON) // 期望下一个字符是冒号 (:)
		t2 := jp.term1()
		// return jp.F.At(pos).Conditional(t, t1, t2).AbstractJCExpression
		return t2
	} else {
		return t
	}
}

func (jp *JavacParser) term2Rest(t *TreeNode, perc int) *TreeNode {

	// TODO
	return t
}

func prec(tk common.TokenKind) int {

	// treeTag := opTag(tk)
	// if treeTag != jc.TREE_TAG_NO_TAG {
	// 	return jc.OpPrec(treeTag)
	// } else {
	// 	return -1
	// }

	// todo
	return -1
}

func opTag(tk common.TokenKind) TreeNodeTag {

	switch tk {
	case common.TOKEN_KIND_BARBAR:
		return Tree_node_tag_or
	case common.TOKEN_KIND_AMPAMP:
		return Tree_node_tag_and
	case common.TOKEN_KIND_BAR:
		return Tree_node_tag_bitor
	case common.TOKEN_KIND_BAREQ:
		return Tree_node_tag_bitor_asg
	case common.TOKEN_KIND_CARET:
		return Tree_node_tag_bitxor
	case common.TOKEN_KIND_CARETEQ:
		return Tree_node_tag_bitxor_asg
	case common.TOKEN_KIND_AMP:
		return Tree_node_tag_bitand
	case common.TOKEN_KIND_AMPEQ:
		return Tree_node_tag_bitand_asg
	case common.TOKEN_KIND_EQEQ:
		return Tree_node_tag_eq
	case common.TOKEN_KIND_BANGEQ:
		return Tree_node_tag_ne
	case common.TOKEN_KIND_LT:
		return Tree_node_tag_lt
	case common.TOKEN_KIND_GT:
		return Tree_node_tag_gt
	case common.TOKEN_KIND_LTEQ:
		return Tree_node_tag_le
	case common.TOKEN_KIND_GTEQ:
		return Tree_node_tag_ge
	case common.TOKEN_KIND_LTLT:
		return Tree_node_tag_sl
	case common.TOKEN_KIND_LTLTEQ:
		return Tree_node_tag_sl_asg
	case common.TOKEN_KIND_GTGT:
		return Tree_node_tag_sr
	case common.TOKEN_KIND_GTGTEQ:
		return Tree_node_tag_sr_asg
	case common.TOKEN_KIND_GTGTGT:
		return Tree_node_tag_usr
	case common.TOKEN_KIND_GTGTGTEQ:
		return Tree_node_tag_usr_asg
	case common.TOKEN_KIND_PLUS:
		return Tree_node_tag_plus
	case common.TOKEN_KIND_PLUSEQ:
		return Tree_node_tag_plus_asg
	case common.TOKEN_KIND_SUB:
		return Tree_node_tag_minus
	case common.TOKEN_KIND_SUBEQ:
		return Tree_node_tag_minus_asg
	case common.TOKEN_KIND_STAR:
		return Tree_node_tag_mul
	case common.TOKEN_KIND_STAREQ:
		return Tree_node_tag_mul_asg
	case common.TOKEN_KIND_SLASH:
		return Tree_node_tag_div
	case common.TOKEN_KIND_SLASHEQ:
		return Tree_node_tag_div_asg
	case common.TOKEN_KIND_PERCENT:
		return Tree_node_tag_mod
	case common.TOKEN_KIND_PERCENTEQ:
		return Tree_node_tag_mod_asg
	case common.TOKEN_KIND_INSTANCEOF:
		return Tree_node_tag_typetest
	default:
		return Tree_node_tag_no_tag
	}
}

/** Skip forward until a suitable stop token is found.
 */
func (jp *JavacParser) skip(stopAtImport bool, stopAtMemberDecl bool,
	stopAtIdentifier bool, stopAtStatement bool) {
	for {
		switch jp.tk {
		case
			common.TOKEN_KIND_SEMI:
			jp.nextToken()
			return
		case common.TOKEN_KIND_PUBLIC, common.TOKEN_KIND_FINAL, common.TOKEN_KIND_ABSTRACT,
			common.TOKEN_KIND_MONKEYS_AT, common.TOKEN_KIND_EOF, common.TOKEN_KIND_CLASS,
			common.TOKEN_KIND_INTERFACE, common.TOKEN_KIND_ENUM:
			return
		case common.TOKEN_KIND_IMPORT:
			if stopAtImport {
				return
			}
			break
		case common.TOKEN_KIND_LBRACE, common.TOKEN_KIND_RBRACE, common.TOKEN_KIND_PRIVATE,
			common.TOKEN_KIND_PROTECTED, common.TOKEN_KIND_STATIC, common.TOKEN_KIND_TRANSIENT,
			common.TOKEN_KIND_NATIVE, common.TOKEN_KIND_VOLATILE, common.TOKEN_KIND_SYNCHRONIZED,
			common.TOKEN_KIND_STRICTFP, common.TOKEN_KIND_LT, common.TOKEN_KIND_BYTE, common.TOKEN_KIND_SHORT,
			common.TOKEN_KIND_CHAR, common.TOKEN_KIND_INT, common.TOKEN_KIND_LONG, common.TOKEN_KIND_FLOAT,
			common.TOKEN_KIND_DOUBLE, common.TOKEN_KIND_BOOLEAN, common.TOKEN_KIND_VOID:
			if stopAtMemberDecl {
				return
			}
			break
		case common.TOKEN_KIND_UNDERSCORE, common.TOKEN_KIND_IDENTIFIER:
			if stopAtIdentifier {
				return
			}
			break
		case common.TOKEN_KIND_CASE, common.TOKEN_KIND_DEF, common.TOKEN_KIND_IF, common.TOKEN_KIND_FOR, common.TOKEN_KIND_WHILE,
			common.TOKEN_KIND_DO, common.TOKEN_KIND_TRY, common.TOKEN_KIND_SWITCH, common.TOKEN_KIND_RETURN, common.TOKEN_KIND_THROW, common.TOKEN_KIND_BREAK,
			common.TOKEN_KIND_CONTINUE, common.TOKEN_KIND_ELSE, common.TOKEN_KIND_FINALLY, common.TOKEN_KIND_CATCH:
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
	if jp.tk == common.TOKEN_KIND_STATIC {
		importStatic = true
		jp.nextToken()
	}
	var pid *TreeNode = GetEmptyTreeNode()
	for {
		pos1 := jp.token.Pos()
		jp.accept(common.TOKEN_KIND_DOT)
		if jp.tk == common.TOKEN_KIND_STAR {
			pid = GetEmptyTreeNode()
			pos1++
			jp.nextToken()
			break
		} else {
			pid = GetEmptyTreeNode()
			pos1++
		}
		if jp.tk != common.TOKEN_KIND_DOT {
			break
		}
	}
	jp.accept(common.TOKEN_KIND_SEMI)
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

	return NewErrorTreeNode(msg)
}

func (jp *JavacParser) basicType() *TreeNode {

	// jp.F.At(jp.token.Pos())
	// tree := jp.F.TypeIdent(typeTag(jp.tk))
	// jp.nextToken()
	// return tree
	return GetEmptyTreeNode()
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
 * 要解析出来注解 ，这里我们暂时不支持注解，先忽略，返回空 TODO annotation
 */
func (jp *JavacParser) typeAnnotationsOpt() *[]TreeNode {

	return GetEmptyTreeNodeArray()
}

/** BracketsSuffixExpr = "." CLASS
 *  BracketsSuffixType =
 *
 *
 *
 * TODO 先不处理
 */
func (jp *JavacParser) bracketsSuffix(opt *TreeNode) *TreeNode {

	if (jp.mode&term_mode_expr) != 0 &&
		jp.tk == common.TOKEN_KIND_DOT {
		jp.mode = term_mode_expr
		// newPos := jp.token.Pos()
		jp.nextToken()
		jp.accept(common.TOKEN_KIND_CLASS)
		// TODO
	} else if (jp.mode & term_mode_type) != 0 {

	} else if jp.tk != common.TOKEN_KIND_COLCOL {
		jp.syntaxError(jp.token.Pos(), "期望.class")
	}
	return opt
}

/**
 * 向前看0个token，是不是指定的token，是就返回true
 */
func (jp *JavacParser) peekToken(tk common.TokenKind) bool {

	lookahead := 0
	return jp.peekTokenLookahead(lookahead, tk)
}

/**
 * 向前看指定数量个token，是不是指定的token，是就返回true
 */
func (jp *JavacParser) peekTokenLookahead(lookahead int, tk common.TokenKind) bool {

	return common.AcceptTokenKind(tk, jp.S.LookAheadByIndex(lookahead+1).GetTokenKind())
}

// 暂时不实现lambda表达式 todo lambda
func (jp *JavacParser) lambdaExpressionOrStatement(hasParens bool, explicitParams bool, pos int) *TreeNode {

	panic("implement me lambda")
}

//记录 pos 到 endPosTable
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
func (jp *JavacParser) typeArgumentsOpt(pm parseMode) *[]TreeNode {

	if jp.tk == common.TOKEN_KIND_LT {
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
func typeArguments(b bool) *[]TreeNode {

	return GetEmptyTreeNodeArray()
}

// 检查是否支持泛型
func (jp *JavacParser) checkGenerics() {

	if jp.allowGenerics == false {
		jp.error(jp.token.Pos(), "不支持泛型")
	}
}

func (jp *JavacParser) term3Rest(t *TreeNode, args *[]TreeNode) *TreeNode {

	// TODO
	return t
}

// { statement }
// 这里还是递归的思路，考虑下嵌套的 { { { } } } 这样的语句
func (jp *JavacParser) parseBlock(father *TreeNode) {

	res := GetEmptyTreeNode()
	for jp.tk != common.TOKEN_KIND_RBRACE {
		res = jp.ParseStatement()
		father.Append(res)
	}
	jp.accept(common.TOKEN_KIND_RBRACE) // 读取到 }
}

/**
 * 转换if语句 if ( condition ) { truePart } else { falsePart }
 */
func (jp *JavacParser) parseIf() *TreeNode {

	jp.accept(common.TOKEN_KIND_LPAREN)
	res := NewIfTreeNode(jp.token)
	condition := jp.parseIfCondition()
	res.Append(condition)
	jp.accept(common.TOKEN_KIND_RPAREN)
	leftBrace := jp.acceptMaybe(common.TOKEN_KIND_LBRACE)

	if leftBrace {
		// 有左括号，终止条件就是右括号
		truePart := NewBlockTreeNode(jp.token)
		jp.parseBlock(truePart)
		res.Append(truePart)
	} else {
		if jp.tk == common.TOKEN_KIND_EOF {
			jp.reportSyntaxError(jp.token.Pos(), "if语句没有truePart", jp.token.GetTokenKind())
		}
		// 没有左括号，最多只读一行代码作为truePart
		truePart := jp.ParseStatement()
		res.Append(truePart)
		jp.nextToken()
	}

	// 看看是否有else
	hasElse := jp.acceptMaybe(common.TOKEN_KIND_ELSE)
	if hasElse {
		leftBrace = jp.acceptMaybe(common.TOKEN_KIND_LBRACE)
		// 如果有左括号，就必须有右括号
		if leftBrace {
			falsePart := NewBlockTreeNode(jp.token)
			jp.parseBlock(falsePart)
			res.Append(falsePart)
		} else {
			// 没有左括号，就只读一行代码
			temp := jp.ParseStatement()
			res.Append(temp)
		}
	} else {
		falsePart := GetEmptyTreeNode()
		res.Append(falsePart)
	}
	return res
}

func (jp *JavacParser) parseIfCondition() *TreeNode {

	left := jp.parseExpression1()
	comparePart := jp.parseCompareExpression()
	res := NewCompareConditionTreeNode(jp.token, comparePart)
	right := jp.parseExpression1()
	res.Append(left)
	res.Append(right)

	// TODO
	// a == 10 a > 100这样的简单的
	//switch jp.tk {
	//case TOKEN_KIND_IDENTIFIER:
	////标识符有很多种情况，有可能是变量，有可能是 函数调用 有可能是对方的属性访问 有可能是对象的函数调用
	//// 需要向前看，预测下一个
	//
	//case TOKEN_KIND_INT_LITERAL,TOKEN_KIND_LONG_LITERAL,
	//TOKEN_KIND_FLOAT_LITERAL,TOKEN_KIND_DOUBLE_LITERAL,TOKEN_KIND_STRING_LITERAL,
	//TOKEN_KIND_CHAR_LITERAL:
	//	// 类似于 100 == xxx 这样的
	//	leftPart := ast.NewLiteralTreeNode(jp.token)
	//	jp.nextToken()
	//	//期望是比较类型的运算符号
	//	comparePart := jp.parseCompareExpression()
	//
	//default:
	//	//错误的
	//	jp.reportSyntaxError(jp.token.Pos(), "错误的if condition语句", jp.tk)
	//	res = ast.NewErrorTreeNode("错误的if condition语句")
	//}

	return res
}

/**
 * 看看是哪种 这个时候就需要决定是向前看，还是向后看了，向后读了之后，还能不能
 * 恢复到原来的状态
 */
func (jp *JavacParser) parseCompareExpression() TreeNodeTag {

	// preToken := jp.peekTokenLookahead(1,TOKEN_KIND_EQ)
	switch jp.tk {

	case common.TOKEN_KIND_EQ:

	}

	jp.nextToken()
	return Tree_node_tag_eq
}

/**
 * 这个要返回一个最小的表达式单元：
 * BNF：
 * identify
 * int_literal|long_literal
 *
 */
func (jp *JavacParser) parseExpression1() *TreeNode {

	res := NewErrorTreeNode("错误的表达式")
	switch jp.tk {
	case common.TOKEN_KIND_IDENTIFIER:
		res = NewIdentifyTreeNode(jp.token)
	case common.TOKEN_KIND_INT_LITERAL, common.TOKEN_KIND_LONG_LITERAL, common.TOKEN_KIND_FLOAT_LITERAL,
		common.TOKEN_KIND_DOUBLE_LITERAL, common.TOKEN_KIND_CHAR_LITERAL, common.TOKEN_KIND_STRING_LITERAL,
		common.TOKEN_KIND_TRUE, common.TOKEN_KIND_FALSE, common.TOKEN_KIND_NULL:
		// 这里都是字面量类型 需要注意，包含了 true false null 这3个，不要漏了
		res = NewLiteralTreeNode(jp.token)
	default:
		// error
	}
	return res
}

// 返回none就是没有类型
func typeTag(tk common.TokenKind) *code.TypeTag {

	switch tk {
	case common.TOKEN_KIND_BYTE:
		return code.TYPE_TAG_BYTE
	case common.TOKEN_KIND_CHAR:
		return code.TYPE_TAG_CHAR
	case common.TOKEN_KIND_SHORT:
		return code.TYPE_TAG_SHORT
	case common.TOKEN_KIND_INT:
		return code.TYPE_TAG_INT
	case common.TOKEN_KIND_LONG:
		return code.TYPE_TAG_LONG
	case common.TOKEN_KIND_FLOAT:
		return code.TYPE_TAG_FLOAT
	case common.TOKEN_KIND_DOUBLE:
		return code.TYPE_TAG_DOUBLE
	case common.TOKEN_KIND_BOOLEAN:
		return code.TYPE_TAG_BOOLEAN
	default:
		return code.TYPE_TAG_NONE
	}
}
