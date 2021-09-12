package ast_tree

/**
 *
 * @author hushengdong
 */
type TreeType string

const (
	JC_CompilationUnit = "JCCompilationUnit"
	JC_Statement       = "JCStatement"
	JC_Expression      = "JCExpression"
	JC_MethodDecl      = "JCMethodDecl"

	JC_Import        = "JCImport"
	JC_Catch         = "JCCatch"
	JC_TypeParameter = "JCTypeParameter"
	JC_Modifiers     = "JCModifiers"
)

/**
 * 抽象语法树的节点的类型，暂时先写个枚举
 */
type AstTreeNodeKind struct {
	Name  string // 名字
	Index int    //索引值
}

var TREE_NIL = &AstTreeNodeKind{"nil", -1}

var TREE_ANNOTATED_TYPE = &AstTreeNodeKind{"annotated_type", 0}
var TREE_ANNOTATION = &AstTreeNodeKind{"annotation", 1}
var TREE_TYPE_ANNOTATION = &AstTreeNodeKind{"type_annotation", 2}
var TREE_ARRAY_ACCESS = &AstTreeNodeKind{"array_access", 3}
var TREE_ARRAY_TYPE = &AstTreeNodeKind{"array_type", 4}
var TREE_ASSERT = &AstTreeNodeKind{"assert", 5}
var TREE_ASSIGNMENT = &AstTreeNodeKind{"assignment", 6}
var TREE_BLOCK = &AstTreeNodeKind{"block", 7}
var TREE_BREAK = &AstTreeNodeKind{"break", 8}
var TREE_CASE = &AstTreeNodeKind{"case", 9}
var TREE_CATCH = &AstTreeNodeKind{"catch", 10}
var TREE_CLASS = &AstTreeNodeKind{"class", 11}
var TREE_COMPILATION_UNIT = &AstTreeNodeKind{"compilation_unit", 12}
var TREE_CONDITIONAL_EXPRESSION = &AstTreeNodeKind{"conditional_expression", 13}
var TREE_CONTINUE = &AstTreeNodeKind{"continue", 14}
var TREE_DO_WHILE_LOOP = &AstTreeNodeKind{"do_while_loop", 15}
var TREE_ENHANCED_FOR_LOOP = &AstTreeNodeKind{"enhanced_for_loop", 16}
var TREE_EXPRESSION_STATEMENT = &AstTreeNodeKind{"expression_statement", 17}
var TREE_MEMBER_SELECT = &AstTreeNodeKind{"member_select", 18}
var TREE_MEMBER_REFERENCE = &AstTreeNodeKind{"member_reference", 19}
var TREE_FOR_LOOP = &AstTreeNodeKind{"for_loop", 20}
var TREE_IDENTIFIER = &AstTreeNodeKind{"identifier", 21}
var TREE_IF = &AstTreeNodeKind{"if", 22}
var TREE_IMPORT = &AstTreeNodeKind{"import", 23}
var TREE_INSTANCE_OF = &AstTreeNodeKind{"instance_of", 24}
var TREE_LABELED_STATEMENT = &AstTreeNodeKind{"labeled_statement", 25}
var TREE_METHOD = &AstTreeNodeKind{"method", 26}
var TREE_METHOD_INVOCATION = &AstTreeNodeKind{"method_invocation", 27}
var TREE_MODIFIERS = &AstTreeNodeKind{"modifiers", 28}
var TREE_NEW_ARRAY = &AstTreeNodeKind{"new_array", 29}
var TREE_NEW_CLASS = &AstTreeNodeKind{"new_class", 30}
var TREE_LAMBDA_EXPRESSION = &AstTreeNodeKind{"lambda_expression", 31}
var TREE_PARENTHESIZED = &AstTreeNodeKind{"parenthesized", 32}
var TREE_PRIMITIVE_TYPE = &AstTreeNodeKind{"primitive_type", 33}
var TREE_RETURN = &AstTreeNodeKind{"return", 34}
var TREE_EMPTY_STATEMENT = &AstTreeNodeKind{"empty_statement", 35}
var TREE_SWITCH = &AstTreeNodeKind{"switch", 36}
var TREE_SYNCHRONIZED = &AstTreeNodeKind{"synchronized", 37}
var TREE_THROW = &AstTreeNodeKind{"throw", 38}
var TREE_TRY = &AstTreeNodeKind{"try", 39}
var TREE_PARAMETERIZED_TYPE = &AstTreeNodeKind{"parameterized_type", 40}
var TREE_UNION_TYPE = &AstTreeNodeKind{"union_type", 41}
var TREE_INTERSECTION_TYPE = &AstTreeNodeKind{"intersection_type", 42}
var TREE_TYPE_CAST = &AstTreeNodeKind{"type_cast", 43}
var TREE_TYPE_PARAMETER = &AstTreeNodeKind{"type_parameter", 44}
var TREE_VARIABLE = &AstTreeNodeKind{"variable", 45}
var TREE_WHILE_LOOP = &AstTreeNodeKind{"while_loop", 46}
var TREE_POSTFIX_INCREMENT = &AstTreeNodeKind{"postfix_increment", 47}
var TREE_POSTFIX_DECREMENT = &AstTreeNodeKind{"postfix_decrement", 48}
var TREE_PREFIX_INCREMENT = &AstTreeNodeKind{"prefix_increment", 49}
var TREE_PREFIX_DECREMENT = &AstTreeNodeKind{"prefix_decrement", 50}
var TREE_UNARY_PLUS = &AstTreeNodeKind{"unary_plus", 51}
var TREE_UNARY_MINUS = &AstTreeNodeKind{"unary_minus", 52}
var TREE_BITWISE_COMPLEMENT = &AstTreeNodeKind{"bitwise_complement", 53}
var TREE_LOGICAL_COMPLEMENT = &AstTreeNodeKind{"logical_complement", 54}
var TREE_MULTIPLY = &AstTreeNodeKind{"multiply", 55}
var TREE_DIVIDE = &AstTreeNodeKind{"divide", 56}
var TREE_REMAINDER = &AstTreeNodeKind{"remainder", 57}
var TREE_PLUS = &AstTreeNodeKind{"plus", 58}
var TREE_MINUS = &AstTreeNodeKind{"minus", 59}
var TREE_LEFT_SHIFT = &AstTreeNodeKind{"left_shift", 60}
var TREE_RIGHT_SHIFT = &AstTreeNodeKind{"right_shift", 61}
var TREE_UNSIGNED_RIGHT_SHIFT = &AstTreeNodeKind{"unsigned_right_shift", 62}
var TREE_LESS_THAN = &AstTreeNodeKind{"less_than", 63}
var TREE_GREATER_THAN = &AstTreeNodeKind{"greater_than", 64}
var TREE_LESS_THAN_EQUAL = &AstTreeNodeKind{"less_than_equal", 65}
var TREE_GREATER_THAN_EQUAL = &AstTreeNodeKind{"greater_than_equal", 66}
var TREE_EQUAL_TO = &AstTreeNodeKind{"equal_to", 67}
var TREE_NOT_EQUAL_TO = &AstTreeNodeKind{"not_equal_to", 68}
var TREE_AND = &AstTreeNodeKind{"and", 69}
var TREE_XOR = &AstTreeNodeKind{"xor", 70}
var TREE_OR = &AstTreeNodeKind{"or", 71}
var TREE_CONDITIONAL_AND = &AstTreeNodeKind{"conditional_and", 72}
var TREE_CONDITIONAL_OR = &AstTreeNodeKind{"conditional_or", 73}
var TREE_MULTIPLY_ASSIGNMENT = &AstTreeNodeKind{"multiply_assignment", 74}
var TREE_DIVIDE_ASSIGNMENT = &AstTreeNodeKind{"divide_assignment", 75}
var TREE_REMAINDER_ASSIGNMENT = &AstTreeNodeKind{"remainder_assignment", 76}
var TREE_PLUS_ASSIGNMENT = &AstTreeNodeKind{"plus_assignment", 77}
var TREE_MINUS_ASSIGNMENT = &AstTreeNodeKind{"minus_assignment", 78}
var TREE_LEFT_SHIFT_ASSIGNMENT = &AstTreeNodeKind{"left_shift_assignment", 79}
var TREE_RIGHT_SHIFT_ASSIGNMENT = &AstTreeNodeKind{"right_shift_assignment", 80}
var TREE_UNSIGNED_RIGHT_SHIFT_ASSIGNMENT = &AstTreeNodeKind{"unsigned_right_shift_assignment", 81}
var TREE_AND_ASSIGNMENT = &AstTreeNodeKind{"and_assignment", 82}
var TREE_XOR_ASSIGNMENT = &AstTreeNodeKind{"xor_assignment", 83}
var TREE_OR_ASSIGNMENT = &AstTreeNodeKind{"or_assignment", 84}
var TREE_INT_LITERAL = &AstTreeNodeKind{"int_literal", 85}
var TREE_LONG_LITERAL = &AstTreeNodeKind{"long_literal", 86}
var TREE_FLOAT_LITERAL = &AstTreeNodeKind{"float_literal", 87}
var TREE_DOUBLE_LITERAL = &AstTreeNodeKind{"double_literal", 88}
var TREE_BOOLEAN_LITERAL = &AstTreeNodeKind{"boolean_literal", 89}
var TREE_CHAR_LITERAL = &AstTreeNodeKind{"char_literal", 90}
var TREE_STRING_LITERAL = &AstTreeNodeKind{"string_literal", 91}
var TREE_NULL_LITERAL = &AstTreeNodeKind{"null_literal", 92}
var TREE_UNBOUNDED_WILDCARD = &AstTreeNodeKind{"unbounded_wildcard", 93}
var TREE_EXTENDS_WILDCARD = &AstTreeNodeKind{"extends_wildcard", 94}
var TREE_SUPER_WILDCARD = &AstTreeNodeKind{"super_wildcard", 95}
var TREE_ERRONEOUS = &AstTreeNodeKind{"erroneous", 96}
var TREE_INTERFACE = &AstTreeNodeKind{"interface", 97}
var TREE_ENUM = &AstTreeNodeKind{"enum", 98}
var TREE_ANNOTATION_TYPE = &AstTreeNodeKind{"annotation_type", 99}
var TREE_OTHER = &AstTreeNodeKind{"other", 100}
