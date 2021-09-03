package parser

/**
 *
 * @author hushengdong
 */

type tokenKind struct {
	Name  string
	Tag   tokenTag
	Index int // 在key 数组中的位置
}

var TOKEN_KIND_EOF = &tokenKind{"eof", TOKEN_TAG_DEFAULT, 0} // EOF特指文件结束
var TOKEN_KIND_ERROR = &tokenKind{"error", TOKEN_TAG_DEFAULT, 1}
var TOKEN_KIND_IDENTIFIER = &tokenKind{"", TOKEN_TAG_NAMED, 2}
var TOKEN_KIND_ABSTRACT = &tokenKind{"abstract", TOKEN_TAG_DEFAULT, 3}
var TOKEN_KIND_ASSERT = &tokenKind{"assert", TOKEN_TAG_NAMED, 4}
var TOKEN_KIND_BOOLEAN = &tokenKind{"boolean", TOKEN_TAG_NAMED, 5}
var TOKEN_KIND_BREAK = &tokenKind{"break", TOKEN_TAG_DEFAULT, 6}
var TOKEN_KIND_BYTE = &tokenKind{"byte", TOKEN_TAG_NAMED, 7}
var TOKEN_KIND_CASE = &tokenKind{"case", TOKEN_TAG_DEFAULT, 8}
var TOKEN_KIND_CATCH = &tokenKind{"catch", TOKEN_TAG_DEFAULT, 9}
var TOKEN_KIND_CHAR = &tokenKind{"char", TOKEN_TAG_NAMED, 10}
var TOKEN_KIND_CLASS = &tokenKind{"class", TOKEN_TAG_DEFAULT, 11}
var TOKEN_KIND_CONST = &tokenKind{"const", TOKEN_TAG_DEFAULT, 12}
var TOKEN_KIND_CONTINUE = &tokenKind{"continue", TOKEN_TAG_DEFAULT, 13}
var TOKEN_KIND_DEF = &tokenKind{"default", TOKEN_TAG_DEFAULT, 14}
var TOKEN_KIND_DO = &tokenKind{"do", TOKEN_TAG_DEFAULT, 15}
var TOKEN_KIND_DOUBLE = &tokenKind{"double", TOKEN_TAG_NAMED, 16}
var TOKEN_KIND_ELSE = &tokenKind{"else", TOKEN_TAG_DEFAULT, 17}
var TOKEN_KIND_ENUM = &tokenKind{"enum", TOKEN_TAG_NAMED, 18}
var TOKEN_KIND_EXTENDS = &tokenKind{"extends", TOKEN_TAG_DEFAULT, 19}
var TOKEN_KIND_FINAL = &tokenKind{"final", TOKEN_TAG_DEFAULT, 20}
var TOKEN_KIND_FINALLY = &tokenKind{"finally", TOKEN_TAG_DEFAULT, 21}
var TOKEN_KIND_FLOAT = &tokenKind{"float", TOKEN_TAG_NAMED, 22}
var TOKEN_KIND_FOR = &tokenKind{"for", TOKEN_TAG_DEFAULT, 23}
var TOKEN_KIND_GOTO = &tokenKind{"goto", TOKEN_TAG_DEFAULT, 24}
var TOKEN_KIND_IF = &tokenKind{"if", TOKEN_TAG_DEFAULT, 25}
var TOKEN_KIND_IMPLEMENTS = &tokenKind{"implements", TOKEN_TAG_DEFAULT, 26}
var TOKEN_KIND_IMPORT = &tokenKind{"import", TOKEN_TAG_DEFAULT, 27}
var TOKEN_KIND_INSTANCEOF = &tokenKind{"instanceof", TOKEN_TAG_DEFAULT, 28}
var TOKEN_KIND_INT = &tokenKind{"int", TOKEN_TAG_NAMED, 29}
var TOKEN_KIND_INTERFACE = &tokenKind{"interface", TOKEN_TAG_DEFAULT, 30}
var TOKEN_KIND_LONG = &tokenKind{"long", TOKEN_TAG_NAMED, 31}
var TOKEN_KIND_NATIVE = &tokenKind{"native", TOKEN_TAG_DEFAULT, 32}
var TOKEN_KIND_NEW = &tokenKind{"new", TOKEN_TAG_DEFAULT, 33}
var TOKEN_KIND_PACKAGE = &tokenKind{"package", TOKEN_TAG_DEFAULT, 34}
var TOKEN_KIND_PRIVATE = &tokenKind{"private", TOKEN_TAG_DEFAULT, 35}
var TOKEN_KIND_PROTECTED = &tokenKind{"protected", TOKEN_TAG_DEFAULT, 36}
var TOKEN_KIND_PUBLIC = &tokenKind{"public", TOKEN_TAG_DEFAULT, 37}
var TOKEN_KIND_RETURN = &tokenKind{"return", TOKEN_TAG_DEFAULT, 38}
var TOKEN_KIND_SHORT = &tokenKind{"short", TOKEN_TAG_NAMED, 39}
var TOKEN_KIND_STATIC = &tokenKind{"static", TOKEN_TAG_DEFAULT, 40}
var TOKEN_KIND_STRICTFP = &tokenKind{"strictfp", TOKEN_TAG_DEFAULT, 41}
var TOKEN_KIND_SUPER = &tokenKind{"super", TOKEN_TAG_NAMED, 42}
var TOKEN_KIND_SWITCH = &tokenKind{"switch", TOKEN_TAG_DEFAULT, 43}
var TOKEN_KIND_SYNCHRONIZED = &tokenKind{"synchronized", TOKEN_TAG_DEFAULT, 44}
var TOKEN_KIND_THIS = &tokenKind{"this", TOKEN_TAG_NAMED, 45}
var TOKEN_KIND_THROW = &tokenKind{"throw", TOKEN_TAG_DEFAULT, 46}
var TOKEN_KIND_THROWS = &tokenKind{"throws", TOKEN_TAG_DEFAULT, 47}
var TOKEN_KIND_TRANSIENT = &tokenKind{"transient", TOKEN_TAG_DEFAULT, 48}
var TOKEN_KIND_TRY = &tokenKind{"try", TOKEN_TAG_DEFAULT, 49}
var TOKEN_KIND_VOID = &tokenKind{"void", TOKEN_TAG_NAMED, 50}
var TOKEN_KIND_VOLATILE = &tokenKind{"volatile", TOKEN_TAG_DEFAULT, 51}
var TOKEN_KIND_WHILE = &tokenKind{"while", TOKEN_TAG_DEFAULT, 52}
var TOKEN_KIND_INT_LITERAL = &tokenKind{"", TOKEN_TAG_NUMERIC, 53}
var TOKEN_KIND_LONG_LITERAL = &tokenKind{"", TOKEN_TAG_NUMERIC, 54}
var TOKEN_KIND_FLOAT_LITERAL = &tokenKind{"", TOKEN_TAG_NUMERIC, 55}
var TOKEN_KIND_DOUBLE_LITERAL = &tokenKind{"", TOKEN_TAG_NUMERIC, 56}
var TOKEN_KIND_CHAR_LITERAL = &tokenKind{"", TOKEN_TAG_NUMERIC, 57}
var TOKEN_KIND_STRING_LITERAL = &tokenKind{"", TOKEN_TAG_STRING, 58}
var TOKEN_KIND_TRUE = &tokenKind{"true", TOKEN_TAG_NAMED, 59}
var TOKEN_KIND_FALSE = &tokenKind{"false", TOKEN_TAG_NAMED, 60}
var TOKEN_KIND_NULL = &tokenKind{"null", TOKEN_TAG_NAMED, 61}
var TOKEN_KIND_UNDERSCORE = &tokenKind{"_", TOKEN_TAG_NAMED, 62}
var TOKEN_KIND_ARROW = &tokenKind{"->", TOKEN_TAG_DEFAULT, 63}
var TOKEN_KIND_COLCOL = &tokenKind{"::", TOKEN_TAG_DEFAULT, 64}
var TOKEN_KIND_LPAREN = &tokenKind{"(", TOKEN_TAG_DEFAULT, 65}
var TOKEN_KIND_RPAREN = &tokenKind{")", TOKEN_TAG_DEFAULT, 66}
var TOKEN_KIND_LBRACE = &tokenKind{"{", TOKEN_TAG_DEFAULT, 67}
var TOKEN_KIND_RBRACE = &tokenKind{"}", TOKEN_TAG_DEFAULT, 68}
var TOKEN_KIND_LBRACKET = &tokenKind{"[", TOKEN_TAG_DEFAULT, 69}
var TOKEN_KIND_RBRACKET = &tokenKind{"]", TOKEN_TAG_DEFAULT, 70}
var TOKEN_KIND_SEMI = &tokenKind{";", TOKEN_TAG_DEFAULT, 71}
var TOKEN_KIND_COMMA = &tokenKind{",", TOKEN_TAG_DEFAULT, 72}
var TOKEN_KIND_DOT = &tokenKind{".", TOKEN_TAG_DEFAULT, 73}
var TOKEN_KIND_ELLIPSIS = &tokenKind{"...", TOKEN_TAG_DEFAULT, 74}
var TOKEN_KIND_EQ = &tokenKind{"=", TOKEN_TAG_DEFAULT, 75}
var TOKEN_KIND_GT = &tokenKind{">", TOKEN_TAG_DEFAULT, 76}
var TOKEN_KIND_LT = &tokenKind{"<", TOKEN_TAG_DEFAULT, 77}
var TOKEN_KIND_BANG = &tokenKind{"!", TOKEN_TAG_DEFAULT, 78}
var TOKEN_KIND_TILDE = &tokenKind{"~", TOKEN_TAG_DEFAULT, 79}
var TOKEN_KIND_QUES = &tokenKind{"?", TOKEN_TAG_DEFAULT, 80}
var TOKEN_KIND_COLON = &tokenKind{":", TOKEN_TAG_DEFAULT, 81}
var TOKEN_KIND_EQEQ = &tokenKind{"==", TOKEN_TAG_DEFAULT, 82}
var TOKEN_KIND_LTEQ = &tokenKind{"<=", TOKEN_TAG_DEFAULT, 83}
var TOKEN_KIND_GTEQ = &tokenKind{">=", TOKEN_TAG_DEFAULT, 84}
var TOKEN_KIND_BANGEQ = &tokenKind{"!=", TOKEN_TAG_DEFAULT, 85}
var TOKEN_KIND_AMPAMP = &tokenKind{"&&", TOKEN_TAG_DEFAULT, 86}
var TOKEN_KIND_BARBAR = &tokenKind{"||", TOKEN_TAG_DEFAULT, 87}
var TOKEN_KIND_PLUSPLUS = &tokenKind{"++", TOKEN_TAG_DEFAULT, 88}
var TOKEN_KIND_SUBSUB = &tokenKind{"--", TOKEN_TAG_DEFAULT, 89}
var TOKEN_KIND_PLUS = &tokenKind{"+", TOKEN_TAG_DEFAULT, 90}
var TOKEN_KIND_SUB = &tokenKind{"-", TOKEN_TAG_DEFAULT, 91}
var TOKEN_KIND_STAR = &tokenKind{"*", TOKEN_TAG_DEFAULT, 92}
var TOKEN_KIND_SLASH = &tokenKind{"/", TOKEN_TAG_DEFAULT, 93}
var TOKEN_KIND_AMP = &tokenKind{"&", TOKEN_TAG_DEFAULT, 94}
var TOKEN_KIND_BAR = &tokenKind{"|", TOKEN_TAG_DEFAULT, 95}
var TOKEN_KIND_CARET = &tokenKind{"^", TOKEN_TAG_DEFAULT, 96}
var TOKEN_KIND_PERCENT = &tokenKind{"%", TOKEN_TAG_DEFAULT, 97}
var TOKEN_KIND_LTLT = &tokenKind{"<<", TOKEN_TAG_DEFAULT, 98}
var TOKEN_KIND_GTGT = &tokenKind{">>", TOKEN_TAG_DEFAULT, 99}
var TOKEN_KIND_GTGTGT = &tokenKind{">>>", TOKEN_TAG_DEFAULT, 100}
var TOKEN_KIND_PLUSEQ = &tokenKind{"+=", TOKEN_TAG_DEFAULT, 101}
var TOKEN_KIND_SUBEQ = &tokenKind{"-=", TOKEN_TAG_DEFAULT, 102}
var TOKEN_KIND_STAREQ = &tokenKind{"*=", TOKEN_TAG_DEFAULT, 103}
var TOKEN_KIND_SLASHEQ = &tokenKind{"/=", TOKEN_TAG_DEFAULT, 104}
var TOKEN_KIND_AMPEQ = &tokenKind{"&=", TOKEN_TAG_DEFAULT, 105}
var TOKEN_KIND_BAREQ = &tokenKind{"|=", TOKEN_TAG_DEFAULT, 106}
var TOKEN_KIND_CARETEQ = &tokenKind{"^=", TOKEN_TAG_DEFAULT, 107}
var TOKEN_KIND_PERCENTEQ = &tokenKind{"%=", TOKEN_TAG_DEFAULT, 108}
var TOKEN_KIND_LTLTEQ = &tokenKind{"<<=", TOKEN_TAG_DEFAULT, 109}
var TOKEN_KIND_GTGTEQ = &tokenKind{">>=", TOKEN_TAG_DEFAULT, 110}
var TOKEN_KIND_GTGTGTEQ = &tokenKind{">>>=", TOKEN_TAG_DEFAULT, 111}
var TOKEN_KIND_MONKEYS_AT = &tokenKind{"@", TOKEN_TAG_DEFAULT, 112}
var TOKEN_KIND_CUSTOM = &tokenKind{"", TOKEN_TAG_DEFAULT, 113}
