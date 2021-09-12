package ast_tree2

/**
 * For example:
 * <pre>
 *   <em>name</em>
 * </pre>
 * @author hushengdong
 */
type IdentifierTreeV2 interface {
	TreeType() TreeType
	ExpressionTreeV2_()
	IdentifierTreeV2_()

	//--
	GetName() string
}
