package ast_tree2

/**
 * For example:
 * <pre>
 *   while ( <em>condition</em> )
 *     <em>statement</em>
 * </pre>
 * @author hushengdong
 */
type WhileLoopTreeV2 interface {
	TreeType() *TreeType
	ExpressionTreeV2_()
	WhileLoopTreeV2_()
	// --
	GetCondition() ExpressionTreeV2
	GetStatement() StatementTreeV2
}
