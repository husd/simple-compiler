package jc

import (
	"husd.com/v0/util"
)

/**
 *
 * @author hushengdong
 */
type JCFieldAccess struct {
	*AbstractJCExpression
	name *util.Name
}

func NewJCFieldAccess(selected *AbstractJCExpression, name *util.Name) *JCFieldAccess {

	res := &JCFieldAccess{selected, name}

	return res
}

func (jc *JCFieldAccess) GetExpression() *AbstractJCExpression {

	return jc.AbstractJCExpression
}
