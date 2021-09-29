package ast_tree2

/**
 * 赋值语句 variable = expression
 * 通过接口的定义，可以看到，有2个节点，1个是 variable 1个是 expression
 * @author hushengdong
 */
type AssignmentTreeV2 interface {
	GetTreeType() TreeType
	ExpressionTreeV2_()

	/**
	 * 变量就是左边的 例如 : a=10; 里的 a
	 */
	GetVariable() ExpressionTreeV2
	/**
	 * 表达式就是右边的表达式 例如： a=10; 里面的10
	 * 可能是字面量，也可能是另外的一个表达式，例如: a = sum(10);
	 */
	GetExpression() ExpressionTreeV2
	AssignmentTreeV2_()
}
