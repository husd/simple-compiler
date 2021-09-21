package ast_tree2

/**
 * 赋值语句 variable = expression
 * 通过接口的定义，可以看到，有2个节点，1个是 variable 1个是 expression
 * @author hushengdong
 */
type AssignmentTreeV2 interface {
	TreeType() *TreeType
	ExpressionTreeV2_()

	GetVariable() ExpressionTreeV2
	GetExpression() ExpressionTreeV2
	AssignmentTreeV2_()
}
