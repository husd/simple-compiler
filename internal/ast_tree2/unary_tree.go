package ast_tree2

/**
 *
 * @author hushengdong
 */
type UnaryTreeV2 interface {
	TreeType() TreeType
	ExpressionTreeV2_()
	UnaryTreeV2_()
	// --
	GetExpression() ExpressionTreeV2
}
