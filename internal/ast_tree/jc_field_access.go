package ast_tree

import "husd.com/v0/util"

/**
 *
 * @author hushengdong
 */
type JCFieldAccess struct {
	expression *JCExpression

	selected *JCExpression
	name     *util.Name
}

func NewJCFieldAccess(selected *JCExpression, name *util.Name) *JCFieldAccess {

	res := &JCFieldAccess{}
	res.selected = selected
	res.name = name
	res.expression = NewJCExpression(selected.jcTree)

	return res
}

func (jc *JCFieldAccess) GetExpression() *JCExpression {

	return jc.selected
}
