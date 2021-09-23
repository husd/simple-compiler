package jc

import (
	"husd.com/v0/lang"
)

/**
 *
 * @author hushengdong
 */
type JCPrimitiveTypeTree struct {
	*AbstractJCExpression
}

func (jc *JCPrimitiveTypeTree) PrimitiveTypeTreeV2_() {

	panic("implement me")
}

func (jc *JCPrimitiveTypeTree) GetPrimitiveTypeKind() lang.TypeKind {

	return lang.TYPE_KIND_INT
}
