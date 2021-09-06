package ast_tree

/**
 * expression
 * @author hushengdong
 */

type JCAnnotation struct {
	expression     *JCExpression
	tag            AstTreeTag     // Either Tag.ANNOTATION or Tag.TYPE_ANNOTATION
	annotationType *Tree          //
	args           []JCExpression //参数
}

func (jc *JCAnnotation) GetKind() *AstTreeNodeKind {

	return treeTag2TreeKind(jc.tag)
}

func (jc *JCAnnotation) Accept(visitor AstTreeVisitor) {
	panic("implement me")
}

func (jc *JCAnnotation) GetTreeType() TreeType {

	return JC_Expression
}
