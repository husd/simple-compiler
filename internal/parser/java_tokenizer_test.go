package parser

import (
	"fmt"
	"husd.com/v0/util"
	"testing"
)

//测试空格
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
		util.AssertEquals(t, "测试数字-整数类型", TOKEN_KIND_INT_LITERAL.Index, tt.GetTokenKind().Index)
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
		util.AssertEquals(t, "测试2进制-整数类型", TOKEN_KIND_INT_LITERAL.Index, tt.GetTokenKind().Index)
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
		util.AssertEquals(t, "测试8进制-整数类型", TOKEN_KIND_INT_LITERAL.Index, tt.GetTokenKind().Index)
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
		util.AssertEquals(t, "测试16进制-整数类型", TOKEN_KIND_INT_LITERAL.Index, tt.GetTokenKind().Index)
		util.AssertEquals(t, "测试16进制-整数数值", "fffe", tt.GetStringVal())
		util.AssertEquals(t, "测试16进制-整数radix", 16, tt.GetRadix())
	}
}

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
