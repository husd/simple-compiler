package ast_tree2

/**
 *
 * @author hushengdong
 */
type LiteralTreeV2 interface {
	GetTreeType() TreeType
	ExpressionTreeV2_()
	LiteralTreeV2_()

	// ---
	GetValue() interface{}
}
