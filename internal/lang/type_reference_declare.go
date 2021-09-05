package lang

/**
 *
 * @author hushengdong
 */
type DeclareReferenceType struct {
}

func (a *DeclareReferenceType) GetTypeKind() TypeKind {

	return TYPE_KIND_DECLARED
}

func (a *DeclareReferenceType) Equals(t TypeMirror) bool {

	return a == t
}

func (a *DeclareReferenceType) GetTypeMirrorGroup() TypeMirrorGroup {

	return Type_group_ReferenceType
}
