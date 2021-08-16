package lexical

import "fmt"

// 这个类主要记录了Java的一些常量 主要是 符号表的常量

// TokenTag 词法分析器的输出结果 Token
type TokenTag struct {
	LineNum   int         // 源代码多少行
	Token     interface{} // 标识符
	TokenType SymbolType  // token的类型
	Index     int         // 符号表里的位置 暂时先定 int  //TODO husd 看看这个符号表位置怎么定
}

//SymbolType 符号表常量
type SymbolType int

//符号表的所有的数据的类型，都统一在这里维护常量，这里不使用 itoa 是为了程序更直观，这里有一个问题还需要确认，就是编译器在哪个阶段
//做的识别的关键字？ 从书上看，在词法分析阶段，并没有识别出关键字，关键字属于语法，应该在语法分析器阶段。
const (
	SymbolTypeId       SymbolType = 1   // 标识符 就是变量 例如 abc123 name 等
	SymbolTypeNum      SymbolType = 2   // 数字 例如 60
	SymbolTypeStr      SymbolType = 3   // 字符串 "abc323"
	SymbolTypeOpNotEq  SymbolType = 101 // 符号  从 101 到 129 一共29个操作符类型
	SymbolTypeKeyword  SymbolType = 201 // 关键字 从 201 开头的 Java的关键字一共有 ? 个 词法分析器不会进行语法分析
	SymbolTypeBaseType SymbolType = 301 // 基本类型 从301 开头
)

// 所有的操作符对应的 SymbolType 的值 22个
//  !=  |  101
//  %=  |  102
//  &=  |  103
//  *=  |  104
//  +=  |  105
//  -=  |  106
//  /=  |  107
//  <=  |  108
//  ==  |  109
//  >=  |  110
//  ^=  |  111
//  |=  |  112
//  ++  |  113
//  --  |  114
//  <<  |  115
//  <<=  |  116
//  >>  |  117
//  >>=  |  118
//  ||  |  119
//  &&  |  120
//  >>>  |  121
//  >>>=  |  122

//基本类型的常量 8个
//  byte  |  301
//  short  |  302
//  int  |  303
//  long  |  304
//  char  |  305
//  float  |  306
//  double  |  307
//  boolean  |  308

var SymbolTypeOpName = [29]string{"+", "-", "*", "/", "&", "|", "%", "!=", "%=", "&=", "*=", "+=", "-=", "/=", "<=", "==", ">=", "^=", "|=", "++", "--", "<<", "<<=", ">>", ">>=", "||", "&&", ">>>", ">>>="}
var SymbolTypeOpType = [29]SymbolType{101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129}

var SymbolTypeBaseTypeName = [8]string{"byte", "short", "int", "long", "char", "float", "double", "boolean"}
var SymbolTypeBaseTypeType = [8]SymbolType{301, 302, 303, 304, 305, 306, 307, 308}

func PrintSymbolTypeOpName() {

	for i := 0; i < 22; i++ {
		fmt.Println("// ", SymbolTypeOpName[i], " | ", SymbolTypeOpType[i])
	}
}

func PrintSymbolTypeBaseType() {

	for i := 0; i < 8; i++ {
		fmt.Println("// ", SymbolTypeBaseTypeName[i], " | ", SymbolTypeBaseTypeType[i])
	}
}
