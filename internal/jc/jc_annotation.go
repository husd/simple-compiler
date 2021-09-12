package jc

/**
 * expression
 * @author hushengdong
 */

type JCAnnotation struct {
	*JCExpression
	tag  JCTreeTag      // Either Tag.ANNOTATION or Tag.TYPE_ANNOTATION
	args []JCExpression //参数
}
