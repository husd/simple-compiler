package ast_tree2

/**
 * For example:
 * <pre>
 *   <em>expression</em> instanceof <em>type</em>
 * </pre>
 * @author hushengdong
 */
type InstanceOfTreeV2 interface {
	GetTreeType() TreeType
	ExpressionTreeV2_()
	InstanceOfTreeV2_()
	// --

	GetExpression() ExpressionTreeV2
	GetType() TreeV2
}
