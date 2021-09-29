package jc

import (
	"husd.com/v0/ast_tree2"
	"husd.com/v0/util"
)

/**
 * Selects through packages and classes
 * @author hushengdong
 */
type JCFieldAccess struct {
	*AbstractJCExpression

	name *util.Name
	sym  *Symbol
}

func NewJCFieldAccess(selected *AbstractJCExpression, name *util.Name, sym *Symbol) *JCFieldAccess {

	res := &JCFieldAccess{selected, name, sym}
	res.getTreeType = func() ast_tree2.TreeType {
		return ast_tree2.TREE_TYPE_MEMBER_SELECT
	}
	res.getTag = func() JCTreeTag {
		return TREE_TAG_SELECT
	}
	return res
}

// memeber select tree

func (jc *JCFieldAccess) MemberSelectTreeV2_() {
}

func (jc *JCFieldAccess) GetExpression() ast_tree2.ExpressionTreeV2 {
	return jc.AbstractJCExpression
}

func (jc *JCFieldAccess) GetIdentifier() *util.Name {
	return jc.name
}

// memeber select tree
