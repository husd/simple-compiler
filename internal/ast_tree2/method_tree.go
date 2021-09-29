package ast_tree2

/**
 * For example:
 * <pre>
 *   <em>modifiers</em> <em>typeParameters</em> <em>type</em> <em>name</em>
 *      ( <em>parameters</em> )
 *      <em>body</em>
 *
 *   <em>modifiers</em> <em>type</em> <em>name</em> () default <em>defaultValue</em>
 * </pre>
 * @author hushengdong
 */
type MethodTreeV2 interface {
	GetTreeType() TreeType
	MethodTreeV2_()

	// --
	GetModifiers() ModifiersTreeV2
	GetName() string
	GetReturnType() TreeV2
	GetTypeParameters() *[]TypeParameterTreeV2
	GetParameters() *[]VariableTreeV2
	GetReceiverParameter() VariableTreeV2
	GetThrows() *[]ExpressionTreeV2
	GetBody() BlockTreeV2
	GetDefaultValue() TreeV2 // 注解类型设计的
}
