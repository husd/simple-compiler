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
	msg string
}

func NewJCError(pos int, msg string) *AbstractJCExpression {

	jcError := &JCErroneous{}
	jcError.AbstractJCExpression = NewJCExpression()
	jcError.getTreeType = func() ast_tree2.TreeType {
		return ast_tree2.TREE_TYPE_ERRONEOUS
	}
	jcError.getTag = func() JCTreeTag {
		return TREE_TAG_ERRONEOUS
	}
	jcError.Pos = pos
	jcError.msg = msg
	return jcError.AbstractJCExpression
}
