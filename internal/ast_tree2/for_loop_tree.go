package ast_tree2

/**
 *
 * @author hushengdong
 */
type ForLoopTreeV2 interface {
	TreeType() TreeType
	StatementTreeV2_()

	// --
	GetInitializer() *[]StatementTreeV2
	GetCondition() ExpressionTreeV2
	GetUpdate() *[]ExpressionStatementTreeV2
	GetStatement() StatementTreeV2
	ForLoopTree_()
}
