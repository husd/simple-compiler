package parser

import "fmt"

/**
 * 所有词法分析器解析出来的符号，都被归类为tokenKind了，肯定属于其中之一
 * 设计程序语言，这里可以是入口。
 * @author hushengdong
 */

type tokenTag int

const (
	TOKEN_TAG_DEFAULT tokenTag = 1
	TOKEN_TAG_NAMED   tokenTag = 2
	TOKEN_TAG_STRING  tokenTag = 3
	TOKEN_TAG_NUMERIC tokenTag = 4
)

func GetTokenKindTag(tk tokenKind) tokenTag {

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

var token_kind_array [118]string = [118]string{}

func init() {

	// token_kind_array[TOKEN_KIND_EOF] = ""
	// token_kind_array[TOKEN_KIND_ERROR] = ""
	token_kind_array[TOKEN_KIND_ABSTRACT] = "abstract"
	token_kind_array[TOKEN_KIND_BREAK] = "break"
	token_kind_array[TOKEN_KIND_CASE] = "case"
	token_kind_array[TOKEN_KIND_CATCH] = "catch"
	token_kind_array[TOKEN_KIND_CLASS] = "class"
	token_kind_array[TOKEN_KIND_CONST] = "const"
	token_kind_array[TOKEN_KIND_CONTINUE] = "continue"
	token_kind_array[TOKEN_KIND_DEF] = "default"
	token_kind_array[TOKEN_KIND_DO] = "do"
	token_kind_array[TOKEN_KIND_ELSE] = "else"
	token_kind_array[TOKEN_KIND_EXTENDS] = "extends"
	token_kind_array[TOKEN_KIND_FINAL] = "final"
	token_kind_array[TOKEN_KIND_FINALLY] = "finally"
	token_kind_array[TOKEN_KIND_FOR] = "for"
	token_kind_array[TOKEN_KIND_GOTO] = "goto"
	token_kind_array[TOKEN_KIND_IF] = "if"
	token_kind_array[TOKEN_KIND_IMPLEMENTS] = "implements"
	token_kind_array[TOKEN_KIND_IMPORT] = "import"
	token_kind_array[TOKEN_KIND_INSTANCEOF] = "instanceof"
	token_kind_array[TOKEN_KIND_INTERFACE] = "interface"
	token_kind_array[TOKEN_KIND_NEW] = "new"
	token_kind_array[TOKEN_KIND_PACKAGE] = "package"
	token_kind_array[TOKEN_KIND_PRIVATE] = "private"
	token_kind_array[TOKEN_KIND_PROTECTED] = "protected"
	token_kind_array[TOKEN_KIND_PUBLIC] = "public"
	token_kind_array[TOKEN_KIND_RETURN] = "return"
	token_kind_array[TOKEN_KIND_NATIVE] = "native"
	token_kind_array[TOKEN_KIND_STATIC] = "static"
	token_kind_array[TOKEN_KIND_STRICTFP] = "strictfp"
	token_kind_array[TOKEN_KIND_SWITCH] = "switch"
	token_kind_array[TOKEN_KIND_SYNCHRONIZED] = "synchronized"
	token_kind_array[TOKEN_KIND_THROWS] = "throws"
	token_kind_array[TOKEN_KIND_TRANSIENT] = "transient"
	token_kind_array[TOKEN_KIND_TRY] = "try"
	token_kind_array[TOKEN_KIND_THROW] = "throw"
	token_kind_array[TOKEN_KIND_VOLATILE] = "volatile"
	token_kind_array[TOKEN_KIND_WHILE] = "while"
	token_kind_array[TOKEN_KIND_ARROW] = "->"
	token_kind_array[TOKEN_KIND_COLCOL] = "::"
	token_kind_array[TOKEN_KIND_LPAREN] = "("
	token_kind_array[TOKEN_KIND_RPAREN] = ")"
	token_kind_array[TOKEN_KIND_LBRACE] = "{"
	token_kind_array[TOKEN_KIND_RBRACE] = "}"
	token_kind_array[TOKEN_KIND_LBRACKET] = "["
	token_kind_array[TOKEN_KIND_RBRACKET] = "]"
	token_kind_array[TOKEN_KIND_SEMI] = ";"
	token_kind_array[TOKEN_KIND_COMMA] = ","
	token_kind_array[TOKEN_KIND_DOT] = "."
	token_kind_array[TOKEN_KIND_ELLIPSIS] = "..."
	token_kind_array[TOKEN_KIND_EQ] = "="
	token_kind_array[TOKEN_KIND_GT] = ">"
	token_kind_array[TOKEN_KIND_LT] = "<"
	token_kind_array[TOKEN_KIND_BANG] = "!"
	token_kind_array[TOKEN_KIND_TILDE] = "~"
	token_kind_array[TOKEN_KIND_QUES] = "?"
	token_kind_array[TOKEN_KIND_COLON] = ":"
	token_kind_array[TOKEN_KIND_EQEQ] = "=="
	token_kind_array[TOKEN_KIND_LTEQ] = "<="
	token_kind_array[TOKEN_KIND_GTEQ] = ">="
	token_kind_array[TOKEN_KIND_BANGEQ] = "!="
	token_kind_array[TOKEN_KIND_AMPAMP] = "&&"
	token_kind_array[TOKEN_KIND_BARBAR] = "||"
	token_kind_array[TOKEN_KIND_PLUSPLUS] = "++"
	token_kind_array[TOKEN_KIND_SUBSUB] = "--"
	token_kind_array[TOKEN_KIND_PLUS] = "+"
	token_kind_array[TOKEN_KIND_SUB] = "-"
	token_kind_array[TOKEN_KIND_STAR] = "*"
	// token_kind_array[TOKEN_KIND_SLASH] = ""
	token_kind_array[TOKEN_KIND_STAREQ] = "="
	token_kind_array[TOKEN_KIND_AMP] = "&"
	token_kind_array[TOKEN_KIND_BAR] = "|"
	token_kind_array[TOKEN_KIND_CARET] = "^"
	token_kind_array[TOKEN_KIND_PERCENT] = "%"
	token_kind_array[TOKEN_KIND_LTLT] = "<<"
	token_kind_array[TOKEN_KIND_GTGT] = ">>"
	token_kind_array[TOKEN_KIND_GTGTGT] = ">>>"
	token_kind_array[TOKEN_KIND_PLUSEQ] = "+="
	token_kind_array[TOKEN_KIND_SUBEQ] = "-="
	token_kind_array[TOKEN_KIND_SLASHEQ] = "="
	token_kind_array[TOKEN_KIND_AMPEQ] = "&="
	token_kind_array[TOKEN_KIND_BAREQ] = "|="
	token_kind_array[TOKEN_KIND_CARETEQ] = "^="
	token_kind_array[TOKEN_KIND_PERCENTEQ] = "%="
	token_kind_array[TOKEN_KIND_LTLTEQ] = "<<="
	token_kind_array[TOKEN_KIND_GTGTEQ] = ">>="
	token_kind_array[TOKEN_KIND_GTGTGTEQ] = ">>>="
	token_kind_array[TOKEN_KIND_MONKEYS_AT] = "@"
	token_kind_array[TOKEN_KIND_CUSTOM] = ","
	// token_kind_array[TOKEN_KIND_IDENTIFIER] = ""
	token_kind_array[TOKEN_KIND_ASSERT] = "assert"
	token_kind_array[TOKEN_KIND_BOOLEAN] = "boolean"
	token_kind_array[TOKEN_KIND_BYTE] = "byte"
	token_kind_array[TOKEN_KIND_CHAR] = "char"
	token_kind_array[TOKEN_KIND_DOUBLE] = "double"
	token_kind_array[TOKEN_KIND_ENUM] = "enum"
	token_kind_array[TOKEN_KIND_FLOAT] = "float"
	token_kind_array[TOKEN_KIND_INT] = "float"
	token_kind_array[TOKEN_KIND_LONG] = "long"
	token_kind_array[TOKEN_KIND_SHORT] = "short"
	token_kind_array[TOKEN_KIND_SUPER] = "super"
	token_kind_array[TOKEN_KIND_THIS] = "this"
	token_kind_array[TOKEN_KIND_VOID] = "void"
	token_kind_array[TOKEN_KIND_TRUE] = "true"
	token_kind_array[TOKEN_KIND_FALSE] = "false"
	token_kind_array[TOKEN_KIND_NULL] = "null"
	token_kind_array[TOKEN_KIND_UNDERSCORE] = "_"
}

type tokenKind int

// default
const TOKEN_KIND_EOF tokenKind = 0         // eof
const TOKEN_KIND_ERROR tokenKind = 1       // error
const TOKEN_KIND_ABSTRACT tokenKind = 2    // abstract
const TOKEN_KIND_BREAK tokenKind = 3       // break
const TOKEN_KIND_CASE tokenKind = 4        // case
const TOKEN_KIND_CATCH tokenKind = 5       // catch
const TOKEN_KIND_CLASS tokenKind = 6       // class
const TOKEN_KIND_CONST tokenKind = 7       // const
const TOKEN_KIND_CONTINUE tokenKind = 8    // continue
const TOKEN_KIND_DEF tokenKind = 9         // default
const TOKEN_KIND_DO tokenKind = 10         // do
const TOKEN_KIND_ELSE tokenKind = 11       // else
const TOKEN_KIND_EXTENDS tokenKind = 12    // extends
const TOKEN_KIND_FINAL tokenKind = 13      // final
const TOKEN_KIND_FINALLY tokenKind = 14    // finally
const TOKEN_KIND_FOR tokenKind = 15        // for
const TOKEN_KIND_GOTO tokenKind = 16       // goto
const TOKEN_KIND_IF tokenKind = 17         // if
const TOKEN_KIND_IMPLEMENTS tokenKind = 18 // implements
const TOKEN_KIND_IMPORT tokenKind = 19     // import
const TOKEN_KIND_INSTANCEOF tokenKind = 20 // instanceof
const TOKEN_KIND_INTERFACE tokenKind = 21  // interface
const TOKEN_KIND_NEW tokenKind = 22        // new
const TOKEN_KIND_PACKAGE tokenKind = 23    // package
const TOKEN_KIND_PRIVATE tokenKind = 24    // private
const TOKEN_KIND_PROTECTED tokenKind = 25  // protected
const TOKEN_KIND_PUBLIC tokenKind = 26     // public
const TOKEN_KIND_RETURN tokenKind = 27     // return
const TOKEN_KIND_NATIVE tokenKind = 28     // native
const TOKEN_KIND_STATIC tokenKind = 29     // static
// java2开始有的一个关键字 声明 类、接口、方法，作用是严格按IEEE-754执行浮点数字计算，否则Java怎么来无法确认
const TOKEN_KIND_STRICTFP tokenKind = 30     // strictfp
const TOKEN_KIND_SWITCH tokenKind = 31       // switch
const TOKEN_KIND_SYNCHRONIZED tokenKind = 32 // synchronized
const TOKEN_KIND_THROWS tokenKind = 33       // throws
const TOKEN_KIND_TRANSIENT tokenKind = 34    // transient
const TOKEN_KIND_TRY tokenKind = 35          // try
const TOKEN_KIND_THROW tokenKind = 36        // throw
const TOKEN_KIND_VOLATILE tokenKind = 37     // volatile
const TOKEN_KIND_WHILE tokenKind = 38        // while
const TOKEN_KIND_ARROW tokenKind = 39        // ->
const TOKEN_KIND_COLCOL tokenKind = 40       // ::
const TOKEN_KIND_LPAREN tokenKind = 41       // (
const TOKEN_KIND_RPAREN tokenKind = 42       // )
const TOKEN_KIND_LBRACE tokenKind = 43       // {
const TOKEN_KIND_RBRACE tokenKind = 44       // }
const TOKEN_KIND_LBRACKET tokenKind = 45     // [
const TOKEN_KIND_RBRACKET tokenKind = 46     // ]
const TOKEN_KIND_SEMI tokenKind = 47         // ;
const TOKEN_KIND_COMMA tokenKind = 48        // ,
const TOKEN_KIND_DOT tokenKind = 49          // .
const TOKEN_KIND_ELLIPSIS tokenKind = 50     // ...
const TOKEN_KIND_EQ tokenKind = 51           // =
const TOKEN_KIND_GT tokenKind = 52           // >
const TOKEN_KIND_LT tokenKind = 53           // <
const TOKEN_KIND_BANG tokenKind = 54         // !
const TOKEN_KIND_TILDE tokenKind = 55        // ~
const TOKEN_KIND_QUES tokenKind = 56         // ?
const TOKEN_KIND_COLON tokenKind = 57        // :
const TOKEN_KIND_EQEQ tokenKind = 58         // ==
const TOKEN_KIND_LTEQ tokenKind = 59         // <=
const TOKEN_KIND_GTEQ tokenKind = 60         // >=
const TOKEN_KIND_BANGEQ tokenKind = 61       // !=
const TOKEN_KIND_AMPAMP tokenKind = 62       // &&
const TOKEN_KIND_BARBAR tokenKind = 63       // ||
const TOKEN_KIND_PLUSPLUS tokenKind = 64     // ++
const TOKEN_KIND_SUBSUB tokenKind = 65       // --
const TOKEN_KIND_PLUS tokenKind = 66         // +
const TOKEN_KIND_SUB tokenKind = 67          // -
const TOKEN_KIND_STAR tokenKind = 68         // *
const TOKEN_KIND_SLASH tokenKind = 69        // /
const TOKEN_KIND_STAREQ tokenKind = 70       // =
const TOKEN_KIND_AMP tokenKind = 71          // &
const TOKEN_KIND_BAR tokenKind = 72          // |
const TOKEN_KIND_CARET tokenKind = 73        // ^
const TOKEN_KIND_PERCENT tokenKind = 74      // %
const TOKEN_KIND_LTLT tokenKind = 75         // <<
const TOKEN_KIND_GTGT tokenKind = 76         // >>
const TOKEN_KIND_GTGTGT tokenKind = 77       // >>>
const TOKEN_KIND_PLUSEQ tokenKind = 78       // +=
const TOKEN_KIND_SUBEQ tokenKind = 79        // -=

const TOKEN_KIND_SLASHEQ tokenKind = 80    // /=
const TOKEN_KIND_AMPEQ tokenKind = 81      // &=
const TOKEN_KIND_BAREQ tokenKind = 82      // |=
const TOKEN_KIND_CARETEQ tokenKind = 83    // ^=
const TOKEN_KIND_PERCENTEQ tokenKind = 84  // %=
const TOKEN_KIND_LTLTEQ tokenKind = 85     // <<=
const TOKEN_KIND_GTGTEQ tokenKind = 86     // >>=
const TOKEN_KIND_GTGTGTEQ tokenKind = 87   // >>>=
const TOKEN_KIND_MONKEYS_AT tokenKind = 88 // @
const TOKEN_KIND_CUSTOM tokenKind = 89     // ,

// named
const TOKEN_KIND_IDENTIFIER tokenKind = 100
const TOKEN_KIND_ASSERT tokenKind = 101     // assert
const TOKEN_KIND_BOOLEAN tokenKind = 102    // boolean
const TOKEN_KIND_BYTE tokenKind = 103       // byte
const TOKEN_KIND_CHAR tokenKind = 104       // char
const TOKEN_KIND_DOUBLE tokenKind = 105     // double
const TOKEN_KIND_ENUM tokenKind = 106       // enum
const TOKEN_KIND_FLOAT tokenKind = 107      // float
const TOKEN_KIND_INT tokenKind = 108        // float
const TOKEN_KIND_LONG tokenKind = 109       // /long
const TOKEN_KIND_SHORT tokenKind = 110      // short
const TOKEN_KIND_SUPER tokenKind = 111      // super
const TOKEN_KIND_THIS tokenKind = 112       // this
const TOKEN_KIND_VOID tokenKind = 113       // void
const TOKEN_KIND_TRUE tokenKind = 114       // true
const TOKEN_KIND_FALSE tokenKind = 115      // false
const TOKEN_KIND_NULL tokenKind = 116       // null
const TOKEN_KIND_UNDERSCORE tokenKind = 117 // _

// string
const TOKEN_KIND_STRING_LITERAL tokenKind = 200 // ""

// numeric
const TOKEN_KIND_INT_LITERAL tokenKind = 201
const TOKEN_KIND_LONG_LITERAL tokenKind = 202
const TOKEN_KIND_FLOAT_LITERAL tokenKind = 203
const TOKEN_KIND_DOUBLE_LITERAL tokenKind = 204
const TOKEN_KIND_CHAR_LITERAL tokenKind = 205

func AcceptTokenKind(expected tokenKind, real tokenKind) bool {

	return expected == real
}
