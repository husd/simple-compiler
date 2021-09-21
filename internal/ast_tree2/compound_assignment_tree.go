package ast_tree2

/**
 * For example:
 * <pre>
 *   <em>variable</em> <em>operator</em> <em>expression</em>
 * </pre>
 * @author hushengdong
 */
type CompoundAssignmentTreeV2 interface {
	TreeType() *TreeType
	ExpressionTreeV2_()

	//---
	GetVariable() ExpressionTreeV2
	GetExpression() ExpressionTreeV2
	CompoundAssignmentTreeV2_()
}
