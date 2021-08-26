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

func (jt *JavaTokenizer) ReadToken() *Token {

	//TODO husd start 核心方法 解析出来下一个token是什么
	pos := int(0)
	reader := jt.reader

	reader.schStart = reader.bp // reset
loop:
	for {
		pos = reader.bp
		switch reader.ch {
		case '\t':
		case ' ':
		case Layout_char_ff: // 空格 tab 换页等
			for {
				reader.ReadRune()
				if reader.ch == '\t' || reader.ch == ' ' || reader.ch == Layout_char_ff {
					break
				}
			}
			//
			break
		case Layout_char_lf: // 换行
			reader.ReadRune()
			break
		case Layout_char_cr: // 回车
			reader.ReadRune()
			if reader.ch == Layout_char_lf { //有的操作系统是： CRLF
				reader.ReadRune()
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
			reader.ReadRune()
			if reader.ch == 'X' || reader.ch == 'x' {
				//16进制
				reader.ReadRune()
				jt.skipUnderLine()
				if reader.ch == '.' {
					jt.scanHexFractionAndSuffix(pos, false)
				}
			}
			break
		case '1':
		case '2':
		case '3':
		case '4':
		case '5':
		case '6':
		case '7':
		case '8':
		case '9':
			jt.scanNumber()
			goto loop
		case '.':
		case ',':
		case ';':
		case '(':
		case ')':
		case '{':
		case '}':
		case '[':
		case ']':
		case '/':
		case '\'':
		case '"':
			fmt.Println("遇到了特殊符号，单独处理")
		default:
			fmt.Println("默认逻辑----------------------")
		}
	}
	t := Token{}
	return &t
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
			reader.ReadRune() //这里直接就忽略了，但是我们的实现，没有忽略这些字符，暂时先这么做，后续优化 TODO husd
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
			reader.ReadRune()
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
				//true if the character is an ignorable control character that may be part of a Java or Unicode identifier; false otherwise.
				isIdentify = false
			}
			if !isIdentify {
				jt.name = reader.name()
				//jt.tk = tokens.lookupKind(name) //TODO husd
				return //end
			}
		}
		// break 之后，执行到这里了，扫描下一个字符
		reader.ReadRune()
	}
}

func (jt *JavaTokenizer) scanNumber() {

	//TODO husd
}

func (jt *JavaTokenizer) skipUnderLine() {

	if jt.reader.ch == '_' {
		jt.lexError(jt.reader.bp, "数字里有无效的下划线")
		for jt.reader.ch == '_' {
			jt.reader.ReadRune()
		}
	}
}

func (jt *JavaTokenizer) lexError(bp int, msg string) {

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
		reader.ReadRune()
		if reader.digit(pos, radix) < 0 && reader.ch == '_' {
			break
		}
	}
	if res == '_' {
		jt.lexError(savedPos, "无效的下划线")
	}
}
