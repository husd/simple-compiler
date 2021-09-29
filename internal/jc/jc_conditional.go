package jc

import (
	"husd.com/v0/ast_tree2"
)

/**
 * A ( ) ? ( ) : ( ) conditional expression
 * @author hushengdong
 */
type JCConditional struct {
	*AbstractJCPolyExpression

	condition *AbstractJCExpression
	truePart  *AbstractJCExpression
	falsePart *AbstractJCExpression
}

func NewJCConditional(condition *AbstractJCExpression,
	truePart *AbstractJCExpression,
	falsePart *AbstractJCExpression) *JCConditional {

	res := &JCConditional{}
	res.AbstractJCPolyExpression = NewJCPolyExpression()
	res.condition = condition
	res.truePart = truePart
	res.falsePart = falsePart

	res.getTreeType = func() ast_tree2.TreeType {
		return ast_tree2.TREE_TYPE_CONDITIONAL_EXPRESSION
	}
	res.getTag = func() JCTreeTag {
		return TREE_TAG_CONDEXPR
	}

	return res
}

// ConditionalExpressionTree

func (jc *JCConditional) GetCondition() ast_tree2.ExpressionTreeV2 {
	return jc.condition
}

func (jc *JCConditional) GetTrueExpression() ast_tree2.ExpressionTreeV2 {
	return jc.truePart
}

func (jc *JCConditional) GetFalseExpression() ast_tree2.ExpressionTreeV2 {
	return jc.falsePart
}

// ConditionalExpressionTree
