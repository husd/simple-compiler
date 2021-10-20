package common

import "fmt"

/**
 * 所有词法分析器解析出来的符号，都被归类为tokenKind了，肯定属于其中之一
 * 设计程序语言，这里可以是入口。
 * @author hushengdong
 */

type TokenTag int

const (
	TOKEN_TAG_DEFAULT TokenTag = 1
	TOKEN_TAG_NAMED   TokenTag = 2
	TOKEN_TAG_STRING  TokenTag = 3
	TOKEN_TAG_NUMERIC TokenTag = 4
)

func GetTokenKindTag(tk TokenKind) TokenTag {

	if tk >= 0 && tk <= 89 {
		return TOKEN_TAG_DEFAULT
	}
	if tk >= 100 && tk <= 117 {
		return TOKEN_TAG_NAMED
	}
	if tk == 200 {
		return TOKEN_TAG_STRING
	}
	if tk >= 201 && tk <= 205 {
		return TOKEN_TAG_NUMERIC
	}
	panic(fmt.Sprintf("无效的token kind: %d", tk))
}

var Token_kind_array [118]string = [118]string{}

func init() {

	// Token_kind_array[TOKEN_KIND_EOF] = ""
	// Token_kind_array[TOKEN_KIND_ERROR] = ""
	Token_kind_array[TOKEN_KIND_ABSTRACT] = "abstract"
	Token_kind_array[TOKEN_KIND_BREAK] = "break"
	Token_kind_array[TOKEN_KIND_CASE] = "case"
	Token_kind_array[TOKEN_KIND_CATCH] = "catch"
	Token_kind_array[TOKEN_KIND_CLASS] = "class"
	Token_kind_array[TOKEN_KIND_CONST] = "const"
	Token_kind_array[TOKEN_KIND_CONTINUE] = "continue"
	Token_kind_array[TOKEN_KIND_DEF] = "default"
	Token_kind_array[TOKEN_KIND_DO] = "do"
	Token_kind_array[TOKEN_KIND_ELSE] = "else"
	Token_kind_array[TOKEN_KIND_EXTENDS] = "extends"
	Token_kind_array[TOKEN_KIND_FINAL] = "final"
	Token_kind_array[TOKEN_KIND_FINALLY] = "finally"
	Token_kind_array[TOKEN_KIND_FOR] = "for"
	Token_kind_array[TOKEN_KIND_GOTO] = "goto"
	Token_kind_array[TOKEN_KIND_IF] = "if"
	Token_kind_array[TOKEN_KIND_IMPLEMENTS] = "implements"
	Token_kind_array[TOKEN_KIND_IMPORT] = "import"
	Token_kind_array[TOKEN_KIND_INSTANCEOF] = "instanceof"
	Token_kind_array[TOKEN_KIND_INTERFACE] = "interface"
	Token_kind_array[TOKEN_KIND_NEW] = "new"
	Token_kind_array[TOKEN_KIND_PACKAGE] = "package"
	Token_kind_array[TOKEN_KIND_PRIVATE] = "private"
	Token_kind_array[TOKEN_KIND_PROTECTED] = "protected"
	Token_kind_array[TOKEN_KIND_PUBLIC] = "public"
	Token_kind_array[TOKEN_KIND_RETURN] = "return"
	Token_kind_array[TOKEN_KIND_NATIVE] = "native"
	Token_kind_array[TOKEN_KIND_STATIC] = "static"
	Token_kind_array[TOKEN_KIND_STRICTFP] = "strictfp"
	Token_kind_array[TOKEN_KIND_SWITCH] = "switch"
	Token_kind_array[TOKEN_KIND_SYNCHRONIZED] = "synchronized"
	Token_kind_array[TOKEN_KIND_THROWS] = "throws"
	Token_kind_array[TOKEN_KIND_TRANSIENT] = "transient"
	Token_kind_array[TOKEN_KIND_TRY] = "try"
	Token_kind_array[TOKEN_KIND_THROW] = "throw"
	Token_kind_array[TOKEN_KIND_VOLATILE] = "volatile"
	Token_kind_array[TOKEN_KIND_WHILE] = "while"
	Token_kind_array[TOKEN_KIND_ARROW] = "->"
	Token_kind_array[TOKEN_KIND_COLCOL] = "::"
	Token_kind_array[TOKEN_KIND_LPAREN] = "("
	Token_kind_array[TOKEN_KIND_RPAREN] = ")"
	Token_kind_array[TOKEN_KIND_LBRACE] = "{"
	Token_kind_array[TOKEN_KIND_RBRACE] = "}"
	Token_kind_array[TOKEN_KIND_LBRACKET] = "["
	Token_kind_array[TOKEN_KIND_RBRACKET] = "]"
	Token_kind_array[TOKEN_KIND_SEMI] = ";"
	Token_kind_array[TOKEN_KIND_COMMA] = ","
	Token_kind_array[TOKEN_KIND_DOT] = "."
	Token_kind_array[TOKEN_KIND_ELLIPSIS] = "..."
	Token_kind_array[TOKEN_KIND_EQ] = "="
	Token_kind_array[TOKEN_KIND_GT] = ">"
	Token_kind_array[TOKEN_KIND_LT] = "<"
	Token_kind_array[TOKEN_KIND_BANG] = "!"
	Token_kind_array[TOKEN_KIND_TILDE] = "~"
	Token_kind_array[TOKEN_KIND_QUES] = "?"
	Token_kind_array[TOKEN_KIND_COLON] = ":"
	Token_kind_array[TOKEN_KIND_EQEQ] = "=="
	Token_kind_array[TOKEN_KIND_LTEQ] = "<="
	Token_kind_array[TOKEN_KIND_GTEQ] = ">="
	Token_kind_array[TOKEN_KIND_BANGEQ] = "!="
	Token_kind_array[TOKEN_KIND_AMPAMP] = "&&"
	Token_kind_array[TOKEN_KIND_BARBAR] = "||"
	Token_kind_array[TOKEN_KIND_PLUSPLUS] = "++"
	Token_kind_array[TOKEN_KIND_SUBSUB] = "--"
	Token_kind_array[TOKEN_KIND_PLUS] = "+"
	Token_kind_array[TOKEN_KIND_SUB] = "-"
	Token_kind_array[TOKEN_KIND_STAR] = "*"
	// Token_kind_array[TOKEN_KIND_SLASH] = ""
	Token_kind_array[TOKEN_KIND_STAREQ] = "*="
	Token_kind_array[TOKEN_KIND_AMP] = "&"
	Token_kind_array[TOKEN_KIND_BAR] = "|"
	Token_kind_array[TOKEN_KIND_CARET] = "^"
	Token_kind_array[TOKEN_KIND_PERCENT] = "%"
	Token_kind_array[TOKEN_KIND_LTLT] = "<<"
	Token_kind_array[TOKEN_KIND_GTGT] = ">>"
	Token_kind_array[TOKEN_KIND_GTGTGT] = ">>>"
	Token_kind_array[TOKEN_KIND_PLUSEQ] = "+="
	Token_kind_array[TOKEN_KIND_SUBEQ] = "-="
	Token_kind_array[TOKEN_KIND_SLASHEQ] = "/="
	Token_kind_array[TOKEN_KIND_AMPEQ] = "&="
	Token_kind_array[TOKEN_KIND_BAREQ] = "|="
	Token_kind_array[TOKEN_KIND_CARETEQ] = "^="
	Token_kind_array[TOKEN_KIND_PERCENTEQ] = "%="
	Token_kind_array[TOKEN_KIND_LTLTEQ] = "<<="
	Token_kind_array[TOKEN_KIND_GTGTEQ] = ">>="
	Token_kind_array[TOKEN_KIND_GTGTGTEQ] = ">>>="
	Token_kind_array[TOKEN_KIND_MONKEYS_AT] = "@"
	Token_kind_array[TOKEN_KIND_CUSTOM] = ","
	// Token_kind_array[TOKEN_KIND_IDENTIFIER] = ""
	Token_kind_array[TOKEN_KIND_ASSERT] = "assert"
	Token_kind_array[TOKEN_KIND_BOOLEAN] = "boolean"
	Token_kind_array[TOKEN_KIND_BYTE] = "byte"
	Token_kind_array[TOKEN_KIND_CHAR] = "char"
	Token_kind_array[TOKEN_KIND_DOUBLE] = "double"
	Token_kind_array[TOKEN_KIND_ENUM] = "enum"
	Token_kind_array[TOKEN_KIND_FLOAT] = "float"
	Token_kind_array[TOKEN_KIND_INT] = "int"
	Token_kind_array[TOKEN_KIND_LONG] = "long"
	Token_kind_array[TOKEN_KIND_SHORT] = "short"
	Token_kind_array[TOKEN_KIND_SUPER] = "super"
	Token_kind_array[TOKEN_KIND_THIS] = "this"
	Token_kind_array[TOKEN_KIND_VOID] = "void"
	Token_kind_array[TOKEN_KIND_TRUE] = "true"
	Token_kind_array[TOKEN_KIND_FALSE] = "false"
	Token_kind_array[TOKEN_KIND_NULL] = "null"
	Token_kind_array[TOKEN_KIND_UNDERSCORE] = "_"
}

type TokenKind int

// default
const TOKEN_KIND_EOF TokenKind = 0         // eof
const TOKEN_KIND_ERROR TokenKind = 1       // error
const TOKEN_KIND_ABSTRACT TokenKind = 2    // abstract
const TOKEN_KIND_BREAK TokenKind = 3       // break
const TOKEN_KIND_CASE TokenKind = 4        // case
const TOKEN_KIND_CATCH TokenKind = 5       // catch
const TOKEN_KIND_CLASS TokenKind = 6       // class
const TOKEN_KIND_CONST TokenKind = 7       // const
const TOKEN_KIND_CONTINUE TokenKind = 8    // continue
const TOKEN_KIND_DEF TokenKind = 9         // default
const TOKEN_KIND_DO TokenKind = 10         // do
const TOKEN_KIND_ELSE TokenKind = 11       // else
const TOKEN_KIND_EXTENDS TokenKind = 12    // extends
const TOKEN_KIND_FINAL TokenKind = 13      // final
const TOKEN_KIND_FINALLY TokenKind = 14    // finally
const TOKEN_KIND_FOR TokenKind = 15        // for
const TOKEN_KIND_GOTO TokenKind = 16       // goto
const TOKEN_KIND_IF TokenKind = 17         // if
const TOKEN_KIND_IMPLEMENTS TokenKind = 18 // implements
const TOKEN_KIND_IMPORT TokenKind = 19     // import
const TOKEN_KIND_INSTANCEOF TokenKind = 20 // instanceof
const TOKEN_KIND_INTERFACE TokenKind = 21  // interface
const TOKEN_KIND_NEW TokenKind = 22        // new
const TOKEN_KIND_PACKAGE TokenKind = 23    // package
const TOKEN_KIND_PRIVATE TokenKind = 24    // private
const TOKEN_KIND_PROTECTED TokenKind = 25  // protected
const TOKEN_KIND_PUBLIC TokenKind = 26     // public
const TOKEN_KIND_RETURN TokenKind = 27     // return
const TOKEN_KIND_NATIVE TokenKind = 28     // native
const TOKEN_KIND_STATIC TokenKind = 29     // static
// java2开始有的一个关键字 声明 类、接口、方法，作用是严格按IEEE-754执行浮点数字计算，否则Java怎么来无法确认
const TOKEN_KIND_STRICTFP TokenKind = 30     // strictfp
const TOKEN_KIND_SWITCH TokenKind = 31       // switch
const TOKEN_KIND_SYNCHRONIZED TokenKind = 32 // synchronized
const TOKEN_KIND_THROWS TokenKind = 33       // throws
const TOKEN_KIND_TRANSIENT TokenKind = 34    // transient
const TOKEN_KIND_TRY TokenKind = 35          // try
const TOKEN_KIND_THROW TokenKind = 36        // throw
const TOKEN_KIND_VOLATILE TokenKind = 37     // volatile
const TOKEN_KIND_WHILE TokenKind = 38        // while
const TOKEN_KIND_ARROW TokenKind = 39        // ->
const TOKEN_KIND_COLCOL TokenKind = 40       // ::
const TOKEN_KIND_LPAREN TokenKind = 41       // (
const TOKEN_KIND_RPAREN TokenKind = 42       // )
const TOKEN_KIND_LBRACE TokenKind = 43       // {
const TOKEN_KIND_RBRACE TokenKind = 44       // }
const TOKEN_KIND_LBRACKET TokenKind = 45     // [
const TOKEN_KIND_RBRACKET TokenKind = 46     // ]
const TOKEN_KIND_SEMI TokenKind = 47         // ;
const TOKEN_KIND_COMMA TokenKind = 48        // ,
const TOKEN_KIND_DOT TokenKind = 49          // .
const TOKEN_KIND_ELLIPSIS TokenKind = 50     // ...
const TOKEN_KIND_EQ TokenKind = 51           // =
const TOKEN_KIND_GT TokenKind = 52           // >
const TOKEN_KIND_LT TokenKind = 53           // <
const TOKEN_KIND_BANG TokenKind = 54         // !
const TOKEN_KIND_TILDE TokenKind = 55        // ~
const TOKEN_KIND_QUES TokenKind = 56         // ?
const TOKEN_KIND_COLON TokenKind = 57        // :
const TOKEN_KIND_EQEQ TokenKind = 58         // ==
const TOKEN_KIND_LTEQ TokenKind = 59         // <=
const TOKEN_KIND_GTEQ TokenKind = 60         // >=
const TOKEN_KIND_BANGEQ TokenKind = 61       // !=
const TOKEN_KIND_AMPAMP TokenKind = 62       // &&
const TOKEN_KIND_BARBAR TokenKind = 63       // ||
const TOKEN_KIND_PLUSPLUS TokenKind = 64     // ++
const TOKEN_KIND_SUBSUB TokenKind = 65       // --

//运算符号
const TOKEN_KIND_PLUS TokenKind = 66    // +
const TOKEN_KIND_SUB TokenKind = 67     // -
const TOKEN_KIND_STAR TokenKind = 68    // *
const TOKEN_KIND_SLASH TokenKind = 69   // /
const TOKEN_KIND_STAREQ TokenKind = 70  // =
const TOKEN_KIND_AMP TokenKind = 71     // &
const TOKEN_KIND_BAR TokenKind = 72     // |
const TOKEN_KIND_CARET TokenKind = 73   // ^
const TOKEN_KIND_PERCENT TokenKind = 74 // %

const TOKEN_KIND_LTLT TokenKind = 75   // <<
const TOKEN_KIND_GTGT TokenKind = 76   // >>
const TOKEN_KIND_GTGTGT TokenKind = 77 // >>>

// AssignmentOperator 78 到 87 都是赋值表达式符号
const TOKEN_KIND_PLUSEQ TokenKind = 78    // +=
const TOKEN_KIND_SUBEQ TokenKind = 79     // -=
const TOKEN_KIND_SLASHEQ TokenKind = 80   // /=
const TOKEN_KIND_AMPEQ TokenKind = 81     // &=
const TOKEN_KIND_BAREQ TokenKind = 82     // |=
const TOKEN_KIND_CARETEQ TokenKind = 83   // ^=
const TOKEN_KIND_PERCENTEQ TokenKind = 84 // %=
const TOKEN_KIND_LTLTEQ TokenKind = 85    // <<=
const TOKEN_KIND_GTGTEQ TokenKind = 86    // >>=
const TOKEN_KIND_GTGTGTEQ TokenKind = 87  // >>>=

const TOKEN_KIND_MONKEYS_AT TokenKind = 88 // @
const TOKEN_KIND_CUSTOM TokenKind = 89     // ,

// named
const TOKEN_KIND_IDENTIFIER TokenKind = 100
const TOKEN_KIND_ASSERT TokenKind = 101     // assert
const TOKEN_KIND_BOOLEAN TokenKind = 102    // boolean
const TOKEN_KIND_BYTE TokenKind = 103       // byte
const TOKEN_KIND_CHAR TokenKind = 104       // char
const TOKEN_KIND_DOUBLE TokenKind = 105     // double
const TOKEN_KIND_ENUM TokenKind = 106       // enum
const TOKEN_KIND_FLOAT TokenKind = 107      // float
const TOKEN_KIND_INT TokenKind = 108        // int
const TOKEN_KIND_LONG TokenKind = 109       // long
const TOKEN_KIND_SHORT TokenKind = 110      // short
const TOKEN_KIND_SUPER TokenKind = 111      // super
const TOKEN_KIND_THIS TokenKind = 112       // this
const TOKEN_KIND_VOID TokenKind = 113       // void
const TOKEN_KIND_TRUE TokenKind = 114       // true
const TOKEN_KIND_FALSE TokenKind = 115      // false
const TOKEN_KIND_NULL TokenKind = 116       // null
const TOKEN_KIND_UNDERSCORE TokenKind = 117 // _

// string
const TOKEN_KIND_STRING_LITERAL TokenKind = 200 // ""

// numeric
const TOKEN_KIND_INT_LITERAL TokenKind = 201
const TOKEN_KIND_LONG_LITERAL TokenKind = 202
const TOKEN_KIND_FLOAT_LITERAL TokenKind = 203
const TOKEN_KIND_DOUBLE_LITERAL TokenKind = 204
const TOKEN_KIND_CHAR_LITERAL TokenKind = 205

func AcceptTokenKind(expected TokenKind, real TokenKind) bool {

	return expected == real
}

func GetTokenString(tk TokenKind) string {

	if tk >= 2 && tk <= 117 {
		return Token_kind_array[int(tk)]
	}
	return ""
}
