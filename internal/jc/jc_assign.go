package jc

import "husd.com/v0/ast_tree2"

/**
 * 赋值语句 a = 10;
 * @author hushengdong
 */
type JCAssign struct {
	*AbstractJCExpression

	left  *AbstractJCExpression
	right *AbstractJCExpression
}

func NewJCAssign(l *AbstractJCExpression, r *AbstractJCExpression) *JCAssign {

	res := &JCAssign{}
	res.AbstractJCExpression = NewJCExpression()
	res.left = l
	res.right = r

	res.getTreeType = func() *ast_tree2.TreeType {
		return ast_tree2.TREE_TYPE_ASSIGNMENT
	}

	res.getTag = func() JCTreeTag {
		return TREE_TAG_ASSIGN
	}

	return res
}

// AssignmentTree

func (jc *JCAssign) GetVariable() ast_tree2.ExpressionTreeV2 {
	return jc.left
}

func (jc *JCAssign) GetExpression() ast_tree2.ExpressionTreeV2 {
	return jc.right
}

func (jc *JCAssign) AssignmentTreeV2_() {
}

// AssignmentTree
