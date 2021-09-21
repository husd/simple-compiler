package jc

/**
 * expression
 * @author hushengdong
 */

type JCAnnotation struct {
	*AbstractJCExpression
	tag  JCTreeTag              // Either Tag.ANNOTATION or Tag.TYPE_ANNOTATION
	args []AbstractJCExpression //参数
}
