package lang

/**
 *
 * @author hushengdong
 */
type NullReferenceType struct {
}

func (a *NullReferenceType) GetTypeKind() TypeKind {

	return TYPE_KIND_NULL
}

func (a *NullReferenceType) Equals(t TypeMirror) bool {

	return a == t
}

func (a *NullReferenceType) GetTypeMirrorGroup() TypeMirrorGroup {

	return Type_group_ReferenceType
}
