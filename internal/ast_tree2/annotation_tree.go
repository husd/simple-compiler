package ast_tree2

/**
 *
 * @author hushengdong
 */
type AnnotationTreeV2 interface {
	GetTreeType() TreeType
	ExpressionTreeV2_()
	AnnotationTreeV2_()
	// --
	GetAnnotationType() TreeV2
	GetArguments() *[]ExpressionTreeV2
}
