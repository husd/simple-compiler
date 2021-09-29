package jc

import (
	"husd.com/v0/ast_tree2"
)

/**
 * 注解
 * @author hushengdong
 */
type JCAnnotation struct {
	*AbstractJCExpression
	tag            JCTreeTag // Either Tag.ANNOTATION or Tag.TYPE_ANNOTATION
	annotationType *AbstractJCTree
	args           *[]ast_tree2.ExpressionTreeV2 // 参数
}

func NewJCAnnotation(tag JCTreeTag, annotationType *AbstractJCTree, args *[]ast_tree2.ExpressionTreeV2) *JCAnnotation {

	res := &JCAnnotation{}
	res.AbstractJCExpression = NewJCExpression()

	res.tag = tag
	res.annotationType = annotationType
	res.args = args

	res.getTreeType = func() ast_tree2.TreeType {
		return ast_tree2.TREE_TYPE_ANNOTATION
	}
	res.getTag = func() JCTreeTag {
		return res.tag
	}
	res.toString = func() string {
		return "annotation"
	}

	return res
}

// AnnotationTree
func (jc *JCAnnotation) AnnotationTreeV2_() {
}

func (jc *JCAnnotation) GetAnnotationType() ast_tree2.TreeV2 {

	return jc.annotationType.getTree()
}

func (jc *JCAnnotation) GetArguments() *[]ast_tree2.ExpressionTreeV2 {

	return jc.args
}

// AnnotationTree
