package lang

/**
 * java里的所有的类型，都在这里了，包含原始类型、声明类型（类和接口）、数组、NULL、Type类型
 * 也包含通配符类型（泛型） ？ 方法签名和返回类型 、包声明的伪类型、void（void比较特殊）它是一种返回类型
 * @author hushengdong
 */
type TypeMirror interface {

	/**
	 * 返回类型的类型
	 */
	GetTypeKind() TypeKind
	/**
	 * 是不是同一个类型
	 */
	Equals() bool
}

//下面的类型，都是TypeMirror类型的子类型

type ReferenceType interface {
	/**
	 * 获取到具体的类型
	 */
	GetTypeMirror() TypeMirror
}

/**
 *
 */
type NoType interface {
	/**
	 * 获取到具体的类型
	 */
	GetTypeMirror() TypeMirror
}

//原始类型
type PrimitiveType interface {
	/**
	 * 获取到具体的类型
	 */
	GetTypeMirror() TypeMirror
}

type UnionType interface {
	/**
	 * 获取到具体的类型
	 */
	GetTypeMirror() TypeMirror
}

/**
 *   ?
 *   ? extends Number
 *   ? super T
 */
type WildcardType interface {
	/**
	 * 获取到具体的类型
	 */
	GetTypeMirror() TypeMirror
}
