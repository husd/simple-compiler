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
}

func (jc *JCLiteral) GetValue() interface{} {

	switch jc.tg {
	case code.TYPE_TAG_BOOLEAN:
		return jc.val.(int) != 0
	case code.TYPE_TAG_INT:
		return jc.val.(int32)
	case code.TYPE_TAG_LONG:
		return jc.val.(int64)
	case code.TYPE_TAG_BYTE:
		return jc.val.(int8)
	case code.TYPE_TAG_CHAR:
		return jc.val.(int16)
	case code.TYPE_TAG_FLOAT:
		return jc.val.(float32)
	case code.TYPE_TAG_DOUBLE:
		return jc.val.(float64)
	default:
		return jc.val
	}
}

func NewJCLiteral(tg *code.TypeTag, v int) *JCLiteral {

	res := &JCLiteral{}
	res.AbstractJCExpression = NewJCExpression()

	res.tg = tg
	res.val = v

	res.getTreeType = func() *ast_tree2.TreeType {
		return ast_tree2.TREE_TYPE_INT_LITERAL
	}
	res.getTag = func() JCTreeTag {
		return TREE_TAG_LITERAL
	}
	return res
}
