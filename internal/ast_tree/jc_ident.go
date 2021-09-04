package ast_tree

import "husd.com/v0/util"

/**
 *
 * @author hushengdong
 */
type JCIdent struct {
	JcTree *JCTree

	pos      int
	treeType TreeType

	name   *util.Name
	symbol *Symbol
}

func NewJCIdent(name *util.Name) *JCIdent {

	ident := &JCIdent{}
	ident.name = name

	return ident
}

func (jc *JCIdent) GetKind() *AstTreeNodeKind {

	return TREE_IDENTIFIER
}

func (jc *JCIdent) Accept(visitor AstTreeVisitor) {
	panic("implement me")
}

func (jc *JCIdent) GetTreeType() TreeType {

	return JC_Expression
}

func (jc *JCIdent) GetName() *util.Name {

	return jc.name
}
