package ast_tree2

/**
 * For example:
 * <pre>
 *   new <em>type</em> <em>dimensions</em> <em>initializers</em>
 *
 *   new <em>type</em> <em>dimensions</em> [ ] <em>initializers</em>
 * </pre>
 * @author hushengdong
 */
type NewArrayTreeV2 interface {
	TreeType() *TreeType
	ExpressionTreeV2_()
	NewArrayTreeV2_()
	// --
	GetType() TreeV2
	GetDimensions() *[]ExpressionTreeV2
	GetInitializers() *[]ExpressionTreeV2
	GetAnnotations() *[]AnnotationTreeV2
	GetDimAnnotations() *[]AnnotationTreeV2
}
