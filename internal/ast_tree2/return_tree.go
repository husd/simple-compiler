package ast_tree2

/**
 * For example:
 * <pre>
 *   return;
 *   return <em>expression</em>;
 * </pre>
 * @author hushengdong
 */
type ReturnTreeV2 interface {
	GetTreeType() TreeType
	StatementTreeV2_()
	ReturnTreeV2_()
	// --
	GetExpression() ExpressionTreeV2
}
