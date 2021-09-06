package lang

/**
 * 元素的类型
 * @author hushengdong
 */
type ElementKindEnum struct {
	Clz   bool
	Inter bool
	Field bool
}

func (e *ElementKindEnum) IsClass() bool {
	return e.Clz
}

func (e *ElementKindEnum) IsInterface() bool {
	return e.Inter
}

func (e *ElementKindEnum) IsField() bool {
	return e.Field
}

var PACKAGE = &ElementKindEnum{false, false, false}
var ENUM = &ElementKindEnum{true, false, false}
var CLASS = &ElementKindEnum{true, false, false}
var ANNOTATION_TYPE = &ElementKindEnum{false, true, false}
var INTERFACE = &ElementKindEnum{false, true, false}
var ENUM_CONSTANT = &ElementKindEnum{false, false, true}
var FIELD = &ElementKindEnum{false, false, true}
var PARAMETER = &ElementKindEnum{false, false, false}
var LOCAL_VARIABLE = &ElementKindEnum{false, false, false}
var EXCEPTION_PARAMETER = &ElementKindEnum{false, false, false}
var METHOD = &ElementKindEnum{false, false, false}
var CONSTRUCTOR = &ElementKindEnum{false, false, false}
var STATIC_INIT = &ElementKindEnum{false, false, false}
var INSTANCE_INIT = &ElementKindEnum{false, false, false}
var TYPE_PARAMETER = &ElementKindEnum{false, false, false}
var OTHER = &ElementKindEnum{false, false, false}
var RESOURCE_VARIABLE = &ElementKindEnum{false, false, false}
