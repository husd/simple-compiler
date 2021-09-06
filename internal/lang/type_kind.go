package lang

/**
 *
 * @author hushengdong
 */
type TypeKind int

const (
	TYPE_KIND_BOOLEAN TypeKind = 0
	/**
	 * 原始类型 {@code byte}.
	 */
	TYPE_KIND_BYTE TypeKind = 1

	/**
	 * 原始类型 {@code short}.
	 */
	TYPE_KIND_SHORT TypeKind = 2

	/**
	 * 原始类型 {@code TypeKind}.
	 */
	TYPE_KIND_INT TypeKind = 3

	/**
	 * 原始类型 {@code long}.
	 */
	TYPE_KIND_LONG TypeKind = 4

	/**
	 * 原始类型 {@code char}.
	 */
	TYPE_KIND_CHAR TypeKind = 5

	/**
	 * 原始类型 {@code float}.
	 */
	TYPE_KIND_FLOAT TypeKind = 6

	/**
	 * 原始类型 {@code double}.
	 */
	TYPE_KIND_DOUBLE TypeKind = 7

	/**
	 * 修饰关键字的伪类型 {@code void}.
	 * @see NoType
	 */
	TYPE_KIND_VOID TypeKind = 8

	/**
	 * 没有合适的类型的话，就用none
	 * @see NoType
	 */
	TYPE_KIND_NONE TypeKind = 9

	/**
	 * null类型 常用
	 */
	TYPE_KIND_NULL TypeKind = 10

	/**
	 * 数组类型
	 */
	TYPE_KIND_ARRAY TypeKind = 11

	/**
	 * 类或者接口类型
	 */
	TYPE_KIND_DECLARED TypeKind = 12

	/**
	 * 不能识别的类或者接口
	 */
	TYPE_KIND_ERROR TypeKind = 13

	/**
	 * 类型变量
	 */
	TYPE_KIND_TYPEVAR TypeKind = 14

	/**
	 * 通配符参数
	 */
	TYPE_KIND_WILDCARD TypeKind = 15

	/**
	 * package
	 * @see NoType
	 */
	TYPE_KIND_PACKAGE TypeKind = 16

	/**
	 * 可执行
	 */
	TYPE_KIND_EXECUTABLE TypeKind = 17

	/**
	 * 保留类型，不要用
	 */
	TYPE_KIND_OTHER TypeKind = 18

	/**
	 * A union type.
	 *
	 * @since 1.7
	 */
	TYPE_KIND_UNION TypeKind = 19

	/**
	 * 交叉类型
	 *
	 * @since 1.8
	 */
	TYPE_KIND_INTERSECTION TypeKind = 20
)

func IsPrimitive(typeKind TypeKind) bool {

	// 原始类型就只有这几种
	switch typeKind {
	case TYPE_KIND_BOOLEAN, TYPE_KIND_BYTE, TYPE_KIND_SHORT,
		TYPE_KIND_INT, TYPE_KIND_LONG, TYPE_KIND_CHAR,
		TYPE_KIND_FLOAT, TYPE_KIND_DOUBLE:
		return true
	}
	return false
}
