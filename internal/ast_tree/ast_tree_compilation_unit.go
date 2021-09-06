package ast_tree

import "husd.com/v0/lang"

/**
 * 源代码里的所有东西在这里都有 这个结构是一个大而全的结构
 * @author hushengdong
 */
type JCCompilationUnit struct {
	jcTree *JCTree
}

func NewJCCompilationUnit(tree *JCTree) *JCCompilationUnit {

	return &JCCompilationUnit{jcTree: tree}
}

func (jc *JCCompilationUnit) SetType(javaType *lang.JavaType) *JCCompilationUnit {

	jc.jcTree.SetJavaType(javaType)
	return jc
}

func (jc *JCCompilationUnit) SetPos(pos int) *JCCompilationUnit {

	jc.jcTree.SetPos(pos)
	return jc
}
