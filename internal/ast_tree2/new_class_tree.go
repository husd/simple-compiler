package ast_tree2

/**
 * For example:
 * <pre>
 *   new <em>identifier</em> ( )
 *
 *   new <em>identifier</em> ( <em>arguments</em> )
 *
 *   new <em>typeArguments</em> <em>identifier</em> ( <em>arguments</em> )
 *       <em>classBody</em>
 *
 *   <em>enclosingExpression</em>.new <em>identifier</em> ( <em>arguments</em> )
 * </pre>
 * @author hushengdong
 */
type NewClassTreeV2 interface {
	TreeType() TreeType
	ExpressionTreeV2_()
	NewClassTreeV2_()
	// --
	GetEnclosingExpression() ExpressionTreeV2
	GetTypeArguments() TreeV2
	GetIdentifier() ExpressionTreeV2
	GetArguments() ExpressionTreeV2
	GetClassBody() ClassTreeV2
}
