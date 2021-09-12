package jc

import (
	"husd.com/v0/ast_tree2"
	"husd.com/v0/util"
)

/**
 *
 * @author hushengdong
 */
type JCIdent struct {
	*JCExpression

	pos    int
	name   *util.Name
	symbol *Symbol
}

func (jc *JCIdent) TreeType() ast_tree2.TreeType {
	panic("implement me")
}

func (jc *JCIdent) ExpressionTreeV2_() {
	panic("implement me")
}

func (jc *JCIdent) GetName() string {
	return jc.name.NameStr
}

func (jc *JCIdent) IdentifierTreeV2_() {
	panic("implement me")
}

func NewJCIdent(name *util.Name) *JCIdent {

	ident := &JCIdent{}
	ident.name = name

	return ident
}
