package parser

/**
 * 转换表达式
 * @author hushengdong
 */

func (jp *JavacParser) term01() *TreeNode {

	return jp.term02()
}

/**
 * 思路要一点一点来做，首先是 解决二义性和左递归，要写成右递归
 * BNF：
 * EXP = EXP1 OP EXP1
 * EXP1 = identify | Literal // 标识符或者常量
 * OP = "==" | “=” | “+” | “-” | “*” | "/" | "%"
 *
 */
func (jp *JavacParser) term02() *TreeNode {

	return jp.parseBinaryOp()
}

/**
 * 转换二元运算表达式 这里先简单的来
 */
func (jp *JavacParser) parseBinaryOp() *TreeNode {

	left := jp.exp1()
	jp.peekTokenLookaheadByFilter(1, acceptSimpleOp)
	treeTag := NewBinaryOpTreeNode(toOpTag(jp.tk))
	right := jp.exp1()
	treeTag.Append(left)
	treeTag.Append(right)
	jp.accept(SEMI) // 分号结尾
	return treeTag
}

func (jp *JavacParser) exp1() *TreeNode {

	switch jp.tk {
	case IDENTIFIER:
		jp.nextToken()
		return NewIdentifyTreeNode(jp.token)
	case INTLITERAL, LONGLITERAL, FLOATLITERAL,
		DOUBLELITERAL, CHARLITERAL, STRINGLITERAL,
		TRUE, FALSE, NULL:
		jp.nextToken()
		return jp.literal(jp.names.Empty, jp.token.Pos())
	default:
		jp.reportSyntaxError(jp.token.Pos(), "错误的类型", jp.tk)
	}
	jp.nextToken()
	return NewErrorTreeNode(jp.token.Pos(), "不支持的token类型")
}

func acceptSimpleOp(tk TokenKind) bool {

	return tk == EQEQ ||
		tk == PLUS || tk == SUB || tk == STAR ||
		tk == SLASH || tk == AMP || tk == BAR ||
		tk == CARET || tk == PERCENT
}
