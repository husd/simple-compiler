package jc

import (
	"husd.com/v0/ast_tree2"
)

/**
 * 访问数组下标
 * @author hushengdong
 */
type JCArrayAccess struct {
	*AbstractJCExpression
	indexed *AbstractJCExpression
	index   *AbstractJCExpression
}

func NewJCArrayAccess(indexed *AbstractJCExpression, index *AbstractJCExpression) *JCArrayAccess {

	res := &JCArrayAccess{}
	res.AbstractJCExpression = NewJCExpression()

	res.indexed = indexed
	res.index = index

	res.getTreeType = func() *ast_tree2.TreeType {
		return ast_tree2.TREE_TYPE_ARRAY_ACCESS
	}
	res.getTag = func() JCTreeTag {
		return TREE_TAG_INDEXED
	}
	res.toString = func() string {
		return res.indexed.toString() + res.index.toString()
	}

	return res
}

// ArrayAccessTree

func (jc *JCArrayAccess) GetExpression() ast_tree2.ExpressionTreeV2 {
	return jc.indexed
}

func (jc *JCArrayAccess) GetIndex() ast_tree2.ExpressionTreeV2 {
	return jc.index
}

// ArrayAccessTree
