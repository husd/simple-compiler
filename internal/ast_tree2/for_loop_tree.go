package ast_tree2

/**
 * For example:
 * <pre>
 *   for ( <em>initializer</em> ; <em>condition</em> ; <em>update</em> )
 *       <em>statement</em>
 * </pre>
 * @author hushengdong
 */
type ForLoopTreeV2 interface {
	GetTreeType() TreeType
	StatementTreeV2_()

	// --
	GetInitializer() *[]StatementTreeV2
	GetCondition() ExpressionTreeV2
	GetUpdate() *[]ExpressionStatementTreeV2
	GetStatement() StatementTreeV2
	ForLoopTree_()
}
