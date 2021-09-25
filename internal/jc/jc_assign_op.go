package jc

import "husd.com/v0/ast_tree2"

/**
 * An assignment with "+=", "|=" ...
 * @author hushengdong
 */
type JCAssignOp struct {
	*AbstractJCExpression

	opCode   JCTreeTag
	left     *AbstractJCExpression
	right    *AbstractJCExpression
	operator *Symbol
}

func NewJCAssignOp(opCode JCTreeTag,
	left *AbstractJCExpression,
	right *AbstractJCExpression,
	operator *Symbol) *JCAssignOp {

	res := &JCAssignOp{}
	res.AbstractJCExpression = NewJCExpression()

	res.opCode = opCode
	res.left = left
	res.right = right
	res.operator = operator

	res.getTag = func() JCTreeTag {
		return res.opCode
	}

	res.getTreeType = func() *ast_tree2.TreeType {
		return treeTag2TreeKind(res.getTag())
	}

	return res
}

// CompoundAssignmentTree
func (jc *JCAssignOp) GetVariable() ast_tree2.ExpressionTreeV2 {
	return jc.left
}

func (jc *JCAssignOp) GetExpression() ast_tree2.ExpressionTreeV2 {
	return jc.right
}

func (jc *JCAssignOp) CompoundAssignmentTreeV2_() {
}

//CompoundAssignmentTree
