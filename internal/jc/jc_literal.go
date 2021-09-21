package jc

import (
	"husd.com/v0/ast_tree2"
	"husd.com/v0/code"
)

/**
 *
 * @author hushengdong
 */
type JCLiteral struct {
	*AbstractJCExpression

	tg  *code.TypeTag
	val interface{}
}

func (jc *JCLiteral) LiteralTreeV2_() {
	//panic("implement me")
}

func (jc *JCLiteral) GetValue() interface{} {

	switch jc.tg {
	case code.TYPE_TAG_BOOL:
		return jc.val.(int) != 0
	default:
		return jc.val
	}
}

func NewJCLiteral(tg *code.TypeTag, v int) *JCLiteral {

	res := &JCLiteral{}
	res.tg = tg
	res.val = v

	res.getTreeType = func() *ast_tree2.TreeType {
		return ast_tree2.AST_TREE_KIND_INT_LITERAL
	}
	res.getTag = func() JCTreeTag {
		return TREE_TAG_LITERAL
	}
	return res
}
