package ast_tree2

/**
 * For example:
 * <pre>
 *   ( <em>type</em> ) <em>expression</em>
 * </pre>
 * @author hushengdong
 */
type TypeCastTreeV2 interface {
	TreeType() TreeType
	ExpressionTreeV2_()
	TypeCastTreeV2_()
	// --
	GetType() TreeV2
	GetExpression() ExpressionTreeV2
}
