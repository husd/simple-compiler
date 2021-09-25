package ast_tree2

/**
 * arr[0] 这样的访问数组数据的表达式
 * @author hushengdong
 */
type ArrayAccessTreeV2 interface {
	TreeType() *TreeType
	ExpressionTreeV2_()

	// --
	GetExpression() ExpressionTreeV2
	GetIndex() ExpressionTreeV2
}
