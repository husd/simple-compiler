package lang

/**
 * 引用类型
 * @author hushengdong
 */

// 下面的接口都是引用类型的子类型
type ArrayType interface {
	GetReferenceType() ReferenceType
}

type DeclareType interface {
	GetReferenceType() ReferenceType
}

type NullType interface {
	GetReferenceType() ReferenceType
}

type TypeVariable interface {
	GetReferenceType() ReferenceType
}
