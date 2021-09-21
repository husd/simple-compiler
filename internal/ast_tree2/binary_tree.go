package ast_tree2

/**
 * a + b 这样的
 * extends ExpressionTreeV2
 * @author hushengdong
 */
type BinaryTreeV2 interface {
	TreeType() *TreeType
	ExpressionTreeV2_()

	//--
	GetLeftOperand() ExpressionTreeV2
	GetRightOperand() ExpressionTreeV2
}
