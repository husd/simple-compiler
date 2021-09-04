package ast_tree

import "husd.com/v0/lang"

/**
 *
 * @author hushengdong
 */
type JCTree struct {
	Pos      int //position in the source file
	JavaType *lang.JavaType
	tag      AstTreeTag
}

func (jc *JCTree) SetJavaType(t *lang.JavaType) *JCTree {

	jc.JavaType = t
	return jc
}

func (jc *JCTree) SetPos(pos int) *JCTree {

	jc.Pos = pos
	return jc
}

func (jc *JCTree) HasTag(tag AstTreeTag) bool {

	return jc.tag == tag
}

func (J *JCTree) GetKind() *AstTreeNodeKind {

	panic("implement me")
}

func (J *JCTree) Accept(visitor AstTreeVisitor) {

	panic("implement me")
}

func (J *JCTree) GetTreeType() TreeType {

	panic("implement me")
}
