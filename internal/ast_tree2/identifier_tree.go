package ast_tree2

/**
 *
 * @author hushengdong
 */
type IdentifierTreeV2 interface {
	TreeType() TreeType
	ExpressionTreeV2_()

	//--
	GetName() string
	IdentifierTreeV2_()
}
