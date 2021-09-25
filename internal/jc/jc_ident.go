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
	*AbstractJCExpression

	pos    int
	name   *util.Name
	symbol *Symbol
}

func (jc *JCIdent) TreeType() ast_tree2.TreeType {
	panic("implement me")
}

func (jc *JCIdent) ExpressionTreeV2_() {
	//panic("implement me")
}

func (jc *JCIdent) GetName() string {
	return jc.name.NameStr
}

func (jc *JCIdent) IdentifierTreeV2_() {
	//panic("implement me")
}

func NewJCIdent(name *util.Name) *JCIdent {

	ident := &JCIdent{}
	ident.name = name

	ident.getTreeType = func() *ast_tree2.TreeType {
		return ast_tree2.TREE_TYPE_IDENTIFIER
	}
	ident.getTag = func() JCTreeTag {
		return TREE_TAG_IDENT
	}
	return ident
}
