package lang

/**
 * 原始类型
 * @author hushengdong
 */
type JCPrimitiveType struct {
	kind TypeKind
}

func NewJCPrimitiveType(k TypeKind) *JCPrimitiveType {

	return &JCPrimitiveType{k}
}

func (J *JCPrimitiveType) GetTypeKind() TypeKind {

	return J.kind
}

func (J *JCPrimitiveType) Equals(t TypeMirror) bool {

	return J == t
}

func (a *JCPrimitiveType) GetTypeMirrorGroup() TypeMirrorGroup {

	return Type_group_PrimitiveType
}
