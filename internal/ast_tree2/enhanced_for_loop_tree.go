package ast_tree2

/**
 *
 * For example:
 * <pre>
 *   for ( <em>variable</em> : <em>expression</em> )
 *       <em>statement</em>
 * </pre>
 * @author hushengdong
 */
type EnhancedForLoopTreeV2 interface {
	TreeType() *TreeType
	StatementTreeV2_()

	//--
	GetVariable() VariableTreeV2
	GetExpression() ExpressionTreeV2
	GetStatement() StatementTreeV2
	EnhancedForLoopTreeV2_()
}
