package jc

import (
	"husd.com/v0/util"
)

/**
 *
 * @author hushengdong
 */
type JCFieldAccess struct {
	*JCExpression
	name *util.Name
}

func NewJCFieldAccess(selected *JCExpression, name *util.Name) *JCFieldAccess {

	res := &JCFieldAccess{selected, name}

	return res
}

func (jc *JCFieldAccess) GetExpression() *JCExpression {

	return jc.JCExpression
}
