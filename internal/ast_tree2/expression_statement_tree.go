package ast_tree2

/**
 * For example:
 * <pre>
 *   <em>expression</em> ;
 * </pre>
 *
 * @jls section 14.8
 * @author hushengdong
 */
type ExpressionStatementTreeV2 interface {
	TreeType() *TreeType
	StatementTreeV2_()

	GetExpression() ExpressionTreeV2
	ExpressionStatementTreeV2_()
}
