package parser

import (
	"fmt"
	"husd.com/v0/code"
	"husd.com/v0/util"
)

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

	typeTag *code.TypeTag // 常量表达式才有的属性
	val     interface{}   // 常量表达式才有的属性 实际的值 需要配合typeTag的类型，确定是 int 还是 char false true等

	pos      int      // token里的pos
	treeType TreeType //树的类型

	symbol *Symbol // 符号
	n      *util.Name
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

var tt_array [102]string = [102]string{}

func init() {

	tt_array[tt_nil] = "nil"
	tt_array[tt_annotated_type] = "annotated_type"
	tt_array[tt_annotation] = "annotation"
	tt_array[tt_type_annotation] = "type_annotation"
	tt_array[tt_array_access] = "array_access"
	tt_array[tt_array_type] = "array_type"
	tt_array[tt_assert] = "assert"
	tt_array[tt_assignment] = "assignment"
	tt_array[tt_block] = "block"
	tt_array[tt_break] = "break"
	tt_array[tt_case] = "case"
	tt_array[tt_catch] = "catch"
	tt_array[tt_class] = "class"
	tt_array[tt_compilation_unit] = "compilation_unit"
	tt_array[tt_conditional_expression] = "conditional_expression"
	tt_array[tt_continue] = "continue"
	tt_array[tt_do_while_loop] = "do_while_loop"
	tt_array[tt_enhanced_for_loop] = "enhanced_for_loop"
	tt_array[tt_expression_statement] = "expression_statement"
	tt_array[tt_member_select] = "member_select"
	tt_array[tt_member_reference] = "member_reference"
	tt_array[tt_for_loop] = "for_loop"
	tt_array[tt_identifier] = "identifier"
	tt_array[tt_if] = "if"
	tt_array[tt_import] = "import"
	tt_array[tt_instance_of] = "instance_of"
	tt_array[tt_labeled_statement] = "labeled_statement"
	tt_array[tt_method] = "method"
	tt_array[tt_method_invocation] = "method_invocation"
	tt_array[tt_modifiers] = "modifiers"
	tt_array[tt_new_array] = "new_array"
	tt_array[tt_new_class] = "new_class"
	tt_array[tt_lambda_expression] = "lambda_expression"
	tt_array[tt_parenthesized] = "parenthesized"
	tt_array[tt_primitive_type] = "primitive_type"
	tt_array[tt_return] = "return"
	tt_array[tt_empty_statement] = "empty_statement"
	tt_array[tt_switch] = "switch"
	tt_array[tt_synchronized] = "synchronized"
	tt_array[tt_throw] = "throw"
	tt_array[tt_try] = "try"
	tt_array[tt_parameterized_type] = "parameterized_type"
	tt_array[tt_union_type] = "union_type"
	tt_array[tt_intersection_type] = "intersection_type"
	tt_array[tt_type_cast] = "type_cast"
	tt_array[tt_type_parameter] = "type_parameter"
	tt_array[tt_variable] = "variable"
	tt_array[tt_while_loop] = "while_loop"
	tt_array[tt_postfix_increment] = "postfix_increment"
	tt_array[tt_postfix_decrement] = "postfix_decrement"
	tt_array[tt_prefix_increment] = "prefix_increment"
	tt_array[tt_prefix_decrement] = "prefix_decrement"
	tt_array[tt_unary_plus] = "unary_plus"
	tt_array[tt_unary_minus] = "unary_minus"
	tt_array[tt_bitwise_complement] = "bitwise_complement"
	tt_array[tt_logical_complement] = "logical_complement"
	tt_array[tt_multiply] = "multiply"
	tt_array[tt_divide] = "divide"
	tt_array[tt_remainder] = "remainder"
	tt_array[tt_plus] = "plus"
	tt_array[tt_minus] = "minus"
	tt_array[tt_left_shift] = "left_shift"
	tt_array[tt_right_shift] = "right_shift"
	tt_array[tt_unsigned_right_shift] = "unsigned_right_shift"
	tt_array[tt_less_than] = "less_than"
	tt_array[tt_greater_than] = "greater_than"
	tt_array[tt_less_than_equal] = "less_than_equal"
	tt_array[tt_greater_than_equal] = "greater_than_equal"
	tt_array[tt_equal_to] = "equal_to"
	tt_array[tt_not_equal_to] = "not_equal_to"
	tt_array[tt_and] = "and"
	tt_array[tt_xor] = "xor"
	tt_array[tt_or] = "or"
	tt_array[tt_conditional_and] = "conditional_and"
	tt_array[tt_conditional_or] = "conditional_or"
	tt_array[tt_multiply_assignment] = "multiply_assignment"
	tt_array[tt_divide_assignment] = "divide_assignment"
	tt_array[tt_remainder_assignment] = "remainder_assignment"
	tt_array[tt_plus_assignment] = "plus_assignment"
	tt_array[tt_minus_assignment] = "minus_assignment"
	tt_array[tt_left_shift_assignment] = "left_shift_assignment"
	tt_array[tt_right_shift_assignment] = "right_shift_assignment"
	tt_array[tt_unsigned_right_shift_assignment] = "unsigned_right_shift_assignment"
	tt_array[tt_and_assignment] = "and_assignment"
	tt_array[tt_xor_assignment] = "xor_assignment"
	tt_array[tt_or_assignment] = "or_assignment"
	tt_array[tt_int_literal] = "int_literal"
	tt_array[tt_long_literal] = "long_literal"
	tt_array[tt_float_literal] = "float_literal"
	tt_array[tt_double_literal] = "double_literal"
	tt_array[tt_boolean_literal] = "boolean_literal"
	tt_array[tt_char_literal] = "char_literal"
	tt_array[tt_string_literal] = "string_literal"
	tt_array[tt_null_literal] = "null_literal"
	tt_array[tt_unbounded_wildcard] = "unbounded_wildcard"
	tt_array[tt_extends_wildcard] = "extends_wildcard"
	tt_array[tt_super_wildcard] = "super_wildcard"
	tt_array[tt_erroneous] = "erroneous"
	tt_array[tt_interface] = "interface"
	tt_array[tt_enum] = "enum"
	tt_array[tt_annotation_type] = "annotation_type"
	tt_array[tt_other] = "other"
}

/**
 * 抽象语法树的节点的类型，暂时先写个枚举
 */
type TreeType int

const (
	tt_nil                             TreeType = 0   // nil
	tt_annotated_type                  TreeType = 1   // annotated_type
	tt_annotation                      TreeType = 2   // annotation
	tt_type_annotation                 TreeType = 3   // type_annotation
	tt_array_access                    TreeType = 4   // array_access
	tt_array_type                      TreeType = 5   // array_type
	tt_assert                          TreeType = 6   // assert
	tt_assignment                      TreeType = 7   // assignment
	tt_block                           TreeType = 8   // block
	tt_break                           TreeType = 9   // break
	tt_case                            TreeType = 10  // case
	tt_catch                           TreeType = 11  // catch
	tt_class                           TreeType = 12  // class
	tt_compilation_unit                TreeType = 13  // compilation_unit
	tt_conditional_expression          TreeType = 14  // conditional_expression
	tt_continue                        TreeType = 15  // continue
	tt_do_while_loop                   TreeType = 16  // do_while_loop
	tt_enhanced_for_loop               TreeType = 17  // enhanced_for_loop
	tt_expression_statement            TreeType = 18  // expression_statement
	tt_member_select                   TreeType = 19  // member_select
	tt_member_reference                TreeType = 20  // member_reference
	tt_for_loop                        TreeType = 21  // for_loop
	tt_identifier                      TreeType = 22  // identifier
	tt_if                              TreeType = 23  // if
	tt_import                          TreeType = 24  // import
	tt_instance_of                     TreeType = 25  // instance_of
	tt_labeled_statement               TreeType = 26  // labeled_statement
	tt_method                          TreeType = 27  // method
	tt_method_invocation               TreeType = 28  // method_invocation
	tt_modifiers                       TreeType = 29  // modifiers
	tt_new_array                       TreeType = 30  // new_array
	tt_new_class                       TreeType = 31  // new_class
	tt_lambda_expression               TreeType = 32  // lambda_expression
	tt_parenthesized                   TreeType = 33  // parenthesized
	tt_primitive_type                  TreeType = 34  // primitive_type
	tt_return                          TreeType = 35  // return
	tt_empty_statement                 TreeType = 36  // empty_statement
	tt_switch                          TreeType = 37  // switch
	tt_synchronized                    TreeType = 38  // synchronized
	tt_throw                           TreeType = 39  // throw
	tt_try                             TreeType = 40  // try
	tt_parameterized_type              TreeType = 41  // parameterized_type
	tt_union_type                      TreeType = 42  // union_type
	tt_intersection_type               TreeType = 43  // intersection_type
	tt_type_cast                       TreeType = 44  // type_cast
	tt_type_parameter                  TreeType = 45  // type_parameter
	tt_variable                        TreeType = 46  // variable
	tt_while_loop                      TreeType = 47  // while_loop
	tt_postfix_increment               TreeType = 48  // postfix_increment
	tt_postfix_decrement               TreeType = 49  // postfix_decrement
	tt_prefix_increment                TreeType = 50  // prefix_increment
	tt_prefix_decrement                TreeType = 51  // prefix_decrement
	tt_unary_plus                      TreeType = 52  // unary_plus
	tt_unary_minus                     TreeType = 53  // unary_minus
	tt_bitwise_complement              TreeType = 54  // bitwise_complement
	tt_logical_complement              TreeType = 55  // logical_complement
	tt_multiply                        TreeType = 56  // multiply
	tt_divide                          TreeType = 57  // divide
	tt_remainder                       TreeType = 58  // remainder
	tt_plus                            TreeType = 59  // plus
	tt_minus                           TreeType = 60  // minus
	tt_left_shift                      TreeType = 61  // left_shift
	tt_right_shift                     TreeType = 62  // right_shift
	tt_unsigned_right_shift            TreeType = 63  // unsigned_right_shift
	tt_less_than                       TreeType = 64  // less_than
	tt_greater_than                    TreeType = 65  // greater_than
	tt_less_than_equal                 TreeType = 66  // less_than_equal
	tt_greater_than_equal              TreeType = 67  // greater_than_equal
	tt_equal_to                        TreeType = 68  // equal_to
	tt_not_equal_to                    TreeType = 69  // not_equal_to
	tt_and                             TreeType = 70  // and
	tt_xor                             TreeType = 71  // xor
	tt_or                              TreeType = 72  // or
	tt_conditional_and                 TreeType = 73  // conditional_and
	tt_conditional_or                  TreeType = 74  // conditional_or
	tt_multiply_assignment             TreeType = 75  // multiply_assignment
	tt_divide_assignment               TreeType = 76  // divide_assignment
	tt_remainder_assignment            TreeType = 77  // remainder_assignment
	tt_plus_assignment                 TreeType = 78  // plus_assignment
	tt_minus_assignment                TreeType = 79  // minus_assignment
	tt_left_shift_assignment           TreeType = 80  // left_shift_assignment
	tt_right_shift_assignment          TreeType = 81  // right_shift_assignment
	tt_unsigned_right_shift_assignment TreeType = 82  // unsigned_right_shift_assignment
	tt_and_assignment                  TreeType = 83  // and_assignment
	tt_xor_assignment                  TreeType = 84  // xor_assignment
	tt_or_assignment                   TreeType = 85  // or_assignment
	tt_int_literal                     TreeType = 86  // int_literal
	tt_long_literal                    TreeType = 87  // long_literal
	tt_float_literal                   TreeType = 88  // float_literal
	tt_double_literal                  TreeType = 89  // double_literal
	tt_boolean_literal                 TreeType = 90  // boolean_literal
	tt_char_literal                    TreeType = 91  // char_literal
	tt_string_literal                  TreeType = 92  // string_literal
	tt_null_literal                    TreeType = 93  // null_literal
	tt_unbounded_wildcard              TreeType = 94  // unbounded_wildcard
	tt_extends_wildcard                TreeType = 95  // extends_wildcard
	tt_super_wildcard                  TreeType = 96  // super_wildcard
	tt_erroneous                       TreeType = 97  // erroneous
	tt_interface                       TreeType = 98  // interface
	tt_enum                            TreeType = 99  // enum
	tt_annotation_type                 TreeType = 100 // annotation_type
	tt_other                           TreeType = 101 // other
)

func GetTreeTypeName(tt TreeType) string {

	if tt >= 0 && tt <= 101 {
		return tt_array[tt]
	}
	return "unknown tree type:" + fmt.Sprint(tt)
}
