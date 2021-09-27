package parser

/**
 * 所有词法分析器解析出来的符号，都被归类为tokenKind了，肯定属于其中之一
 * 设计程序语言，这里可以是入口。
 * @author hushengdong
 */

// default 0 - 999
// named 1000 - 1999
// string 2000 - 2999
// num  3000 - 3999

// default
const TOKEN_KIND_2_EOF int = 0         // eof
const TOKEN_KIND_2_ERROR int = 1       // error
const TOKEN_KIND_2_ABSTRACT int = 2    // abstract
const TOKEN_KIND_2_BREAK int = 3       // break
const TOKEN_KIND_2_CASE int = 4        // case
const TOKEN_KIND_2_CATCH int = 5       // catch
const TOKEN_KIND_2_CLASS int = 6       // class
const TOKEN_KIND_2_CONST int = 7       // const
const TOKEN_KIND_2_CONTINUE int = 8    // continue
const TOKEN_KIND_2_DEF int = 9         // default
const TOKEN_KIND_2_DO int = 10         // do
const TOKEN_KIND_2_ELSE int = 11       // else
const TOKEN_KIND_2_EXTENDS int = 12    // extends
const TOKEN_KIND_2_FINAL int = 13      // final
const TOKEN_KIND_2_FINALLY int = 14    // finally
const TOKEN_KIND_2_FOR int = 15        // for
const TOKEN_KIND_2_GOTO int = 16       // goto
const TOKEN_KIND_2_IF int = 17         // if
const TOKEN_KIND_2_IMPLEMENTS int = 18 // implements
const TOKEN_KIND_2_IMPORT int = 19     // import
const TOKEN_KIND_2_INSTANCEOF int = 20 // instanceof
const TOKEN_KIND_2_INTERFACE int = 21  // interface
const TOKEN_KIND_2_NEW int = 22        // new
const TOKEN_KIND_2_PACKAGE int = 23    // package
const TOKEN_KIND_2_PRIVATE int = 24    // private
const TOKEN_KIND_2_PROTECTED int = 25  // protected
const TOKEN_KIND_2_PUBLIC int = 26     // public
const TOKEN_KIND_2_RETURN int = 27     // return
const TOKEN_KIND_2_NATIVE int = 28     // native
const TOKEN_KIND_2_STATIC int = 29     // static
// java2开始有的一个关键字 声明 类、接口、方法，作用是严格按IEEE-754执行浮点数字计算，否则Java怎么来无法确认
const TOKEN_KIND_2_STRICTFP int = 30     // strictfp
const TOKEN_KIND_2_SWITCH int = 31       // switch
const TOKEN_KIND_2_SYNCHRONIZED int = 32 // synchronized
const TOKEN_KIND_2_THROWS int = 33       // throws
const TOKEN_KIND_2_TRANSIENT int = 34    // transient
const TOKEN_KIND_2_TRY int = 35          // try
const TOKEN_KIND_2_THROW int = 36        // throw
const TOKEN_KIND_2_VOLATILE int = 37     // volatile
const TOKEN_KIND_2_WHILE int = 38        // while
const TOKEN_KIND_2_ARROW int = 39        // ->
const TOKEN_KIND_2_COLCOL int = 40       // ::
const TOKEN_KIND_2_LPAREN int = 41       // (
const TOKEN_KIND_2_RPAREN int = 42       // )
const TOKEN_KIND_2_LBRACE int = 43       // {
const TOKEN_KIND_2_RBRACE int = 44       // }
const TOKEN_KIND_2_LBRACKET int = 45     // [
const TOKEN_KIND_2_RBRACKET int = 46     // ]
const TOKEN_KIND_2_SEMI int = 47         // ;
const TOKEN_KIND_2_COMMA int = 48        // ,
const TOKEN_KIND_2_DOT int = 49          // .
const TOKEN_KIND_2_ELLIPSIS int = 50     // ...
const TOKEN_KIND_2_EQ int = 51           // int =
const TOKEN_KIND_2_GT int = 52           // >
const TOKEN_KIND_2_LT int = 53           // <
const TOKEN_KIND_2_BANG int = 54         // !
const TOKEN_KIND_2_TILDE int = 55        // ~
const TOKEN_KIND_2_QUES int = 56         // ?
const TOKEN_KIND_2_COLON int = 57        // :
const TOKEN_KIND_2_EQEQ int = 58         // int =  int =
const TOKEN_KIND_2_LTEQ int = 59         // <int =
const TOKEN_KIND_2_GTEQ int = 60         // >int =
const TOKEN_KIND_2_BANGEQ int = 61       // !int =
const TOKEN_KIND_2_AMPAMP int = 62       // &&
const TOKEN_KIND_2_BARBAR int = 63       // ||
const TOKEN_KIND_2_PLUSPLUS int = 64     // ++
const TOKEN_KIND_2_SUBSUB int = 65       // --
const TOKEN_KIND_2_PLUS int = 66         // +
const TOKEN_KIND_2_SUB int = 67          // -
const TOKEN_KIND_2_STAR int = 68         // *
const TOKEN_KIND_2_SLASH int = 69        // /
const TOKEN_KIND_2_STAREQ int = 70       // *int =
const TOKEN_KIND_2_AMP int = 71          // &
const TOKEN_KIND_2_BAR int = 72          // |
const TOKEN_KIND_2_CARET int = 73        // ^
const TOKEN_KIND_2_PERCENT int = 74      // %
const TOKEN_KIND_2_LTLT int = 75         // <<
const TOKEN_KIND_2_GTGT int = 76         // >>
const TOKEN_KIND_2_GTGTGT int = 77       // >>>
const TOKEN_KIND_2_PLUSEQ int = 78       // +int =
const TOKEN_KIND_2_SUBEQ int = 79        // -int =

const TOKEN_KIND_2_SLASHEQ int = 80    // /int =
const TOKEN_KIND_2_AMPEQ int = 81      // &int =
const TOKEN_KIND_2_BAREQ int = 82      // |int =
const TOKEN_KIND_2_CARETEQ int = 83    // ^int =
const TOKEN_KIND_2_PERCENTEQ int = 84  // %int =
const TOKEN_KIND_2_LTLTEQ int = 85     // <<int =
const TOKEN_KIND_2_GTGTEQ int = 86     // >>int =
const TOKEN_KIND_2_GTGTGTEQ int = 87   // >>>int =
const TOKEN_KIND_2_MONKEYS_AT int = 88 // @
const TOKEN_KIND_2_CUSTOM int = 89     // ,

// named
const TOKEN_KIND_2_IDENTIFIER int = 1000
const TOKEN_KIND_2_ASSERT int = 1001     // assert
const TOKEN_KIND_2_BOOLEAN int = 1002    // boolean
const TOKEN_KIND_2_BYTE int = 1003       // byte
const TOKEN_KIND_2_CHAR int = 1004       // char
const TOKEN_KIND_2_DOUBLE int = 1005     // double
const TOKEN_KIND_2_ENUM int = 1006       // enum
const TOKEN_KIND_2_FLOAT int = 1007      // float
const TOKEN_KIND_2_INT int = 1008        // float
const TOKEN_KIND_2_LONG int = 1009       // /long
const TOKEN_KIND_2_SHORT int = 1010      // short
const TOKEN_KIND_2_SUPER int = 1011      // super
const TOKEN_KIND_2_THIS int = 1012       // this
const TOKEN_KIND_2_VOID int = 1013       // void
const TOKEN_KIND_2_TRUE int = 1014       // true
const TOKEN_KIND_2_FALSE int = 1015      // false
const TOKEN_KIND_2_NULL int = 1016       // null
const TOKEN_KIND_2_UNDERSCORE int = 1017 // _

// string
const TOKEN_KIND_2_STRING_LITERAL int = 2000

// numeric
const TOKEN_KIND_2_INT_LITERAL int = 3001
const TOKEN_KIND_2_LONG_LITERAL int = 3002
const TOKEN_KIND_2_FLOAT_LITERAL int = 3003
const TOKEN_KIND_2_DOUBLE_LITERAL int = 3004
const TOKEN_KIND_2_CHAR_LITERAL int = 3005
