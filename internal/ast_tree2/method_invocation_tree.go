package ast_tree2

/**
 *
 * @author hushengdong
 */
type MethodInvocationTreeV2 interface {
	GetTreeType() TreeType
	ExpressionTreeV2_()
	MethodInvocationTreeV2_()

	// --
	GetTypeArguments() *[]TreeV2
	GetMethodSelect() ExpressionTreeV2
	GetArguments() ExpressionTreeV2
}
