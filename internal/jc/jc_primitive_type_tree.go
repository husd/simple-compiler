package jc

import (
	"husd.com/v0/ast_tree2"
	"husd.com/v0/code"
	"husd.com/v0/lang"
)

/**
 * 原声类型  树结构
 * @author hushengdong
 */
type JCPrimitiveTypeTree struct {
	*AbstractJCExpression
	tg *code.TypeTag
}

func NewJCPrimitiveTypeTree(tg *code.TypeTag) *JCPrimitiveTypeTree {

	res := &JCPrimitiveTypeTree{}
	res.tg = tg
	res.AbstractJCExpression = NewJCExpression()

	res.getTreeType = func() ast_tree2.TreeType {
		return ast_tree2.TREE_TYPE_PRIMITIVE_TYPE
	}
	res.getTag = func() JCTreeTag {
		return TREE_TAG_IDENT
	}

	return res
}

func (jc *JCPrimitiveTypeTree) PrimitiveTypeTreeV2_() {
}

func (jc *JCPrimitiveTypeTree) GetPrimitiveTypeKind() lang.TypeKind {

	return code.GetTypeKindByTypeTag(jc.tg)
}
