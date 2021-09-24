package ast_tree2

/**
 *
 * @author hushengdong
 */
type AnnotationTreeV2 interface {
	TreeType() *TreeType
	ExpressionTreeV2_()
	AnnotationTreeV2_()
	// --
	GetAnnotationType() TreeV2
	GetArguments() *[]ExpressionTreeV2
}
