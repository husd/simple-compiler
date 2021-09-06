package lang

/**
 * 通配符类型的类似 List<? extends Person> 这样的
 * @author hushengdong
 */
type VariableReferenceType struct {
}

func (a *VariableReferenceType) GetTypeKind() TypeKind {

	return TYPE_KIND_VOID
}

func (a *VariableReferenceType) Equals(t TypeMirror) bool {

	return a == t
}

func (a *VariableReferenceType) GetTypeMirrorGroup() TypeMirrorGroup {

	return Type_group_ReferenceType
}
