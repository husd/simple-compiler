package lang

/**
 * java里的所有的类型，都在这里了，包含原始类型、声明类型（类和接口）、数组、NULL、Type类型
 * 也包含通配符类型（泛型） ？ 方法签名和返回类型 、包声明的伪类型、void（void比较特殊）它是一种返回类型
 * @author hushengdong
 */
type TypeMirror interface {
	/**
	 * 返回类型的类型 这个是具体的类型
	 */
	GetTypeKind() TypeKind
	/**
	 * 是不是同一个类型
	 */
	Equals(t TypeMirror) bool
}

//引用类型
type TypeMirrorReference interface {
	/**
	 * 属于哪几种类型
	 */
	GetTypeMirrorGroup() TypeMirrorGroup
}

type TypeMirrorGroup string

const (
	Type_group_ReferenceType TypeMirrorGroup = "ReferenceType"
	Type_group_NoType        TypeMirrorGroup = "NoType"
	Type_group_PrimitiveType TypeMirrorGroup = "PrimitiveType"
	Type_group_UnionType     TypeMirrorGroup = "UnionType"
	Type_group_WildcardType  TypeMirrorGroup = "WildcardType" // ? extends Number or ? super T
)
