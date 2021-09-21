package ast_tree2

/**
 * For example:
 * <pre>
 *   throw <em>expression</em>;
 * </pre>
 * @author hushengdong
 */
type ThrowTreeV2 interface {
	TreeType() *TreeType
	StatementTreeV2_()
	ThrowTreeV2_()
	// --
	GetExpression() ExpressionTreeV2
}
