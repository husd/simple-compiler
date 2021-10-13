package parser

import "fmt"

type TreeNodeType int

const node_type_expression TreeNodeType = 0
const node_type_statement TreeNodeType = 1
const node_type_unknown TreeNodeType = 1

/**
 * 重新设计的ast的节点数据 抽象语法树统一的设计为二叉树，这样是否能好一点？
 * @author hushengdong
 */
type TreeNode struct {
	index         int          // 符号表里的索引
	tag           TreeNodeTag  // ast有有限的几种类型
	childrenCount int          // 子树的数量
	children      []*TreeNode  // 表示子树的集合 这里用一个切片来表示 长度一般都不超过6个 再考虑这个属性要不要有
	name          string       // 树的一些基本信息 ，属于备注信息，后续可以去掉
	expr_or_state TreeNodeType // 0:expression 1:statement -1:未知
	// 补充一点位置信息，在源代码中的位置
}

func (this *TreeNode) Append(child *TreeNode) {

	if this == child {
		fmt.Println("------------树节点增加错误，自己不可以添加自己为子节点------------")
		return
	}
	this.ensureCapacity(this.childrenCount, 1)
	this.children[this.childrenCount] = child
	this.childrenCount++
}

func (this *TreeNode) GetFirstChildren() *TreeNode {

	if this.childrenCount <= 0 {
		return nil
	}
	return this.children[0]
}

/**
 * @param spos 当前容量
 * @param need 需要的容量
 */
func (this *TreeNode) ensureCapacity(spos int, need int) {

	currentCap := cap(this.children)
	if spos+need > currentCap {
		newCap := calcNewLength(currentCap, spos+need)
		newArray := make([]*TreeNode, newCap, newCap) // len设置为cap，这样才可以在任意位置写入
		copy(newArray, this.children)
		this.children = newArray
	}
}

type ITreeType interface {

	/**
	 * 树的类型
	 */
	GetTreeType() TreeType

	TreeType_()
}

var TT_array [102]string = [102]string{}

func init() {

	TT_array[TT_NIL] = "nil"
	TT_array[TT_ANNOTATED_TYPE] = "annotated_type"
	TT_array[TT_ANNOTATION] = "annotation"
	TT_array[TT_TYPE_ANNOTATION] = "type_annotation"
	TT_array[TT_ARRAY_ACCESS] = "array_access"
	TT_array[TT_ARRAY_TYPE] = "array_type"
	TT_array[TT_ASSERT] = "assert"
	TT_array[TT_ASSIGNMENT] = "assignment"
	TT_array[TT_BLOCK] = "block"
	TT_array[TT_BREAK] = "break"
	TT_array[TT_CASE] = "case"
	TT_array[TT_CATCH] = "catch"
	TT_array[TT_CLASS] = "class"
	TT_array[TT_COMPILATION_UNIT] = "compilation_unit"
	TT_array[TT_CONDITIONAL_EXPRESSION] = "conditional_expression"
	TT_array[TT_CONTINUE] = "continue"
	TT_array[TT_DO_WHILE_LOOP] = "do_while_loop"
	TT_array[TT_ENHANCED_FOR_LOOP] = "enhanced_for_loop"
	TT_array[TT_EXPRESSION_STATEMENT] = "expression_statement"
	TT_array[TT_MEMBER_SELECT] = "member_select"
	TT_array[TT_MEMBER_REFERENCE] = "member_reference"
	TT_array[TT_FOR_LOOP] = "for_loop"
	TT_array[TT_IDENTIFIER] = "identifier"
	TT_array[TT_IF] = "if"
	TT_array[TT_IMPORT] = "import"
	TT_array[TT_INSTANCE_OF] = "instance_of"
	TT_array[TT_LABELED_STATEMENT] = "labeled_statement"
	TT_array[TT_METHOD] = "method"
	TT_array[TT_METHOD_INVOCATION] = "method_invocation"
	TT_array[TT_MODIFIERS] = "modifiers"
	TT_array[TT_NEW_ARRAY] = "new_array"
	TT_array[TT_NEW_CLASS] = "new_class"
	TT_array[TT_LAMBDA_EXPRESSION] = "lambda_expression"
	TT_array[TT_PARENTHESIZED] = "parenthesized"
	TT_array[TT_PRIMITIVE_TYPE] = "primitive_type"
	TT_array[TT_RETURN] = "return"
	TT_array[TT_EMPTY_STATEMENT] = "empty_statement"
	TT_array[TT_SWITCH] = "switch"
	TT_array[TT_SYNCHRONIZED] = "synchronized"
	TT_array[TT_THROW] = "throw"
	TT_array[TT_TRY] = "try"
	TT_array[TT_PARAMETERIZED_TYPE] = "parameterized_type"
	TT_array[TT_UNION_TYPE] = "union_type"
	TT_array[TT_INTERSECTION_TYPE] = "intersection_type"
	TT_array[TT_TYPE_CAST] = "type_cast"
	TT_array[TT_TYPE_PARAMETER] = "type_parameter"
	TT_array[TT_VARIABLE] = "variable"
	TT_array[TT_WHILE_LOOP] = "while_loop"
	TT_array[TT_POSTFIX_INCREMENT] = "postfix_increment"
	TT_array[TT_POSTFIX_DECREMENT] = "postfix_decrement"
	TT_array[TT_PREFIX_INCREMENT] = "prefix_increment"
	TT_array[TT_PREFIX_DECREMENT] = "prefix_decrement"
	TT_array[TT_UNARY_PLUS] = "unary_plus"
	TT_array[TT_UNARY_MINUS] = "unary_minus"
	TT_array[TT_BITWISE_COMPLEMENT] = "bitwise_complement"
	TT_array[TT_LOGICAL_COMPLEMENT] = "logical_complement"
	TT_array[TT_MULTIPLY] = "multiply"
	TT_array[TT_DIVIDE] = "divide"
	TT_array[TT_REMAINDER] = "remainder"
	TT_array[TT_PLUS] = "plus"
	TT_array[TT_MINUS] = "minus"
	TT_array[TT_LEFT_SHIFT] = "left_shift"
	TT_array[TT_RIGHT_SHIFT] = "right_shift"
	TT_array[TT_UNSIGNED_RIGHT_SHIFT] = "unsigned_right_shift"
	TT_array[TT_LESS_THAN] = "less_than"
	TT_array[TT_GREATER_THAN] = "greater_than"
	TT_array[TT_LESS_THAN_EQUAL] = "less_than_equal"
	TT_array[TT_GREATER_THAN_EQUAL] = "greater_than_equal"
	TT_array[TT_EQUAL_TO] = "equal_to"
	TT_array[TT_NOT_EQUAL_TO] = "not_equal_to"
	TT_array[TT_AND] = "and"
	TT_array[TT_XOR] = "xor"
	TT_array[TT_OR] = "or"
	TT_array[TT_CONDITIONAL_AND] = "conditional_and"
	TT_array[TT_CONDITIONAL_OR] = "conditional_or"
	TT_array[TT_MULTIPLY_ASSIGNMENT] = "multiply_assignment"
	TT_array[TT_DIVIDE_ASSIGNMENT] = "divide_assignment"
	TT_array[TT_REMAINDER_ASSIGNMENT] = "remainder_assignment"
	TT_array[TT_PLUS_ASSIGNMENT] = "plus_assignment"
	TT_array[TT_MINUS_ASSIGNMENT] = "minus_assignment"
	TT_array[TT_LEFT_SHIFT_ASSIGNMENT] = "left_shift_assignment"
	TT_array[TT_RIGHT_SHIFT_ASSIGNMENT] = "right_shift_assignment"
	TT_array[TT_UNSIGNED_RIGHT_SHIFT_ASSIGNMENT] = "unsigned_right_shift_assignment"
	TT_array[TT_AND_ASSIGNMENT] = "and_assignment"
	TT_array[TT_XOR_ASSIGNMENT] = "xor_assignment"
	TT_array[TT_OR_ASSIGNMENT] = "or_assignment"
	TT_array[TT_INT_LITERAL] = "int_literal"
	TT_array[TT_LONG_LITERAL] = "long_literal"
	TT_array[TT_FLOAT_LITERAL] = "float_literal"
	TT_array[TT_DOUBLE_LITERAL] = "double_literal"
	TT_array[TT_BOOLEAN_LITERAL] = "boolean_literal"
	TT_array[TT_CHAR_LITERAL] = "char_literal"
	TT_array[TT_STRING_LITERAL] = "string_literal"
	TT_array[TT_NULL_LITERAL] = "null_literal"
	TT_array[TT_UNBOUNDED_WILDCARD] = "unbounded_wildcard"
	TT_array[TT_EXTENDS_WILDCARD] = "extends_wildcard"
	TT_array[TT_SUPER_WILDCARD] = "super_wildcard"
	TT_array[TT_ERRONEOUS] = "erroneous"
	TT_array[TT_INTERFACE] = "interface"
	TT_array[TT_ENUM] = "enum"
	TT_array[TT_ANNOTATION_TYPE] = "annotation_type"
	TT_array[TT_OTHER] = "other"
}

/**
 * 抽象语法树的节点的类型，暂时先写个枚举
 */
type TreeType int

const (
	TT_NIL                             TreeType = 0   // nil
	TT_ANNOTATED_TYPE                  TreeType = 1   // annotated_type
	TT_ANNOTATION                      TreeType = 2   // annotation
	TT_TYPE_ANNOTATION                 TreeType = 3   // type_annotation
	TT_ARRAY_ACCESS                    TreeType = 4   // array_access
	TT_ARRAY_TYPE                      TreeType = 5   // array_type
	TT_ASSERT                          TreeType = 6   // assert
	TT_ASSIGNMENT                      TreeType = 7   // assignment
	TT_BLOCK                           TreeType = 8   // block
	TT_BREAK                           TreeType = 9   // break
	TT_CASE                            TreeType = 10  // case
	TT_CATCH                           TreeType = 11  // catch
	TT_CLASS                           TreeType = 12  // class
	TT_COMPILATION_UNIT                TreeType = 13  // compilation_unit
	TT_CONDITIONAL_EXPRESSION          TreeType = 14  // conditional_expression
	TT_CONTINUE                        TreeType = 15  // continue
	TT_DO_WHILE_LOOP                   TreeType = 16  // do_while_loop
	TT_ENHANCED_FOR_LOOP               TreeType = 17  // enhanced_for_loop
	TT_EXPRESSION_STATEMENT            TreeType = 18  // expression_statement
	TT_MEMBER_SELECT                   TreeType = 19  // member_select
	TT_MEMBER_REFERENCE                TreeType = 20  // member_reference
	TT_FOR_LOOP                        TreeType = 21  // for_loop
	TT_IDENTIFIER                      TreeType = 22  // identifier
	TT_IF                              TreeType = 23  // if
	TT_IMPORT                          TreeType = 24  // import
	TT_INSTANCE_OF                     TreeType = 25  // instance_of
	TT_LABELED_STATEMENT               TreeType = 26  // labeled_statement
	TT_METHOD                          TreeType = 27  // method
	TT_METHOD_INVOCATION               TreeType = 28  // method_invocation
	TT_MODIFIERS                       TreeType = 29  // modifiers
	TT_NEW_ARRAY                       TreeType = 30  // new_array
	TT_NEW_CLASS                       TreeType = 31  // new_class
	TT_LAMBDA_EXPRESSION               TreeType = 32  // lambda_expression
	TT_PARENTHESIZED                   TreeType = 33  // parenthesized
	TT_PRIMITIVE_TYPE                  TreeType = 34  // primitive_type
	TT_RETURN                          TreeType = 35  // return
	TT_EMPTY_STATEMENT                 TreeType = 36  // empty_statement
	TT_SWITCH                          TreeType = 37  // switch
	TT_SYNCHRONIZED                    TreeType = 38  // synchronized
	TT_THROW                           TreeType = 39  // throw
	TT_TRY                             TreeType = 40  // try
	TT_PARAMETERIZED_TYPE              TreeType = 41  // parameterized_type
	TT_UNION_TYPE                      TreeType = 42  // union_type
	TT_INTERSECTION_TYPE               TreeType = 43  // intersection_type
	TT_TYPE_CAST                       TreeType = 44  // type_cast
	TT_TYPE_PARAMETER                  TreeType = 45  // type_parameter
	TT_VARIABLE                        TreeType = 46  // variable
	TT_WHILE_LOOP                      TreeType = 47  // while_loop
	TT_POSTFIX_INCREMENT               TreeType = 48  // postfix_increment
	TT_POSTFIX_DECREMENT               TreeType = 49  // postfix_decrement
	TT_PREFIX_INCREMENT                TreeType = 50  // prefix_increment
	TT_PREFIX_DECREMENT                TreeType = 51  // prefix_decrement
	TT_UNARY_PLUS                      TreeType = 52  // unary_plus
	TT_UNARY_MINUS                     TreeType = 53  // unary_minus
	TT_BITWISE_COMPLEMENT              TreeType = 54  // bitwise_complement
	TT_LOGICAL_COMPLEMENT              TreeType = 55  // logical_complement
	TT_MULTIPLY                        TreeType = 56  // multiply
	TT_DIVIDE                          TreeType = 57  // divide
	TT_REMAINDER                       TreeType = 58  // remainder
	TT_PLUS                            TreeType = 59  // plus
	TT_MINUS                           TreeType = 60  // minus
	TT_LEFT_SHIFT                      TreeType = 61  // left_shift
	TT_RIGHT_SHIFT                     TreeType = 62  // right_shift
	TT_UNSIGNED_RIGHT_SHIFT            TreeType = 63  // unsigned_right_shift
	TT_LESS_THAN                       TreeType = 64  // less_than
	TT_GREATER_THAN                    TreeType = 65  // greater_than
	TT_LESS_THAN_EQUAL                 TreeType = 66  // less_than_equal
	TT_GREATER_THAN_EQUAL              TreeType = 67  // greater_than_equal
	TT_EQUAL_TO                        TreeType = 68  // equal_to
	TT_NOT_EQUAL_TO                    TreeType = 69  // not_equal_to
	TT_AND                             TreeType = 70  // and
	TT_XOR                             TreeType = 71  // xor
	TT_OR                              TreeType = 72  // or
	TT_CONDITIONAL_AND                 TreeType = 73  // conditional_and
	TT_CONDITIONAL_OR                  TreeType = 74  // conditional_or
	TT_MULTIPLY_ASSIGNMENT             TreeType = 75  // multiply_assignment
	TT_DIVIDE_ASSIGNMENT               TreeType = 76  // divide_assignment
	TT_REMAINDER_ASSIGNMENT            TreeType = 77  // remainder_assignment
	TT_PLUS_ASSIGNMENT                 TreeType = 78  // plus_assignment
	TT_MINUS_ASSIGNMENT                TreeType = 79  // minus_assignment
	TT_LEFT_SHIFT_ASSIGNMENT           TreeType = 80  // left_shift_assignment
	TT_RIGHT_SHIFT_ASSIGNMENT          TreeType = 81  // right_shift_assignment
	TT_UNSIGNED_RIGHT_SHIFT_ASSIGNMENT TreeType = 82  // unsigned_right_shift_assignment
	TT_AND_ASSIGNMENT                  TreeType = 83  // and_assignment
	TT_XOR_ASSIGNMENT                  TreeType = 84  // xor_assignment
	TT_OR_ASSIGNMENT                   TreeType = 85  // or_assignment
	TT_INT_LITERAL                     TreeType = 86  // int_literal
	TT_LONG_LITERAL                    TreeType = 87  // long_literal
	TT_FLOAT_LITERAL                   TreeType = 88  // float_literal
	TT_DOUBLE_LITERAL                  TreeType = 89  // double_literal
	TT_BOOLEAN_LITERAL                 TreeType = 90  // boolean_literal
	TT_CHAR_LITERAL                    TreeType = 91  // char_literal
	TT_STRING_LITERAL                  TreeType = 92  // string_literal
	TT_NULL_LITERAL                    TreeType = 93  // null_literal
	TT_UNBOUNDED_WILDCARD              TreeType = 94  // unbounded_wildcard
	TT_EXTENDS_WILDCARD                TreeType = 95  // extends_wildcard
	TT_SUPER_WILDCARD                  TreeType = 96  // super_wildcard
	TT_ERRONEOUS                       TreeType = 97  // erroneous
	TT_INTERFACE                       TreeType = 98  // interface
	TT_ENUM                            TreeType = 99  // enum
	TT_ANNOTATION_TYPE                 TreeType = 100 // annotation_type
	TT_OTHER                           TreeType = 101 // other
)

func GetTreeTypeName(tt TreeType) string {

	if tt >= 0 && tt <= 101 {
		return TT_array[tt]
	}
	return "unknown tree type:" + fmt.Sprint(tt)
}
