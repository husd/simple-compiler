package ast_tree2

/**
 * AST的抽象定义，这里的代码应该是和语言无关的，理论上可以轻松迁移到其它语言。
 * 所有的Tree都必须实现这个接口。
 * @author hushengdong
 */
type TreeV2 interface {
	/**
	 * 树的节点，有不同的类型
	 */
	TreeType() *TreeType
}

/**
 * 抽象语法树的节点的类型，暂时先写个枚举
 */
type TreeType struct {
	Name  string // 名字
	Index int    // 索引值
}

var TREE_TYPE_NIL = &TreeType{"nil", -1}
var TREE_TYPE_ANNOTATED_TYPE = &TreeType{"annotated_type", 0}
var TREE_TYPE_ANNOTATION = &TreeType{"annotation", 1}
var TREE_TYPE_TYPE_ANNOTATION = &TreeType{"type_annotation", 2}
var TREE_TYPE_ARRAY_ACCESS = &TreeType{"array_access", 3}
var TREE_TYPE_ARRAY_TYPE = &TreeType{"array_type", 4}
var TREE_TYPE_ASSERT = &TreeType{"assert", 5}
var TREE_TYPE_ASSIGNMENT = &TreeType{"assignment", 6}
var TREE_TYPE_BLOCK = &TreeType{"block", 7}
var TREE_TYPE_BREAK = &TreeType{"break", 8}
var TREE_TYPE_CASE = &TreeType{"case", 9}
var TREE_TYPE_CATCH = &TreeType{"catch", 10}
var TREE_TYPE_CLASS = &TreeType{"class", 11}
var TREE_TYPE_COMPILATION_UNIT = &TreeType{"compilation_unit", 12}
var TREE_TYPE_CONDITIONAL_EXPRESSION = &TreeType{"conditional_expression", 13}
var TREE_TYPE_CONTINUE = &TreeType{"continue", 14}
var TREE_TYPE_DO_WHILE_LOOP = &TreeType{"do_while_loop", 15}
var TREE_TYPE_ENHANCED_FOR_LOOP = &TreeType{"enhanced_for_loop", 16}
var TREE_TYPE_EXPRESSION_STATEMENT = &TreeType{"expression_statement", 17}
var TREE_TYPE_MEMBER_SELECT = &TreeType{"member_select", 18}
var TREE_TYPE_MEMBER_REFERENCE = &TreeType{"member_reference", 19}
var TREE_TYPE_FOR_LOOP = &TreeType{"for_loop", 20}
var TREE_TYPE_IDENTIFIER = &TreeType{"identifier", 21}
var TREE_TYPE_IF = &TreeType{"if", 22}
var TREE_TYPE_IMPORT = &TreeType{"import", 23}
var TREE_TYPE_INSTANCE_OF = &TreeType{"instance_of", 24}
var TREE_TYPE_LABELED_STATEMENT = &TreeType{"labeled_statement", 25}
var TREE_TYPE_METHOD = &TreeType{"method", 26}
var TREE_TYPE_METHOD_INVOCATION = &TreeType{"method_invocation", 27}
var TREE_TYPE_MODIFIERS = &TreeType{"modifiers", 28}
var TREE_TYPE_NEW_ARRAY = &TreeType{"new_array", 29}
var TREE_TYPE_NEW_CLASS = &TreeType{"new_class", 30}
var TREE_TYPE_LAMBDA_EXPRESSION = &TreeType{"lambda_expression", 31}
var TREE_TYPE_PARENTHESIZED = &TreeType{"parenthesized", 32}
var TREE_TYPE_PRIMITIVE_TYPE = &TreeType{"primitive_type", 33}
var TREE_TYPE_RETURN = &TreeType{"return", 34}
var TREE_TYPE_EMPTY_STATEMENT = &TreeType{"empty_statement", 35}
var TREE_TYPE_SWITCH = &TreeType{"switch", 36}
var TREE_TYPE_SYNCHRONIZED = &TreeType{"synchronized", 37}
var TREE_TYPE_THROW = &TreeType{"throw", 38}
var TREE_TYPE_TRY = &TreeType{"try", 39}
var TREE_TYPE_PARAMETERIZED_TYPE = &TreeType{"parameterized_type", 40}
var TREE_TYPE_UNION_TYPE = &TreeType{"union_type", 41}
var TREE_TYPE_INTERSECTION_TYPE = &TreeType{"intersection_type", 42}
var TREE_TYPE_TYPE_CAST = &TreeType{"type_cast", 43}
var TREE_TYPE_TYPE_PARAMETER = &TreeType{"type_parameter", 44}
var TREE_TYPE_VARIABLE = &TreeType{"variable", 45}
var TREE_TYPE_WHILE_LOOP = &TreeType{"while_loop", 46}
var TREE_TYPE_POSTFIX_INCREMENT = &TreeType{"postfix_increment", 47}
var TREE_TYPE_POSTFIX_DECREMENT = &TreeType{"postfix_decrement", 48}
var TREE_TYPE_PREFIX_INCREMENT = &TreeType{"prefix_increment", 49}
var TREE_TYPE_PREFIX_DECREMENT = &TreeType{"prefix_decrement", 50}
var TREE_TYPE_UNARY_PLUS = &TreeType{"unary_plus", 51}
var TREE_TYPE_UNARY_MINUS = &TreeType{"unary_minus", 52}
var TREE_TYPE_BITWISE_COMPLEMENT = &TreeType{"bitwise_complement", 53}
var TREE_TYPE_LOGICAL_COMPLEMENT = &TreeType{"logical_complement", 54}
var TREE_TYPE_MULTIPLY = &TreeType{"multiply", 55}
var TREE_TYPE_DIVIDE = &TreeType{"divide", 56}
var TREE_TYPE_REMAINDER = &TreeType{"remainder", 57}
var TREE_TYPE_PLUS = &TreeType{"plus", 58}
var TREE_TYPE_MINUS = &TreeType{"minus", 59}
var TREE_TYPE_LEFT_SHIFT = &TreeType{"left_shift", 60}
var TREE_TYPE_RIGHT_SHIFT = &TreeType{"right_shift", 61}
var TREE_TYPE_UNSIGNED_RIGHT_SHIFT = &TreeType{"unsigned_right_shift", 62}
var TREE_TYPE_LESS_THAN = &TreeType{"less_than", 63}
var TREE_TYPE_GREATER_THAN = &TreeType{"greater_than", 64}
var TREE_TYPE_LESS_THAN_EQUAL = &TreeType{"less_than_equal", 65}
var TREE_TYPE_GREATER_THAN_EQUAL = &TreeType{"greater_than_equal", 66}
var TREE_TYPE_EQUAL_TO = &TreeType{"equal_to", 67}
var TREE_TYPE_NOT_EQUAL_TO = &TreeType{"not_equal_to", 68}
var TREE_TYPE_AND = &TreeType{"and", 69}
var TREE_TYPE_XOR = &TreeType{"xor", 70}
var TREE_TYPE_OR = &TreeType{"or", 71}
var TREE_TYPE_CONDITIONAL_AND = &TreeType{"conditional_and", 72}
var TREE_TYPE_CONDITIONAL_OR = &TreeType{"conditional_or", 73}
var TREE_TYPE_MULTIPLY_ASSIGNMENT = &TreeType{"multiply_assignment", 74}
var TREE_TYPE_DIVIDE_ASSIGNMENT = &TreeType{"divide_assignment", 75}
var TREE_TYPE_REMAINDER_ASSIGNMENT = &TreeType{"remainder_assignment", 76}
var TREE_TYPE_PLUS_ASSIGNMENT = &TreeType{"plus_assignment", 77}
var TREE_TYPE_MINUS_ASSIGNMENT = &TreeType{"minus_assignment", 78}
var TREE_TYPE_LEFT_SHIFT_ASSIGNMENT = &TreeType{"left_shift_assignment", 79}
var TREE_TYPE_RIGHT_SHIFT_ASSIGNMENT = &TreeType{"right_shift_assignment", 80}
var TREE_TYPE_UNSIGNED_RIGHT_SHIFT_ASSIGNMENT = &TreeType{"unsigned_right_shift_assignment", 81}
var TREE_TYPE_AND_ASSIGNMENT = &TreeType{"and_assignment", 82}
var TREE_TYPE_XOR_ASSIGNMENT = &TreeType{"xor_assignment", 83}
var TREE_TYPE_OR_ASSIGNMENT = &TreeType{"or_assignment", 84}
var TREE_TYPE_INT_LITERAL = &TreeType{"int_literal", 85}
var TREE_TYPE_LONG_LITERAL = &TreeType{"long_literal", 86}
var TREE_TYPE_FLOAT_LITERAL = &TreeType{"float_literal", 87}
var TREE_TYPE_DOUBLE_LITERAL = &TreeType{"double_literal", 88}
var TREE_TYPE_BOOLEAN_LITERAL = &TreeType{"boolean_literal", 89}
var TREE_TYPE_CHAR_LITERAL = &TreeType{"char_literal", 90}
var TREE_TYPE_STRING_LITERAL = &TreeType{"string_literal", 91}
var TREE_TYPE_NULL_LITERAL = &TreeType{"null_literal", 92}
var TREE_TYPE_UNBOUNDED_WILDCARD = &TreeType{"unbounded_wildcard", 93}
var TREE_TYPE_EXTENDS_WILDCARD = &TreeType{"extends_wildcard", 94}
var TREE_TYPE_SUPER_WILDCARD = &TreeType{"super_wildcard", 95}
var TREE_TYPE_ERRONEOUS = &TreeType{"erroneous", 96}
var TREE_TYPE_INTERFACE = &TreeType{"interface", 97}
var TREE_TYPE_ENUM = &TreeType{"enum", 98}
var TREE_TYPE_ANNOTATION_TYPE = &TreeType{"annotation_type", 99}
var TREE_TYPE_OTHER = &TreeType{"other", 100}

// 树的节点
type AstTreeNode struct {
	nodeType TreeType
	msg      string // 描述信息
}
