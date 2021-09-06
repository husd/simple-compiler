package ast_tree2

/**
 * 类声明 	//TODO next
 * @author hushengdong
 */
type ClassTreeV2 interface {
	TreeType() TreeType
	ExpressionTreeV2_()

	//-
	GetModifiers() ModifiersTreeV2
	GetSimpleName() string
	//List<? extends TypeParameterTree> getTypeParameters();
	GetExtendsClause() TreeV2
	GetImplementsClause() *[]TreeV2
	GetMembers() *[]TreeV2
}
