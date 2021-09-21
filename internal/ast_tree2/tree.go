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
	Index int    //索引值
}

var AST_TREE_KIND_NIL = &TreeType{"nil", -1}

var AST_TREE_KIND_ANNOTATED_TYPE = &TreeType{"annotated_type", 0}
var AST_TREE_KIND_ANNOTATION = &TreeType{"annotation", 1}
var AST_TREE_KIND_TYPE_ANNOTATION = &TreeType{"type_annotation", 2}
var AST_TREE_KIND_ARRAY_ACCESS = &TreeType{"array_access", 3}
var AST_TREE_KIND_ARRAY_TYPE = &TreeType{"array_type", 4}
var AST_TREE_KIND_ASSERT = &TreeType{"assert", 5}
var AST_TREE_KIND_ASSIGNMENT = &TreeType{"assignment", 6}
var AST_TREE_KIND_BLOCK = &TreeType{"block", 7}
var AST_TREE_KIND_BREAK = &TreeType{"break", 8}
var AST_TREE_KIND_CASE = &TreeType{"case", 9}
var AST_TREE_KIND_CATCH = &TreeType{"catch", 10}
var AST_TREE_KIND_CLASS = &TreeType{"class", 11}
var AST_TREE_KIND_COMPILATION_UNIT = &TreeType{"compilation_unit", 12}
var AST_TREE_KIND_CONDITIONAL_EXPRESSION = &TreeType{"conditional_expression", 13}
var AST_TREE_KIND_CONTINUE = &TreeType{"continue", 14}
var AST_TREE_KIND_DO_WHILE_LOOP = &TreeType{"do_while_loop", 15}
var AST_TREE_KIND_ENHANCED_FOR_LOOP = &TreeType{"enhanced_for_loop", 16}
var AST_TREE_KIND_EXPRESSION_STATEMENT = &TreeType{"expression_statement", 17}
var AST_TREE_KIND_MEMBER_SELECT = &TreeType{"member_select", 18}
var AST_TREE_KIND_MEMBER_REFERENCE = &TreeType{"member_reference", 19}
var AST_TREE_KIND_FOR_LOOP = &TreeType{"for_loop", 20}
var AST_TREE_KIND_IDENTIFIER = &TreeType{"identifier", 21}
var AST_TREE_KIND_IF = &TreeType{"if", 22}
var AST_TREE_KIND_IMPORT = &TreeType{"import", 23}
var AST_TREE_KIND_INSTANCE_OF = &TreeType{"instance_of", 24}
var AST_TREE_KIND_LABELED_STATEMENT = &TreeType{"labeled_statement", 25}
var AST_TREE_KIND_METHOD = &TreeType{"method", 26}
var AST_TREE_KIND_METHOD_INVOCATION = &TreeType{"method_invocation", 27}
var AST_TREE_KIND_MODIFIERS = &TreeType{"modifiers", 28}
var AST_TREE_KIND_NEW_ARRAY = &TreeType{"new_array", 29}
var AST_TREE_KIND_NEW_CLASS = &TreeType{"new_class", 30}
var AST_TREE_KIND_LAMBDA_EXPRESSION = &TreeType{"lambda_expression", 31}
var AST_TREE_KIND_PARENTHESIZED = &TreeType{"parenthesized", 32}
var AST_TREE_KIND_PRIMITIVE_TYPE = &TreeType{"primitive_type", 33}
var AST_TREE_KIND_RETURN = &TreeType{"return", 34}
var AST_TREE_KIND_EMPTY_STATEMENT = &TreeType{"empty_statement", 35}
var AST_TREE_KIND_SWITCH = &TreeType{"switch", 36}
var AST_TREE_KIND_SYNCHRONIZED = &TreeType{"synchronized", 37}
var AST_TREE_KIND_THROW = &TreeType{"throw", 38}
var AST_TREE_KIND_TRY = &TreeType{"try", 39}
var AST_TREE_KIND_PARAMETERIZED_TYPE = &TreeType{"parameterized_type", 40}
var AST_TREE_KIND_UNION_TYPE = &TreeType{"union_type", 41}
var AST_TREE_KIND_INTERSECTION_TYPE = &TreeType{"intersection_type", 42}
var AST_TREE_KIND_TYPE_CAST = &TreeType{"type_cast", 43}
var AST_TREE_KIND_TYPE_PARAMETER = &TreeType{"type_parameter", 44}
var AST_TREE_KIND_VARIABLE = &TreeType{"variable", 45}
var AST_TREE_KIND_WHILE_LOOP = &TreeType{"while_loop", 46}
var AST_TREE_KIND_POSTFIX_INCREMENT = &TreeType{"postfix_increment", 47}
var AST_TREE_KIND_POSTFIX_DECREMENT = &TreeType{"postfix_decrement", 48}
var AST_TREE_KIND_PREFIX_INCREMENT = &TreeType{"prefix_increment", 49}
var AST_TREE_KIND_PREFIX_DECREMENT = &TreeType{"prefix_decrement", 50}
var AST_TREE_KIND_UNARY_PLUS = &TreeType{"unary_plus", 51}
var AST_TREE_KIND_UNARY_MINUS = &TreeType{"unary_minus", 52}
var AST_TREE_KIND_BITWISE_COMPLEMENT = &TreeType{"bitwise_complement", 53}
var AST_TREE_KIND_LOGICAL_COMPLEMENT = &TreeType{"logical_complement", 54}
var AST_TREE_KIND_MULTIPLY = &TreeType{"multiply", 55}
var AST_TREE_KIND_DIVIDE = &TreeType{"divide", 56}
var AST_TREE_KIND_REMAINDER = &TreeType{"remainder", 57}
var AST_TREE_KIND_PLUS = &TreeType{"plus", 58}
var AST_TREE_KIND_MINUS = &TreeType{"minus", 59}
var AST_TREE_KIND_LEFT_SHIFT = &TreeType{"left_shift", 60}
var AST_TREE_KIND_RIGHT_SHIFT = &TreeType{"right_shift", 61}
var AST_TREE_KIND_UNSIGNED_RIGHT_SHIFT = &TreeType{"unsigned_right_shift", 62}
var AST_TREE_KIND_LESS_THAN = &TreeType{"less_than", 63}
var AST_TREE_KIND_GREATER_THAN = &TreeType{"greater_than", 64}
var AST_TREE_KIND_LESS_THAN_EQUAL = &TreeType{"less_than_equal", 65}
var AST_TREE_KIND_GREATER_THAN_EQUAL = &TreeType{"greater_than_equal", 66}
var AST_TREE_KIND_EQUAL_TO = &TreeType{"equal_to", 67}
var AST_TREE_KIND_NOT_EQUAL_TO = &TreeType{"not_equal_to", 68}
var AST_TREE_KIND_AND = &TreeType{"and", 69}
var AST_TREE_KIND_XOR = &TreeType{"xor", 70}
var AST_TREE_KIND_OR = &TreeType{"or", 71}
var AST_TREE_KIND_CONDITIONAL_AND = &TreeType{"conditional_and", 72}
var AST_TREE_KIND_CONDITIONAL_OR = &TreeType{"conditional_or", 73}
var AST_TREE_KIND_MULTIPLY_ASSIGNMENT = &TreeType{"multiply_assignment", 74}
var AST_TREE_KIND_DIVIDE_ASSIGNMENT = &TreeType{"divide_assignment", 75}
var AST_TREE_KIND_REMAINDER_ASSIGNMENT = &TreeType{"remainder_assignment", 76}
var AST_TREE_KIND_PLUS_ASSIGNMENT = &TreeType{"plus_assignment", 77}
var AST_TREE_KIND_MINUS_ASSIGNMENT = &TreeType{"minus_assignment", 78}
var AST_TREE_KIND_LEFT_SHIFT_ASSIGNMENT = &TreeType{"left_shift_assignment", 79}
var AST_TREE_KIND_RIGHT_SHIFT_ASSIGNMENT = &TreeType{"right_shift_assignment", 80}
var AST_TREE_KIND_UNSIGNED_RIGHT_SHIFT_ASSIGNMENT = &TreeType{"unsigned_right_shift_assignment", 81}
var AST_TREE_KIND_AND_ASSIGNMENT = &TreeType{"and_assignment", 82}
var AST_TREE_KIND_XOR_ASSIGNMENT = &TreeType{"xor_assignment", 83}
var AST_TREE_KIND_OR_ASSIGNMENT = &TreeType{"or_assignment", 84}
var AST_TREE_KIND_INT_LITERAL = &TreeType{"int_literal", 85}
var AST_TREE_KIND_LONG_LITERAL = &TreeType{"long_literal", 86}
var AST_TREE_KIND_FLOAT_LITERAL = &TreeType{"float_literal", 87}
var AST_TREE_KIND_DOUBLE_LITERAL = &TreeType{"double_literal", 88}
var AST_TREE_KIND_BOOLEAN_LITERAL = &TreeType{"boolean_literal", 89}
var AST_TREE_KIND_CHAR_LITERAL = &TreeType{"char_literal", 90}
var AST_TREE_KIND_STRING_LITERAL = &TreeType{"string_literal", 91}
var AST_TREE_KIND_NULL_LITERAL = &TreeType{"null_literal", 92}
var AST_TREE_KIND_UNBOUNDED_WILDCARD = &TreeType{"unbounded_wildcard", 93}
var AST_TREE_KIND_EXTENDS_WILDCARD = &TreeType{"extends_wildcard", 94}
var AST_TREE_KIND_SUPER_WILDCARD = &TreeType{"super_wildcard", 95}
var AST_TREE_KIND_ERRONEOUS = &TreeType{"erroneous", 96}
var AST_TREE_KIND_INTERFACE = &TreeType{"interface", 97}
var AST_TREE_KIND_ENUM = &TreeType{"enum", 98}
var AST_TREE_KIND_ANNOTATION_TYPE = &TreeType{"annotation_type", 99}
var AST_TREE_KIND_OTHER = &TreeType{"other", 100}

//树的节点
type AstTreeNode struct {
	nodeType TreeType
	msg      string //描述信息
}
