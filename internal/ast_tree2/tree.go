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
	GetTreeType() TreeType
}

var Tree_type_array [102]string = [102]string{}

func init() {

	Tree_type_array[TREE_TYPE_NIL] = "nil"
	Tree_type_array[TREE_TYPE_ANNOTATED_TYPE] = "annotated_type"
	Tree_type_array[TREE_TYPE_ANNOTATION] = "annotation"
	Tree_type_array[TREE_TYPE_TYPE_ANNOTATION] = "type_annotation"
	Tree_type_array[TREE_TYPE_ARRAY_ACCESS] = "array_access"
	Tree_type_array[TREE_TYPE_ARRAY_TYPE] = "array_type"
	Tree_type_array[TREE_TYPE_ASSERT] = "assert"
	Tree_type_array[TREE_TYPE_ASSIGNMENT] = "assignment"
	Tree_type_array[TREE_TYPE_BLOCK] = "block"
	Tree_type_array[TREE_TYPE_BREAK] = "break"
	Tree_type_array[TREE_TYPE_CASE] = "case"
	Tree_type_array[TREE_TYPE_CATCH] = "catch"
	Tree_type_array[TREE_TYPE_CLASS] = "class"
	Tree_type_array[TREE_TYPE_COMPILATION_UNIT] = "compilation_unit"
	Tree_type_array[TREE_TYPE_CONDITIONAL_EXPRESSION] = "conditional_expression"
	Tree_type_array[TREE_TYPE_CONTINUE] = "continue"
	Tree_type_array[TREE_TYPE_DO_WHILE_LOOP] = "do_while_loop"
	Tree_type_array[TREE_TYPE_ENHANCED_FOR_LOOP] = "enhanced_for_loop"
	Tree_type_array[TREE_TYPE_EXPRESSION_STATEMENT] = "expression_statement"
	Tree_type_array[TREE_TYPE_MEMBER_SELECT] = "member_select"
	Tree_type_array[TREE_TYPE_MEMBER_REFERENCE] = "member_reference"
	Tree_type_array[TREE_TYPE_FOR_LOOP] = "for_loop"
	Tree_type_array[TREE_TYPE_IDENTIFIER] = "identifier"
	Tree_type_array[TREE_TYPE_IF] = "if"
	Tree_type_array[TREE_TYPE_IMPORT] = "import"
	Tree_type_array[TREE_TYPE_INSTANCE_OF] = "instance_of"
	Tree_type_array[TREE_TYPE_LABELED_STATEMENT] = "labeled_statement"
	Tree_type_array[TREE_TYPE_METHOD] = "method"
	Tree_type_array[TREE_TYPE_METHOD_INVOCATION] = "method_invocation"
	Tree_type_array[TREE_TYPE_MODIFIERS] = "modifiers"
	Tree_type_array[TREE_TYPE_NEW_ARRAY] = "new_array"
	Tree_type_array[TREE_TYPE_NEW_CLASS] = "new_class"
	Tree_type_array[TREE_TYPE_LAMBDA_EXPRESSION] = "lambda_expression"
	Tree_type_array[TREE_TYPE_PARENTHESIZED] = "parenthesized"
	Tree_type_array[TREE_TYPE_PRIMITIVE_TYPE] = "primitive_type"
	Tree_type_array[TREE_TYPE_RETURN] = "return"
	Tree_type_array[TREE_TYPE_EMPTY_STATEMENT] = "empty_statement"
	Tree_type_array[TREE_TYPE_SWITCH] = "switch"
	Tree_type_array[TREE_TYPE_SYNCHRONIZED] = "synchronized"
	Tree_type_array[TREE_TYPE_THROW] = "throw"
	Tree_type_array[TREE_TYPE_TRY] = "try"
	Tree_type_array[TREE_TYPE_PARAMETERIZED_TYPE] = "parameterized_type"
	Tree_type_array[TREE_TYPE_UNION_TYPE] = "union_type"
	Tree_type_array[TREE_TYPE_INTERSECTION_TYPE] = "intersection_type"
	Tree_type_array[TREE_TYPE_TYPE_CAST] = "type_cast"
	Tree_type_array[TREE_TYPE_TYPE_PARAMETER] = "type_parameter"
	Tree_type_array[TREE_TYPE_VARIABLE] = "variable"
	Tree_type_array[TREE_TYPE_WHILE_LOOP] = "while_loop"
	Tree_type_array[TREE_TYPE_POSTFIX_INCREMENT] = "postfix_increment"
	Tree_type_array[TREE_TYPE_POSTFIX_DECREMENT] = "postfix_decrement"
	Tree_type_array[TREE_TYPE_PREFIX_INCREMENT] = "prefix_increment"
	Tree_type_array[TREE_TYPE_PREFIX_DECREMENT] = "prefix_decrement"
	Tree_type_array[TREE_TYPE_UNARY_PLUS] = "unary_plus"
	Tree_type_array[TREE_TYPE_UNARY_MINUS] = "unary_minus"
	Tree_type_array[TREE_TYPE_BITWISE_COMPLEMENT] = "bitwise_complement"
	Tree_type_array[TREE_TYPE_LOGICAL_COMPLEMENT] = "logical_complement"
	Tree_type_array[TREE_TYPE_MULTIPLY] = "multiply"
	Tree_type_array[TREE_TYPE_DIVIDE] = "divide"
	Tree_type_array[TREE_TYPE_REMAINDER] = "remainder"
	Tree_type_array[TREE_TYPE_PLUS] = "plus"
	Tree_type_array[TREE_TYPE_MINUS] = "minus"
	Tree_type_array[TREE_TYPE_LEFT_SHIFT] = "left_shift"
	Tree_type_array[TREE_TYPE_RIGHT_SHIFT] = "right_shift"
	Tree_type_array[TREE_TYPE_UNSIGNED_RIGHT_SHIFT] = "unsigned_right_shift"
	Tree_type_array[TREE_TYPE_LESS_THAN] = "less_than"
	Tree_type_array[TREE_TYPE_GREATER_THAN] = "greater_than"
	Tree_type_array[TREE_TYPE_LESS_THAN_EQUAL] = "less_than_equal"
	Tree_type_array[TREE_TYPE_GREATER_THAN_EQUAL] = "greater_than_equal"
	Tree_type_array[TREE_TYPE_EQUAL_TO] = "equal_to"
	Tree_type_array[TREE_TYPE_NOT_EQUAL_TO] = "not_equal_to"
	Tree_type_array[TREE_TYPE_AND] = "and"
	Tree_type_array[TREE_TYPE_XOR] = "xor"
	Tree_type_array[TREE_TYPE_OR] = "or"
	Tree_type_array[TREE_TYPE_CONDITIONAL_AND] = "conditional_and"
	Tree_type_array[TREE_TYPE_CONDITIONAL_OR] = "conditional_or"
	Tree_type_array[TREE_TYPE_MULTIPLY_ASSIGNMENT] = "multiply_assignment"
	Tree_type_array[TREE_TYPE_DIVIDE_ASSIGNMENT] = "divide_assignment"
	Tree_type_array[TREE_TYPE_REMAINDER_ASSIGNMENT] = "remainder_assignment"
	Tree_type_array[TREE_TYPE_PLUS_ASSIGNMENT] = "plus_assignment"
	Tree_type_array[TREE_TYPE_MINUS_ASSIGNMENT] = "minus_assignment"
	Tree_type_array[TREE_TYPE_LEFT_SHIFT_ASSIGNMENT] = "left_shift_assignment"
	Tree_type_array[TREE_TYPE_RIGHT_SHIFT_ASSIGNMENT] = "right_shift_assignment"
	Tree_type_array[TREE_TYPE_UNSIGNED_RIGHT_SHIFT_ASSIGNMENT] = "unsigned_right_shift_assignment"
	Tree_type_array[TREE_TYPE_AND_ASSIGNMENT] = "and_assignment"
	Tree_type_array[TREE_TYPE_XOR_ASSIGNMENT] = "xor_assignment"
	Tree_type_array[TREE_TYPE_OR_ASSIGNMENT] = "or_assignment"
	Tree_type_array[TREE_TYPE_INT_LITERAL] = "int_literal"
	Tree_type_array[TREE_TYPE_LONG_LITERAL] = "long_literal"
	Tree_type_array[TREE_TYPE_FLOAT_LITERAL] = "float_literal"
	Tree_type_array[TREE_TYPE_DOUBLE_LITERAL] = "double_literal"
	Tree_type_array[TREE_TYPE_BOOLEAN_LITERAL] = "boolean_literal"
	Tree_type_array[TREE_TYPE_CHAR_LITERAL] = "char_literal"
	Tree_type_array[TREE_TYPE_STRING_LITERAL] = "string_literal"
	Tree_type_array[TREE_TYPE_NULL_LITERAL] = "null_literal"
	Tree_type_array[TREE_TYPE_UNBOUNDED_WILDCARD] = "unbounded_wildcard"
	Tree_type_array[TREE_TYPE_EXTENDS_WILDCARD] = "extends_wildcard"
	Tree_type_array[TREE_TYPE_SUPER_WILDCARD] = "super_wildcard"
	Tree_type_array[TREE_TYPE_ERRONEOUS] = "erroneous"
	Tree_type_array[TREE_TYPE_INTERFACE] = "interface"
	Tree_type_array[TREE_TYPE_ENUM] = "enum"
	Tree_type_array[TREE_TYPE_ANNOTATION_TYPE] = "annotation_type"
	Tree_type_array[TREE_TYPE_OTHER] = "other"
}

/**
 * 抽象语法树的节点的类型，暂时先写个枚举
 */
type TreeType int

const (
	TREE_TYPE_NIL                             TreeType = 0   //nil
	TREE_TYPE_ANNOTATED_TYPE                  TreeType = 1   //annotated_type
	TREE_TYPE_ANNOTATION                      TreeType = 2   //annotation
	TREE_TYPE_TYPE_ANNOTATION                 TreeType = 3   //type_annotation
	TREE_TYPE_ARRAY_ACCESS                    TreeType = 4   //array_access
	TREE_TYPE_ARRAY_TYPE                      TreeType = 5   //array_type
	TREE_TYPE_ASSERT                          TreeType = 6   //assert
	TREE_TYPE_ASSIGNMENT                      TreeType = 7   //assignment
	TREE_TYPE_BLOCK                           TreeType = 8   //block
	TREE_TYPE_BREAK                           TreeType = 9   //break
	TREE_TYPE_CASE                            TreeType = 10  //case
	TREE_TYPE_CATCH                           TreeType = 11  //catch
	TREE_TYPE_CLASS                           TreeType = 12  //class
	TREE_TYPE_COMPILATION_UNIT                TreeType = 13  //compilation_unit
	TREE_TYPE_CONDITIONAL_EXPRESSION          TreeType = 14  //conditional_expression
	TREE_TYPE_CONTINUE                        TreeType = 15  //continue
	TREE_TYPE_DO_WHILE_LOOP                   TreeType = 16  //do_while_loop
	TREE_TYPE_ENHANCED_FOR_LOOP               TreeType = 17  //enhanced_for_loop
	TREE_TYPE_EXPRESSION_STATEMENT            TreeType = 18  //expression_statement
	TREE_TYPE_MEMBER_SELECT                   TreeType = 19  //member_select
	TREE_TYPE_MEMBER_REFERENCE                TreeType = 20  //member_reference
	TREE_TYPE_FOR_LOOP                        TreeType = 21  //for_loop
	TREE_TYPE_IDENTIFIER                      TreeType = 22  //identifier
	TREE_TYPE_IF                              TreeType = 23  //if
	TREE_TYPE_IMPORT                          TreeType = 24  //import
	TREE_TYPE_INSTANCE_OF                     TreeType = 25  //instance_of
	TREE_TYPE_LABELED_STATEMENT               TreeType = 26  //labeled_statement
	TREE_TYPE_METHOD                          TreeType = 27  //method
	TREE_TYPE_METHOD_INVOCATION               TreeType = 28  //method_invocation
	TREE_TYPE_MODIFIERS                       TreeType = 29  //modifiers
	TREE_TYPE_NEW_ARRAY                       TreeType = 30  //new_array
	TREE_TYPE_NEW_CLASS                       TreeType = 31  //new_class
	TREE_TYPE_LAMBDA_EXPRESSION               TreeType = 32  //lambda_expression
	TREE_TYPE_PARENTHESIZED                   TreeType = 33  //parenthesized
	TREE_TYPE_PRIMITIVE_TYPE                  TreeType = 34  //primitive_type
	TREE_TYPE_RETURN                          TreeType = 35  //return
	TREE_TYPE_EMPTY_STATEMENT                 TreeType = 36  //empty_statement
	TREE_TYPE_SWITCH                          TreeType = 37  //switch
	TREE_TYPE_SYNCHRONIZED                    TreeType = 38  //synchronized
	TREE_TYPE_THROW                           TreeType = 39  //throw
	TREE_TYPE_TRY                             TreeType = 40  //try
	TREE_TYPE_PARAMETERIZED_TYPE              TreeType = 41  //parameterized_type
	TREE_TYPE_UNION_TYPE                      TreeType = 42  //union_type
	TREE_TYPE_INTERSECTION_TYPE               TreeType = 43  //intersection_type
	TREE_TYPE_TYPE_CAST                       TreeType = 44  //type_cast
	TREE_TYPE_TYPE_PARAMETER                  TreeType = 45  //type_parameter
	TREE_TYPE_VARIABLE                        TreeType = 46  //variable
	TREE_TYPE_WHILE_LOOP                      TreeType = 47  //while_loop
	TREE_TYPE_POSTFIX_INCREMENT               TreeType = 48  //postfix_increment
	TREE_TYPE_POSTFIX_DECREMENT               TreeType = 49  //postfix_decrement
	TREE_TYPE_PREFIX_INCREMENT                TreeType = 50  //prefix_increment
	TREE_TYPE_PREFIX_DECREMENT                TreeType = 51  //prefix_decrement
	TREE_TYPE_UNARY_PLUS                      TreeType = 52  //unary_plus
	TREE_TYPE_UNARY_MINUS                     TreeType = 53  //unary_minus
	TREE_TYPE_BITWISE_COMPLEMENT              TreeType = 54  //bitwise_complement
	TREE_TYPE_LOGICAL_COMPLEMENT              TreeType = 55  //logical_complement
	TREE_TYPE_MULTIPLY                        TreeType = 56  //multiply
	TREE_TYPE_DIVIDE                          TreeType = 57  //divide
	TREE_TYPE_REMAINDER                       TreeType = 58  //remainder
	TREE_TYPE_PLUS                            TreeType = 59  //plus
	TREE_TYPE_MINUS                           TreeType = 60  //minus
	TREE_TYPE_LEFT_SHIFT                      TreeType = 61  //left_shift
	TREE_TYPE_RIGHT_SHIFT                     TreeType = 62  //right_shift
	TREE_TYPE_UNSIGNED_RIGHT_SHIFT            TreeType = 63  //unsigned_right_shift
	TREE_TYPE_LESS_THAN                       TreeType = 64  //less_than
	TREE_TYPE_GREATER_THAN                    TreeType = 65  //greater_than
	TREE_TYPE_LESS_THAN_EQUAL                 TreeType = 66  //less_than_equal
	TREE_TYPE_GREATER_THAN_EQUAL              TreeType = 67  //greater_than_equal
	TREE_TYPE_EQUAL_TO                        TreeType = 68  //equal_to
	TREE_TYPE_NOT_EQUAL_TO                    TreeType = 69  //not_equal_to
	TREE_TYPE_AND                             TreeType = 70  //and
	TREE_TYPE_XOR                             TreeType = 71  //xor
	TREE_TYPE_OR                              TreeType = 72  //or
	TREE_TYPE_CONDITIONAL_AND                 TreeType = 73  //conditional_and
	TREE_TYPE_CONDITIONAL_OR                  TreeType = 74  //conditional_or
	TREE_TYPE_MULTIPLY_ASSIGNMENT             TreeType = 75  //multiply_assignment
	TREE_TYPE_DIVIDE_ASSIGNMENT               TreeType = 76  //divide_assignment
	TREE_TYPE_REMAINDER_ASSIGNMENT            TreeType = 77  //remainder_assignment
	TREE_TYPE_PLUS_ASSIGNMENT                 TreeType = 78  //plus_assignment
	TREE_TYPE_MINUS_ASSIGNMENT                TreeType = 79  //minus_assignment
	TREE_TYPE_LEFT_SHIFT_ASSIGNMENT           TreeType = 80  //left_shift_assignment
	TREE_TYPE_RIGHT_SHIFT_ASSIGNMENT          TreeType = 81  //right_shift_assignment
	TREE_TYPE_UNSIGNED_RIGHT_SHIFT_ASSIGNMENT TreeType = 82  //unsigned_right_shift_assignment
	TREE_TYPE_AND_ASSIGNMENT                  TreeType = 83  //and_assignment
	TREE_TYPE_XOR_ASSIGNMENT                  TreeType = 84  //xor_assignment
	TREE_TYPE_OR_ASSIGNMENT                   TreeType = 85  //or_assignment
	TREE_TYPE_INT_LITERAL                     TreeType = 86  //int_literal
	TREE_TYPE_LONG_LITERAL                    TreeType = 87  //long_literal
	TREE_TYPE_FLOAT_LITERAL                   TreeType = 88  //float_literal
	TREE_TYPE_DOUBLE_LITERAL                  TreeType = 89  //double_literal
	TREE_TYPE_BOOLEAN_LITERAL                 TreeType = 90  //boolean_literal
	TREE_TYPE_CHAR_LITERAL                    TreeType = 91  //char_literal
	TREE_TYPE_STRING_LITERAL                  TreeType = 92  //string_literal
	TREE_TYPE_NULL_LITERAL                    TreeType = 93  //null_literal
	TREE_TYPE_UNBOUNDED_WILDCARD              TreeType = 94  //unbounded_wildcard
	TREE_TYPE_EXTENDS_WILDCARD                TreeType = 95  //extends_wildcard
	TREE_TYPE_SUPER_WILDCARD                  TreeType = 96  //super_wildcard
	TREE_TYPE_ERRONEOUS                       TreeType = 97  //erroneous
	TREE_TYPE_INTERFACE                       TreeType = 98  //interface
	TREE_TYPE_ENUM                            TreeType = 99  //enum
	TREE_TYPE_ANNOTATION_TYPE                 TreeType = 100 //annotation_type
	TREE_TYPE_OTHER                           TreeType = 101 //other
)

func GetTreeTypeName(tt TreeType) string {

	if tt >= 0 && tt <= 101 {
		return Tree_type_array[tt]
	}
	return "unknown tree type:" + string(tt)
}
