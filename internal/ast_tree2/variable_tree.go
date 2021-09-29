package ast_tree2

/** extends StatementTree
 * For example:
 * <pre>
 *   <em>modifiers</em> <em>type</em> <em>name</em> <em>initializer</em> ;
 *   <em>modifiers</em> <em>type</em> <em>qualified-name</em>.this
 * </pre>
 * @author hushengdong
 */
type VariableTreeV2 interface {
	GetTreeType() TreeType
	StatementTreeV2_()
	VariableTreeV2_()
	// --
	GetModifier() ModifiersTreeV2
	GetName() string
	GetNameExpression() ExpressionTreeV2
	GetVarType() TreeV2
	GetInit() ExpressionTreeV2
}
