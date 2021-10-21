package parser

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

	// Token_kind_array[EOF] = ""
	// Token_kind_array[ERROR] = ""
	Token_kind_array[ABSTRACT] = "abstract"
	Token_kind_array[BREAK] = "break"
	Token_kind_array[CASE] = "case"
	Token_kind_array[CATCH] = "catch"
	Token_kind_array[CLASS] = "class"
	Token_kind_array[CONST] = "const"
	Token_kind_array[CONTINUE] = "continue"
	Token_kind_array[DEF] = "default"
	Token_kind_array[DO] = "do"
	Token_kind_array[ELSE] = "else"
	Token_kind_array[EXTENDS] = "extends"
	Token_kind_array[FINAL] = "final"
	Token_kind_array[FINALLY] = "finally"
	Token_kind_array[FOR] = "for"
	Token_kind_array[GOTO] = "goto"
	Token_kind_array[IF] = "if"
	Token_kind_array[IMPLEMENTS] = "implements"
	Token_kind_array[IMPORT] = "import"
	Token_kind_array[INSTANCEOF] = "instanceof"
	Token_kind_array[INTERFACE] = "interface"
	Token_kind_array[NEW] = "new"
	Token_kind_array[PACKAGE] = "package"
	Token_kind_array[PRIVATE] = "private"
	Token_kind_array[PROTECTED] = "protected"
	Token_kind_array[PUBLIC] = "public"
	Token_kind_array[RETURN] = "return"
	Token_kind_array[NATIVE] = "native"
	Token_kind_array[STATIC] = "static"
	Token_kind_array[STRICTFP] = "strictfp"
	Token_kind_array[SWITCH] = "switch"
	Token_kind_array[SYNCHRONIZED] = "synchronized"
	Token_kind_array[THROWS] = "throws"
	Token_kind_array[TRANSIENT] = "transient"
	Token_kind_array[TRY] = "try"
	Token_kind_array[THROW] = "throw"
	Token_kind_array[VOLATILE] = "volatile"
	Token_kind_array[WHILE] = "while"
	Token_kind_array[ARROW] = "->"
	Token_kind_array[COLCOL] = "::"
	Token_kind_array[LPAREN] = "("
	Token_kind_array[RPAREN] = ")"
	Token_kind_array[LBRACE] = "{"
	Token_kind_array[RBRACE] = "}"
	Token_kind_array[LBRACKET] = "["
	Token_kind_array[RBRACKET] = "]"
	Token_kind_array[SEMI] = ";"
	Token_kind_array[COMMA] = ","
	Token_kind_array[DOT] = "."
	Token_kind_array[ELLIPSIS] = "..."
	Token_kind_array[EQ] = "="
	Token_kind_array[GT] = ">"
	Token_kind_array[LT] = "<"
	Token_kind_array[BANG] = "!"
	Token_kind_array[TILDE] = "~"
	Token_kind_array[QUES] = "?"
	Token_kind_array[COLON] = ":"
	Token_kind_array[EQEQ] = "=="
	Token_kind_array[LTEQ] = "<="
	Token_kind_array[GTEQ] = ">="
	Token_kind_array[BANGEQ] = "!="
	Token_kind_array[AMPAMP] = "&&"
	Token_kind_array[BARBAR] = "||"
	Token_kind_array[PLUSPLUS] = "++"
	Token_kind_array[SUBSUB] = "--"
	Token_kind_array[PLUS] = "+"
	Token_kind_array[SUB] = "-"
	Token_kind_array[STAR] = "*"
	// Token_kind_array[SLASH] = ""
	Token_kind_array[STAREQ] = "*="
	Token_kind_array[AMP] = "&"
	Token_kind_array[BAR] = "|"
	Token_kind_array[CARET] = "^"
	Token_kind_array[PERCENT] = "%"
	Token_kind_array[LTLT] = "<<"
	Token_kind_array[GTGT] = ">>"
	Token_kind_array[GTGTGT] = ">>>"
	Token_kind_array[PLUSEQ] = "+="
	Token_kind_array[SUBEQ] = "-="
	Token_kind_array[SLASHEQ] = "/="
	Token_kind_array[AMPEQ] = "&="
	Token_kind_array[BAREQ] = "|="
	Token_kind_array[CARETEQ] = "^="
	Token_kind_array[PERCENTEQ] = "%="
	Token_kind_array[LTLTEQ] = "<<="
	Token_kind_array[GTGTEQ] = ">>="
	Token_kind_array[GTGTGTEQ] = ">>>="
	Token_kind_array[MONKEYS_AT] = "@"
	Token_kind_array[CUSTOM] = ","
	// Token_kind_array[IDENTIFIER] = ""
	Token_kind_array[ASSERT] = "assert"
	Token_kind_array[BOOLEAN] = "boolean"
	Token_kind_array[BYTE] = "byte"
	Token_kind_array[CHAR] = "char"
	Token_kind_array[DOUBLE] = "double"
	Token_kind_array[ENUM] = "enum"
	Token_kind_array[FLOAT] = "float"
	Token_kind_array[INT] = "int"
	Token_kind_array[LONG] = "long"
	Token_kind_array[SHORT] = "short"
	Token_kind_array[SUPER] = "super"
	Token_kind_array[THIS] = "this"
	Token_kind_array[VOID] = "void"
	Token_kind_array[TRUE] = "true"
	Token_kind_array[FALSE] = "false"
	Token_kind_array[NULL] = "null"
	Token_kind_array[UNDERSCORE] = "_"
}

type TokenKind int

// default
const EOF TokenKind = 0         // eof
const ERROR TokenKind = 1       // error
const ABSTRACT TokenKind = 2    // abstract
const BREAK TokenKind = 3       // break
const CASE TokenKind = 4        // case
const CATCH TokenKind = 5       // catch
const CLASS TokenKind = 6       // class
const CONST TokenKind = 7       // const
const CONTINUE TokenKind = 8    // continue
const DEF TokenKind = 9         // default
const DO TokenKind = 10         // do
const ELSE TokenKind = 11       // else
const EXTENDS TokenKind = 12    // extends
const FINAL TokenKind = 13      // final
const FINALLY TokenKind = 14    // finally
const FOR TokenKind = 15        // for
const GOTO TokenKind = 16       // goto
const IF TokenKind = 17         // if
const IMPLEMENTS TokenKind = 18 // implements
const IMPORT TokenKind = 19     // import
const INSTANCEOF TokenKind = 20 // instanceof
const INTERFACE TokenKind = 21  // interface
const NEW TokenKind = 22        // new
const PACKAGE TokenKind = 23    // package
const PRIVATE TokenKind = 24    // private
const PROTECTED TokenKind = 25  // protected
const PUBLIC TokenKind = 26     // public
const RETURN TokenKind = 27     // return
const NATIVE TokenKind = 28     // native
const STATIC TokenKind = 29     // static
// java2开始有的一个关键字 声明 类、接口、方法，作用是严格按IEEE-754执行浮点数字计算，否则Java怎么来无法确认
const STRICTFP TokenKind = 30     // strictfp
const SWITCH TokenKind = 31       // switch
const SYNCHRONIZED TokenKind = 32 // synchronized
const THROWS TokenKind = 33       // throws
const TRANSIENT TokenKind = 34    // transient
const TRY TokenKind = 35          // try
const THROW TokenKind = 36        // throw
const VOLATILE TokenKind = 37     // volatile
const WHILE TokenKind = 38        // while
const ARROW TokenKind = 39        // ->
const COLCOL TokenKind = 40       // ::
const LPAREN TokenKind = 41       // (
const RPAREN TokenKind = 42       // )
const LBRACE TokenKind = 43       // {
const RBRACE TokenKind = 44       // }
const LBRACKET TokenKind = 45     // [
const RBRACKET TokenKind = 46     // ]
const SEMI TokenKind = 47         // ;
const COMMA TokenKind = 48        // ,
const DOT TokenKind = 49          // .
const ELLIPSIS TokenKind = 50     // ...
const EQ TokenKind = 51           // =
const GT TokenKind = 52           // >
const LT TokenKind = 53           // <
const BANG TokenKind = 54         // !
const TILDE TokenKind = 55        // ~
const QUES TokenKind = 56         // ?
const COLON TokenKind = 57        // :
const EQEQ TokenKind = 58         // ==
const LTEQ TokenKind = 59         // <=
const GTEQ TokenKind = 60         // >=
const BANGEQ TokenKind = 61       // !=
const AMPAMP TokenKind = 62       // &&
const BARBAR TokenKind = 63       // ||
const PLUSPLUS TokenKind = 64     // ++
const SUBSUB TokenKind = 65       // --

// 运算符号
const PLUS TokenKind = 66    // +
const SUB TokenKind = 67     // -
const STAR TokenKind = 68    // *
const SLASH TokenKind = 69   // /
const STAREQ TokenKind = 70  // =
const AMP TokenKind = 71     // &
const BAR TokenKind = 72     // |
const CARET TokenKind = 73   // ^
const PERCENT TokenKind = 74 // %

const LTLT TokenKind = 75   // <<
const GTGT TokenKind = 76   // >>
const GTGTGT TokenKind = 77 // >>>

// AssignmentOperator 78 到 87 都是赋值表达式符号
const PLUSEQ TokenKind = 78    // +=
const SUBEQ TokenKind = 79     // -=
const SLASHEQ TokenKind = 80   // /=
const AMPEQ TokenKind = 81     // &=
const BAREQ TokenKind = 82     // |=
const CARETEQ TokenKind = 83   // ^=
const PERCENTEQ TokenKind = 84 // %=
const LTLTEQ TokenKind = 85    // <<=
const GTGTEQ TokenKind = 86    // >>=
const GTGTGTEQ TokenKind = 87  // >>>=

const MONKEYS_AT TokenKind = 88 // @
const CUSTOM TokenKind = 89     // ,

// named
const IDENTIFIER TokenKind = 100
const ASSERT TokenKind = 101     // assert
const BOOLEAN TokenKind = 102    // boolean
const BYTE TokenKind = 103       // byte
const CHAR TokenKind = 104       // char
const DOUBLE TokenKind = 105     // double
const ENUM TokenKind = 106       // enum
const FLOAT TokenKind = 107      // float
const INT TokenKind = 108        // int
const LONG TokenKind = 109       // long
const SHORT TokenKind = 110      // short
const SUPER TokenKind = 111      // super
const THIS TokenKind = 112       // this
const VOID TokenKind = 113       // void
const TRUE TokenKind = 114       // true
const FALSE TokenKind = 115      // false
const NULL TokenKind = 116       // null
const UNDERSCORE TokenKind = 117 // _

// string
const STRINGLITERAL TokenKind = 200 // ""

// numeric
const INTLITERAL TokenKind = 201
const LONGLITERAL TokenKind = 202
const FLOATLITERAL TokenKind = 203
const DOUBLELITERAL TokenKind = 204
const CHARLITERAL TokenKind = 205

func AcceptTokenKind(expected TokenKind, real TokenKind) bool {

	return expected == real
}

func GetTokenString(tk TokenKind) string {

	if tk >= 2 && tk <= 117 {
		return Token_kind_array[int(tk)]
	}
	return ""
}
