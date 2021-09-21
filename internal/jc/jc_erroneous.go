package jc

import (
	"husd.com/v0/ast_tree2"
)

/**
 * 错误信息
 * @author hushengdong
 */
type JCErroneous struct {
	*AbstractJCExpression
}

func NewJCError() *AbstractJCExpression {

	jcError := &JCErroneous{}
	jcError.AbstractJCExpression = NewJCExpression()
	jcError.getTreeType = func() *ast_tree2.TreeType {
		return ast_tree2.AST_TREE_KIND_ERRONEOUS
	}
	jcError.getTag = func() JCTreeTag {
		return TREE_TAG_ERRONEOUS
	}
	return jcError.AbstractJCExpression
}