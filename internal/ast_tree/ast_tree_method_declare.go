package ast_tree

import "husd.com/v0/lang"

/**
 * 这个主要是方法声明
 * @author hushengdong
 */
type JCMethodDecl struct {
	jcTree *JCTree
}

func NewJCMethodDecl(tree *JCTree) *JCMethodDecl {

	return &JCMethodDecl{jcTree: tree}
}

func (jc *JCMethodDecl) SetType(javaType *lang.JavaType) *JCMethodDecl {

	jc.jcTree.SetJavaType(javaType)
	return jc
}

func (jc *JCMethodDecl) SetPos(pos int) *JCMethodDecl {

	jc.jcTree.SetPos(pos)
	return jc
}
