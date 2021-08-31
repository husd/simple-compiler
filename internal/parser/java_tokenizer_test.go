package parser

import (
	"husd.com/v0/util"
	"testing"
)

func TestJavaTokenizer_readToken(t *testing.T) {

	c := util.NewContext()
	str := []string{
		"/** 多行注释 **/",
		"/** **/",
		"/**" +
			"* 1234 " +
			"// sdf " +
			" **/",
		"//",
		"////",
		"// public static void main //",
		"// sdf */", "// /* ",
	}
	for _, s := range str {
		tokenizer := NewJavaTokenizerWithString(s, c)
		tk := tokenizer.readToken()
		util.AssertEquals(t, "测试注释", TOKEN_KIND_ERROR, tk.GetTokenKind())
	}
}

func TestJavaTokenizer_readToken_num(t *testing.T) {

	c := util.NewContext()
	str := " 100 "
	tokenizer := NewJavaTokenizerWithString(str, c)
	var tt token

	tt = tokenizer.readToken()
	//tt = tokenizer.readToken()
	util.AssertEquals(t, "测试数字1", TOKEN_KIND_INT, tt.GetTokenKind())
	util.AssertEquals(t, "测试数字1", "100", tt.GetStringVal())
}

func TestJavaTokenizer_readToken2(t *testing.T) {

	c := util.NewContext()
	str := "   int a  = 10;   "
	tokenizer := NewJavaTokenizerWithString(str, c)
	var tt token

	tt = tokenizer.readToken()
	//tt = tokenizer.readToken()
	util.AssertEquals(t, "测试数字1", TOKEN_KIND_INT, tt.GetTokenKind())
	util.AssertEquals(t, "测试数字1", "int", tt.GetStringVal())

	tt = tokenizer.readToken()
	util.AssertEquals(t, "测试数字2", TOKEN_KIND_IDENTIFIER, tt.GetTokenKind())
	util.AssertEquals(t, "测试数字2", "a", tt.GetStringVal())

	tt = tokenizer.readToken()
	util.AssertEquals(t, "测试数字3", TOKEN_KIND_EQ, tt.GetTokenKind())
	util.AssertEquals(t, "测试数字3", "=", tt.GetStringVal())

	tt = tokenizer.readToken()
	util.AssertEquals(t, "测试数字4", TOKEN_KIND_INT, tt.GetTokenKind())
	util.AssertEquals(t, "测试数字4", "10", tt.GetStringVal())

	tt = tokenizer.readToken()
	util.AssertEquals(t, "测试数字5", TOKEN_KIND_INT, tt.GetTokenKind())
	util.AssertEquals(t, "测试数字5", "10", tt.GetStringVal())
}
