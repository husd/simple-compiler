package parser

import (
	"fmt"
	"husd.com/v0/util"
	"testing"
)

/**
 *
 * @author hushengdong
 */

// 测试空格
func TestJavaTokenizer_readToken_blank(t *testing.T) {

	c := util.NewContext()
	s := "      "
	tokenizer := NewJavaTokenizerWithString(s, c)
	tk := tokenizer.readToken()
	util.AssertEquals(t, "测试空格", TOKEN_KIND_EOF, tk.GetTokenKind())
}

//测试注释
func TestJavaTokenizer_readToken(t *testing.T) {

	c := util.NewContext()
	str := []string{
		"/** 多行注释 **/     ",
		"/** sadf     **/    ",
		`/**
			"* 1234 
			"// sdf 
			" **/     `,
		"//",
		"////",
		"// public static void main //    ",
		"// sdf */",
		"// /* ",
	}
	for inx, s := range str {
		tokenizer := NewJavaTokenizerWithString(s, c)
		tk := tokenizer.readToken()
		util.AssertEquals(t, fmt.Sprintf("测试注释 index:%d", inx), TOKEN_KIND_EOF, tk.GetTokenKind())
	}
}

//测试10进制数字
func TestJavaTokenizer_readToken_num(t *testing.T) {

	var tt Token
	//var str string
	c := util.NewContext()

	strArr := []string{"100", "10_0", "1_0_0"}

	for _, str := range strArr {
		tokenizer := NewJavaTokenizerWithString(str, c)
		tt = tokenizer.readToken()
		util.AssertEquals(t, "测试数字-整数类型", TOKEN_KIND_INT_LITERAL, tt.GetTokenKind())
		util.AssertEquals(t, "测试数字-整数数值", "100", tt.GetStringVal())
		util.AssertEquals(t, "测试数字-整数radix", 10, tt.GetRadix())
	}

}

//测试2进制数字
func TestJavaTokenizer_readToken_num2(t *testing.T) {

	var tt Token
	c := util.NewContext()
	strArr := []string{"0b1010", "0b1_010", "0b_1010"}
	for _, str := range strArr {
		tokenizer := NewJavaTokenizerWithString(str, c)
		tt = tokenizer.readToken()
		util.AssertEquals(t, "测试2进制-整数类型", TOKEN_KIND_INT_LITERAL, tt.GetTokenKind())
		util.AssertEquals(t, "测试2进制-整数数值", "1010", tt.GetStringVal())
		util.AssertEquals(t, "测试2进制-整数radix", 2, tt.GetRadix())
	}
}

//测试8进制数字
func TestJavaTokenizer_readToken_num8(t *testing.T) {

	var tt Token
	c := util.NewContext()
	strArr := []string{"01072", "01_072", "0_1072"}
	for _, str := range strArr {
		tokenizer := NewJavaTokenizerWithString(str, c)
		tt = tokenizer.readToken()
		util.AssertEquals(t, "测试8进制-整数类型", TOKEN_KIND_INT_LITERAL, tt.GetTokenKind())
		util.AssertEquals(t, "测试8进制-整数数值", "01072", tt.GetStringVal())
		util.AssertEquals(t, "测试8进制-整数radix", 8, tt.GetRadix())
	}
}

//测试16进制数字
func TestJavaTokenizer_readToken_num16(t *testing.T) {

	var tt Token
	c := util.NewContext()
	strArr := []string{"0Xfffe", "0xf_ffe", "0X_fffe"}
	for _, str := range strArr {
		tokenizer := NewJavaTokenizerWithString(str, c)
		tt = tokenizer.readToken()
		util.AssertEquals(t, "测试16进制-整数类型", TOKEN_KIND_INT_LITERAL, tt.GetTokenKind())
		util.AssertEquals(t, "测试16进制-整数数值", "fffe", tt.GetStringVal())
		util.AssertEquals(t, "测试16进制-整数radix", 16, tt.GetRadix())
	}
}

// 测试 expression
func TestJavaTokenizer_readToken2(t *testing.T) {

	c := util.NewContext()
	str := "   int a  = 10;   "
	tokenizer := NewJavaTokenizerWithString(str, c)
	var tt Token

	tt = tokenizer.readToken()
	//tt = tokenizer.readToken()
	util.AssertEquals(t, "测试数字1 type", TOKEN_KIND_INT, tt.GetTokenKind())
	util.AssertEquals(t, "测试数字1", "int", tt.GetStringVal())

	tt = tokenizer.readToken()
	util.AssertEquals(t, "测试数字2 type", TOKEN_KIND_IDENTIFIER, tt.GetTokenKind())
	util.AssertEquals(t, "测试数字2", "a", tt.GetStringVal())

	tt = tokenizer.readToken()
	util.AssertEquals(t, "测试数字3 type", TOKEN_KIND_EQ, tt.GetTokenKind())
	util.AssertEquals(t, "测试数字3", "=", tt.GetStringVal())

	tt = tokenizer.readToken()
	util.AssertEquals(t, "测试数字4 type", TOKEN_KIND_INT_LITERAL, tt.GetTokenKind())
	util.AssertEquals(t, "测试数字4", "10", tt.GetStringVal())

	tt = tokenizer.readToken()
	util.AssertEquals(t, "测试数字5 type", TOKEN_KIND_SEMI, tt.GetTokenKind())
	util.AssertEquals(t, "测试数字5", ";", tt.GetStringVal())
}

// 测试 expression
func TestJavaTokenizer_readToken3(t *testing.T) {

	c := util.NewContext()
	str := "   String a=\"abc\";   "
	tokenizer := NewJavaTokenizerWithString(str, c)
	var tt Token

	tt = tokenizer.readToken()
	//tt = tokenizer.readToken()
	util.AssertEquals(t, "测试数字1 type", TOKEN_KIND_IDENTIFIER, tt.GetTokenKind())
	util.AssertEquals(t, "测试数字1", "String", tt.GetStringVal())

	tt = tokenizer.readToken()
	util.AssertEquals(t, "测试数字2 type", TOKEN_KIND_IDENTIFIER, tt.GetTokenKind())
	util.AssertEquals(t, "测试数字2", "a", tt.GetStringVal())

	tt = tokenizer.readToken()
	util.AssertEquals(t, "测试数字3 type", TOKEN_KIND_EQ, tt.GetTokenKind())
	util.AssertEquals(t, "测试数字3", "=", tt.GetStringVal())

	tt = tokenizer.readToken()
	util.AssertEquals(t, "测试数字4 type", TOKEN_KIND_STRING_LITERAL, tt.GetTokenKind())
	util.AssertEquals(t, "测试数字4", "abc", tt.GetStringVal())

	tt = tokenizer.readToken()
	util.AssertEquals(t, "测试数字5 type", TOKEN_KIND_SEMI, tt.GetTokenKind())
	util.AssertEquals(t, "测试数字5", ";", tt.GetStringVal())
}

func TestJavaTokenizer_readToken4(t *testing.T) {

	c := util.NewContext()
	str := `
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
`

	str = `
	var TREE_TYPE_NIL = &TreeType{"nil", -1}
`
	tokenizer := NewJavaTokenizerWithString(str, c)
	var tt Token

	tt = tokenizer.readToken()
	for tt.GetTokenKind() != TOKEN_KIND_EOF {
		fmt.Println(tt.GetStringVal())
		tt = tokenizer.readToken()
	}
}
