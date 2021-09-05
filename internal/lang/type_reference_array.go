package lang

/**
 *
 * @author hushengdong
 */
type ArrayReferenceType struct {
}

func (a *ArrayReferenceType) GetTypeKind() TypeKind {

	return TYPE_KIND_ARRAY
}

func (a *ArrayReferenceType) Equals(t TypeMirror) bool {

	return a == t
}

func (a *ArrayReferenceType) GetTypeMirrorGroup() TypeMirrorGroup {

	return Type_group_ReferenceType
}
