package ast_tree2

/**
 *
 * @author hushengdong
 */
type IfTreeV2 interface {
	TreeType() TreeType
	StatementTreeV2_()

	// --
	GetCondition() ExpressionTreeV2
	GetThenStatement() StatementTreeV2
	GetElseStatement() StatementTreeV2
}
