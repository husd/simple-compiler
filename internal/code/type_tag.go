package code

import (
	"fmt"
	"husd.com/v0/lang"
)

/**
 * 这里定义了基本的类型 Java里的所有的类型
 * @author hushengdong
 */

type TypeTag struct {
	numberClass int
	superClass  int
	primitive   bool
	name        string
}

var TYPE_TAG_BYTE = &TypeTag{byte_class, byte_superclasses, true, "byte"}
var TYPE_TAG_CHAR = &TypeTag{char_class, char_superclasses, true, "char"}
var TYPE_TAG_SHORT = &TypeTag{short_class, short_superclasses, true, "short"}
var TYPE_TAG_LONG = &TypeTag{long_class, long_superclasses, true, "long"}
var TYPE_TAG_FLOAT = &TypeTag{float_class, float_superclasses, true, "float"}
var TYPE_TAG_INT = &TypeTag{int_class, int_superclasses, true, "int"}
var TYPE_TAG_DOUBLE = &TypeTag{double_class, double_class, true, "double"}
var TYPE_TAG_BOOLEAN = &TypeTag{0, 0, true, "boolean"}

/** The tag of the type `void'.
 */
var TYPE_TAG_VOID = &TypeTag{0, 0, false, "void"}

/** The tag of all class and interface types.
 */
var TYPE_TAG_CLASS = &TypeTag{0, 0, false, "class"}

// 数组
var TYPE_TAG_ARRAY = &TypeTag{0, 0, false, "array"}

/** The tag of all (monomorphic) method types. 单态方法类型
 */
var TYPE_TAG_METHOD = &TypeTag{0, 0, false, "method"}

/** The tag of all package "types".
 */
var TYPE_TAG_PACKAGE = &TypeTag{0, 0, false, "package"}

/** The tag of all (source-level) type variables.
 */
var TYPE_TAG_TYPEVAR = &TypeTag{0, 0, false, "typevar"}

/** The tag of all type arguments. 范型类型
 */
var TYPE_TAG_WILDCARD = &TypeTag{0, 0, false, "wildcard"}

/** The tag of all polymorphic (method-) types. 多态方法类型
 */
var TYPE_TAG_FORALL = &TypeTag{0, 0, false, "forall"}

/** The tag of deferred expression types in method context
 */
var TYPE_TAG_DEFERRED = &TypeTag{0, 0, false, "deferred"}

/** The tag of the bottom type {@code <null>}.
 */
var TYPE_TAG_BOT = &TypeTag{0, 0, false, "bot"}
var TYPE_TAG_NONE = &TypeTag{0, 0, false, "none"}
var TYPE_TAG_ERROR = &TypeTag{0, 0, false, "error"}
var TYPE_TAG_UNKNOWN = &TypeTag{0, 0, false, "unknown"}
var TYPE_TAG_UNDETVAR = &TypeTag{0, 0, false, "undetvar"}

/** Pseudo-types, these are special tags
 */
var TYPE_TAG_UNINITIALIZED_THIS = &TypeTag{0, 0, false, "TYPE_TAG_UNINITIALIZED_THIS"}
var TYPE_TAG_UNINITIALIZED_OBJECT = &TypeTag{0, 0, false, "TYPE_TAG_UNINITIALIZED_OBJECT"}

// 这个规则决定了隐士类型转换 byte可以自动转 short int 等它的父亲类型，反之则不可以
const byte_class = 1
const char_class = 2
const short_class = 4
const int_class = 8
const long_class = 16
const float_class = 32
const double_class = 64

const byte_superclasses = byte_class | short_class | int_class | long_class | float_class | double_class
const char_superclasses = char_class | int_class | long_class | float_class | double_class
const short_superclasses = short_class | int_class | long_class | float_class | double_class
const int_superclasses = int_class | long_class | float_class | double_class
const long_superclasses = long_class | float_class | double_class
const float_superclasses = float_class | double_class

func GetTypeKindByTypeTag(tg *TypeTag) lang.TypeKind {

	switch tg {
	case TYPE_TAG_BOOLEAN:
		return lang.TYPE_KIND_BOOLEAN
	case TYPE_TAG_BYTE:
		return lang.TYPE_KIND_BYTE
	case TYPE_TAG_SHORT:
		return lang.TYPE_KIND_SHORT
	case TYPE_TAG_INT:
		return lang.TYPE_KIND_INT
	case TYPE_TAG_LONG:
		return lang.TYPE_KIND_LONG
	case TYPE_TAG_CHAR:
		return lang.TYPE_KIND_CHAR
	case TYPE_TAG_FLOAT:
		return lang.TYPE_KIND_FLOAT
	case TYPE_TAG_DOUBLE:
		return lang.TYPE_KIND_DOUBLE
	case TYPE_TAG_VOID:
		return lang.TYPE_KIND_VOID
	default:
		panic(fmt.Sprintf("unknown primitive type : %v", tg))
	}
}
