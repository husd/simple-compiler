package parser

type tokenTag int

const (
	DEFAULT tokenTag = 1
	NAMED   tokenTag = 2
	STRING  tokenTag = 3
	NUMERIC tokenTag = 4
)

type tokenKind struct {
	Name string
	Tag  tokenTag
}

var TOKEN_KIND_EOF = tokenKind{"eof", DEFAULT}
var TOKEN_KIND_ERROR = tokenKind{"error", DEFAULT}
var TOKEN_KIND_IDENTIFIER = tokenKind{"", NAMED}
var TOKEN_KIND_ABSTRACT = tokenKind{"abstract", DEFAULT}
var TOKEN_KIND_ASSERT = tokenKind{"assert", NAMED}
var TOKEN_KIND_BOOLEAN = tokenKind{"boolean", NAMED}
var TOKEN_KIND_BREAK = tokenKind{"break", DEFAULT}
var TOKEN_KIND_BYTE = tokenKind{"byte", NAMED}
var TOKEN_KIND_CASE = tokenKind{"case", DEFAULT}
var TOKEN_KIND_CATCH = tokenKind{"catch", DEFAULT}
var TOKEN_KIND_CHAR = tokenKind{"char", NAMED}
var TOKEN_KIND_CLASS = tokenKind{"class", DEFAULT}
var TOKEN_KIND_CONST = tokenKind{"const", DEFAULT}
var TOKEN_KIND_CONTINUE = tokenKind{"continue", DEFAULT}
var TOKEN_KIND_DEF = tokenKind{"default", DEFAULT}
var TOKEN_KIND_DO = tokenKind{"do", DEFAULT}
var TOKEN_KIND_DOUBLE = tokenKind{"double", NAMED}
var TOKEN_KIND_ELSE = tokenKind{"else", DEFAULT}
var TOKEN_KIND_ENUM = tokenKind{"enum", NAMED}
var TOKEN_KIND_EXTENDS = tokenKind{"extends", DEFAULT}
var TOKEN_KIND_FINAL = tokenKind{"final", DEFAULT}
var TOKEN_KIND_FINALLY = tokenKind{"finally", DEFAULT}
var TOKEN_KIND_FLOAT = tokenKind{"float", NAMED}
var TOKEN_KIND_FOR = tokenKind{"for", DEFAULT}
var TOKEN_KIND_GOTO = tokenKind{"goto", DEFAULT}
var TOKEN_KIND_IF = tokenKind{"if", DEFAULT}
var TOKEN_KIND_IMPLEMENTS = tokenKind{"implements", DEFAULT}
var TOKEN_KIND_IMPORT = tokenKind{"import", DEFAULT}
var TOKEN_KIND_INSTANCEOF = tokenKind{"instanceof", DEFAULT}
var TOKEN_KIND_INT = tokenKind{"int", NAMED}
var TOKEN_KIND_INTERFACE = tokenKind{"interface", DEFAULT}
var TOKEN_KIND_LONG = tokenKind{"long", NAMED}
var TOKEN_KIND_NATIVE = tokenKind{"native", DEFAULT}
var TOKEN_KIND_NEW = tokenKind{"new", DEFAULT}
var TOKEN_KIND_PACKAGE = tokenKind{"package", DEFAULT}
var TOKEN_KIND_PRIVATE = tokenKind{"private", DEFAULT}
var TOKEN_KIND_PROTECTED = tokenKind{"protected", DEFAULT}
var TOKEN_KIND_PUBLIC = tokenKind{"public", DEFAULT}
var TOKEN_KIND_RETURN = tokenKind{"return", DEFAULT}
var TOKEN_KIND_SHORT = tokenKind{"short", NAMED}
var TOKEN_KIND_STATIC = tokenKind{"static", DEFAULT}
var TOKEN_KIND_STRICTFP = tokenKind{"strictfp", DEFAULT}
var TOKEN_KIND_SUPER = tokenKind{"super", NAMED}
var TOKEN_KIND_SWITCH = tokenKind{"switch", DEFAULT}
var TOKEN_KIND_SYNCHRONIZED = tokenKind{"synchronized", DEFAULT}
var TOKEN_KIND_THIS = tokenKind{"this", NAMED}
var TOKEN_KIND_THROW = tokenKind{"throw", DEFAULT}
var TOKEN_KIND_THROWS = tokenKind{"throws", DEFAULT}
var TOKEN_KIND_TRANSIENT = tokenKind{"transient", DEFAULT}
var TOKEN_KIND_TRY = tokenKind{"try", DEFAULT}
var TOKEN_KIND_VOID = tokenKind{"void", NAMED}
var TOKEN_KIND_VOLATILE = tokenKind{"volatile", DEFAULT}
var TOKEN_KIND_WHILE = tokenKind{"while", DEFAULT}
var TOKEN_KIND_INT_LITERAL = tokenKind{"", NUMERIC}
var TOKEN_KIND_LONG_LITERAL = tokenKind{"", NUMERIC}
var TOKEN_KIND_FLOAT_LITERAL = tokenKind{"", NUMERIC}
var TOKEN_KIND_DOUBLE_LITERAL = tokenKind{"", NUMERIC}
var TOKEN_KIND_CHAR_LITERAL = tokenKind{"", NUMERIC}
var TOKEN_KIND_STRING_LITERAL = tokenKind{"", STRING}
var TOKEN_KIND_TRUE = tokenKind{"true", NAMED}
var TOKEN_KIND_FALSE = tokenKind{"false", NAMED}
var TOKEN_KIND_NULL = tokenKind{"null", NAMED}
var TOKEN_KIND_UNDERSCORE = tokenKind{"_", NAMED}
var TOKEN_KIND_ARROW = tokenKind{"->", DEFAULT}
var TOKEN_KIND_COLCOL = tokenKind{"::", DEFAULT}
var TOKEN_KIND_LPAREN = tokenKind{"(", DEFAULT}
var TOKEN_KIND_RPAREN = tokenKind{")", DEFAULT}
var TOKEN_KIND_LBRACE = tokenKind{"{", DEFAULT}
var TOKEN_KIND_RBRACE = tokenKind{"}", DEFAULT}
var TOKEN_KIND_LBRACKET = tokenKind{"[", DEFAULT}
var TOKEN_KIND_RBRACKET = tokenKind{"]", DEFAULT}
var TOKEN_KIND_SEMI = tokenKind{";", DEFAULT}
var TOKEN_KIND_COMMA = tokenKind{",", DEFAULT}
var TOKEN_KIND_DOT = tokenKind{".", DEFAULT}
var TOKEN_KIND_ELLIPSIS = tokenKind{"...", DEFAULT}
var TOKEN_KIND_EQ = tokenKind{"=", DEFAULT}
var TOKEN_KIND_GT = tokenKind{">", DEFAULT}
var TOKEN_KIND_LT = tokenKind{"<", DEFAULT}
var TOKEN_KIND_BANG = tokenKind{"!", DEFAULT}
var TOKEN_KIND_TILDE = tokenKind{"~", DEFAULT}
var TOKEN_KIND_QUES = tokenKind{"?", DEFAULT}
var TOKEN_KIND_COLON = tokenKind{":", DEFAULT}
var TOKEN_KIND_EQEQ = tokenKind{"==", DEFAULT}
var TOKEN_KIND_LTEQ = tokenKind{"<=", DEFAULT}
var TOKEN_KIND_GTEQ = tokenKind{">=", DEFAULT}
var TOKEN_KIND_BANGEQ = tokenKind{"!=", DEFAULT}
var TOKEN_KIND_AMPAMP = tokenKind{"&&", DEFAULT}
var TOKEN_KIND_BARBAR = tokenKind{"||", DEFAULT}
var TOKEN_KIND_PLUSPLUS = tokenKind{"++", DEFAULT}
var TOKEN_KIND_SUBSUB = tokenKind{"--", DEFAULT}
var TOKEN_KIND_PLUS = tokenKind{"+", DEFAULT}
var TOKEN_KIND_SUB = tokenKind{"-", DEFAULT}
var TOKEN_KIND_STAR = tokenKind{"*", DEFAULT}
var TOKEN_KIND_SLASH = tokenKind{"/", DEFAULT}
var TOKEN_KIND_AMP = tokenKind{"&", DEFAULT}
var TOKEN_KIND_BAR = tokenKind{"|", DEFAULT}
var TOKEN_KIND_CARET = tokenKind{"^", DEFAULT}
var TOKEN_KIND_PERCENT = tokenKind{"%", DEFAULT}
var TOKEN_KIND_LTLT = tokenKind{"<<", DEFAULT}
var TOKEN_KIND_GTGT = tokenKind{">>", DEFAULT}
var TOKEN_KIND_GTGTGT = tokenKind{">>>", DEFAULT}
var TOKEN_KIND_PLUSEQ = tokenKind{"+=", DEFAULT}
var TOKEN_KIND_SUBEQ = tokenKind{"-=", DEFAULT}
var TOKEN_KIND_STAREQ = tokenKind{"*=", DEFAULT}
var TOKEN_KIND_SLASHEQ = tokenKind{"/=", DEFAULT}
var TOKEN_KIND_AMPEQ = tokenKind{"&=", DEFAULT}
var TOKEN_KIND_BAREQ = tokenKind{"|=", DEFAULT}
var TOKEN_KIND_CARETEQ = tokenKind{"^=", DEFAULT}
var TOKEN_KIND_PERCENTEQ = tokenKind{"%=", DEFAULT}
var TOKEN_KIND_LTLTEQ = tokenKind{"<<=", DEFAULT}
var TOKEN_KIND_GTGTEQ = tokenKind{">>=", DEFAULT}
var TOKEN_KIND_GTGTGTEQ = tokenKind{">>>=", DEFAULT}
var TOKEN_KIND_MONKEYS_AT = tokenKind{"@", DEFAULT}
var TOKEN_KIND_CUSTOM = tokenKind{"", DEFAULT}
