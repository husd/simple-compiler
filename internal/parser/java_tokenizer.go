package parser

import (
	"fmt"
	"husd.com/v0/code"
	"husd.com/v0/compiler"
	"husd.com/v0/util"
)

/**
 *
 * 可以把源代码里的 String name = "xiaoming 你好" 分解成:
 * String
 * name
 * =
 * "xiaoming 你好"
 *
 * 可以看出，并不是简单的按某个字符分割，还需要过滤出注释 包括单行注释和多行注释。
 *
 */

type JavaTokenizer struct {
	reader *UnicodeReader // reader
	source code.JVersion  // jdk版本
	tk     *tokenKind     // 当前token的类型
	name   *util.Name     // identify name

	errPos int // 错误地址
	radix  int // 进制 set by nextToken().
}

func NewJavaTokenizer(path string) *JavaTokenizer {

	v := code.JDK8
	javaTokenizer := JavaTokenizer{}
	javaTokenizer.reader = NewUnicodeReaderFromFile(path)
	javaTokenizer.source = v
	javaTokenizer.tk = TOKEN_KIND_ERROR

	return &javaTokenizer
}

func (jt *JavaTokenizer) readToken() *Token {

	//TODO husd start 核心方法 解析出来下一个token是什么
	pos := 0
	reader := jt.reader

	reader.spos = 0 // reset spos
loop:
	for {
		pos = reader.bp
		switch reader.ch {
		case '\t':
		case ' ':
		case Layout_char_ff: // 空格 tab 换页等
			for {
				reader.scanRune()
				if reader.ch == '\t' || reader.ch == ' ' || reader.ch == Layout_char_ff {
					break
				}
			}
			//
			break
		case Layout_char_lf: // 换行
			reader.scanRune()
			break
		case Layout_char_cr: // 回车
			reader.scanRune()
			if reader.ch == Layout_char_lf { //有的操作系统是： CRLF
				reader.scanRune()
			}
			break
		case 'A':
		case 'B':
		case 'C':
		case 'D':
		case 'E':
		case 'F':
		case 'G':
		case 'H':
		case 'I':
		case 'J':
		case 'K':
		case 'L':
		case 'M':
		case 'N':
		case 'O':
		case 'P':
		case 'Q':
		case 'R':
		case 'S':
		case 'T':
		case 'U':
		case 'V':
		case 'W':
		case 'X':
		case 'Y':
		case 'Z':
		case 'a':
		case 'b':
		case 'c':
		case 'd':
		case 'e':
		case 'f':
		case 'g':
		case 'h':
		case 'i':
		case 'j':
		case 'k':
		case 'l':
		case 'm':
		case 'n':
		case 'o':
		case 'p':
		case 'q':
		case 'r':
		case 's':
		case 't':
		case 'u':
		case 'v':
		case 'w':
		case 'x':
		case 'y':
		case 'z':
		case '$':
		case '_': //java的标识符，只能以这些字符开头
			jt.scanIdentify()
			goto loop
		case '0': //0比较特殊，例如 0xF 0b10 等数字，需要单独处理
			//TODO husd
			reader.scanRune()
			if reader.ch == 'X' || reader.ch == 'x' {
				//16进制
				reader.scanRune()
				jt.skipUnderLine()
				if reader.ch == '.' {
					jt.scanHexFractionAndSuffix(pos, false)
				} else if reader.digit(pos, 16) < 0 {
					jt.lexError(pos, "无效的16进制数字")
				} else {
					jt.scanNumber(pos, 16)
				}
			} else if reader.ch == 'b' || reader.ch == 'B' {
				if !allowBinaryLiterals {
					jt.lexError(pos, "不允许使用二进制字面量", jt.source)
				}
				reader.scanRune()
				jt.skipUnderLine()
				if reader.digit(pos, 2) < 0 {
					jt.lexError(pos, "无效的二进制数字")
				} else {
					jt.scanNumber(pos, 2)
				}
			} else {
				reader.putChar('0')
				if reader.ch == '_' {
					savedPos := reader.bp
					for {
						reader.scanRune()
						if reader.ch != '_' {
							break
						}
					}
					if reader.digit(pos, 10) < 0 {
						jt.lexError(savedPos, "无效的下划线")
					}
				}
				jt.scanNumber(pos, 8)
			}
			goto loop
		case '1':
		case '2':
		case '3':
		case '4':
		case '5':
		case '6':
		case '7':
		case '8':
		case '9':
			jt.scanNumber(pos, 10)
			goto loop
		case '.':
			//Java里对于点的这种开头的数字
			reader.scanRune()
			if reader.ch >= '0' && reader.ch <= '9' {
				reader.putChar('.')
				jt.scanFractionAndSuffix(pos)
			} else if reader.ch == '.' {
				savedPos := reader.bp
				reader.putChar('.')
				reader.putRuneChar('.', true)
				if reader.ch == '.' {
					reader.scanRune()
					reader.putChar('.')
					jt.tk = TOKEN_KIND_ELLIPSIS
				} else {
					jt.lexError(savedPos, "无效的点 . ")
				}
			} else {
				jt.tk = TOKEN_KIND_DOT
			}
			goto loop
		case ',':
			reader.scanRune()
			jt.tk = TOKEN_KIND_COMMA
			goto loop
		case ';':
			reader.scanRune()
			jt.tk = TOKEN_KIND_SEMI
			goto loop
		case '(':
			reader.scanRune()
			jt.tk = TOKEN_KIND_LPAREN
			goto loop
		case ')':
			reader.scanRune()
			jt.tk = TOKEN_KIND_RPAREN
			goto loop
		case '{':
			reader.scanRune()
			jt.tk = TOKEN_KIND_LBRACE
			goto loop
		case '}':
			reader.scanRune()
			jt.tk = TOKEN_KIND_RBRACE
			goto loop
		case '[':
			reader.scanRune()
			jt.tk = TOKEN_KIND_LBRACKET
			goto loop
		case ']':
			reader.scanRune()
			jt.tk = TOKEN_KIND_RBRACKET
			goto loop
		case '/': // 这里涉及到了注释的解析
			reader.scanRune()
			if reader.ch == '/' { //读到了单行注释
				for {
					reader.scanRune()
					if reader.ch == Layout_char_cr ||
						reader.ch == Layout_char_lf ||
						reader.bp >= reader.size { // 一直读到换行或者结束
						break
					}
					// 我们不保留注释
					break
				}
			} else if reader.ch == '*' { //读到了多行注释
				for reader.bp < reader.size {
					reader.scanRune()
					if reader.ch == '*' {
						// 看下一个是不是 /
						reader.scanRune()
						if reader.ch == '/' { //找到了结束符号
							reader.scanRune()
							break
						}
					} else {
						//不是 * 就继续找
						reader.scanRune()
					}
				}
				if reader.bp >= reader.size {
					//找到结束，都没有找到多行注释的结束符号，直接报错
					jt.lexError(pos, "多行注释没有找到结束符号")
				}
			} else if reader.ch == '=' {
				// 对于Java语法来说，就是遇到了 /= 这种符号
				jt.tk = TOKEN_KIND_SLASHEQ
				reader.scanRune()
			} else {
				jt.tk = TOKEN_KIND_SLASH //单纯就是 /
			}
			goto loop
		case '\'': //遇到了单引号
			reader.scanRune()
			if reader.ch == '\'' { // ''
				jt.lexError(pos, "空 '' ")
			} else {
				if reader.ch == Layout_char_cr ||
					reader.ch == Layout_char_lf {
					jt.lexError(pos, "无效的换行")
				}
				jt.scanLitChar(pos)
				if reader.ch == '\'' { // 找到了 ''
					reader.scanRune()
					jt.tk = TOKEN_KIND_CHAR_LITERAL
				} else {
					jt.lexError(pos, "单引号没有关闭")
				}
			}
			goto loop
		case '"':
			reader.scanRune()
			for reader.ch != '"' &&
				reader.ch != Layout_char_cr &&
				reader.ch != Layout_char_lf &&
				reader.bp < reader.size {
				jt.scanLitChar(pos)
			}
			if reader.ch == '"' {
				jt.tk = TOKEN_KIND_STRING_LITERAL
				reader.scanRune()
			} else {
				jt.lexError(pos, "字符串没有关闭 只有一个双引号")
			}
		default:
			if isSpecialOperator(reader.ch) {
				jt.scanOperator()
			} else {
				//走到这里，默认就假设是标识符了 ，按标识符解析
				isJavaIdentiryStart := false
				if reader.ch < '\u0080' { // 小于 \u0080 的都是ascii码，已经处理过了，所以肯定不是标识符
					isJavaIdentiryStart = false
				} else {
					//TODO husd 判断是不是标识符
					isJavaIdentiryStart = true
				}
				if isJavaIdentiryStart {
					jt.scanIdentify()
				} else if reader.bp == reader.size ||
					reader.ch == Layout_char_eoi ||
					reader.bp+1 == reader.size { //JTS 3.5 主要说的EOI的问题
					jt.tk = TOKEN_KIND_EOF
					pos = reader.size
				} else {
					jt.lexError(pos, "无效的字符:", reader.ch)
					reader.scanRune()
				}
			}
			goto loop
		}
	}
	t := Token{}
	return &t
}

func isSpecialOperator(ch rune) bool {

	switch ch {
	case '!':
	case '%':
	case '&':
	case '*':
	case '?':
	case '+':
	case '-':
	case ':':
	case '<':
	case '=':
	case '>':
	case '^':
	case '|':
	case '~':
	case '@':
		return true
	default:
		return false
	}
	return false
}

func (jt *JavaTokenizer) scanIdentify() {

	reader := jt.reader
	isIdentify := false
	//var high rune

	reader.putRune(true)
	for {
		switch reader.ch {
		case 'A':
		case 'B':
		case 'C':
		case 'D':
		case 'E':
		case 'F':
		case 'G':
		case 'H':
		case 'I':
		case 'J':
		case 'K':
		case 'L':
		case 'M':
		case 'N':
		case 'O':
		case 'P':
		case 'Q':
		case 'R':
		case 'S':
		case 'T':
		case 'U':
		case 'V':
		case 'W':
		case 'X':
		case 'Y':
		case 'Z':
		case 'a':
		case 'b':
		case 'c':
		case 'd':
		case 'e':
		case 'f':
		case 'g':
		case 'h':
		case 'i':
		case 'j':
		case 'k':
		case 'l':
		case 'm':
		case 'n':
		case 'o':
		case 'p':
		case 'q':
		case 'r':
		case 's':
		case 't':
		case 'u':
		case 'v':
		case 'w':
		case 'x':
		case 'y':
		case 'z':
		case '$':
		case '_':
		case '0':
		case '1':
		case '2':
		case '3':
		case '4':
		case '5':
		case '6':
		case '7':
		case '8':
		case '9':
			break // continue to find next one
		case '\u0000':
		case '\u0001':
		case '\u0002':
		case '\u0003':
		case '\u0004':
		case '\u0005':
		case '\u0006':
		case '\u0007':
		case '\u0008':
		case '\u000E':
		case '\u000F':
		case '\u0010':
		case '\u0011':
		case '\u0012':
		case '\u0013':
		case '\u0014':
		case '\u0015':
		case '\u0016':
		case '\u0017':
		case '\u0018':
		case '\u0019':
		case '\u001B':
		case '\u007F':
			reader.scanRune() //这里直接就忽略了，但是我们的实现，没有忽略这些字符，暂时先这么做，后续优化 TODO husd
			continue
		case '\u001A': // EOI 也是一个有效的标识符 That's the Ctrl+Z control code.
			if reader.bp >= reader.size {
				// 这个名字，可能需要缓存，因为关键字很多，没必要非得转字符串，把字节计算一下hash
				// 能提高下性能，这里要存一下符号表了，可以指定一下字节数组的位置和长度，就行了。
				jt.name = reader.name() // 不仅仅有标识符的名字，还有其它属性 TODO husd
				//jt.tk = tokens.lookupKind(name) //TODO husd
				if compiler.DEBUG {
					fmt.Println("name is :", jt.name)
				}
				return
			}
			reader.scanRune()
			continue
		default:
			if reader.ch < '\u0080' {
				isIdentify = false //ascii已经考虑过了，直接结束
			} else {
				//这里有点复杂，不属于任何一个情况 为了简化功能，这里直接就是结束扫描了
				// TODO 判断其它是变量的标识符的情况，如果是，设置 isIdentify = true
				// 就会继续扫描下去
				// Determines if the specified character (Unicode code point) should be regarded as an ignorable character in a Java identifier or a Unicode identifier.
				//The following Unicode characters are ignorable in a Java identifier or a Unicode identifier:
				//ISO control characters that are not whitespace
				//'\u0000' through '\u0008'
				//'\u000E' through '\u001B'
				//'\u007F' through '\u009F'
				//all characters that have the FORMAT general category value
				//Params:
				//codePoint – the character (Unicode code point) to be tested.
				//Returns:
				//true if the character is an ignorable control character that may be part of a Java or Unicode identifier false otherwise.
				isIdentify = false
			}
			if !isIdentify {
				jt.name = reader.name()
				//jt.tk = tokens.lookupKind(name) //TODO husd
				return //end
			}
		}
		// break 之后，执行到这里了，扫描下一个字符
		reader.scanRune()
	}
}

func (jt *JavaTokenizer) scanNumber(pos int, radix int) {

	//TODO husd
}

func (jt *JavaTokenizer) skipUnderLine() {

	if jt.reader.ch == '_' {
		jt.lexError(jt.reader.bp, "数字里有无效的下划线")
		for jt.reader.ch == '_' {
			jt.reader.scanRune()
		}
	}
}

func (jt *JavaTokenizer) lexError(bp int, msg ...interface{}) {

	fmt.Println("词法解析错误，位置：", bp, " msg:", msg)
	jt.tk = TOKEN_KIND_ERROR
	jt.errPos = bp
}

// Read fractional part and 'd' or 'f' suffix of floating point number.
func (jt *JavaTokenizer) scanHexFractionAndSuffix(pos int, seenDigit bool) {

	jt.radix = 16
	reader := jt.reader
	reader.putRune(true)
	jt.skipUnderLine()
	if reader.digit(pos, 16) >= 0 {
		seenDigit = true
		jt.scanDigits(pos, 16)
	}
	if !seenDigit {
		jt.lexError(pos, "无效的16进制数字")
	} else {
		jt.scanHexExponentAndSuffix(pos)
	}
}

// Read fractional part of hexadecimal floating point number.
func (jt *JavaTokenizer) scanHexExponentAndSuffix(pos int) {

	reader := jt.reader
	if reader.ch == 'p' || reader.ch == 'P' {
		reader.putRune(true)
		jt.skipUnderLine()
		if reader.ch == '+' || reader.ch == '-' {
			reader.putRune(true)
		}
		jt.skipUnderLine()
		if reader.ch >= 0 && reader.ch <= 9 {
			jt.scanDigits(pos, 10)
			if !allowHexFloats {
				jt.lexError(pos, "不允许十六进制浮点文本")
			} else if !hexFloatsWork {
				jt.lexError(pos, "不允许十六进制浮点文本")
			}
		} else {
			jt.lexError(pos, "fp.lt 格式不正确")
		}
	} else {
		jt.lexError(pos, "fp.lt 格式不正确")
	}

	if reader.ch == 'f' || reader.ch == 'F' {
		reader.putRune(true)
		jt.tk = TOKEN_KIND_FLOAT_LITERAL
		jt.radix = 16
	} else {
		if reader.ch == 'd' || reader.ch == 'D' {
			reader.putRune(true)
			jt.tk = TOKEN_KIND_DOUBLE_LITERAL
			jt.radix = 16
		}
	}
}

// 扫描数字
func (jt *JavaTokenizer) scanDigits(pos int, radix int) {

	reader := jt.reader
	var res rune
	var savedPos int
	for {
		if reader.ch != '_' {
			reader.putRune(false)
		} else {
			if !allowUnderscoresInLiterals {
				jt.lexError(pos, "unsupported.underscore.li")
			}
			res = reader.ch
			savedPos = reader.bp
		}
		reader.scanRune()
		if reader.digit(pos, radix) < 0 && reader.ch == '_' {
			break
		}
	}
	if res == '_' {
		jt.lexError(savedPos, "无效的下划线")
	}
}

// read next char or string literal and copy into sbuf
func (jt *JavaTokenizer) scanLitChar(pos int) {

	reader := jt.reader
	if reader.ch == '\\' {
		if reader.peekChar() == '\\' && !reader.isUnicode() {
			reader.skipChar()
			reader.putRuneChar('\\', true)
		} else {
			reader.scanRune()
			switch reader.ch {
			case '0':
			case '1':
			case '2':
			case '3':
			case '4':
			case '5':
			case '6':
			case '7':
				leadCh := reader.ch
				oct := reader.digit(pos, 8)
				reader.scanRune()
				if reader.ch >= '0' && reader.ch <= '7' { //解析八进制的数字
					oct = oct*8 + reader.digit(pos, 8)
					reader.scanRune()
					if leadCh <= '3' && reader.ch >= '0' && reader.ch <= '7' {
						oct = oct*8 + reader.digit(pos, 8)
						reader.scanRune()
					}
				}
				reader.putChar(oct)
				break
			case 'b':
				reader.putRuneChar('\b', true)
				break
			case 't':
				reader.putRuneChar('\t', true)
				break
			case 'n':
				reader.putRuneChar('\n', true)
				break
			case 'f':
				reader.putRuneChar('\f', true)
				break
			case 'r':
				reader.putRuneChar('\r', true)
				break
			case '\'':
				reader.putRuneChar('\'', true)
				break
			case '"':
				reader.putRuneChar('"', true)
				break
			case '\\':
				reader.putRuneChar('\\', true)
				break
			default:
				jt.lexError(reader.bp, "illegal.esc.char")
			}
		}
	} else if reader.bp != reader.size {
		reader.putRune(true)
	}
}

// 扫描操作符号 需要尽可能的长
func (jt *JavaTokenizer) scanOperator() {

	reader := jt.reader
	for {
		reader.putRune(false)
		newName := reader.name()
		tk1 := lookupTokenKind(newName)
		if tk1 == TOKEN_KIND_IDENTIFIER { //?
			reader.spos--
			break
		}
		jt.tk = tk1
		reader.scanRune()
		if !isSpecialOperator(reader.ch) {
			break
		}
	}
}

func (jt *JavaTokenizer) scanFractionAndSuffix(pos int) {

	//TODO husd
}
