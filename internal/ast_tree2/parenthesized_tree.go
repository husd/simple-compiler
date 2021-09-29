package ast_tree2

/**
 * For example:
 * <pre>
 *   ( <em>expression</em> )
 * </pre>
 * @author hushengdong
 */
type ParenthesizedTreeV2 interface {
	GetTreeType() TreeType
	ExpressionTreeV2_()
	ParenthesizedTreeV2_()

	// --
	GetExpression() ExpressionTreeV2
}
