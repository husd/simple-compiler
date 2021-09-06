package ast_tree2

/**
 * For example:
 * <pre>
 *   <em>expression</em> . <em>identifier</em>
 * </pre>
 * @author hushengdong
 */
type MemberSelectTreeV2 interface {
	TreeType() TreeType
	ExpressionTreeV2_()
	MemberSelectTreeV2_()

	// --
	GetExpression() ExpressionTreeV2
	GetIdentifier() string
}
