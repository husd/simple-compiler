package ast_tree2

/**
 * For example:
 * <pre>
 *   do
 *       <em>statement</em>
 *   while ( <em>expression</em> );
 * </pre>
 * @author hushengdong
 */
type DoWhileLoopTreeV2 interface {
	GetTreeType() TreeType
	StatementTreeV2_()

	//-
	GetCondition() ExpressionTreeV2
	GetStatement() StatementTreeV2
	DoWhileLoopTreeV2_()
}
