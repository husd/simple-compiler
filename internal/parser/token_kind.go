package parser

/**
 * 所有词法分析器解析出来的符号，都被归类为tokenKind2了，肯定属于其中之一
 * 设计程序语言，这里可以是入口。
 * @author hushengdong
 */

type tokenKind2 struct {
	Name  string
	Tag   tokenTag
	Index int // 在key 数组中的位置
}

var TOKEN_KIND_2_EOF = &tokenKind2{"eof", TOKEN_TAG_DEFAULT, 0} // EOF特指文件结束
var TOKEN_KIND_2_ERROR = &tokenKind2{"error", TOKEN_TAG_DEFAULT, 1}
var TOKEN_KIND_2_IDENTIFIER = &tokenKind2{"", TOKEN_TAG_NAMED, 2}
var TOKEN_KIND_2_ABSTRACT = &tokenKind2{"abstract", TOKEN_TAG_DEFAULT, 3}
var TOKEN_KIND_2_ASSERT = &tokenKind2{"assert", TOKEN_TAG_NAMED, 4}
var TOKEN_KIND_2_BOOLEAN = &tokenKind2{"boolean", TOKEN_TAG_NAMED, 5}
var TOKEN_KIND_2_BREAK = &tokenKind2{"break", TOKEN_TAG_DEFAULT, 6}
var TOKEN_KIND_2_BYTE = &tokenKind2{"byte", TOKEN_TAG_NAMED, 7}
var TOKEN_KIND_2_CASE = &tokenKind2{"case", TOKEN_TAG_DEFAULT, 8}
var TOKEN_KIND_2_CATCH = &tokenKind2{"catch", TOKEN_TAG_DEFAULT, 9}
var TOKEN_KIND_2_CHAR = &tokenKind2{"char", TOKEN_TAG_NAMED, 10}
var TOKEN_KIND_2_CLASS = &tokenKind2{"class", TOKEN_TAG_DEFAULT, 11}
var TOKEN_KIND_2_CONST = &tokenKind2{"const", TOKEN_TAG_DEFAULT, 12}
var TOKEN_KIND_2_CONTINUE = &tokenKind2{"continue", TOKEN_TAG_DEFAULT, 13}
var TOKEN_KIND_2_DEF = &tokenKind2{"default", TOKEN_TAG_DEFAULT, 14}
var TOKEN_KIND_2_DO = &tokenKind2{"do", TOKEN_TAG_DEFAULT, 15}
var TOKEN_KIND_2_DOUBLE = &tokenKind2{"double", TOKEN_TAG_NAMED, 16}
var TOKEN_KIND_2_ELSE = &tokenKind2{"else", TOKEN_TAG_DEFAULT, 17}
var TOKEN_KIND_2_ENUM = &tokenKind2{"enum", TOKEN_TAG_NAMED, 18}
var TOKEN_KIND_2_EXTENDS = &tokenKind2{"extends", TOKEN_TAG_DEFAULT, 19}
var TOKEN_KIND_2_FINAL = &tokenKind2{"final", TOKEN_TAG_DEFAULT, 20}
var TOKEN_KIND_2_FINALLY = &tokenKind2{"finally", TOKEN_TAG_DEFAULT, 21}
var TOKEN_KIND_2_FLOAT = &tokenKind2{"float", TOKEN_TAG_NAMED, 22}
var TOKEN_KIND_2_FOR = &tokenKind2{"for", TOKEN_TAG_DEFAULT, 23}
var TOKEN_KIND_2_GOTO = &tokenKind2{"goto", TOKEN_TAG_DEFAULT, 24}
var TOKEN_KIND_2_IF = &tokenKind2{"if", TOKEN_TAG_DEFAULT, 25}
var TOKEN_KIND_2_IMPLEMENTS = &tokenKind2{"implements", TOKEN_TAG_DEFAULT, 26}
var TOKEN_KIND_2_IMPORT = &tokenKind2{"import", TOKEN_TAG_DEFAULT, 27}
var TOKEN_KIND_2_INSTANCEOF = &tokenKind2{"instanceof", TOKEN_TAG_DEFAULT, 28}
var TOKEN_KIND_2_INT = &tokenKind2{"int", TOKEN_TAG_NAMED, 29}
var TOKEN_KIND_2_INTERFACE = &tokenKind2{"interface", TOKEN_TAG_DEFAULT, 30}
var TOKEN_KIND_2_LONG = &tokenKind2{"long", TOKEN_TAG_NAMED, 31}
var TOKEN_KIND_2_NATIVE = &tokenKind2{"native", TOKEN_TAG_DEFAULT, 32}
var TOKEN_KIND_2_NEW = &tokenKind2{"new", TOKEN_TAG_DEFAULT, 33}
var TOKEN_KIND_2_PACKAGE = &tokenKind2{"package", TOKEN_TAG_DEFAULT, 34}
var TOKEN_KIND_2_PRIVATE = &tokenKind2{"private", TOKEN_TAG_DEFAULT, 35}
var TOKEN_KIND_2_PROTECTED = &tokenKind2{"protected", TOKEN_TAG_DEFAULT, 36}
var TOKEN_KIND_2_PUBLIC = &tokenKind2{"public", TOKEN_TAG_DEFAULT, 37}
var TOKEN_KIND_2_RETURN = &tokenKind2{"return", TOKEN_TAG_DEFAULT, 38}
var TOKEN_KIND_2_SHORT = &tokenKind2{"short", TOKEN_TAG_NAMED, 39}
var TOKEN_KIND_2_STATIC = &tokenKind2{"static", TOKEN_TAG_DEFAULT, 40}

// java2开始有的一个关键字 声明 类、接口、方法，作用是严格按IEEE-754执行浮点数字计算，否则Java怎么来无法确认
var TOKEN_KIND_2_STRICTFP = &tokenKind2{"strictfp", TOKEN_TAG_DEFAULT, 41}
var TOKEN_KIND_2_SUPER = &tokenKind2{"super", TOKEN_TAG_NAMED, 42}
var TOKEN_KIND_2_SWITCH = &tokenKind2{"switch", TOKEN_TAG_DEFAULT, 43}
var TOKEN_KIND_2_SYNCHRONIZED = &tokenKind2{"synchronized", TOKEN_TAG_DEFAULT, 44}
var TOKEN_KIND_2_THIS = &tokenKind2{"this", TOKEN_TAG_NAMED, 45}
var TOKEN_KIND_2_THROW = &tokenKind2{"throw", TOKEN_TAG_DEFAULT, 46}
var TOKEN_KIND_2_THROWS = &tokenKind2{"throws", TOKEN_TAG_DEFAULT, 47}
var TOKEN_KIND_2_TRANSIENT = &tokenKind2{"transient", TOKEN_TAG_DEFAULT, 48}
var TOKEN_KIND_2_TRY = &tokenKind2{"try", TOKEN_TAG_DEFAULT, 49}
var TOKEN_KIND_2_VOID = &tokenKind2{"void", TOKEN_TAG_NAMED, 50}
var TOKEN_KIND_2_VOLATILE = &tokenKind2{"volatile", TOKEN_TAG_DEFAULT, 51}
var TOKEN_KIND_2_WHILE = &tokenKind2{"while", TOKEN_TAG_DEFAULT, 52}
var TOKEN_KIND_2_INT_LITERAL = &tokenKind2{"", TOKEN_TAG_NUMERIC, 53}
var TOKEN_KIND_2_LONG_LITERAL = &tokenKind2{"", TOKEN_TAG_NUMERIC, 54}
var TOKEN_KIND_2_FLOAT_LITERAL = &tokenKind2{"", TOKEN_TAG_NUMERIC, 55}
var TOKEN_KIND_2_DOUBLE_LITERAL = &tokenKind2{"", TOKEN_TAG_NUMERIC, 56}
var TOKEN_KIND_2_CHAR_LITERAL = &tokenKind2{"", TOKEN_TAG_NUMERIC, 57}
var TOKEN_KIND_2_STRING_LITERAL = &tokenKind2{"", TOKEN_TAG_STRING, 58}
var TOKEN_KIND_2_TRUE = &tokenKind2{"true", TOKEN_TAG_NAMED, 59}
var TOKEN_KIND_2_FALSE = &tokenKind2{"false", TOKEN_TAG_NAMED, 60}
var TOKEN_KIND_2_NULL = &tokenKind2{"null", TOKEN_TAG_NAMED, 61}
var TOKEN_KIND_2_UNDERSCORE = &tokenKind2{"_", TOKEN_TAG_NAMED, 62}
var TOKEN_KIND_2_ARROW = &tokenKind2{"->", TOKEN_TAG_DEFAULT, 63}
var TOKEN_KIND_2_COLCOL = &tokenKind2{"::", TOKEN_TAG_DEFAULT, 64}
var TOKEN_KIND_2_LPAREN = &tokenKind2{"(", TOKEN_TAG_DEFAULT, 65}
var TOKEN_KIND_2_RPAREN = &tokenKind2{")", TOKEN_TAG_DEFAULT, 66}
var TOKEN_KIND_2_LBRACE = &tokenKind2{"{", TOKEN_TAG_DEFAULT, 67}
var TOKEN_KIND_2_RBRACE = &tokenKind2{"}", TOKEN_TAG_DEFAULT, 68}
var TOKEN_KIND_2_LBRACKET = &tokenKind2{"[", TOKEN_TAG_DEFAULT, 69}
var TOKEN_KIND_2_RBRACKET = &tokenKind2{"]", TOKEN_TAG_DEFAULT, 70}
var TOKEN_KIND_2_SEMI = &tokenKind2{";", TOKEN_TAG_DEFAULT, 71}
var TOKEN_KIND_2_COMMA = &tokenKind2{",", TOKEN_TAG_DEFAULT, 72}
var TOKEN_KIND_2_DOT = &tokenKind2{".", TOKEN_TAG_DEFAULT, 73}
var TOKEN_KIND_2_ELLIPSIS = &tokenKind2{"...", TOKEN_TAG_DEFAULT, 74}
var TOKEN_KIND_2_EQ = &tokenKind2{"=", TOKEN_TAG_DEFAULT, 75}
var TOKEN_KIND_2_GT = &tokenKind2{">", TOKEN_TAG_DEFAULT, 76}
var TOKEN_KIND_2_LT = &tokenKind2{"<", TOKEN_TAG_DEFAULT, 77}
var TOKEN_KIND_2_BANG = &tokenKind2{"!", TOKEN_TAG_DEFAULT, 78}
var TOKEN_KIND_2_TILDE = &tokenKind2{"~", TOKEN_TAG_DEFAULT, 79}
var TOKEN_KIND_2_QUES = &tokenKind2{"?", TOKEN_TAG_DEFAULT, 80}
var TOKEN_KIND_2_COLON = &tokenKind2{":", TOKEN_TAG_DEFAULT, 81}
var TOKEN_KIND_2_EQEQ = &tokenKind2{"==", TOKEN_TAG_DEFAULT, 82}
var TOKEN_KIND_2_LTEQ = &tokenKind2{"<=", TOKEN_TAG_DEFAULT, 83}
var TOKEN_KIND_2_GTEQ = &tokenKind2{">=", TOKEN_TAG_DEFAULT, 84}
var TOKEN_KIND_2_BANGEQ = &tokenKind2{"!=", TOKEN_TAG_DEFAULT, 85}
var TOKEN_KIND_2_AMPAMP = &tokenKind2{"&&", TOKEN_TAG_DEFAULT, 86}
var TOKEN_KIND_2_BARBAR = &tokenKind2{"||", TOKEN_TAG_DEFAULT, 87}
var TOKEN_KIND_2_PLUSPLUS = &tokenKind2{"++", TOKEN_TAG_DEFAULT, 88}
var TOKEN_KIND_2_SUBSUB = &tokenKind2{"--", TOKEN_TAG_DEFAULT, 89}
var TOKEN_KIND_2_PLUS = &tokenKind2{"+", TOKEN_TAG_DEFAULT, 90}
var TOKEN_KIND_2_SUB = &tokenKind2{"-", TOKEN_TAG_DEFAULT, 91}
var TOKEN_KIND_2_STAR = &tokenKind2{"*", TOKEN_TAG_DEFAULT, 92}
var TOKEN_KIND_2_SLASH = &tokenKind2{"/", TOKEN_TAG_DEFAULT, 93}
var TOKEN_KIND_2_AMP = &tokenKind2{"&", TOKEN_TAG_DEFAULT, 94}
var TOKEN_KIND_2_BAR = &tokenKind2{"|", TOKEN_TAG_DEFAULT, 95}
var TOKEN_KIND_2_CARET = &tokenKind2{"^", TOKEN_TAG_DEFAULT, 96}
var TOKEN_KIND_2_PERCENT = &tokenKind2{"%", TOKEN_TAG_DEFAULT, 97}
var TOKEN_KIND_2_LTLT = &tokenKind2{"<<", TOKEN_TAG_DEFAULT, 98}
var TOKEN_KIND_2_GTGT = &tokenKind2{">>", TOKEN_TAG_DEFAULT, 99}
var TOKEN_KIND_2_GTGTGT = &tokenKind2{">>>", TOKEN_TAG_DEFAULT, 100}
var TOKEN_KIND_2_PLUSEQ = &tokenKind2{"+=", TOKEN_TAG_DEFAULT, 101}
var TOKEN_KIND_2_SUBEQ = &tokenKind2{"-=", TOKEN_TAG_DEFAULT, 102}
var TOKEN_KIND_2_STAREQ = &tokenKind2{"*=", TOKEN_TAG_DEFAULT, 103}
var TOKEN_KIND_2_SLASHEQ = &tokenKind2{"/=", TOKEN_TAG_DEFAULT, 104}
var TOKEN_KIND_2_AMPEQ = &tokenKind2{"&=", TOKEN_TAG_DEFAULT, 105}
var TOKEN_KIND_2_BAREQ = &tokenKind2{"|=", TOKEN_TAG_DEFAULT, 106}
var TOKEN_KIND_2_CARETEQ = &tokenKind2{"^=", TOKEN_TAG_DEFAULT, 107}
var TOKEN_KIND_2_PERCENTEQ = &tokenKind2{"%=", TOKEN_TAG_DEFAULT, 108}
var TOKEN_KIND_2_LTLTEQ = &tokenKind2{"<<=", TOKEN_TAG_DEFAULT, 109}
var TOKEN_KIND_2_GTGTEQ = &tokenKind2{">>=", TOKEN_TAG_DEFAULT, 110}
var TOKEN_KIND_2_GTGTGTEQ = &tokenKind2{">>>=", TOKEN_TAG_DEFAULT, 111}
var TOKEN_KIND_2_MONKEYS_AT = &tokenKind2{"@", TOKEN_TAG_DEFAULT, 112}
var TOKEN_KIND_2_CUSTOM = &tokenKind2{"", TOKEN_TAG_DEFAULT, 113}

// Filter接口 ---
func (t2 *tokenKind2) Accept(t interface{}) bool {

	tk := t.(tokenKind2)
	return t2.Index == tk.Index
}
func (t2 *tokenKind2) Filter_() {}

// Filter接口 ---
