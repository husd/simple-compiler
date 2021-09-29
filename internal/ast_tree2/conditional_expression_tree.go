package ast_tree2

/**
 * For example:
 * <pre>
 *   <em>condition</em> ? <em>trueExpression</em> : <em>falseExpression</em>
 * </pre>
 * @author hushengdong
 */
type ConditionalExpressionTreeV2 interface {
	GetTreeType() TreeType
	ExpressionTreeV2_()

	//--
	GetCondition() ExpressionTreeV2
	GetTrueExpression() ExpressionTreeV2
	GetFalseExpression() ExpressionTreeV2
}
